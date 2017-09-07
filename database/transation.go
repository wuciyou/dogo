package database

import (
	"database/sql"
)

type Transation struct {
	db            *sql.DB
	tx            *sql.Tx
	transationNum int
}

func NewModel(t *Transation) *Model {
	model := &Model{transation: t}
	return model
}

func NewTransation(db *sql.DB) *Transation {
	t := &Transation{db: db}
	return t
}
func A() {

}

/**
 * 开启事务
 *
 */
func (t *Transation) Begin() *sql.Tx {
	t.transationNum++
	if t.tx != nil {
		return t.tx
	}
	tx, err := t.db.Begin()
	if err != nil {
		panic(err)
		return nil
	}
	t.tx = tx
	return t.tx
}

/**
 * 提交事务
 *
 */
func (t *Transation) Commit() {
	t.transationNum--
	if t.transationNum == 0 {
		t.tx.Commit()
	}
}

/**
 * 回滚事务
 *
 */
func (t *Transation) Rollback() {
	t.transationNum--
	t.tx.Rollback()
}
