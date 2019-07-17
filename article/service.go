package article

import (
	"context"
	"github.com/bbcyyb/pcrs/models"
)

type Service interface {
	GetAll(ctx context.Context) ([]*models.Article, error)
	GetByID(ctx context.Context, id int64) (*models.Article, error)
	Update(ctx context.Context, ar *models.Article) error
	GetByTitle(ctx context.Context, title string) (*models.Article, error)
	Create(context.Context, *models.Article) error
	Delete(ctx context.Context, id int64) error
}
