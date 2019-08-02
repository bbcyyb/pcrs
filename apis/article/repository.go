package article

import (
	"context"
	"github.com/bbcyyb/pcrs/models"
	"github.com/bbcyyb/pcrs/pkg/database"
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
	Conn *database.Conn
}

func NewArticleRepository() IRepository {
	return &ArticleRepository{database.DbConn}
}

func (a *ArticleRepository) GetAll(ctx context.Context) (res []*models.Article, err error) {
	query := `SELECT [id]
      ,[tag_id]
      ,[title]
      ,[desc]
      ,[content]
      ,[created_on]
      ,[created_by]
      ,[modified_on]
      ,[modified_by]
      ,[deleted_on]
      ,[state]
      ,[cover_image_url]
     FROM [blog_article]`
	err = a.Conn.QueryMany(res, query)
	return
}

func (a *ArticleRepository) GetByID(ctx context.Context, id int64) (res *models.Article, err error) {
	query := `SELECT [id]
      ,[tag_id]
      ,[title]
      ,[desc]
      ,[content]
      ,[created_on]
      ,[created_by]
      ,[modified_on]
      ,[modified_by]
      ,[deleted_on]
      ,[state]
      ,[cover_image_url]
     FROM [blog_article] where [id] = ?`
	err = a.Conn.QuerySingle(res, query, id)
	return
}

func (a *ArticleRepository) GetByTitle(ctx context.Context, title string) (res *models.Article, err error) {
	query := `SELECT [id]
      ,[tag_id]
      ,[title]
      ,[desc]
      ,[content]
      ,[created_on]
      ,[created_by]
      ,[modified_on]
      ,[modified_by]
      ,[deleted_on]
      ,[state]
      ,[cover_image_url]
     FROM [blog_article] where [title] = ?`
	err = a.Conn.QuerySingle(res, query, title)
	return
}

func (a *ArticleRepository) Create(ctx context.Context, art *models.Article) (err error) {
	sql := `INSERT INTO [blog_article]
           ([tag_id]
           ,[title]
           ,[desc]
           ,[content]
           ,[created_on]
           ,[created_by]
           ,[modified_on]
           ,[modified_by]
           ,[deleted_on]
           ,[state]
           ,[cover_image_url])
     VALUES
           (?,?,?,?,?,?,?,?,?,?,?)`
	_, err = a.Conn.Exec(sql, art.TagID, art.Title, art.Desc, art.Content,
		art.CreatedOn, art.ModifiedOn, art.ModifiedBy, art.DeletedOn, art.State, art.CoverImageUrl)
	return
}

func (a *ArticleRepository) Delete(ctx context.Context, id int64) (err error) {
	_, err = a.Conn.Exec("DELETE FROM [blog_article] WHERE id = ?", id)
	return
}

func (a *ArticleRepository) Update(ctx context.Context, art *models.Article) (err error) {
	sql := `UPDATE [blog_article]
	   SET [title] = ?
	      ,[desc] = ?
	      ,[content] = ?
	      ,[created_on] = ?
	      ,[created_by] = ?
	      ,[modified_on] = ?
	      ,[modified_by] = ?
	      ,[deleted_on] = ?
	      ,[state] = ?
	      ,[cover_image_url] = ?
         WHERE [id] = ?`
	_, err = a.Conn.Exec(sql, art.Title, art.Desc, art.Content, art.CreatedOn, art.CreatedBy,
		art.ModifiedOn, art.ModifiedBy, art.DeletedOn, art.State, art.CoverImageUrl, art.ID)
	return
}
