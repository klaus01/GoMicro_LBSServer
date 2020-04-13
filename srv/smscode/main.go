package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	smscode "github.com/klaus01/GoMicro_LBSServer/srv/smscode/proto"
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

const gServiceName = "go.micro.srv.smscode"

// Smscode srv
type Smscode struct {
	dbCollection *mongo.Collection
	dbContext    context.Context
}

// CreateVerificationCode 生成短信验证码
func (s *Smscode) CreateVerificationCode(ctx context.Context, req *smscode.CreateVerificationCodeRequest, rep *smscode.CreateVerificationCodeResult) error {
	const method string = "createVerificationCode"
	const id string = gServiceName + "." + method

	if len(req.PhoneNumber) <= 0 {
		return errors.BadRequest(id, "缺少手机号")
	}

	filter := bson.M{"phoneNumber": req.PhoneNumber}
	update := bson.M{"$set": bson.M{"createAt": time.Now()}}
	var smscodeModel smscode.SmscodeModel
	err := s.dbCollection.FindOneAndUpdate(s.dbContext, filter, update).Decode(&smscodeModel)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Println("[ERROR]", "查询", req.PhoneNumber, "短信失败", err)
			return err
		}

		smscodeModel.PhoneNumber = req.PhoneNumber
		smscodeModel.Code = fmt.Sprintf("%d", 1000+utils.RandomInt(8999))
		smscodeModel.CreateAt = &timestamp.Timestamp{Seconds: time.Now().Unix()}
		_, err := s.dbCollection.InsertOne(s.dbContext, smscodeModel)
		if err != nil {
			log.Println("[ERROR]", "插入短信失败", err)
			return err
		}
	}

	rep.Code = smscodeModel.Code

	return nil
}

// CheckVerificationCode 校验短信验证码
func (s *Smscode) CheckVerificationCode(ctx context.Context, req *smscode.CheckVerificationCodeRequest, rep *smscode.CheckVerificationCodeResult) error {
	const method string = "checkVerificationCode"
	const id string = gServiceName + "." + method

	if len(req.PhoneNumber) <= 0 {
		return errors.BadRequest(id, "缺少手机号")
	}
	if len(req.Code) <= 0 {
		return errors.BadRequest(id, "缺少验证码")
	}

	filter := bson.M{"phoneNumber": req.PhoneNumber, "code": req.Code}
	err := s.dbCollection.FindOneAndDelete(s.dbContext, filter).Err()
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Println("[ERROR]", "查询", req.PhoneNumber, "短信失败", err)
			return err
		}

		rep.Success = false
	} else {
		rep.Success = true
	}

	return nil
}

func getCollection(db *mongo.Database, smsCodeExpireAfterSeconds int32) *mongo.Collection {
	collectionName := "smscodes"
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
	dbConfig, yuntongxunConfig := loadConfigs()
	db := getDB(dbConfig.Get("uri").String(""), dbConfig.Get("dbName").String(""))
	smsCodeExpireAfterSeconds := int32(yuntongxunConfig.Get("smsCodeExpireAfterSeconds").Int(600))
	dbCollection := getCollection(db, smsCodeExpireAfterSeconds)

	service := micro.NewService(micro.Name(gServiceName))
	smscode.RegisterSmscodeHandler(service.Server(), &Smscode{dbCollection, context.Background()})
	service.Init()
	service.Run()
}
