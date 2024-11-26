package repositories

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"auth-service/models"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestNewUserRepository(t *testing.T) {
	mockedDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(mockedDB *sql.DB) {
		closeErr := mockedDB.Close()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when closing a stub database connection", closeErr)
		}

	}(mockedDB)

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want UserRepository
	}{
		{
			name: "Successfully create new UserRepository",
			args: args{
				db: mockedDB,
			},
			want: &userRepository{db: mockedDB},
		},
		{
			name: "Fail to create new UserRepository",
			args: args{
				db: nil,
			},
			want: &userRepository{},
		},
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
	mockedDB, sqMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockedDB.Close()

	type fields struct {
		db *sql.DB
	}

	type args struct {
		user *models.User
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		mockFunc func()
	}{
		// {
		// 	name:    "Successfully create new User",
		// 	fields:  fields{db: mockedDB},
		// 	args:    args{user: &models.User{}},
		// 	wantErr: false,
		// 	mockFunc: func() {
		// 		sqlQueryRegexStr := `^INSERT INTO users \(email, password, first_name, last_name\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING id$`
		// 		sqMock.ExpectQuery(sqlQueryRegexStr).
		// 			WithArgs("test@example.com", "password", "John", "Doe").
		// 			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		// 	},
		// },
		{
			name:   "Successfully create new User",
			fields: fields{db: mockedDB},
			args: args{user: &models.User{
				Email:     "test@example.com",
				Password:  "password",
				FirstName: "John",
				LastName:  "Doe",
			}},
			wantErr: false,
			mockFunc: func() {
				sqlQueryRegexStr := `^INSERT INTO users \(email, password, first_name, last_name\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING id$`
				sqMock.ExpectExec(sqlQueryRegexStr).
					WithArgs("test@example.com", "password", "John", "Doe").
					WillReturnResult(sqlmock.NewResult(1, 1)) // Assuming 1 is the ID of the newly created user
			},
		},
		{
			name:    "Fail to create user",
			fields:  fields{db: mockedDB},
			args:    args{user: &models.User{Email: "test@example.com", Password: "password", FirstName: "John", LastName: "Doe"}},
			wantErr: true,
			mockFunc: func() {
				sqMock.ExpectQuery(`^INSERT INTO "users" \(`).
					WithArgs("test@example.com", "password", "John", "Doe").
					WillReturnError(sql.ErrConnDone)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc() // Call this to set up the expectations
			r := userRepository{
				db: tt.fields.db,
			}
			if err := r.Create(context.Background(), tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("User repo create error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_GetByUserEmail(t *testing.T) {
	type fields struct {
		db  *sql.DB
		log *logrus.Logger
	}
	type args struct {
		email string
		ctx   context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
		mockFn  func(sqlMockObj sqlmock.Sqlmock)
	}{
		{
			name: "Successfully User found",
			args: args{
				email: "test@example.com",
				ctx:   context.Background(),
			},
			want: models.User{
				ID: 1,
			},
			wantErr: false,
			mockFn: func(sqlMockObj sqlmock.Sqlmock) {
				userRows := sqlmock.NewRows(
					[]string{
						"id",
						"email",
						"password",
						"first_name",
						"last_name",
					},
				).AddRow(
					1, "", "", "", "",
				)

				sqlStr := `^SELECT id, email, password, first_name, last_name FROM users WHERE email = \$1$`
				sqlMockObj.ExpectQuery(sqlStr).
					WithArgs("test@example.com").
					WillReturnRows(userRows)
			},
		},
		{
			name: "User  not found",
			args: args{
				email: "notfound@example.com",
				ctx:   context.Background(),
			},
			want:    models.User{},
			wantErr: true,
			mockFn: func(sqlMockObj sqlmock.Sqlmock) {
				// Use the exact query string instead of regex
				// sqlMockObj.ExpectQuery("SELECT id, email, first_name, last_name FROM users WHERE email = \\$1").
				sqlStr := `^SELECT id, email, password, first_name, last_name FROM users WHERE email = \$1$`

				sqlMockObj.ExpectQuery(sqlStr).
					WithArgs("notfound@example.com").
					WillReturnError(sql.ErrNoRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock db
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			r := &userRepository{db: db}

			tt.mockFn(mock)
			got, err := r.GetByUserEmail(tt.args.ctx, tt.args.email)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tt.want.ID, got.ID)
			}
			// check got and want

			// assert.True(t, !reflect.DeepEqual(got, tt.want))
			assert.Equal(t, tt.want, got)

			if mErr := mock.ExpectationsWereMet(); mErr != nil {
				t.Errorf("there were unfulfilled expectations: %v", mErr)
			}
		})
	}
}
