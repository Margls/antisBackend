package repositories

import(
	"context"
	"time"
	"database/sql"
    "antis/backend/models"
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

	query := `SELECT id, name, description, created_at`
	err := r.db.GetContext(ctx, &user, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func (r *ItemRepository) GetAllItems(ctx context.Context)([]models.Item, error) {
	var items []models.Item
	query := `SELECT id, name, description, created_at FROM items ORDER BY id`
	err := r.db.GetContext(ctx, &items, query )

	if err != nil {
		 return nil,nil
	}

	if items == nil {
		return []models.Item{}, nil
	}

	 return items, nil
}