package services

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"auth-service/config"
	configmocks "auth-service/config/mocks"
	"auth-service/models"
	"auth-service/repositories"
	"auth-service/repositories/mocks"

	"github.com/stretchr/testify/mock"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		repo repositories.UserRepository
		conf config.Configuration
	}
	tests := []struct {
		name string
		args args
		want *userService
	}{
		{
			name: "create new UserService",
			args: args{
				repo: mocks.NewUserRepository(t),
				conf: configmocks.NewConfiguration(t),
			},
			want: &userService{
				repo: mocks.NewUserRepository(t),
				conf: configmocks.NewConfiguration(t),
			},
		},
		{
			name: "create new UserService with nil repository",
			args: args{
				repo: nil,
				conf: configmocks.NewConfiguration(t),
			},
			want: &userService{conf: configmocks.NewConfiguration(t)},
		},
		{
			name: "create new UserService with nil configuration",
			args: args{
				repo: mocks.NewUserRepository(t),
				conf: nil,
			},
			want: &userService{conf: nil, repo: mocks.NewUserRepository(t)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUserService(tt.args.repo, tt.args.conf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Login(t *testing.T) {
	type fields struct {
		repo *mocks.UserRepository
		conf *configmocks.Configuration
	}
	type args struct {
		ctx      context.Context
		loginReq models.LoginRequest
	}
	sampleLoginReq := models.LoginRequest{
		Email:    "asif@example.com",
		Password: "123145",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(*fields)
		want    models.LoginResponse
		wantErr bool
	}{
		{
			name:   "UserService login fails when user not found",
			fields: fields{},
			args: args{
				ctx:      context.Background(),
				loginReq: sampleLoginReq,
			},
			prepare: func(f *fields) {
				f.repo.On(
					"GetByUserEmail",
					mock.Anything,
					mock.Anything).
					Return(
						models.User{},
						errors.New("not found"),
					)
			},

			want:    models.LoginResponse{},
			wantErr: true,
		},
		{
			name:   "UserService login fails due to CompareHashAndPassword error",
			fields: fields{},
			args: args{
				ctx:      context.Background(),
				loginReq: sampleLoginReq,
			},
			prepare: func(fields *fields) {
				fields.repo.On("GetByUserEmail", mock.Anything, mock.Anything).Return(models.User{
					Email:    sampleLoginReq.Email,
					Password: "wrongPass",
				}, nil)
			},
			want:    models.LoginResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields = fields{
				repo: &mocks.UserRepository{},
				conf: &configmocks.Configuration{},
			}
			if tt.prepare != nil {
				tt.prepare(&tt.fields)
			}
			u := userService{
				repo: tt.fields.repo,
				conf: tt.fields.conf,
			}
			_, got, err := u.Login(tt.args.ctx, tt.args.loginReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Register(t *testing.T) {
	type fields struct {
		repo repositories.UserRepository
		conf config.Configuration
	}
	type args struct {
		ctx  context.Context
		user *models.User
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userService{
				repo: tt.fields.repo,
				conf: tt.fields.conf,
			}
			if _, err := u.Register(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userService_Verify(t *testing.T) {
	type fields struct {
		repo repositories.UserRepository
		conf config.Configuration
	}
	type args struct {
		in0 context.Context
		req models.VerifyRequest
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userService{
				repo: tt.fields.repo,
				conf: tt.fields.conf,
			}
			if _, err := u.VerifyToken(tt.args.in0, tt.args.req.Token); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
