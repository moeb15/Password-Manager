package database

import (
	"context"
	"log"
	"os"
	"pwdmanager_api/internal/helpers"
	"pwdmanager_api/pkg/models"
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
	insert, err := user_coll.InsertOne(ctx, bson.D{
		{Key: "name", Value: user.Username},
		{Key: "password", Value: helpers.HashPwd(user.Password)},
		{Key: "masterkey", Value: helpers.HashPwd(user.MasterKey)}})

	if err != nil {
		log.Fatal(err)
	}

	insert_id := insert.InsertedID.(primitive.ObjectID).Hex()
	user.ID = insert_id
	returned_user := models.User{ID: insert_id, Username: user.Username}
	return &returned_user
}

func (db *DB) FindUser(id string) (*models.User, error) {
	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &models.User{}, err
	}

	user_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	res := user_coll.FindOne(ctx, bson.M{"_id": obj_id})
	user := models.User{ID: id}

	res.Decode(&user)
	return &user, nil
}

func (db *DB) FindUserByName(name string) (*models.User, error) {
	user_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	res := user_coll.FindOne(ctx, bson.M{"name": name})
	user := models.User{Username: name}

	res.Decode(&user)
	return &user, nil
}

func (db *DB) CreatePassword(pwd models.Password, user models.User) *models.Password {
	pwd_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("passwords")
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	insert, err := pwd_coll.InsertOne(ctx, bson.D{
		{Key: "userid", Value: user.ID},
		{Key: "application", Value: pwd.Application},
		{Key: "password", Value: helpers.HashPwd(pwd.Password)},
	})

	if err != nil {
		log.Fatal(err)
	}

	insert_id := insert.InsertedID.(primitive.ObjectID).Hex()
	returned_pwd := models.Password{ID: insert_id, Application: pwd.Application}
	return &returned_pwd
}

func (db *DB) FindPasswords(user_id string) ([]*models.Password, error) {
	pwd_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("passwords")
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	cur, err := pwd_coll.Find(ctx, bson.M{"userid": user_id})
	if err != nil {
		return []*models.Password{}, err
	}

	var pwds []*models.Password
	for cur.Next(ctx) {
		sus, err := cur.Current.Elements()
		if err != nil {
			log.Fatal(err)
		}

		pwd := models.Password{ID: (sus[0].String()), UserID: (sus[1].String()),
			Application: (sus[2].String()), Password: (sus[3].String())}

		pwds = append(pwds, &pwd)
	}

	return pwds, nil
}
