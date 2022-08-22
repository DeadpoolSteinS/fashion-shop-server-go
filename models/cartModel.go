package models

type Cart struct {
	Product  string `bson:"product"`
	Quantity int    `bson:"quantity"`
}
