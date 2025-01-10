package repository

import (
	"context"
	"errors"
	"event-service/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// import "go.mongodb.org/mongo-driver/v2/mongo"
type EventRepository interface {
	Create(event *domain.Event) error
	GetById(id string) (*domain.Event, error)
	GetByUserId(id string) ([]*domain.Event, error)
	GetAll() ([]*domain.Event, error)
	Update(event *domain.Event) error
	DeleteById(id string) error
}
type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Client, dbName string, collectionName string) *MongoRepository {
	return &MongoRepository{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

func (r *MongoRepository) Create(event *domain.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//hceck if event already exists
	filter := bson.M{"id": event.ID}
	var existingEvent domain.Event
	err := r.collection.FindOne(ctx, filter).Decode(&existingEvent)
	if err != nil {
		return errors.New("user already exists")
	}
	_, err = r.collection.InsertOne(ctx, event)
	return err

}

func (r *MongoRepository) GetById(id string) (*domain.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": id}

	var event domain.Event
	err := r.collection.FindOne(ctx, filter).Decode(&event)
	if err != nil {
		return nil, errors.New("event not found")
	}
	return &event, nil
}

func (r *MongoRepository) GetByUserId(id string) ([]*domain.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": id}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.New("no events found")
	}
	defer cursor.Close(ctx)
	var events []*domain.Event
	for cursor.Next(ctx) {
		var event domain.Event
		if err := cursor.Decode(&event); err != nil {
			return nil, errors.New("ran into an error while fetching your events")
		}
		events = append(events, &event)
	}
	return events, nil
}

func (r *MongoRepository) GetAll() ([]*domain.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("no events found")
	}
	defer cursor.Close(ctx)

	var events []*domain.Event
	for cursor.Next(ctx) {
		var event domain.Event
		if err := cursor.Decode(&event); err != nil {
			return nil, errors.New("ran into an error while fetching your events")
		}
		events = append(events, &event)
	}
	return events, nil
}
func (r *MongoRepository) Update(event *domain.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": event.ID}
	update := bson.M{"$set": event}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("could not update event")
	}
	if result.MatchedCount == 0 {
		return errors.New("event not found")
	}
	return nil
}
func (r *MongoRepository) DeleteById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": id}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New("could not delete event")
	}
	if result.DeletedCount == 0 {
		return errors.New("event not found")
	}
	return nil
}
