package document

import "go.mongodb.org/mongo-driver/v2/bson"


type Document struct {
	ID      bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Title   string        `json:"title" bson:"title"`
	Content string        `json:"content" bson:"content"`
}


type DocumentEvent struct {
	DocumentID bson.ObjectID `json:"document_id" bson:"document_id"`
	EventType  string        `json:"event_type" bson:"event_type"`
	Timestamp  int64         `json:"timestamp" bson:"timestamp"`
}