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
	return &repository{
		client: client,
		config: config,
	}
}

func (r *repository) getCollection() *mongo.Collection {
	return r.client.Database(r.config.MongoDatabase).Collection(r.config.MongoCollection)
}

func (r *repository) Create(note *Note, ctx context.Context) (*Note, error) {

	collection := r.getCollection()

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, ErrCreatingNoteId
	}
	note.ID = id
	note.CreatedAt = time.Now().UTC()

	_, err = collection.InsertOne(ctx, note)

	if err != nil {
		return nil, ErrCreatingNote
	}

	return note, nil
}

func (r *repository) GetById(id uuid.UUID, ctx context.Context) (*Note, error) {
	var note Note

	collection := r.getCollection()
	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(&note)
	if err != nil {
		return nil, ErrFoundingNote
	}

	return &note, nil
}

func (r *repository) GetAll(ctx context.Context) ([]Note, error) {
	var notes []Note

	findOptions := options.Find()
	findOptions.SetLimit(100)

	collection := r.getCollection()

	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, ErrFoundingNote
	}

	for cur.Next(ctx) {
		var note Note
		if err := cur.Decode(&note); err != nil {
			return nil, ErrDecodingNote
		}

		notes = append(notes, note)
	}
	cur.Close(ctx)

	if notes == nil {
		return nil, ErrFoundingNote
	}

	return notes, nil
}

func (r *repository) Update(note *Note, ctx context.Context) (*Note, error) {
	collection := r.getCollection()
	filter := bson.M{"_id": note.ID}

	note.UpdatedAt = time.Now().UTC()

	update := bson.M{
		"$set": note,
	}

	result, err := collection.UpdateOne(ctx, filter, update)

	if result.MatchedCount == 0 {
		return nil, ErrFoundingNote
	}

	if err != nil {
		return nil, ErrUpdatingNote
	}

	return note, nil
}
