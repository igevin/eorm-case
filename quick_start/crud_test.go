package quick_start

import (
	"github.com/gotomicro/eorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPersonInserterSql(t *testing.T) {
	db := memoryDB(t)
	p := &Person{
		Id:   1000,
		Name: "Tom",
	}
	q, err := PersonInserterSql(db, p)
	require.NoError(t, err)
	t.Log(q.SQL, q.Args)
	expectedSql := "INSERT INTO `person`(`id`,`name`) VALUES(?,?);"
	expectedArgs := []any{int64(1000), "Tom"}
	assert.Equal(t, expectedSql, q.SQL)
	assert.Equal(t, expectedArgs, q.Args)
}

func memoryDB(t *testing.T, opts ...eorm.DBOption) *eorm.DB {
	db, err := eorm.Open("sqlite3",
		"file:test.db?cache=shared&mode=memory",
		// 仅仅用于单元测试，不会发起真的查询
		opts...)
	require.NoError(t, err)
	return db
}
