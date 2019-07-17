package repository

import (
	"context"
	"github.com/bbcyyb/pcrs/article"
	"github.com/bbcyyb/pcrs/models"
	"github.com/jinzhu/gorm"
)

type articleRepository struct {
	Conn *gorm.DB
}

// NewArticleRepository will create an object that represent the article.Repository interface
func NewArticleRepository(Conn *gorm.DB) article.Repository {
	return &articleRepository{Conn}
}

func (a *articleRepository) GetAll(ctx context.Context) (res []*models.Article, err error) {
	//err = a.Conn.Preload("Tag").Find(&res).Error
	err = a.Conn.Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return res, nil
}

func (a *articleRepository) GetByID(ctx context.Context, id int64) (res *models.Article, err error) {
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

func (a *articleRepository) GetByTitle(ctx context.Context, title string) (res *models.Article, err error) {
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

func (a *articleRepository) Create(ctx context.Context, article *models.Article) error {
	if err := a.Conn.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func (a *articleRepository) Delete(ctx context.Context, id int64) error {
	if err := a.Conn.Exec("DELETE FROM [blog_article] WHERE id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (a *articleRepository) Update(ctx context.Context, article *models.Article) error {
	id := &article.ID
	if err := a.Conn.Model(&models.Article{}).Where("id = ?", id, 0).Updates(&article).Error;
		err != nil {
		return err
	}
	return nil
}
