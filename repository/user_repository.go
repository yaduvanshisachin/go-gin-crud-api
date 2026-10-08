package repository

import (
	"context"

	"gin-crud/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create query method.

func (r *UserRepository) Create(
	ctx context.Context,
	user *model.User,
) error {

	query := `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id
	`

	return r.db.QueryRow(
		ctx,
		query,
		user.Name,
		user.Email,
	).Scan(&user.ID)
}

func (r *UserRepository) GetAll(
	ctx context.Context,
) ([]model.User, error) {

	query := `
		SELECT id, name, email
		FROM users
		ORDER BY id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, rows.Err()
}

func (r *UserRepository) GetByID(
	ctx context.Context,
	id int,
) (*model.User, error) {

	query := `
		SELECT id, name, email
		FROM users
		WHERE id = $1
	`

	var user model.User

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update user
func (r *UserRepository) Update(
	ctx context.Context,
	user *model.User,
) error {

	query := `
		UPDATE users
		SET name = $1, email = $2
		WHERE id = $3
	`

	_, err := r.db.Exec(
		ctx,
		query,
		user.Name,
		user.Email,
		user.ID,
	)

	return err
}

// delete user
func (r *UserRepository) Delete(
	ctx context.Context,
	id int,
) error {

	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, id)

	return err
}
