package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Config() *viper.Viper {
	vi := viper.New()
	vi.AddConfigPath("./config")
	vi.SetConfigName("config")
	vi.SetConfigType("yaml")
	err := vi.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return vi
}

func Conectmongo() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI(Config().GetString("MONGODB_HOST")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	usersCollection := client.Database(Config().GetString("MONGODB_DATABASE")).Collection(Config().GetString("MONGODB_COLLECTION"))
	return usersCollection
}
