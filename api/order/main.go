package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	api_order "github.com/klaus01/GoMicro_LBSServer/api/order/proto"
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
func (s *Order) Get(context context.Context, req *api_order.GetRequest, rep *api_order.OrderModel) error {
	const method string = "get"
	const id string = gServiceName + "." + method

	ctx, tr := utils.CreateTracing(context, gServiceName, method)
	defer tr.Finish()

	orderClient := srv_order.NewOrderService("go.micro.srv.order", s.client)
	srvReq := srv_order.GetRequest{OrderId: req.OrderId}
	srvRep, err := orderClient.Get(ctx, &srvReq)
	if err == nil {
		convertSrvOrderToAPIOrder(srvRep, rep)
	}
	return err
}

// Search Search
func (s *Order) Search(context context.Context, req *api_order.SearchRequest, rep *api_order.SearchResult) error {
	const method string = "search"
	const id string = gServiceName + "." + method

	ctx, tr := utils.CreateTracing(context, gServiceName, method)
	defer tr.Finish()

	orderClient := srv_order.NewOrderService("go.micro.srv.order", s.client)
	srvReq := srv_order.SearchRequest{
		PageNo:        req.PageNo,
		PageSize:      req.PageSize,
		BeginDateTime: req.BeginDateTime,
		EndDateTime:   req.EndDateTime,
		PhoneNumber:   req.PhoneNumber,
		PayStatus:     srv_order.OrderPayStatus(req.PayStatus),
		IsShipped:     req.IsShipped,
	}
	srvRep, err := orderClient.Search(ctx, &srvReq)
	if err == nil {
		rep.PageNo = srvRep.PageNo
		rep.PageTotal = srvRep.PageTotal
		for _, data := range srvRep.Datas {
			newData := &api_order.OrderModel{}
			convertSrvOrderToAPIOrder(data, newData)
			rep.Datas = append(rep.Datas, newData)
		}
	}
	return err
}

// Create Create
func (s *Order) Create(context context.Context, req *api_order.CreateRequest, rep *api_order.CreateResult) error {
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

	orderClient := srv_order.NewOrderService("go.micro.srv.order", s.client)
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
	srvRep, err := orderClient.Create(ctx, &srvReq)
	if err == nil {
		rep.OrderId = srvRep.OrderId
	}
	return err
}

// SetDeliveryInfo SetDeliveryInfo
func (s *Order) SetDeliveryInfo(context context.Context, req *api_order.SetDeliveryInfoRequest, rep *empty.Empty) error {
	const method string = "search"
	const id string = gServiceName + "." + method

	ctx, tr := utils.CreateTracing(context, gServiceName, method)
	defer tr.Finish()

	orderClient := srv_order.NewOrderService("go.micro.srv.order", s.client)
	srvReq := srv_order.SetDeliveryInfoRequest{
		OrderId:        req.OrderId,
		CourierCompany: req.CourierCompany,
		WaybillNumber:  req.WaybillNumber,
	}
	_, err := orderClient.SetDeliveryInfo(ctx, &srvReq)
	return err
}

func convertSrvOrderToAPIOrder(srvOrder *srv_order.OrderModel, apiOrder *api_order.OrderModel) {
	apiOrder.OrderId = srvOrder.OrderId
	apiOrder.CreateAt = srvOrder.CreateAt
	apiOrder.ProductName = srvOrder.ProductName
	apiOrder.ProductAmount = srvOrder.ProductAmount
	apiOrder.Name = srvOrder.Name
	apiOrder.PhoneNumber = srvOrder.PhoneNumber
	apiOrder.Province = srvOrder.Province
	apiOrder.City = srvOrder.City
	apiOrder.District = srvOrder.District
	apiOrder.Address = srvOrder.Address
	apiOrder.PayStatus = api_order.OrderPayStatus(srvOrder.PayStatus)
	if srvOrder.PayInfo != nil {
		apiOrder.PayInfo.ModeName = srvOrder.PayInfo.ModeName
		apiOrder.PayInfo.Money = srvOrder.PayInfo.Money
		apiOrder.PayInfo.CreateAt = srvOrder.PayInfo.CreateAt
	}
	if srvOrder.DeliveryInfo != nil {
		apiOrder.DeliveryInfo.CourierCompany = srvOrder.DeliveryInfo.CourierCompany
		apiOrder.DeliveryInfo.WaybillNumber = srvOrder.DeliveryInfo.WaybillNumber
		apiOrder.DeliveryInfo.CreateAt = srvOrder.DeliveryInfo.CreateAt
	}
}

func main() {
	service := micro.NewService(micro.Name(gServiceName))
	service.Init()
	if err := api_order.RegisterOrderHandler(service.Server(), &Order{service.Client()}); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
