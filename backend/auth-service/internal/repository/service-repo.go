package repository

import "github.com/Lemper29/Jarcy/auth-service/internal/database"

type Repository struct {
	db *database.PostgresDatabase
}

func NewRepo(db *database.PostgresDatabase) *Repository {
	return &Repository{
		db: db,
	}
}
