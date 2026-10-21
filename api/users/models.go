package users

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
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
func (u *User) Prepare(stage string) error {
	if err := u.validate(stage); err != nil {
		return err
	}
	if err := u.format(stage); err != nil {
		return err
	}

	return nil
}

// validate checks if the required fields is filled
func (u *User) validate(stage string) error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("Name is required")
	}

	if strings.TrimSpace(u.Email) == "" {
		return errors.New("Email is required")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("E-mail valid format")
	}

	if strings.TrimSpace(u.Password) == "" && stage == "register" {
		return errors.New("Password is required")
	}

	return nil
}

// format remove the blanks spaces and format username
func (u *User) format(stage string) error {
	u.Name = strings.TrimSpace(u.Name)

	if u.Username == "" {
		u.Username = strings.Split(u.Email, "@")[0]
		u.Username = strings.TrimSpace(u.Username)
	}

	u.Email = strings.TrimSpace(u.Email)

	if stage == "register" {
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(u.Password), bcrypt.DefaultCost,
		)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}
