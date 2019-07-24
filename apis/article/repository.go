package article

import (
	"context"
	"github.com/bbcyyb/pcrs/models"
)

type IRepository interface {
	GetAll(context.Context) ([]*models.Article, error)
	GetByID(context.Context, int64) (*models.Article, error)
	GetByTitle(context.Context, string) (*models.Article, error)
	Update(context.Context, *models.Article) error
	Create(context.Context, *models.Article) error
	Delete(context.Context, int64) error
}

type ArticleRepository struct {
	Conn *gorm.DB
}

// NewArticleRepository will create an object that represent the article.Repository interface
func NewArticleRepository(conn *gorm.DB) ArticleRepository {
	return &articleRepository{conn}
}

func (a *ArticleRepository) GetAll(ctx context.Context) (res []*models.Article, err error) {
	//err = a.Conn.Preload("Tag").Find(&res).Error
	err = a.Conn.Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return res, nil
}

func (a *ArticleRepository) GetByID(ctx context.Context, id int64) (res *models.Article, err error) {
	var article models.Article
	err = a.Conn.Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	//err = a.Conn.Model(&article).Related(&article.Tag).Error
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return nil, err
	//}

	return &article, nil
}

func (a *ArticleRepository) GetByTitle(ctx context.Context, title string) (res *models.Article, err error) {
	var article models.Article
	err = a.Conn.Where("title = ?", title).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	//err = a.Conn.Model(&article).Related(&article.Tag).Error
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return nil, err
	//}

	return &article, nil
}

func (a *ArticleRepository) Create(ctx context.Context, article *models.Article) error {
	if err := a.Conn.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func (a *ArticleRepository) Delete(ctx context.Context, id int64) error {
	if err := a.Conn.Exec("DELETE FROM [blog_article] WHERE id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (a *ArticleRepository) Update(ctx context.Context, article *models.Article) error {
	id := &article.ID
	if err := a.Conn.Model(&models.Article{}).Where("id = ?", id, 0).Updates(&article).Error; err != nil {
		return err
	}
	return nil
}
