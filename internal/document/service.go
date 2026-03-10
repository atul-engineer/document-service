package document

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DocumentService struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewDocumentService(client *mongo.Client) *DocumentService {
	return &DocumentService{
		client: client,
		col: client.Database("documentsdb").Collection("documents"),
	}
}

func (srv *DocumentService) Insert(ctx context.Context, document *Document) (*mongo.InsertOneResult, error) {
	res, err := srv.col.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}
	return res, nil
}
