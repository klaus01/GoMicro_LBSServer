package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	sms "github.com/klaus01/GoMicro_LBSServer/srv/sms/proto"
	"github.com/klaus01/GoMicro_LBSServer/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

var (
	dbConfig         config.Config
	yuntongxunConfig config.Config
)

// Sms srv
type Sms struct {
	db *mongo.Database
}

// CreateVerificationCode 生成短信验证码
func (s *Sms) CreateVerificationCode(ctx context.Context, req *sms.CreateVerificationCodeRequest, rep *sms.CreateVerificationCodeResult) error {
	return nil
}

// CheckVerificationCode 校验短信验证码
func (s *Sms) CheckVerificationCode(ctx context.Context, req *sms.CheckVerificationCodeRequest, rep *sms.CheckVerificationCodeResult) error {
	return nil
}

// SendVerificationCode 发送验证码
func (s *Sms) SendVerificationCode(ctx context.Context, req *sms.SendVerificationCodeRequest, rsp *empty.Empty) error {
	expirationTime := yuntongxunConfig.Get("smsCodeExpireAfterSeconds").Int(600) / 60
	templateID := yuntongxunConfig.Get("templateIDs", "verificationCode").String("")
	return postSendSMS(req.PhoneNumber, templateID, []string{req.Code, strconv.Itoa(expirationTime)})
}

// SendOrderConfirmation 发送订单确认短信
func (s *Sms) SendOrderConfirmation(ctx context.Context, req *sms.SendOrderConfirmationRequest, rsp *empty.Empty) error {
	templateID := yuntongxunConfig.Get("templateIDs", "orderConfirmation").String("")
	return postSendSMS(req.PhoneNumber, templateID, []string{req.Name, req.OrderId})
}

// SendShippingNotice 发送发货通知短信
func (s *Sms) SendShippingNotice(ctx context.Context, req *sms.SendShippingNoticeRequest, rsp *empty.Empty) error {
	templateID := yuntongxunConfig.Get("templateIDs", "shippingNotice").String("")
	return postSendSMS(req.PhoneNumber, templateID, []string{req.CourierCompany, req.WaybillNumber})
}

func postSendSMS(phoneNumber string, templateID string, parameters []string) error {
	accountSID := yuntongxunConfig.Get("accountSID").String("")
	authToken := yuntongxunConfig.Get("authToken").String("")
	appID := yuntongxunConfig.Get("appID").String("")

	timeNow := time.Now().Format("20060102030405")
	sig := fmt.Sprintf("%s%s%s", accountSID, authToken, timeNow)
	sig = strings.ToUpper(utils.Md5(sig))
	authorization := fmt.Sprintf("%s:%s", accountSID, timeNow)
	authorization = base64.URLEncoding.EncodeToString([]byte(authorization))
	postJSON := map[string]interface{}{"to": phoneNumber, "appId": appID, "templateId": templateID, "datas": parameters}
	postData, err := json.Marshal(postJSON)
	if err != nil {
		log.Println("[ERROR]", "JSON 序列化失败", err)
		return err
	}

	url := fmt.Sprintf("https://app.cloopen.com:8883/2013-12-26/Accounts/%s/SMS/TemplateSMS?sig=%s", accountSID, sig)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
	if err != nil {
		log.Println("[ERROR]", "NewRequest 失败", err)
		return err
	}
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR]", "client.Do 失败", err)
		return err
	}
	defer resp.Body.Close()

	httpCode := strings.TrimSpace(resp.Status)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR]", "httpCode", httpCode, "read body error", err)
		return err
	}
	var bodyMap map[string]interface{}
	err = json.Unmarshal(b, &bodyMap)
	if err != nil {
		log.Println("[ERROR]", "json.Unmarshal", err, "json string:", string(b))
		return err
	}
	ytxCode := bodyMap["statusCode"].(string)
	if httpCode != "200" || ytxCode != "000000" {
		log.Println("[ERROR]", "send sms fail", string(b))
		message := "发送短信失败"
		if len(ytxCode) > 0 || bodyMap["statusMsg"] != nil {
			var ytxMsg string
			if bodyMap["statusMsg"] != nil {
				ytxMsg = bodyMap["statusMsg"].(string)
			}
			message = fmt.Sprintf("%s %s%s", message, ytxCode, ytxMsg)
		} else if len(httpCode) > 0 {
			message = fmt.Sprintf("%s code:%s", message, httpCode)
		}
		return errors.New(message)
	}
	log.Println("[INFO]", "send sms success", string(b))
	return nil
}

func checkCollection(db *mongo.Database, smsCodeExpireAfterSeconds int32) {
	collectionName := "checkCollection"
	collection := db.Collection(collectionName)
	indexes := collection.Indexes()
	ctx := context.Background()
	cur, err := indexes.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	createAtIndexExist, phoneNumberIndexExist := false, false
	for cur.Next(ctx) {
		index := bson.D{}
		err := cur.Decode(&index)
		if err != nil {
			log.Fatal(err)
		}
		var name interface{}
		for _, item := range index {
			if item.Key == "name" {
				name = item.Value
				break
			}
		}
		if name == "createAt_1" {
			createAtIndexExist = true
			// log.Println(collectionName, "索引 createAt 存在")
		} else if name == "phoneNumber_1" {
			phoneNumberIndexExist = true
			// log.Println(collectionName, "索引 phoneNumber 存在")
		}
	}

	if !createAtIndexExist {
		log.Println(collectionName, "创建 createAt 索引")
		_, err := indexes.CreateOne(ctx, mongo.IndexModel{Keys: bsonx.Doc{{Key: "createAt", Value: bsonx.Int32(1)}}, Options: &options.IndexOptions{ExpireAfterSeconds: &smsCodeExpireAfterSeconds}})
		if err != nil {
			log.Fatal(err)
		}
	}
	if !phoneNumberIndexExist {
		log.Println(collectionName, "创建 phoneNumber 索引")
		unique := true
		_, err := indexes.CreateOne(ctx, mongo.IndexModel{Keys: bsonx.Doc{{Key: "phoneNumber", Value: bsonx.Int32(1)}}, Options: &options.IndexOptions{Unique: &unique}})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getDB(uri string, name string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(name)
}

func loadConfigs() (config.Config, config.Config) {
	dbConfig, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := dbConfig.Load(file.NewSource(
		file.WithPath("srv/config/db-dev.yaml"),
	)); err != nil {
		log.Fatal(err)
	}

	yuntongxunConfig, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := yuntongxunConfig.Load(file.NewSource(
		file.WithPath("srv/config/yuntongxun-dev.yaml"),
	)); err != nil {
		log.Fatal(err)
	}
	return dbConfig, yuntongxunConfig
}

func main() {
	dbConfig, yuntongxunConfig = loadConfigs()
	db := getDB(dbConfig.Get("uri").String(""), dbConfig.Get("dbName").String(""))
	smsCodeExpireAfterSeconds := int32(yuntongxunConfig.Get("smsCodeExpireAfterSeconds").Int(600))
	checkCollection(db, smsCodeExpireAfterSeconds)

	service := micro.NewService(micro.Name("go.micro.srv.sms"))
	sms.RegisterSmsHandler(service.Server(), &Sms{db: db})
	service.Init()
	service.Run()
}
