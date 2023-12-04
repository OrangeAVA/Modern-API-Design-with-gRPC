package configs

import (
	"time"

	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/repo"
	"github.com/HiteshRepo/Modern-API-Design-with-gRPC/books-app/internal/pkg/service"
)

const (
	SecretKey     = "secret"
	TokenDuration = 15 * time.Minute

	bookServicePath       = "/prot.BookService/"
	bookReviewServicePath = "/prot.ReviewService/"
	bookInfoServicePath   = "/prot.BookInfoService/"
)

func SeedUsers(userStore repo.UserStore) error {
	err := createUser(userStore, "admin1", "12345", "admin")
	if err != nil {
		return err
	}

	return createUser(userStore, "user1", "54321", "user")

}

func AccessibleRoles() map[string][]string {
	return map[string][]string{
		bookServicePath + "AddBook":    {"admin"},
		bookServicePath + "UpdateBook": {"admin"},
		bookServicePath + "RemoveBook": {"admin"},

		bookServicePath + "GetBook":                    {"admin", "user"},
		bookReviewServicePath + "GetBookReviews":       {"admin", "user"},
		bookReviewServicePath + "SubmitReviews":        {"admin", "user"},
		bookInfoServicePath + "GetBookInfoWithReviews": {"admin", "user"},
	}
}

func createUser(userStore repo.UserStore, username, password, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return err
	}

	return userStore.Save(user)
}
