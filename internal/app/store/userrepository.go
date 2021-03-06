package store

import (
	"database/sql"
	"http-rest-api/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email,encrypted_password) VALUES ($1, $2) RETURNING id", u.Email, u.EncryptorPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptorPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return u, nil
}
