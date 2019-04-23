# go-iris-curd

#### 介绍

    go iris curd，参考
    https://gitee.com/yhm_my/go-iris

#### 软件架构

- golang
- iris
- casbin

#### 安装教程

1. 修改 conf 下的 app.yml 和 db.yml
2. 初始化数据库，sql 文件在 doc/casbin.sql

3. 使用 go-bindata 生成配置文件

```
    go-bindata -pkg conf -o main/inits/bindata/conf/conf-data.go conf
    go-bindata -pkg static -o main/inits/bindata/static/static-data.go resources/...
```

4. 使用 govendor

```
    govendor init 初始化，建立vendor.json
    govendor add +external 抓取文件
```

#### 使用说明

1. 运行 go run main/main.go
2. 打开浏览器 http://localhost:8088
3. xxxx
