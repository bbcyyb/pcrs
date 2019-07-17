package repository

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bbcyyb/pcrs/article"
	"github.com/bbcyyb/pcrs/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository article.Repository
	article    models.Article
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
	s.repository = NewArticleRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGetAll() {
	rows := sqlmock.NewRows([]string{"id", "tag_id", "title", "desc", "content", "created_by", "state", "cover_image_url"}).
		AddRow(1, 1, "title 1", "desc 1", "Content 1", "user1", 1, "http://book.com.cn").
		AddRow(2, 2, "title 2", "desc 2", "Content 2", "user2", 1, "http://book.com.cn")

	query := `SELECT * FROM "articles"`
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	list, err := s.repository.GetAll(context.TODO())
	require.NoError(s.T(), err)
	require.Len(s.T(), list, 2)
}

func (s *Suite) TestGetById() {
	id := int64(12)
	rows := sqlmock.NewRows([]string{"id", "tag_id", "title", "desc", "content", "created_by", "state", "cover_image_url"}).
		AddRow(id, 1, "title 1", "desc 1", "Content 1", "Jane", 1, "http://book.com.cn")

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "articles" WHERE (id = $1)`)).
		WithArgs(id).
		WillReturnRows(rows)

	article, err := s.repository.GetByID(context.TODO(), id)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), article)
}

func (s *Suite) TestGetByTitle() {
	title := "title 1"

	rows := sqlmock.NewRows([]string{"id", "tag_id", "title", "desc", "content", "created_by", "state", "cover_image_url"}).
		AddRow(1, 1, title, "desc 1", "Content 1", "Jane", 1, "http://book.com.cn")
	query := `SELECT * FROM "articles" WHERE (title = $1)`

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(title).WillReturnRows(rows)
	article, err := s.repository.GetByTitle(context.TODO(), title)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), article)
}

func (s *Suite) TestDelete() {
	idNum := int64(12)
	query := `DELETE FROM [blog_article] WHERE id = $1`
	s.mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(idNum).WillReturnResult(sqlmock.NewResult(idNum, 1))
	err := s.repository.Delete(context.TODO(), idNum)
	require.NoError(s.T(), err)
}

//// TODO
//func (s *Suite) TestCreate() {
//
//	ar := &models.Article{
//		TagID:         1,
//		Title:         "title1",
//		Desc:          "desc1",
//		Content:       "content1",
//		CreatedBy:     "user1",
//		State:         1,
//		CoverImageUrl: "http://book.com.cn",
//	}
//	query := `INSERT INTO "articles" ("tag_id", "title", "desc", "content", "created_by", "state", "cover_image_url") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "articles"."TagID"`
//
//	rows := sqlmock.NewRows([]string{"tag_id"}).
//		AddRow(ar.TagID)
//
//	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
//		WithArgs(ar.TagID, ar.Title, ar.Desc, ar.Content, ar.CreatedBy, ar.State, ar.CoverImageUrl).
//		WillReturnRows(rows)
//
//	//prep := s.mock.ExpectPrepare(query)
//	//prep.ExpectExec().WithArgs(ar.TagID, ar.Title, ar.Desc, ar.Content, ar.CreatedBy, ar.State, ar.CoverImageUrl)
//
//	err := s.repository.Create(context.TODO(), ar)
//	require.NoError(s.T(), err)
//}

// TODO
//func (s *Suite) TestUpdate() {
//	model := &models.Model{
//		ID: 1,
//	}
//	ar := &models.Article{
//		Model:         *model,
//		TagID:         1,
//		Title:         "Pride",
//		Desc:          "It is a article name.",
//		Content:       "It is about social story.",
//		CreatedBy:     "Jane",
//		State:         1,
//		CoverImageUrl: "http://book.com.cn",
//	}
//	query := `UPDATE [blog_article] SET tag_id=?, title=?, desc=?, content=?, createdBy=?, state=?, coverImageUrl=? WHERE id = ?`
//
//	prep := s.mock.ExpectPrepare(query)
//	prep.ExpectExec().WithArgs(ar.TagID, ar.Title, ar.Desc, ar.Content, ar.CreatedBy, ar.State, ar.CoverImageUrl, ar.ID).
//		WillReturnResult(sqlmock.NewResult(12, 1))
//
//	err := s.repository.Update(context.TODO(), ar)
//	require.NoError(s.T(), err)
//}
