package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/wuciyou/dogo"
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

func NewModel(t *Transation) *Model {
	model := &Model{transation: t}
	return model
}

func (m *Model) clear() {
	m.table = ""
	m.where = nil
	m.paramData = nil
	m.limit = ""
	m.order = ""
}

func (m *Model) Table(name string) *Model {
	m.table = name
	return m
}

func (m *Model) checkTransation() (*Transation, error) {
	if m.transation == nil {
		return nil, errors.New("Can't transation is null")
	}
	return m.transation, nil
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

func (m *Model) selectSql(fields ...string) (string, []interface{}) {

	var field string
	where, dataParam := m.buildWhere()

	if len(fields) == 0 {
		field = "*"
	} else {
		field = strings.Join(fields, ",")
	}

	sql := fmt.Sprintf("SELECT %s FROM %s %s", field, m.table, where)
	return sql, dataParam

}

func (m *Model) Select(fields ...string) (*sql.Rows, error) {
	defer m.clear()
	transation, err := m.checkTransation()
	if err != nil {
		panic(err)
	}
	sql, dataParam := m.selectSql(fields...)
	tx := transation.Begin()

	rows, err := tx.Query(sql, dataParam...)
	if err != nil {
		transation.Rollback()
		dogo.Dglog.Errorf("获取数据失败 err:%v", err)
	} else {
		transation.Commit()
	}
	return rows, err
}

func (m *Model) SelectOne(fields ...string) *sql.Row {
	defer m.clear()
	transation, err := m.checkTransation()
	if err != nil {
		panic(err)
	}
	m.Limit(1, 1)
	sql, dataParam := m.selectSql(fields...)
	tx := transation.Begin()
	row := tx.QueryRow(sql, dataParam...)
	transation.Commit()
	return row
}

func (m *Model) Update(datas map[string]interface{}) (sql.Result, error) {

	defer m.clear()
	transation, err := m.checkTransation()
	if err != nil {
		panic(err)
	}
	where, whereParam := m.buildWhere()
	var updateFileds []string
	var updateValues []interface{}
	for filedName, data := range datas {
		updateFileds = append(updateFileds, fmt.Sprintf("`%s` = ? ", filedName))
		updateValues = append(updateValues, data)
	}

	sql := fmt.Sprintf("UPDATE %s SET %s %s", m.table, strings.Join(updateFileds, ","), where)

	tx := transation.Begin()
	dogo.Dglog.Debugf("update sql :%s", sql)
	result, err := tx.Exec(sql, append(updateValues, whereParam...)...)
	if err != nil {
		transation.Rollback()
		dogo.Dglog.Errorf("数据更新失败:err:%+v", err)
		return nil, err
	}
	if affected_num, err := result.RowsAffected(); affected_num <= 0 || err != nil {
		transation.Rollback()
		dogo.Dglog.Errorf("数据更新失败:Affected :%d ,err:%+v", affected_num, err)
		return nil, err
	}
	transation.Commit()
	return result, nil
}

func (m *Model) Insert(datas map[string]interface{}) (sql.Result, error) {
	defer m.clear()
	if len(datas) == 0 {
		return nil, errors.New("Need fields data")
	}
	transation, err := m.checkTransation()
	if err != nil {
		panic(err)
		return nil, err
	}

	var insert_fields []string
	var insert_placeholders []string
	var insert_values []interface{}

	for field_name, filed_value := range datas {
		insert_fields = append(insert_fields, fmt.Sprintf("`%s`", field_name))
		insert_placeholders = append(insert_placeholders, "?")
		insert_values = append(insert_values, filed_value)
	}

	sql := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", m.table, strings.Join(insert_fields, ","), strings.Join(insert_placeholders, ","))

	dogo.Dglog.Debugf("insert sql :%s", sql)
	tx := transation.Begin()

	result, err := tx.Exec(sql, insert_values...)

	if err != nil {
		transation.Rollback()
		dogo.Dglog.Errorf("数据插入失败:err:%+v", err)
		return nil, err
	}
	if affected_num, err := result.RowsAffected(); affected_num <= 0 || err != nil {
		transation.Rollback()
		dogo.Dglog.Errorf("数据插入失败:Affected :%d ,err:%+v", affected_num, err)
		return nil, err
	}
	transation.Commit()
	return result, nil
}

func (m *Model) Delete() (sql.Result, error) {
	defer m.clear()

	transation, err := m.checkTransation()

	if err != nil {
		panic(err)
		return nil, err
	}

	where, where_param := m.buildWhere()
	sql := fmt.Sprintf("DELETE FROM `%s` %s", m.table, where)
	dogo.Dglog.Debugf("delete sql :%s", sql)
	tx := transation.Begin()
	result, err := tx.Exec(sql, where_param...)
	if err != nil {
		transation.Rollback()
		dogo.Dglog.Errorf("数据删除失败:err:%+v", err)
		return nil, err
	}
	if affected_num, err := result.RowsAffected(); affected_num <= 0 || err != nil {
		transation.Rollback()
		dogo.Dglog.Errorf("数据删除失败:Affected :%d ,err:%+v", affected_num, err)
		return nil, err
	}
	transation.Commit()
	return result, nil

}

func (m *Model) buildWhere() (whereSql string, data []interface{}) {
	if m.where == nil {
		return
	}
	return m.where.buildWhere()
}
