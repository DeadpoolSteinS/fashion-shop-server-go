package models

type Product struct {
	Id    string `bson:"_id"`
	Title string `bson:"title"`
	Price int    `bson:"price"`
	Image string `bson:"image"`
}
