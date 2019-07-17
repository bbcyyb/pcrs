package article

import (
	"context"
	"github.com/bbcyyb/pcrs/models"
)

// Repository represent the article's repository contract
type Repository interface {
	GetAll(ctx context.Context) ([]*models.Article, error)
	GetByID(ctx context.Context, id int64) (*models.Article, error)
	GetByTitle(ctx context.Context, title string) (*models.Article, error)
	Update(ctx context.Context, ar *models.Article) error
	Create(ctx context.Context, a *models.Article) error
	Delete(ctx context.Context, id int64) error
}
