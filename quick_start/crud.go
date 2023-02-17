package quick_start

import (
	"context"
	"github.com/gotomicro/eorm"
)

func PersonInserterSql(db *eorm.DB, person *Person) (*eorm.Query, error) {
	i := eorm.NewInserter[Person](db)
	return i.Values(person).Build()

}

func InsertPerson(ctx context.Context, db *eorm.DB, person *Person) eorm.Result {
	i := eorm.NewInserter[Person](db)
	return i.Values(person).Exec(ctx)
}