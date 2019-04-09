go-bindata -pkg conf -o main/inits/bindata/conf/conf-data.go conf
go-bindata -pkg static -o main/inits/bindata/static/static-data.go resources/...
cd main
go build
cd ..


