package db

import "go-training/season2/db/mysql"

//只是为了模拟db调用,重点关注error handle
type CURD interface {
	Insert() (int, error)
	Query() (int, error)
	Delete() (int, error)
	Update() (int, error)
}

func GetCURDInstance() CURD {
	return mysql.NewEngine()
}
