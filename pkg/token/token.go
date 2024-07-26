// Copyright 2024 summingyu(余苏明) <summingbest@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/summingyu/miniblog.

package token

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

// Config 包含token包的配置选项
type Config struct {
	key         string
	identityKey string
}

var ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")

var (
	config = Config{"changethis", "identityKey"}
	once   sync.Once
)

// Init 设置包级别的配置 config, config 会用于本包后面的token生成和解析
func Init(key string, identityKey string) {
	once.Do(func() {
		if key != "" {
			config.key = key
		}
		if identityKey != "" {
			config.identityKey = identityKey
		}
		config = Config{key, identityKey}
	})
}

// Parse 使用指定的密钥 key 解析 token, 解析成功返回 token 上下文, 否则报错.
func Parse(tokenString string, key string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}

	var identityKey string

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		identityKey = claims[config.identityKey].(string)
	}
	return identityKey, nil
}

// ParseRequest 从请求头中获取令牌, 并将其传递给 Parse 函数以解析令牌
func ParseRequest(c *gin.Context) (string, error) {
	header := c.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return "", ErrMissingHeader
	}

	var t string
	fmt.Sscanf(header, "Bearer %s", &t)

	return Parse(t, config.key)
}

// Sign 使用jwtSecret 签发 token, token的claims中会存放传入的subject.
func Sign(identityKey string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identityKey,
		"nbf":              time.Now().Unix(),
		"iat":              time.Now().Unix(),
		"exp":              time.Now().Add(100000 * time.Hour).Unix(),
	})

	tokenString, err = token.SignedString([]byte(config.key))
	return
}
