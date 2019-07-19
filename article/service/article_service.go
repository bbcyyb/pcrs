package service

import (
	"context"
	"errors"
	"github.com/bbcyyb/pcrs/article"
	"github.com/bbcyyb/pcrs/models"
)

type ArticleService struct {
	articleRepo article.Repository
}

// NewArticleService will create new an articleService object representation of article.Service interface
func NewArticleService(a *article.Repository) *article.ArticleService {
	return &ArticleService{
		articleRepo: a,
	}
}

func (a *articleService) GetAll(c context.Context) ([]*models.Article, error) {

	listArticle, err := a.articleRepo.GetAll(c)
	if err != nil {
		return nil, err
	}
	return listArticle, nil
}

func (a *articleService) GetByID(c context.Context, id int64) (*models.Article, error) {
	res, err := a.articleRepo.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *articleService) Update(c context.Context, article *models.Article) error {
	return a.articleRepo.Update(c, article)
}

func (a *articleService) GetByTitle(c context.Context, title string) (*models.Article, error) {
	res, err := a.articleRepo.GetByTitle(c, title)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *articleService) Create(c context.Context, article *models.Article) error {
	existedArticle, _ := a.GetByTitle(c, article.Title)
	if existedArticle != nil {
		return errors.New("item already exist")
	}

	err := a.articleRepo.Create(c, article)
	if err != nil {
		return err
	}
	return nil
}

func (a *articleService) Delete(c context.Context, id int64) error {
	existedArticle, err := a.articleRepo.GetByID(c, id)
	if err != nil {
		return err
	}
	if existedArticle == nil {
		return errors.New("item not found")
	}
	return a.articleRepo.Delete(c, id)
}
