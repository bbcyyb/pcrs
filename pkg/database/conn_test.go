package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type Suite struct {
	suite.Suite
	DB     *gorm.DB
	mockDB sqlmock.Sqlmock
	c      *Conn

	expected     TestData
	expectedMany []TestData
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mockDB, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	s.c = &Conn{
		s.DB, nil,
	}
	require.NoError(s.T(), err)

	s.expected = TestData{"one", "two", "three"}
	s.expectedMany = []TestData{
		{"one", "two", "three"},
	}
}

func (s *Suite) TestQuery() {
	rows := sqlmock.NewRows([]string{"one", "two", "three"}).
		AddRow("one", "two", "three")
	query := `SELECT * FROM articles`
	s.mockDB.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	resList, err := s.c.Query(query)
	s.NoError(err)
	s.NotNil(resList)
}

func (s *Suite) TestQueryRow() {
	query := `SELECT * FROM articles`
	res := s.c.QueryRow(query)
	s.NotNil(res)
}

func (s *Suite) TestExec() {
	query := `Delete articles Where [id] = 1"`
	s.mockDB.ExpectExec("Delete articles").WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := s.c.Exec(query)
	s.Nil(err)
	s.NotNil(res)
	id, _ := res.LastInsertId()
	rowAffected, _ := res.RowsAffected()
	s.Equal(1, int(id))
	s.Equal(1, int(rowAffected))
}

type TestData struct {
	One   string `json:"one"`
	Two   string `json:"two"`
	Three string `json:"three"`
}

type testSQLRows struct {
	iteration int
}

func (t *testSQLRows) Next() bool {
	return t.iteration == 1
}

func (t *testSQLRows) Columns() ([]string, error) {
	return []string{"one", "two", "three"}, nil
}

func (t *testSQLRows) Close() error {
	return nil
}

func (t *testSQLRows) Scan(args ...interface{}) error {
	one := args[0].(*interface{})
	two := args[1].(*interface{})
	three := args[2].(*interface{})
	t.iteration++

	*one = "one"
	*two = "two"
	*three = "three"
	return nil
}

func (s *Suite) TestScanRows() {
	s.c.rows = &testSQLRows{1}
	tester := []TestData{}
	err := s.c.scanRows(&tester)
	s.Nil(err)
	s.Equal(1, len(tester))
	s.Equal(s.expectedMany, tester)
}

func (s *Suite) TestScanRow() {
	rows := &testSQLRows{1}
	s.c.rows = rows
	tester := TestData{}
	err := s.c.scanRow(&tester)
	s.Nil(err)
	s.Equal(s.expected, tester)
	s.Equal(tester.One, "one")
	s.Equal(tester.Two, "two")
	s.Equal(tester.Three, "three")
}
