package main

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	sms "github.com/klaus01/GoMicro_LBSServer/srv/sms/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/metadata"
	"golang.org/x/net/trace"
)

// Sms srv
type Sms struct {
	db *mongo.Database
}

// SendVerificationCode 发送验证码
func (s *Sms) SendVerificationCode(ctx context.Context, req *sms.SendVerificationCodeRequest, rsp *empty.Empty) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}
	return nil
}

// SendOrderConfirmation 发送订单确认短信
func (s *Sms) SendOrderConfirmation(ctx context.Context, req *sms.SendOrderConfirmationRequest, rsp *empty.Empty) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}
	return nil
}

// SendShippingNotice 发送发货通知短信
func (s *Sms) SendShippingNotice(ctx context.Context, req *sms.SendShippingNoticeRequest, rsp *empty.Empty) error {
	return nil
}

// CheckVerificationCode 校验短信验证码
func (s *Sms) CheckVerificationCode(ctx context.Context, req *sms.CheckVerificationCodeRequest, rep *sms.CheckVerificationCodeResult) error {
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
			log.Println(collectionName, "索引 createAt 存在")
		} else if name == "phoneNumber_1" {
			phoneNumberIndexExist = true
			log.Println(collectionName, "索引 phoneNumber 存在")
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

func getConfig() (string, string, int32) {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := conf.Load(file.NewSource(
		file.WithPath("srv/config/db-dev.yml"),
	)); err != nil {
		log.Fatal(err)
	}
	dbURI := conf.Get("uri").String("")
	dbName := conf.Get("dbName").String("")

	if err := conf.Load(file.NewSource(
		file.WithPath("srv/config/yuntongxun-dev.yml"),
	)); err != nil {
		log.Fatal(err)
	}
	smsCodeExpireAfterSeconds := int32(conf.Get("yuntongxun", "smsCodeExpireAfterSeconds").Int(600))
	return dbURI, dbName, smsCodeExpireAfterSeconds
}

func main() {
	dbURI, dbName, smsCodeExpireAfterSeconds := getConfig()
	db := getDB(dbURI, dbName)
	checkCollection(db, smsCodeExpireAfterSeconds)

	service := micro.NewService(micro.Name("go.micro.srv.sms"))
	sms.RegisterSmsHandler(service.Server(), &Sms{db: db})
	service.Init()
	service.Run()
}
