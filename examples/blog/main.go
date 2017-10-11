package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wuciyou/dogo"
	"github.com/wuciyou/dogo/database"
	"time"
)

func databaseDelete(ctx *dogo.Context) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ithome?charset=utf8&loc=Asia%2FShanghai")
	if err != nil {
		panic(err)
	}
	transation := database.NewTransation(db)
	transation.Begin()
	model := database.NewModel(transation)
	result, err := model.Table("it_chat").Where(func(w *database.Where) {
		w.And("id", "=?", "d766e06c-6f15-c533-c229-74902115109a")
		w.Or("uid", "=?", "99493904-baf8-82d4-aa13-72ecb40fcd22")
	}).Delete()
	if err != nil {
		panic(err)
	}
	dogo.Dglog.Debugf("删除数据：%+v", result)
	transation.Commit()
}

func databaseInsert(ctx *dogo.Context) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ithome?charset=utf8&loc=Asia%2FShanghai")
	if err != nil {
		panic(err)
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")

	transation := database.NewTransation(db)
	transation.Begin()
	model := database.NewModel(transation)
	result, err := model.Table("it_chat").Insert(map[string]interface{}{
		"id":          "dwwuciyfIDIDDI",
		"uid":         "wuciyou的账号",
		"chat_type":   "ONE_TO_ONE",
		"create_time": time.Now().In(loc),
	})

	if err != nil {
		panic(err)
	}
	dogo.Dglog.Debugf("插入数据：%+v time:%s", result, time.Now().In(loc).String())
	transation.Commit()
}

func main() {
	app := dogo.App()
	app.Route().Filter(func() {
		app.Route().Get("/database/insert", databaseInsert)
	}, func(ctx *dogo.Context) bool {
		ctx.W.Write([]byte("安全认证失败"))
		ctx.W.Send()
		return false
	})

	app.Route().Get("/database/delete", databaseDelete)
	app.Run()
}
