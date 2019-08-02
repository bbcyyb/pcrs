package article

import (
	"context"
	"errors"
	"github.com/bbcyyb/pcrs/models"
)

type IService interface {
	GetAll(context.Context) ([]*models.Article, error)
	GetByID(context.Context, int64) (*models.Article, error)
	Update(context.Context, *models.Article) error
	GetByTitle(context.Context, string) (*models.Article, error)
	Create(context.Context, *models.Article) error
	Delete(context.Context, int64) error
}

type ArticleService struct {
	Repo IRepository
}

// NewArticleService will create new an ArticleService object representation of article.Service interface
func NewArticleService(repo IRepository) IService {
	return &ArticleService{
		Repo: repo,
	}
}

func (a *ArticleService) GetAll(c context.Context) ([]*models.Article, error) {
	listArticle, err := a.Repo.GetAll(c)

	if err != nil {
		return nil, err
	}
	return listArticle, nil
}

func (a *ArticleService) GetByID(c context.Context, id int64) (*models.Article, error) {
	res, err := a.Repo.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *ArticleService) Update(c context.Context, article *models.Article) error {
	return a.Repo.Update(c, article)
}

func (a *ArticleService) GetByTitle(c context.Context, title string) (*models.Article, error) {
	res, err := a.Repo.GetByTitle(c, title)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *ArticleService) Create(c context.Context, article *models.Article) error {
	existedArticle, _ := a.GetByTitle(c, article.Title)
	if existedArticle != nil {
		return errors.New("item already exist")
	}

	err := a.Repo.Create(c, article)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleService) Delete(c context.Context, id int64) error {
	existedArticle, err := a.Repo.GetByID(c, id)
	if err != nil {
		return err
	}
	if existedArticle == nil {
		return errors.New("item not found")
	}
	return a.Repo.Delete(c, id)
}
