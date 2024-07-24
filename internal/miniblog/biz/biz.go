package biz

import (
	"github.com/summingyu/miniblog/internal/miniblog/biz/user"
	"github.com/summingyu/miniblog/internal/miniblog/store"
)

// IBiz 定义了 Biz 层需要实现的方法
type IBiz interface {
	Users() user.UserBiz
}

// biz 是 IBiz 的实现
type biz struct {
	ds store.IStore
}

var _ IBiz = (*biz)(nil)

func NewBiz(ds store.IStore) IBiz {
	return &biz{
		ds: ds,
	}
}

func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}
