package mongoPkg

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"music-streaming/config"
	"music-streaming/vo"
)

type Wrap struct {
	client *mongo.Client
}

func NewMongoWrap(cfg config.Mongo) *Wrap {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d/", cfg.Host, cfg.Port))
	connect, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = connect.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return &Wrap{
		client: connect,
	}
}

func (obj *Wrap) AppCollection(collection vo.Collection) *mongo.Collection {
	return obj.client.Database(vo.App.String()).Collection(collection.String())
}
