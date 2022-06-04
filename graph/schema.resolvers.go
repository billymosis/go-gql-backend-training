package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"example.com/billy/graph/generated"
	"example.com/billy/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.AddUser) (*model.User, error) {
	user := &model.User{ID: input.Username, Name: input.Name}
	r.DB.MustExec("INSERT INTO users (user_id, email) values ($1, $2)", user.ID, user.Name)
	return user, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.AddTodo) (*model.Todo, error) {
	user := &model.User{ID: input.User.Name, Name: input.User.Username}
	todo := &model.Todo{ID: input.ID, Text: input.Text, Done: false, UserID: input.User.Username, User: user}
	r.DB.MustExec("INSERT INTO todos (todos_id, text, done, user_id) values ($1, $2, $3, $4)", todo.ID, todo.Text, todo.Done, todo.UserID)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todo model.Todo
	var user model.User
	var todos []*model.Todo

	rows, err := r.DB.Queryx("SELECT todos_id, text, done, todos.user_id as tid, users.user_id as uid, users.email FROM todos LEFT JOIN users ON todos.user_id = users.user_id")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Text, &todo.Done, &todo.UserID, &user.ID, &user.Name)
		todo.User = &user
		fmt.Println(todo)
		fmt.Println(user)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, &todo)
	}

	fmt.Println(todos)

	return todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return obj.User, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
