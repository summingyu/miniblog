package store

import (
	"sync"

	"gorm.io/gorm"
)

type datastore struct {
	// 数据库连接池
	db *gorm.DB
}

var (
	once sync.Once
	// 全局变量, 方便其他包直接调用已初始化好的 S 实例
	S *datastore
)

// Istore 定义了 Store 层需要实现的接口
type IStore interface {
	Users() UserStore
}

// 确保 datastore 实现了 IStore 接口
var _ IStore = (*datastore)(nil)

func (ds *datastore) Users() UserStore {
	return newUsers(ds.db)
}

func NewStore(db *gorm.DB) IStore {
	// 确保 S 只被初始化一次
	once.Do(func() {
		S = &datastore{
			db: db,
		}
	})
	return S
}