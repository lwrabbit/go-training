# go-training
geektime 

## season2 Error handling
```go
$ cd $GOPATH/src/go-training
$ go run season2/main.go
```
```text
Output:
main: sql: no rows in result set
mysql: engine err 
go-training/season2/db/mysql.(*engine).Query
        $GOPATH/src/go-training/season2/db/mysql/mysql.go:19
main.main
        $GOPATH/src/go-training/season2/main.go:10
runtime.main
        $GOROOT/src/runtime/proc.go:204
runtime.goexit
        $GOROOT/src/runtime/asm_amd64.s:1374
```
