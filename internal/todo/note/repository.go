package note

import (
	"context"
	"time"
	"todo-api-golang/internal/config"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	client *mongo.Client
	config *config.Config
}

func NewRepository(client *mongo.Client, config *config.Config) Repository {
	return &repository{client, config}
}

func (repository *repository) getCollection() *mongo.Collection {

	collection := repository.client.Database(repository.config.MongoDatabase).Collection(repository.config.MongoCollection)

	return collection
}

func (repository *repository) Create(note *Note, ctx context.Context) (*Note, error) {

	collection := repository.getCollection()

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	note.ID = id

	note.CreatedAt = time.Now().UTC()
	_, err = collection.InsertOne(ctx, note)

	if err != nil {
		return nil, err
	}

	return note, nil
}

func (repository *repository) GetById(id uuid.UUID, ctx context.Context) (*Note, error) {
	var note Note

	collection := repository.getCollection()

	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(&note)
	if err != nil {
		return nil, ErrNoteNotFound
	}

	return &note, nil
}

func (repository *repository) GetAll(ctx context.Context) ([]Note, error) {

	findOptions := options.Find()
	findOptions.SetLimit(100)

	var notes []Note

	collection := repository.getCollection()

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, ErrNoteNotFound
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var note Note
		if err := cur.Decode(&note); err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	cur.Close(ctx)

	if notes == nil {
		return nil, ErrNoteNotFound
	}

	return notes, nil
}

func (repository *repository) Update(note *Note, ctx context.Context) (*Note, error) {

	collection := repository.getCollection()

	filter := bson.M{"_id": note.ID}

	note.UpdatedAt = time.Now().UTC()

	update := bson.M{
		"$set": note,
	}

	result, err := collection.UpdateOne(ctx, filter, update)

	if result.MatchedCount == 0 {
		return nil, ErrNoteNotFound
	}

	if err != nil {
		return nil, err
	}

	return note, nil
}
