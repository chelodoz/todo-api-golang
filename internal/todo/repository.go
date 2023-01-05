package todo

import (
	"context"
	"log"
	"time"
	"todo-api-golang/internal/apperror"
	"todo-api-golang/internal/config"
	"todo-api-golang/internal/entity"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepository struct {
	client *mongo.Client
	config *config.Config
}

func NewTodoRepository(client *mongo.Client, config *config.Config) TodoRepository {
	return &todoRepository{client, config}
}

func (todoRepository *todoRepository) getCollection() *mongo.Collection {

	collection := todoRepository.client.Database(todoRepository.config.DBName).Collection(todoRepository.config.DBCollection)

	return collection
}

func (todoRepository *todoRepository) CreateTodo(todo *entity.Todo, ctx context.Context) (*entity.Todo, error) {

	collection := todoRepository.getCollection()
	todo.ID = uuid.New()
	todo.CreatedAt = time.Now().UTC()
	_, err := collection.InsertOne(ctx, todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (todoRepository *todoRepository) GetTodoById(id uuid.UUID, ctx context.Context) (*entity.Todo, error) {
	var todo entity.Todo

	collection := todoRepository.getCollection()

	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(&todo)
	if err != nil {
		return nil, apperror.ErrTodoNotFound
	}

	return &todo, nil
}

func (todoRepository *todoRepository) GetTodos(ctx context.Context) ([]entity.Todo, error) {

	findOptions := options.Find()
	findOptions.SetLimit(100)

	todos := []entity.Todo{}

	collection := todoRepository.getCollection()

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
		return nil, apperror.ErrTodoNotFound
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var todo entity.Todo
		if err := cur.Decode(&todo); err != nil {
			log.Fatal(err)
			return nil, err
		}

		todos = append(todos, todo)
	}

	cur.Close(ctx)

	return todos, nil
}

func (todoRepository *todoRepository) UpdateTodo(todo *entity.Todo, ctx context.Context) (*entity.Todo, error) {

	collection := todoRepository.getCollection()

	filter := bson.M{"_id": todo.ID}

	todo.UpdatedAt = time.Now().UTC()

	update := bson.M{
		"$set": todo,
	}

	result, err := collection.UpdateOne(ctx, filter, update)

	if result.MatchedCount == 0 {
		return nil, apperror.ErrTodoNotFound
	}

	if err != nil {
		return nil, err
	}

	return todo, nil
}
