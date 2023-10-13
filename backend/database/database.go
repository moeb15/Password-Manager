package database

import (
	"context"
	"log"
	"os"
	"pwdmanager_api/helpers"
	"pwdmanager_api/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect(dbURL string) *DB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURL))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{client: client}
}

func (db *DB) CreateUser(user models.User) *models.User {
	user_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	insert, err := user_coll.InsertOne(ctx, bson.D{{Key: "name", Value: user.Username},
		{Key: "password", Value: helpers.HashPwd(user.Password)},
		{Key: "masterkey", Value: helpers.HashPwd(user.MasterKey)}})

	if err != nil {
		log.Fatal(err)
	}

	insert_id := insert.InsertedID.(primitive.ObjectID).Hex()
	returned_user := models.User{ID: insert_id, Username: user.Username}
	return &returned_user
}
