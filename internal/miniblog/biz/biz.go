// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

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
