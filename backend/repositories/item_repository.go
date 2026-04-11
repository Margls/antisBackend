package repositories

import (
	"antis/backend/models"
	"context"
	"database/sql"
	"fmt"
	"time"
	"github.com/jmoiron/sqlx"
)


type ItemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(ctx context.Context, item *models.Item) error {
	query := `
		INSERT INTO items (name, description, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
		`
		now := time.Now()
		item.CreatedAt = now

		return r.db.QueryRowContext(ctx, query, item.Name, item.Description, now).Scan(&item.ID)

}

func (r *ItemRepository) GetByID(ctx context.Context, id int64) (*models.Item, error) {
	
	var user models.Item

	query := `SELECT id, name, description, created_at FROM items WHERE id = $1`
	err := r.db.GetContext(ctx, &user, query, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func (r *ItemRepository) GetAllItems(ctx context.Context)([]models.Item, error) {
	var items []models.Item
	query := `SELECT id, name, description, created_at FROM items ORDER BY id`
	err := r.db.SelectContext(ctx, &items, query )

	if err != nil {
		 return nil, fmt.Errorf("failed to get items: %w", err)
	}

	if items == nil {
		return []models.Item{}, nil
	}

	 return items, nil
}