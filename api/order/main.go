package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	order "github.com/klaus01/GoMicro_LBSServer/api/order/proto"
	srv_order "github.com/klaus01/GoMicro_LBSServer/srv/order/proto"
	smscode "github.com/klaus01/GoMicro_LBSServer/srv/smscode/proto"
	"github.com/klaus01/GoMicro_LBSServer/utils"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/errors"
)

const gServiceName = "go.micro.api.order"

// Order api
type Order struct {
	client client.Client
}

// Get Get
func (s *Order) Get(context context.Context, req *srv_order.GetRequest, rep *srv_order.OrderModel) error {
	const method string = "get"
	const id string = gServiceName + "." + method

	ctx, tr := utils.CreateTracing(context, gServiceName, method)
	defer tr.Finish()

	orderClient := srv_order.NewOrderService("go.micro.srv.order", s.client)
	srvRep, err := orderClient.Get(ctx, req)
	if err == nil {
		rep.OrderId = srvRep.OrderId
		rep.CreateAt = srvRep.CreateAt
		rep.ProductName = srvRep.ProductName
		rep.ProductAmount = srvRep.ProductAmount
		rep.Name = srvRep.Name
		rep.PhoneNumber = srvRep.PhoneNumber
		rep.Province = srvRep.Province
		rep.City = srvRep.City
		rep.District = srvRep.District
		rep.Address = srvRep.Address
		rep.PayStatus = srvRep.PayStatus
		rep.PayInfo = srvRep.PayInfo
		rep.DeliveryInfo = srvRep.DeliveryInfo
	}
	return err
}

// Search Search
func (s *Order) Search(context context.Context, req *srv_order.SearchRequest, rep *srv_order.SearchResult) error {
	const method string = "search"
	const id string = gServiceName + "." + method

	ctx, tr := utils.CreateTracing(context, gServiceName, method)
	defer tr.Finish()

	orderClient := srv_order.NewOrderService("go.micro.srv.order", s.client)
	srvRep, err := orderClient.Search(ctx, req)
	if err == nil {
		rep.PageNo = srvRep.PageNo
		rep.PageTotal = srvRep.PageTotal
		rep.Datas = srvRep.Datas
	}
	return err
}

// Create Create
func (s *Order) Create(context context.Context, req *order.APICreateRequest, rep *srv_order.CreateResult) error {
	const method string = "create"
	const id string = gServiceName + "." + method

	ctx, tr := utils.CreateTracing(context, gServiceName, method)
	defer tr.Finish()

	sig := fmt.Sprintf("CREATE%v%v%v%v%v%v%vORDER", req.ProductName, req.ProductAmount, req.Name, req.PhoneNumber, req.SmsCode, req.Address, req.Time)
	if utils.Md5(sig) != req.Sign {
		return errors.BadRequest(id, "sign 错误")
	}

	smscodeClient := smscode.NewSmscodeService("go.micro.srv.smscode", s.client)
	cvcRep, err := smscodeClient.CheckVerificationCode(ctx, &smscode.CheckVerificationCodeRequest{
		PhoneNumber: req.PhoneNumber,
		Code:        req.SmsCode,
	})
	if err != nil {
		return err
	}
	if !cvcRep.Success {
		return errors.BadRequest(id, "手机验证码错误或已过期")
	}

	srvReq := srv_order.CreateRequest{
		ProductName:   req.ProductName,
		ProductAmount: req.ProductAmount,
		Name:          req.Name,
		PhoneNumber:   req.PhoneNumber,
		Province:      req.Province,
		City:          req.City,
		District:      req.District,
		Address:       req.Address,
	}
	orderClient := srv_order.NewOrderService("go.micro.srv.order", s.client)
	srvRep, err := orderClient.Create(ctx, &srvReq)
	if err == nil {
		rep.OrderId = srvRep.OrderId
	}
	return err
}

// SetDeliveryInfo SetDeliveryInfo
func (s *Order) SetDeliveryInfo(context context.Context, req *srv_order.SetDeliveryInfoRequest, rep *empty.Empty) error {
	const method string = "search"
	const id string = gServiceName + "." + method

	ctx, tr := utils.CreateTracing(context, gServiceName, method)
	defer tr.Finish()

	orderClient := srv_order.NewOrderService("go.micro.srv.order", s.client)
	_, err := orderClient.SetDeliveryInfo(ctx, req)
	return err
}

func main() {
	service := micro.NewService(micro.Name(gServiceName))
	order.RegisterOrderAPIHandler(service.Server(), &Order{service.Client()})
	service.Init()
	service.Run()
}
