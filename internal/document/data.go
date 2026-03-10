package document

import "go.mongodb.org/mongo-driver/v2/bson"


type Document struct {
	ID      bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Title   string        `json:"title" bson:"title"`
	Content string        `json:"content" bson:"content"`
}
