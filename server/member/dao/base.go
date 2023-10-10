package dao

import (
	"context"
	gormBase "github.com/lwzphper/go-mall/pkg/db/mysql/gorm"
	"github.com/lwzphper/go-mall/server/member/global"
	"gorm.io/gorm"
)

func GetDB(ctx context.Context) *gorm.DB {
	var db *gorm.DB

	fromCtx, err := gormBase.GetTransactionFromCtx(ctx)
	if err != nil {
		global.Logger.Panicf("get transition from ctx error：%v", err)
		return nil
	}

	if fromCtx != nil { // 事务操作
		db = fromCtx
	} else {
		db = global.DB.WithContext(ctx)
	}
	return db
}

func Transaction(ctx context.Context, fc func(txctx context.Context) error) error {
	db := global.DB.WithContext(ctx)

	return db.Transaction(func(tx *gorm.DB) error {
		txctx := gormBase.CtxWithTransaction(ctx, tx)
		return fc(txctx)
	})
}
