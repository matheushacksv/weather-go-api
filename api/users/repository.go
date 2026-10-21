package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

// CreateUser creates a User in database
func (u UserRepository) CreateUser(ctx context.Context, user User) (User, error) {
	err := u.db.QueryRow(
		ctx,
		`INSERT INTO users (name, username, email, password)
		VALUES ($1,$2,$3,$4) RETURNING id, name, username, email, joined_at`,
		user.Name, user.Username, user.Email, user.Password,
	).Scan(
		&user.ID, &user.Name, &user.Username, &user.Email, &user.JoinedAt,
	)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// ListUser get all users from database
func (u UserRepository) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := u.db.Query(
		ctx,
		"SELECT id, name, username, email, joined_at FROM users",
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(
			&user.ID, &user.Name, &user.Username, &user.Email, &user.JoinedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetUser get a user by ID from database
func (u UserRepository) GetUser(ctx context.Context, ID uint64) (User, error) {
	var user User
	err := u.db.QueryRow(
		ctx,
		"SELECT id, name, username, email, joined_at FROM users WHERE id = $1",
		ID,
	).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.JoinedAt)

	if err != nil {
		return User{}, err
	}
	return user, nil
}

// UpdateUser updates data of a user from database
func (u UserRepository) UpdateUser(ctx context.Context, ID uint64, name, username, email string) (User, error) {
	var user User
	err := u.db.QueryRow(
		ctx,
		`UPDATE users SET name = $1, username = $2, email = $3
		WHERE id = $4 RETURNING id, name, username, email, joined_at`,
		name, username, email, ID,
	).Scan(
		&user.ID, &user.Name, &user.Username, &user.Email, &user.JoinedAt,
	)

	if err != nil {
		return User{}, err
	}
	return user, nil
}

// DeleteUser delete a user from the database
func (u UserRepository) DeleteUser(ctx context.Context, ID uint64) error {
	_, err := u.db.Exec(
		ctx,
		"DELETE FROM users WHERE id = $1",
		ID,
	)
	if err != nil {
		return err
	}
	return nil
}
