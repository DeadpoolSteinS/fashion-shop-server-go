package initializers

import (
	"context"
	"fashion-shop/constant"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ProductColl, CartColl *mongo.Collection

func ConnectToDB() {
	// get database url from .env file
	dns := os.Getenv("DB_URL")

	// connect to database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dns))
	constant.CheckError(err)
	// defer client.Disconnect(context.TODO())

	// get/create products collections
	ProductColl = client.Database("Fashion-Shop").Collection("products")
	CartColl = client.Database("Fashion-Shop").Collection("carts")

	// var result bson.M
	// keySearch := bson.D{{Key: "title", Value: "Zara - Jacket With Pockets"}}

	// err = ProductColl.FindOne(context.TODO(), keySearch).Decode(&result)
	// constant.CheckError(err)
	// fmt.Println(result)
}
