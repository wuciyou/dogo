package database

import (
	"testing"
)

func TestWhere(t *testing.T) {
	w := &Where{}
	w.And("name", ">?", "wuciyou")
	w.WhereFunc(func(newWhwer *Where) {
		newWhwer.Or("age", ">?", 10)
		newWhwer.WhereFunc(func(myWhere *Where) {
			myWhere.Or("time", "=?", 1000).Or("isShow", "=?", true)
			myWhere.WhereFunc(func(myWhere *Where) {
				myWhere.And("time", "=?", 1000).And("isShow", "=?", true)
			})
		})
	})
	sql, data := w.buildWhere()
	t.Logf("sql:%s , data:%+v", sql, data)
}

func TestModel(t *testing.T) {
	m := &Model{}
	sql := m.Table("it_artice").Where(func(w *Where) {
		w.And("name", "=?", "wuciyou").Or("age", "<= ?", 100)
	}).Select()
	t.Logf("sql:%s", sql)
}
