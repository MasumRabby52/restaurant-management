package repository

import (
	"context"
	"fmt"
	"golang-restaurant-management/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(pool *pgxpool.Pool, user *models.User) (*models.User, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
	INSERT INTO users(first_name,last_name,password,email,phone)
	VALUES($1,$2,$3,$4,$5)
	RETURNING first_name,last_name,email,phone,created_at,updated_at,user_id
	`
	err := pool.QueryRow(ctx, query, user.First_name, user.Last_name, user.Password, user.Email, user.Phone).Scan(
		&user.First_name,
		&user.Last_name,
		&user.Email,
		&user.Phone,
		&user.Created_at,
		&user.Updated_at,
		&user.User_id,
	)
	if err != nil {
		fmt.Println("Internal server error")
		return nil, err
	}
	return user, nil
}
func GetUserByEmail(pool *pgxpool.Pool, email string) (*models.User, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
	SELECT *
	FROM users
	WHERE email=$1
	`
	var user models.User
	err := pool.QueryRow(ctx, query, email).Scan(
		&user.First_name,
		&user.Last_name,
		&user.Password,
		&user.Email,
		&user.Phone,
		&user.Created_at,
		&user.Updated_at,
		&user.User_id,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
