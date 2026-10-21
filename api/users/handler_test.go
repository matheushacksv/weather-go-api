package users

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestCreateInvalidJSON(t *testing.T) {
	router := gin.New()
	router.POST("/users", UserHandler{}.Create)

	body := strings.NewReader(`{`)
	req := httptest.NewRequest(http.MethodPost, "/users", body)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", rec.Code)
	}
}

func TestCreateSuccess(t *testing.T) {
	pool := setupDB(t)

	router := gin.New()
	router.POST("/users", UserHandler{pool}.Create)

	body := strings.NewReader(
		`{
			"name": "Ana",
			"email": "ana@mail.com",
			"password": "123"
		}`,
	)
	req := httptest.NewRequest(http.MethodPost, "/users", body)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("status = %d, want 201", rec.Code)
	}

	var got struct {
		User User `json:"user"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.User.Name != "Ana" {
		t.Errorf("name = %q, want %q", got.User.Name, "Ana")
	}
	if got.User.Email != "ana@mail.com" {
		t.Errorf("email = %q, want %q", got.User.Email, "ana@mail.com")
	}
	if got.User.Username != "ana" {
		t.Errorf("username = %q, want %q", got.User.Username, "ana")
	}
	// if got.User.Password != "123" {
	// 	t.Errorf("password = %q, want %q", got.User.Password, "123")
	// }
}

func TestListEmpty(t *testing.T) {
	pool := setupDB(t)

	router := gin.New()
	router.GET("/users", UserHandler{pool}.List)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", rec.Code)
	}
	if rec.Body.String() != `{"users":[]}` {
		t.Errorf("body = %s", rec.Body.String())
	}
}

func TestGetUserSuccess(t *testing.T) {
	pool := setupDB(t)

	router := gin.New()
	router.GET("/users/:ID", UserHandler{pool}.Get)

	createTestUser(t, pool)

	req := httptest.NewRequest(
		http.MethodGet, "/users/1", nil,
	)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", rec.Code)
	}

}

func TestUpdateUserSuccess(t *testing.T) {
	pool := setupDB(t)

	router := gin.New()
	router.PUT("/users/:ID", UserHandler{pool}.Update)

	createTestUser(t, pool)

	body := strings.NewReader(
		`{
			"name": "Ana2",
			"username": "ana2",
			"email": "ana2@mail.com",
			"password": "12345"
		}`,
	)

	req := httptest.NewRequest(
		http.MethodPut, "/users/1", body,
	)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", rec.Code)
	}

	var got User
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.Name != "Ana2" {
		t.Errorf("name = %q, want %q", got.Name, "Ana2")
	}
	if got.Username != "ana2" {
		t.Errorf("username = %q, want %q", got.Username, "ana2")
	}
	if got.Email != "ana2@mail.com" {
		t.Errorf("email = %q, want %q", got.Email, "ana2@mail.com")
	}
	if got.Password != "" {
		t.Errorf("password = %q, want %q", got.Password, "12345")
	}

}

func TestDeleteUserSucess(t *testing.T) {
	pool := setupDB(t)

	router := gin.New()
	router.DELETE("/users/:ID", UserHandler{pool}.Delete)

	createTestUser(t, pool)

	req := httptest.NewRequest(
		http.MethodDelete, "/users/1", nil,
	)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("status = %d, want 204", rec.Code)
	}

}

// Helpers

func setupDB(t *testing.T) *pgxpool.Pool {
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)

	_, err = pool.Exec(
		context.Background(),
		"TRUNCATE users RESTART IDENTITY",
	)
	if err != nil {
		t.Fatal(err)
	}
	return pool
}

func createTestUser(t *testing.T, pool *pgxpool.Pool) {
	_, err := pool.Exec(
		context.Background(),
		"INSERT INTO users (name, username, email, password) VALUES ('Ana', 'ana', 'ana@mail.com', '123')",
	)
	if err != nil {
		t.Error(err)
	}
}
