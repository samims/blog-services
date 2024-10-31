package repositories

import (
	"reflect"
	"testing"

	"auth-service/models"
)

func TestNewUserRepository(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want UserRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Create(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				DB: tt.fields.DB,
			}
			if err := r.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_GetByUserEmail(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				DB: tt.fields.DB,
			}
			got, err := r.GetByUserEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUserEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUserEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
