package data

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type Data struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) *Data {
	return &Data{db: db}
}

type contextTxKey struct {
}

type DBOption func(*gorm.DB) *gorm.DB

type BaseRepo interface {
	// WithByID 公用方法
	WithByID(id int64) DBOption
	// Paginate 分页
	Paginate(page, size int) DBOption
	// UpdateData 公用方法，更新数据
	UpdateData(column string, value interface{}) DBOption
	// OrderBy 排序，传入数据列名和排序方法（asc/desc）
	OrderBy(column, orderType string) DBOption
	// Where 查询条件
	Where(query interface{}, args ...interface{}) DBOption
	// ExecTx 事务处理
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}

func (d *Data) DB(ctx ...context.Context) *gorm.DB {
	if len(ctx) == 1 {
		tx, ok := ctx[0].Value(contextTxKey{}).(*gorm.DB)
		if ok {
			return tx
		}
	}
	return d.db
}

func (d *Data) optionDB(ctx context.Context, opts ...DBOption) *gorm.DB {
	db := d.DB(ctx).WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	return db
}

func (d *Data) WithByID(id int64) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

func (d *Data) UpdateData(column string, value interface{}) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Update(column, value)
	}
}

func (d *Data) Where(query interface{}, args ...interface{}) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args)
	}
}

func (d *Data) Paginate(page, size int) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 10
		}
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

func (d *Data) OrderBy(column, orderType string) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", column, orderType))
	}
}

func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}
