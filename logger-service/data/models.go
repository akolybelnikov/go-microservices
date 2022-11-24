package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	Data      string    `json:"data,omitempty" bson:"data,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		LogEntry: LogEntry{},
	}
}

func (l *LogEntry) Insert(entry LogEntry) error {
	collection := client.Database("logger").Collection("log_entries")

	_, err := collection.InsertOne(context.Background(), LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Could not insert log entry: ", err)
		return err
	}

	return nil
}

func (l *LogEntry) FindAll() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("logger").Collection("log_entries")

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		log.Println("Could not find log entries: ", err)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println("Could not close cursor: ", err)
		}
	}(cursor, ctx)

	var entries []*LogEntry
	for cursor.Next(ctx) {
		var entry LogEntry
		err = cursor.Decode(&entry)
		if err != nil {
			log.Println("Could not decode log entry: ", err)
			return nil, err
		}

		entries = append(entries, &entry)
	}

	return entries, nil
}

func (l *LogEntry) getOne(id string) (*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("logger").Collection("log_entries")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Could not get log entry: ", err)
		return nil, err
	}

	var entry LogEntry
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry)
	if err != nil {
		log.Println("Could not get log entry: ", err)
		return nil, err
	}

	return &entry, nil
}

func (l *LogEntry) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("logger").Collection("log_entries")

	return collection.Drop(ctx)
}

func (l *LogEntry) Update() (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("logger").Collection("log_entries")

	docID, err := primitive.ObjectIDFromHex(l.ID)
	if err != nil {
		log.Println("Could not get log entry: ", err)
		return nil, err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{"$set", bson.D{
				{"name", l.Name},
				{"data", l.Data},
				{"updated_at", time.Now()},
			}},
		})
	if err != nil {
		log.Println("Could not update log entry: ", err)
		return nil, err
	}

	return result, nil
}