package reposiotry

import (
	belajargolangdatabase "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// anggap ini bisnis logic nya
func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@gmil.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())
	ctx := context.Background()

	comment, err := commentRepository.FindByID(ctx, 90)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindALL(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())
	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}

}
