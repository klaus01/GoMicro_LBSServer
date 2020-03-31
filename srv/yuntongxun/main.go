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
	yuntongxun "github.com/klaus01/GoMicro_LBSServer/srv/yuntongxun/proto"
	"github.com/klaus01/GoMicro_LBSServer/utils"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

// Yuntongxun srv
type Yuntongxun struct {
	config config.Config
}

// SendVerificationCode 发送验证码
func (s *Yuntongxun) SendVerificationCode(ctx context.Context, req *yuntongxun.SendVerificationCodeRequest, rsp *empty.Empty) error {
	expirationTime := s.config.Get("smsCodeExpireAfterSeconds").Int(600) / 60
	templateID := s.config.Get("templateIDs", "verificationCode").String("")
	return postSendSMS(s.config, req.PhoneNumber, templateID, []string{req.Code, strconv.Itoa(expirationTime)})
}

// SendOrderConfirmation 发送订单确认短信
func (s *Yuntongxun) SendOrderConfirmation(ctx context.Context, req *yuntongxun.SendOrderConfirmationRequest, rsp *empty.Empty) error {
	templateID := s.config.Get("templateIDs", "orderConfirmation").String("")
	return postSendSMS(s.config, req.PhoneNumber, templateID, []string{req.Name, req.OrderId})
}

// SendShippingNotice 发送发货通知短信
func (s *Yuntongxun) SendShippingNotice(ctx context.Context, req *yuntongxun.SendShippingNoticeRequest, rsp *empty.Empty) error {
	templateID := s.config.Get("templateIDs", "shippingNotice").String("")
	return postSendSMS(s.config, req.PhoneNumber, templateID, []string{req.CourierCompany, req.WaybillNumber})
}

func postSendSMS(yuntongxunConfig config.Config, phoneNumber string, templateID string, parameters []string) error {
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

func loadConfigs() config.Config {
	result, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := result.Load(file.NewSource(
		file.WithPath("srv/config/yuntongxun-dev.yaml"),
	)); err != nil {
		log.Fatal(err)
	}
	return result
}

func main() {
	config := loadConfigs()

	service := micro.NewService(micro.Name("go.micro.srv.yuntongxun"))
	yuntongxun.RegisterYuntongxunHandler(service.Server(), &Yuntongxun{config})
	service.Init()
	service.Run()
}
