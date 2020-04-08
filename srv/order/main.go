package main

import (
	"context"
	"errors"
	"log"
	"math"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	order "github.com/klaus01/GoMicro_LBSServer/srv/order/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

// Order srv
type Order struct {
	dbCollection *mongo.Collection
	dbContext    context.Context
}

// Get 获取订单信息
func (s *Order) Get(ctx context.Context, req *order.GetRequest, rep *order.OrderModel) error {
	if len(req.OrderId) <= 0 {
		return errors.New("缺少 OrderId")
	}

	filter := bson.M{"orderId": req.OrderId}
	err := s.dbCollection.FindOne(s.dbContext, filter).Decode(rep)
	if err != nil {
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
	return nil
}

// SetDeliveryInfo SetDeliveryInfo
func (s *Order) SetDeliveryInfo(ctx context.Context, req *order.SetDeliveryInfoRequest, rep *empty.Empty) error {
	return nil
}

// SetPayInfo SetPayInfo
func (s *Order) SetPayInfo(ctx context.Context, req *order.SetPayInfoRequest, rep *empty.Empty) error {
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

	service := micro.NewService(micro.Name("go.micro.srv.order"))
	order.RegisterOrderHandler(service.Server(), &Order{dbCollection, context.Background()})
	service.Init()
	service.Run()
}
