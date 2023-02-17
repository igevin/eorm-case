package quick_start

import (
	"context"
	"database/sql"
	"github.com/gotomicro/eorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type crudSuite struct {
	suite.Suite

	driver string
	dsn    string

	db *eorm.DB
}

func (s *crudSuite) SetupSuite() {
	//db, err := eorm.Open("sqlite3",
	//	"file:test.db?cache=shared&mode=memory")
	db, err := eorm.Open(s.driver, s.dsn)
	s.Require().NoError(err)
	s.db = db
	s.createTables()
}

func (s *crudSuite) createTables() {
	q := `create table if not exists person
(
    id       integer not null primary key autoincrement,
    name     text    not null,
    birthday datetime
);
`
	db, err := sql.Open(s.driver, s.dsn)
	s.Require().NoError(err)
	_, err = db.Exec(q)
	s.Require().NoError(err)
}

func TestCrud(t *testing.T) {
	s := &crudSuite{
		driver: "sqlite3",
		dsn:    "file:test.db?cache=shared&mode=memory",
	}
	suite.Run(t, s)
}

func (s *crudSuite) TestInsertPerson() {
	p := &Person{
		Id:   1000,
		Name: "Tom",
	}
	res := InsertPerson(context.Background(), s.db, p)
	t := s.T()
	require.NoError(t, res.Err())
	affected, err := res.RowsAffected()
	assert.Equal(t, int64(1), affected)
	require.NoError(t, err)
	lastId, err := res.LastInsertId()
	require.NoError(t, err)
	assert.Equal(t, int64(1000), lastId)
}
