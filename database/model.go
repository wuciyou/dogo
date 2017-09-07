package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type WhereRule struct {
	field           string
	rule            string
	data            []interface{}
	isAnd           bool
	openBracketNum  int
	closeBracketNum int
}

type Where struct {
	rules []*WhereRule
}

func (w *Where) And(field string, rule string, data ...interface{}) *Where {
	whereRule := &WhereRule{field: field, rule: rule, data: data, isAnd: true}
	w.AddRule(whereRule)
	return w
}

func (w *Where) Or(field string, rule string, data ...interface{}) *Where {
	whereRule := &WhereRule{field: field, rule: rule, data: data, isAnd: false}
	w.AddRule(whereRule)
	return w
}

func (w *Where) WhereFunc(wf func(w *Where)) {
	newWhere := &Where{}
	wf(newWhere)
	if newWhere.rules != nil && len(newWhere.rules) > 0 {
		newWhere.rules[0].openBracketNum++
		newWhere.rules[len(newWhere.rules)-1].closeBracketNum++
	}
	w.AddRule(newWhere.rules...)
}

func (w *Where) AddRule(whereRule ...*WhereRule) {

	if w.rules == nil {
		w.rules = []*WhereRule{}
	}
	w.rules = append(w.rules, whereRule...)
}

func (w *Where) buildWhere() (whereSql string, data []interface{}) {

	for k, r := range w.rules {
		if k != 0 {
			if r.isAnd {
				whereSql += " AND "
			} else {

				whereSql += " OR "
			}
		} else {
			whereSql += " WHERE "
		}
		if r.openBracketNum > 0 {
			for r.openBracketNum > 0 {
				whereSql += " ( "
				r.openBracketNum--
			}
		}
		whereSql += fmt.Sprintf("%s %s", r.field, r.rule)
		if r.closeBracketNum > 0 {
			for r.closeBracketNum > 0 {
				whereSql += " ) "
				r.closeBracketNum--
			}
		}
		data = append(data, r.data...)

	}

	if strings.Count(whereSql, "?") != len(data) {
		panic(fmt.Sprintf("Can't use data:%+v query sql:%s, because data len is %d ,whereSql ? len is %d", data, whereSql, len(data), strings.Count(whereSql, "?")))
	}

	return

}

type Model struct {
	// 表名
	table      string
	transation *Transation
	where      *Where
	paramData  []interface{}
	limit      string
	order      string
}

func (m *Model) Table(name string) *Model {
	m.table = name
	return m
}

func (m *Model) Where(whereFunc func(w *Where)) *Model {
	if m.where == nil {
		m.where = &Where{}
	}
	whereFunc(m.where)
	return m
}

func (m *Model) Limit(start, length int) *Model {
	m.limit = fmt.Sprintf(" LIMIT %d,%d", start, length)
	return m
}

func (m *Model) Order(field, order string) *Model {
	m.order = fmt.Sprintf(" ORDER BY %s %s ", field, order)
	return m
}

func (m *Model) Select(fields ...string) (*sql.Rows, error) {

	transation, err := m.checkTransation()
	if err != nil {
		panic(err)
	}

	var field string
	where, dataParam := m.buildWhere()

	if len(fields) == 0 {
		field = "*"
	} else {
		field = strings.Join(fields, ",")
	}

	sql := fmt.Sprintf("SELECT %s FROM %s %s", field, m.table, where)

	tx := transation.Begin()
	return tx.Query(sql, dataParam...)
}

func (m *Model) checkTransation() (*Transation, error) {
	if m.transation == nil {
		return nil, errors.New("Can't transation is null")
	}
	return m.transation, nil
}

func (m *Model) SelectOne(fields ...string) (*sql.Rows, error) {
	m.Limit(1, 1)
	return m.Select(fields...)
}

func (m *Model) buildWhere() (whereSql string, data []interface{}) {
	if m.where == nil {
		return
	}
	return m.where.buildWhere()
}

func (m *Model) buildSql() string {
	return ""
}
