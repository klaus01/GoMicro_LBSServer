package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	order "github.com/klaus01/GoMicro_LBSServer/srv/order/proto"
	"github.com/klaus01/GoMicro_LBSServer/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/errors"
)

const gServiceName = "go.micro.srv.order"

// Order srv
type Order struct {
	dbCollection *mongo.Collection
	dbContext    context.Context
}

// Get 获取订单信息
func (s *Order) Get(ctx context.Context, req *order.GetRequest, rep *order.OrderModel) error {
	const method string = "get"
	const id string = gServiceName + "." + method

	if len(req.OrderId) <= 0 {
		return errors.BadRequest(id, "缺少订单号")
	}

	filter := bson.M{"orderId": req.OrderId}
	err := s.dbCollection.FindOne(s.dbContext, filter).Decode(rep)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.NotFound(id, "订单号不存在")
		}
		return err
	}

	return nil
}

// Search 查询订单
func (s *Order) Search(ctx context.Context, req *order.SearchRequest, rep *order.SearchResult) error {
	if req.PageNo < 1 {
		req.PageNo = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 20
	} else if req.PageSize > 100 {
		req.PageSize = 100
	}

	filter := bson.M{}
	if req.BeginDateTime != nil && req.EndDateTime != nil {
		filter["createAt"] = bson.M{"$gte": req.BeginDateTime.Seconds, "$lte": req.EndDateTime.Seconds}
	} else if req.BeginDateTime != nil {
		filter["createAt"] = bson.M{"$gte": req.BeginDateTime.Seconds}
	} else if req.EndDateTime != nil {
		filter["createAt"] = bson.M{"$lte": req.EndDateTime.Seconds}
	}
	if len(req.OrderId) > 0 {
		filter["orderId"] = req.OrderId
	}
	if len(req.PhoneNumber) > 0 {
		filter["phoneNumber"] = req.PhoneNumber
	}

	count, err := s.dbCollection.CountDocuments(s.dbContext, filter)
	if err != nil {
		return err
	}
	cur, err := s.dbCollection.Find(s.dbContext, filter)
	if err != nil {
		return err
	}
	defer cur.Close(s.dbContext)
	for cur.Next(s.dbContext) {
		var orderModel order.OrderModel
		if err := cur.Decode(&orderModel); err != nil {
			return err
		}
		rep.Datas = append(rep.Datas, &orderModel)
	}
	if err := cur.Err(); err != nil {
		return err
	}

	rep.PageNo = req.PageNo
	rep.PageTotal = uint32(math.Ceil(float64(count) / float64(req.PageSize)))

	return nil
}

// Create Create
func (s *Order) Create(ctx context.Context, req *order.CreateRequest, rep *order.CreateResult) error {
	const method string = "create"
	const id string = gServiceName + "." + method

	if len(req.ProductName) <= 0 {
		return errors.BadRequest(id, "缺少商品名称")
	}
	if req.ProductAmount <= 0 {
		return errors.BadRequest(id, "缺少商品价格")
	}
	if len(req.Name) <= 0 {
		return errors.BadRequest(id, "缺少收货人姓名")
	}
	if len(req.PhoneNumber) <= 0 {
		return errors.BadRequest(id, "缺少收货人手机号")
	}
	if len(req.Province) <= 0 {
		return errors.BadRequest(id, "缺少省名称")
	}
	if len(req.City) <= 0 {
		return errors.BadRequest(id, "缺少市名称")
	}
	if len(req.Address) <= 0 {
		return errors.BadRequest(id, "缺少收货地址")
	}

	orderId := fmt.Sprintf("%v_%v", time.Now().UnixNano(), 1000+utils.RandomInt(8999))
	orderModel := order.OrderModel{
		OrderId:       orderId,
		CreateAt:      &timestamp.Timestamp{Seconds: time.Now().Unix()},
		ProductName:   req.ProductName,
		ProductAmount: req.ProductAmount,
		Name:          req.Name,
		PhoneNumber:   req.PhoneNumber,
		Province:      req.Province,
		City:          req.City,
		District:      req.District,
		Address:       req.Address,
		PayStatus:     order.OrderPayStatus_BE_PAID,
	}
	if _, err := s.dbCollection.InsertOne(s.dbContext, orderModel); err != nil {
		log.Println("[ERROR]", "插入订单失败", err)
		return err
	}

	return nil
}

// SetDeliveryInfo SetDeliveryInfo
func (s *Order) SetDeliveryInfo(ctx context.Context, req *order.SetDeliveryInfoRequest, rep *empty.Empty) error {
	const method string = "setDeliveryInfo"
	const id string = gServiceName + "." + method

	if len(req.OrderId) <= 0 {
		return errors.BadRequest(id, "缺少订单号")
	}
	if len(req.CourierCompany) <= 0 {
		return errors.BadRequest(id, "缺少快递公司")
	}
	if len(req.WaybillNumber) <= 0 {
		return errors.BadRequest(id, "缺少运单号")
	}

	deliveryInfo := order.OrderDeliveryInfo{
		CourierCompany: req.CourierCompany,
		WaybillNumber:  req.WaybillNumber,
		CreateAt:       &timestamp.Timestamp{Seconds: time.Now().Unix()},
	}
	filter := bson.M{"orderId": req.OrderId, "deliveryInfo": nil}
	update := bson.M{"deliveryInfo": deliveryInfo}
	updateResult, err := s.dbCollection.UpdateOne(s.dbContext, filter, update)
	if err != nil {
		log.Println("[ERROR]", "提交发货信息出错", req, err)
		return err
	}
	if updateResult.ModifiedCount <= 0 {
		return errors.BadRequest(id, "订单不存在或已填写发货信息")
	}

	return nil
}

// SetPayInfo SetPayInfo
func (s *Order) SetPayInfo(ctx context.Context, req *order.SetPayInfoRequest, rep *empty.Empty) error {
	const method string = "setPayInfo"
	const id string = gServiceName + "." + method

	if len(req.OrderId) <= 0 {
		return errors.BadRequest(id, "缺少订单号")
	}
	if len(req.ModeName) <= 0 {
		return errors.BadRequest(id, "缺少支付方式")
	}

	filter := bson.M{"orderId": req.OrderId}
	orderModel := order.OrderModel{}
	err := s.dbCollection.FindOne(s.dbContext, filter).Decode(&orderModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.BadRequest(id, fmt.Sprintf("订单%v不存在", req.OrderId))
		}
		return err
	}
	if orderModel.PayStatus != order.OrderPayStatus_BE_PAID {
		return errors.BadRequest(id, fmt.Sprintf("订单%v已支付过了", req.OrderId))
	}
	var payStatus order.OrderPayStatus
	if orderModel.ProductAmount == req.Money {
		payStatus = order.OrderPayStatus_PAID
	} else {
		payStatus = order.OrderPayStatus_PAY_EXCEPTION
	}

	payInfo := order.OrderPayInfo{
		ModeName: req.ModeName,
		Money:    req.Money,
		CreateAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
	}
	filter = bson.M{"orderId": req.OrderId}
	update := bson.M{"payStatus": payStatus, "payInfo": payInfo}
	if _, err := s.dbCollection.UpdateOne(s.dbContext, filter, update); err != nil {
		log.Println("[ERROR]", "提交发货信息出错", req, err)
		return err
	}

	return nil
}

func getCollection(db *mongo.Database) *mongo.Collection {
	collectionName := "orders"
	collection := db.Collection(collectionName)
	indexes := collection.Indexes()
	ctx := context.Background()
	cur, err := indexes.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	orderIDIndexExist := false
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
		if name == "orderId_1" {
			orderIDIndexExist = true
		}
	}

	if !orderIDIndexExist {
		log.Println(collectionName, "创建 orderId 索引")
		unique := true
		_, err := indexes.CreateOne(ctx, mongo.IndexModel{Keys: bsonx.Doc{{Key: "orderId", Value: bsonx.Int32(1)}}, Options: &options.IndexOptions{Unique: &unique}})
		if err != nil {
			log.Fatal(err)
		}
	}
	return collection
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

func loadConfigs() config.Config {
	dbConfig, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := dbConfig.Load(file.NewSource(
		file.WithPath("srv/config/db-dev.yaml"),
	)); err != nil {
		log.Fatal(err)
	}
	return dbConfig
}

func main() {
	dbConfig := loadConfigs()
	db := getDB(dbConfig.Get("uri").String(""), dbConfig.Get("dbName").String(""))
	dbCollection := getCollection(db)

	service := micro.NewService(micro.Name(gServiceName))
	order.RegisterOrderHandler(service.Server(), &Order{dbCollection, context.Background()})
	service.Init()
	service.Run()
}
