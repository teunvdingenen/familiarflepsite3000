package data

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/teunvdingenen/familiarflepsite3000/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type MongoDatastore struct {
	db      *mongo.Database
	Session *mongo.Client
	logger  *logrus.Logger
}

func (store *MongoDatastore) GetUserRepository() *UserRepository {
	return &UserRepository{store: store}
}

func (store *MongoDatastore) GetCollection(name string) *mongo.Collection {
	return store.db.Collection(name)
}

func NewDatastore(config *config.GeneralConfig, logger *logrus.Logger) *MongoDatastore {
	var mongoDataStore *MongoDatastore
	db, session := connect(config, logger)
	if db != nil && session != nil {
		mongoDataStore = new(MongoDatastore)
		mongoDataStore.db = db
		mongoDataStore.logger = logger
		mongoDataStore.Session = session
		return mongoDataStore
	}
	logger.Fatal("Failed to connect to Database")
	return nil
}

func connect(generalConfig *config.GeneralConfig, logger *logrus.Logger) (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo(generalConfig, logger)
	})
	return db, session
}

func connectToMongo(generalConfig *config.GeneralConfig, logger *logrus.Logger) (a *mongo.Database, b *mongo.Client) {
	var err error
	session, err := mongo.NewClient(options.Client().ApplyURI(generalConfig.DB.DatabaseHost))
	if err != nil {
		logger.Fatal(err)
	}
	session.Connect(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}
	var DB = session.Database(generalConfig.DB.DatabaseName)
	logger.Info("Connected to database ", generalConfig.DB.DatabaseName)
	return DB, session
}
