package users

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents the user of the system
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
	JoinedAt time.Time `json:"joined_at,omitzero"`
}

// Prepare formats and validate the required fields for User
func (u *User) Prepare() error {
	u.format()

	if err := u.validate(); err != nil {
		return err
	}

	return nil
}

// validate checks if the required fields is filled
func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("Name is required")
	}

	if u.Email == "" {
		return errors.New("Email is required")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("E-mail valid format")
	}

	if u.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}

// format remove the blanks spaces and format username
func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)

	if u.Username == "" {
		u.Username = strings.Split(u.Email, "@")[0]
		u.Username = strings.TrimSpace(u.Username)
	}

	u.Email = strings.TrimSpace(u.Email)
}
