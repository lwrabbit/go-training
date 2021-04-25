package mysql

import (
	"database/sql"
	"github.com/pkg/errors"
	xerrors "github.com/pkg/errors"
)

var myErr = errors.New("mysql: engine err ")

type engine struct{}

//NewEngine return a mock implement instance
func NewEngine() *engine {
	return &engine{}
}

func (e *engine) Query() (affectRows int, err error) {
	return -1, xerrors.Wrapf(sql.ErrNoRows, myErr.Error())
	//bad case: return -1, xerrors.Wrapf(myErr,sql.ErrNoRows.Error())
}

func (e *engine) Delete() (affectRows int, err error) {
	return affectRows, err
}
func (e *engine) Update() (affectRows int, err error) {
	return affectRows, err
}

func (e *engine) Insert() (effectRows int, err error) {
	return effectRows, err
}
