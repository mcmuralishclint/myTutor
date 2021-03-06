package models

import (
	"context"

	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var LecturerInfoCollection *mongo.Collection
var SkillsCollection *mongo.Collection
var AdminCollection *mongo.Collection
var LecturerSkillCollection *mongo.Collection
var Client *mongo.Client

func ConnectDB() {
	Client, err := config.InitMongo()
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	LecturerInfoCollection = Client.Database("lecturer").Collection("lecturer-info")
	SkillsCollection = Client.Database("lecturer").Collection("skills")
	LecturerSkillCollection = Client.Database("lecturer").Collection("lecturer-skills")
	AdminCollection = Client.Database("adminUser").Collection("admin")
}
