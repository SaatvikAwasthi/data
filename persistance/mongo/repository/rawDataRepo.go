package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"tester/persistance/mongo/dao"
)

const (
	collectionName = "rawData"
)

type rawDataRepo struct {
	db *mongo.Database
}

func NewDataRepo(db *mongo.Database) *rawDataRepo {
	return &rawDataRepo{
		db: db,
	}
}

func (repo *rawDataRepo) Add(ctx context.Context, data dao.RawData) error {
	collection := repo.db.Collection(collectionName)
	session, err := repo.db.Client().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	if _, err := session.WithTransaction(ctx, func(sessCtx context.Context) (interface{}, error) {
		return collection.InsertOne(ctx, data)
	}); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("duplicate raw data entry")
		}
		return err
	}

	return nil
}

func (repo *rawDataRepo) GetAll(ctx context.Context) ([]dao.RawData, error) {
	collection := repo.db.Collection(collectionName)
	cursor, err := collection.Find(ctx, bson.D{{}}, options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []dao.RawData
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (repo *rawDataRepo) GetByID(ctx context.Context, id string) (*dao.RawData, error) {
	var result dao.RawData
	collection := repo.db.Collection(collectionName)
	bsonID, _ := bson.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, bson.M{"_id": bsonID}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (repo *rawDataRepo) GetByTime(ctx context.Context, createdAt string) (*dao.RawData, error) {
	collection := repo.db.Collection(collectionName)
	var result dao.RawData
	err := collection.FindOne(ctx, bson.M{"created_at": createdAt}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
