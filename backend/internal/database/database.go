package database

import (
	"context"
	"errors"
	"log"
	"os"
	"pwdmanager_api/internal/helpers"
	"pwdmanager_api/pkg/models"
	"time"

	"slices"

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

	// creates user collection ,if it does not exist, on initial connection
	names, err := client.Database(os.Getenv("DB_NAME")).ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if !slices.Contains(names, "users") {
		err = client.Database(os.Getenv("DB_NAME")).CreateCollection(context.TODO(), "users")
		if err != nil {
			log.Fatal(err)
		}
	}
	// creates unique index in users collection
	mod := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	user_coll := client.Database(os.Getenv("DB_NAME")).Collection("users")
	_, err = user_coll.Indexes().CreateOne(context.TODO(), mod)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{client: client}
}

func (db *DB) CreateUser(user models.User) (*models.User, error) {
	user_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	insert, err := user_coll.InsertOne(ctx, bson.D{
		{Key: "name", Value: user.Username},
		{Key: "email", Value: user.Email},
		{Key: "password", Value: helpers.HashPwd(user.Password)},
		{Key: "masterkey", Value: helpers.HashPwd(user.MasterKey)}})

	if err != nil {
		return &models.User{}, err
	}

	insert_id := insert.InsertedID.(primitive.ObjectID).Hex()
	user.ID = insert_id
	returned_user := models.User{ID: insert_id, Username: user.Username, Email: user.Email}
	return &returned_user, nil
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

func (db *DB) FindUserByEmail(email string) (*models.User, error) {
	user_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	res := user_coll.FindOne(ctx, bson.M{"email": email})
	user := models.User{Email: email}

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
		{Key: "username", Value: pwd.Username},
		{Key: "password", Value: pwd.Password},
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
		var pwd models.Password
		err := cur.Decode(&pwd)
		if err != nil {
			return []*models.Password{}, err
		}
		pwds = append(pwds, &pwd)
	}

	return pwds, nil
}

func (db *DB) DeleteByApp(app_name, user_id string) (int, error) {
	pwd_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("passwords")
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	res, err := pwd_coll.DeleteOne(ctx, bson.D{
		{Key: "application", Value: app_name},
		{Key: "userid", Value: user_id}})
	if err != nil {
		return 0, err
	}
	return int(res.DeletedCount), nil
}

func (db *DB) RetrieveByApp(app_name, user_id string) models.Password {
	pwd_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("passwords")
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	res := pwd_coll.FindOne(ctx, bson.D{
		{Key: "application", Value: app_name},
		{Key: "userid", Value: user_id}})
	var pwd models.Password
	res.Decode(&pwd)

	return pwd
}

func (db *DB) UpdatePassword(app_name, new_pwd string, user models.User) (int, error) {
	pwd_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("passwords")
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	res, err := pwd_coll.UpdateOne(ctx, bson.D{
		{Key: "application", Value: app_name},
		{Key: "userid", Value: user.ID},
	}, bson.M{"$set": bson.D{{Key: "password", Value: new_pwd}}})
	if err != nil {
		return 0, err
	}
	if res.MatchedCount == 0 {
		return 0, errors.New("no matching applications")
	}
	return int(res.MatchedCount), nil
}

func (db *DB) DeleteAccount(user models.User) error {
	user_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("users")
	pwd_coll := db.client.Database(os.Getenv("DB_NAME")).Collection("passwords")

	obj_id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	_, err = pwd_coll.DeleteMany(ctx, bson.M{"userid": user.ID})
	if err != nil {
		return err
	}
	_, err = user_coll.DeleteOne(ctx, bson.M{"_id": obj_id})
	if err != nil {
		return err
	}

	return nil
}
