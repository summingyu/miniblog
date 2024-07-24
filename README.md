## miniblog 项目

### db2struct使用

```bash
db2struct --gorm --no-json -H 127.0.0.1 -d miniblog -t user --package model --struct UserM -u miniblog -p 'xxxxxx' --target=internal/pkg/model/user.go
db2struct --gorm --no-json -H 127.0.0.1 -d miniblog -t post --package model --struct PostM -u miniblog -p 'xxxxxx' --target=internal/pkg/model/post.go

```
