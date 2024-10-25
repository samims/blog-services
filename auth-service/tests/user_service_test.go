package tests

import (
	"auth-service/models"
	"auth-service/repositories"
	"auth-service/services"
	"database/sql"
	"testing"
)

func TestUserService_Register(t *testing.T) {
	db, err := sql.Open(
		"postgres",
		"postgresql://postgres@localhost:5432/postgres?sslmode=disable",
	)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)

	user := models.User{
		ID:       123,
		Username: "abc",
		Password: "abc",
	}
	err = service.Register(user)
	if err != nil {
		t.Fatal(err)
	}

}
