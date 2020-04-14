package main

import (
	"context"
	"fmt"
	"log"

	sms "github.com/klaus01/GoMicro_LBSServer/api/sms/proto"
	smscode "github.com/klaus01/GoMicro_LBSServer/srv/smscode/proto"
	yuntongxun "github.com/klaus01/GoMicro_LBSServer/srv/yuntongxun/proto"
	"github.com/klaus01/GoMicro_LBSServer/utils"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/errors"
)

const gServiceName = "go.micro.api.sms"

// Sms api
type Sms struct {
	client client.Client
}

// SendVerificationCode 发送验证码
func (s *Sms) SendVerificationCode(context context.Context, req *sms.Request, rep *sms.Response) error {
	const method string = "sendVerificationCode"
	const id string = gServiceName + "." + method

	ctx, tr := utils.CreateTracing(context, gServiceName, method)
	defer tr.Finish()

	if len(req.PhoneNumber) <= 0 {
		return errors.BadRequest(id, "缺少手机号")
	}
	if len(req.Time) <= 0 {
		return errors.BadRequest(id, "缺少参数 time")
	}
	if len(req.Sign) <= 0 {
		return errors.BadRequest(id, "缺少参数 sign")
	}
	sig := fmt.Sprintf("SMS%sCODE%sS", req.PhoneNumber, req.Time)
	if utils.Md5(sig) != req.Sign {
		return errors.BadRequest(id, "sign 错误")
	}

	smscodeClient := smscode.NewSmscodeService("go.micro.srv.smscode", s.client)
	cvcRep, err := smscodeClient.CreateVerificationCode(ctx, &smscode.CreateVerificationCodeRequest{PhoneNumber: req.PhoneNumber})
	if err != nil {
		return err
	}

	yuntongxunClient := yuntongxun.NewYuntongxunService("go.micro.srv.yuntongxun", s.client)
	if _, err := yuntongxunClient.SendVerificationCode(ctx, &yuntongxun.SendVerificationCodeRequest{PhoneNumber: req.PhoneNumber, Code: cvcRep.Code}); err != nil {
		return err
	}

	return nil
}

func main() {
	service := micro.NewService(micro.Name(gServiceName))
	service.Init()
	if err := sms.RegisterSmsHandler(service.Server(), &Sms{service.Client()}); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
