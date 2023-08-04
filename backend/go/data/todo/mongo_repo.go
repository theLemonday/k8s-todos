package todo

import (
	"context"
	"time"

	"github.com/theLemonday/k8s-todos/api/presenter"
	"github.com/theLemonday/k8s-todos/data/entities"
	usererror "github.com/theLemonday/k8s-todos/user_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDbRepository struct {
	coll *mongo.Collection
}

func NewMongoDbRepository(coll *mongo.Collection) Repository {
	return &mongoDbRepository{coll: coll}
}

func (r *mongoDbRepository) CreateTodo(todo *entities.Todo) (*entities.Todo, error) {
	todo.ID = primitive.NewObjectID()
	todo.CreatedAt = time.Now()
	todo.UpdateAt = time.Now()

	_, err := r.coll.InsertOne(context.Background(), todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *mongoDbRepository) GetTodoById(id string) (*presenter.Todo, error) {
	todoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var todo presenter.Todo
	if err := r.coll.FindOne(context.TODO(), bson.M{"_id": todoId}).Decode(&todo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, usererror.ErrNoDocumentFound
		}
		return nil, err
	}

	return &todo, nil
}

func (r *mongoDbRepository) GetAllTodos() (*[]presenter.Todo, error) {
	cur, err := r.coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var todos []presenter.Todo
	if err := cur.All(context.TODO(), &todos); err != nil {
		return nil, err
	}

	return &todos, nil
}

func (r *mongoDbRepository) UpdateTodo(todo *entities.Todo) (*entities.Todo, error) {
	todo.UpdateAt = time.Now()

	if _, err := r.coll.UpdateOne(context.TODO(), todo.ID, bson.M{"$set": todo}); err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *mongoDbRepository) DeleteTodo(id string) error {
	todoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if _, err := r.coll.DeleteOne(context.TODO(), bson.M{"_id": todoId}); err != nil {
		return err
	}

	return nil
}
