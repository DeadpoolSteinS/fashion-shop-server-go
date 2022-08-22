package controllers

import (
	"context"
	"encoding/json"
	"fashion-shop/constant"
	"fashion-shop/initializers"
	"fashion-shop/models"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

// func CreateProduct(c *gin.Context) {
// 	var body struct {
// 		Title string
// 		Price string
// 	}
// 	c.Bind(&body)

// 	product := models.Product{
// 		Title: body.Title,
// 		Price: body.Price,
// 	}
// 	result, err := initializers.ProductColl.InsertOne(context.TODO(), product)
// 	constant.CheckError(err)
// 	fmt.Println(result.InsertedID)

// 	c.JSON(200, gin.H{
// 		"product": product,
// 	})
// }

func GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var products []models.Product

	result, err := initializers.ProductColl.Find(context.TODO(), bson.D{{}})
	constant.CheckError(err)

	for result.Next(context.TODO()) {
		var elem models.Product
		err := result.Decode(&elem)
		constant.CheckError(err)
		products = append(products, elem)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
	// c.JSON(products)
}

func AddToCart(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var res map[string]interface{}
	json.NewDecoder(r.Body).Decode(&res)

	id := res["id"]
	fmt.Println(id)
	bsonId := bson.D{{Key: "title", Value: id}}

	var product models.Product
	err := initializers.ProductColl.FindOne(context.TODO(), bsonId).Decode(&product)
	fmt.Println(product)
	constant.CheckError(err)

	// get all product in carts collection
	var carts []models.Cart
	result, err := initializers.CartColl.Find(context.TODO(), bson.D{{}})
	constant.CheckError(err)

	for result.Next(context.TODO()) {
		var elem models.Cart
		err := result.Decode(&elem)
		constant.CheckError(err)
		carts = append(carts, elem)
	}

	// insert product to carts collection
	if len(carts) == 0 {
		cart := models.Cart{Product: product.Id, Quantity: 1}
		initializers.CartColl.InsertOne(context.TODO(), cart)
	} else {
		for i := 0; i <= len(carts); i++ {
			if i == len(carts) {
				cart := models.Cart{Product: product.Id, Quantity: 1}
				initializers.CartColl.InsertOne(context.TODO(), cart)
			} else if carts[i].Id == product.Id {
				update := bson.D{{Key: "$set", Value: bson.D{{Key: "quantity", Value: carts[i].Quantity + 1}}}}
				initializers.CartColl.UpdateOne(context.TODO(), id, update)
				break
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	// json.NewEncoder(w).Encode()
}
