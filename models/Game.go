package models

import ("go.mongodb.org/mongo-driver/bson/primitive")

type Game struct {
	ID                       primitive.ObjectID `json:"_id" bson:"_id, omitempty"`
	NAME                     string             `json:"name" bson:"name"`
	LINK                     string             `json:"link" bson:"link"`
	IMG                      string             `json:"img" bson:"img"`
	DISCOUNT                 string             `json:"discount" bson:"discount"`
	PRICE_WITHOUT_DISCOUNTED string             `json:"price_without_discounted" bson:"price_without_discounted"`
	PRICE_WITH_DISCOUNTED    string             `json:"price_with_discounted" bson:"price_with_discounted"`
}
