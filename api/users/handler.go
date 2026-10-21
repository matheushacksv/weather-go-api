package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserHandler represents the repository with Pool to database
type UserHandler struct {
	Pool *pgxpool.Pool
}

// Create handles the HTTP and calls repository to create User
func (u UserHandler) Create(c *gin.Context) {
	var data User

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repository := UserRepository{u.Pool}

	if err := data.Prepare(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.CreateUser(
		c.Request.Context(),
		data,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// Get handles the HTTP and call repository to retrievies a single User
func (u UserHandler) Get(c *gin.Context) {
	param := c.Param("ID")
	userID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not get user"})
		return
	}

	repository := UserRepository{u.Pool}

	user, err := repository.GetUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// List handles the HTTP and call repository to retrievies all Users
func (u UserHandler) List(c *gin.Context) {
	repository := UserRepository{u.Pool}

	users, err := repository.ListUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// Update handles the http and call repository to update User
func (u UserHandler) Update(c *gin.Context) {
	var data User

	param := c.Param("ID")
	err := c.ShouldBindJSON(&data)

	userID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := data.Prepare(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	repository := UserRepository{u.Pool}

	user, err := repository.UpdateUser(
		c.Request.Context(),
		userID,
		data.Name,
		data.Username,
		data.Email,
	)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "could not reach the service"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Delete handles the http and call repository to delete User
func (u UserHandler) Delete(c *gin.Context) {
	param := c.Param("ID")
	userID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repository := UserRepository{u.Pool}

	if err := repository.DeleteUser(c.Request.Context(), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete user"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
