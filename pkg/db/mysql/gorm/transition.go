package gorm

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"reflect"
)

// 事务key
type ctxTransactionKey struct{}

func CtxWithTransaction(c context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(c, ctxTransactionKey{}, tx)
}

func GetTransactionFromCtx(c context.Context) (*gorm.DB, error) {
	t := c.Value(ctxTransactionKey{})
	if t != nil {
		tx, ok := t.(*gorm.DB)
		if !ok {
			return nil, errors.Errorf("unexpect context value type: %s", reflect.TypeOf(tx))
		}
		return tx, nil
	}
	return nil, nil
}
