package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertData(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@gmail.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	result, err := CommentRepository.FindById(ctx, 32)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	result, err := CommentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
