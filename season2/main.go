package main

import (
	"fmt"
	//"github.com/pkg/errors"
	"go-training/season2/db"
)

func main() {
	curd := db.GetCURDInstance()
	if res, err := curd.Query(); err != nil {
		//fmt.Printf("main: %T %v\n", errors.Cause(err),errors.Cause(err))
		fmt.Printf("main: %+v\n", err)
	} else {
		//handle res ...
		_ = res
	}
}
