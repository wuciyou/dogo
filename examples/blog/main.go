package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wuciyou/dogo"
	"github.com/wuciyou/dogo/database"
)

func myFunc(c *dogo.Context) {
	dogo.Debug("有新的页面请求")
	dogo.Dglog.Debugf("新的请求 headear:%s", c.R.Header)
	name := c.Get("name")
	name1 := c.Post("name1")
	ip, port := c.ClientIp()
	c.SetCookie("name", "wuciyou")
	c.SetCookie("name", "wuwuciyou")
	token := c.GetCookie("token")
	c.W.Write([]byte(fmt.Sprintf("namde:%s-name1:%s ,ip:%s port:%d token:%s", name, name1, ip, port, token)))
	c.W.Send()
}

func mysqlFunc(ctx *dogo.Context) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ithome")
	if err != nil {
		dogo.Dglog.Errorf("连接数据库失败：%v", err)
		panic(err)
	}
	dogo.Dglog.Debugf("连接数据库成功：%v", db)
	transation := database.NewTransation(db)
	model := database.NewModel(transation)
	rows, err := model.Table("it_article").Where(func(w *database.Where) {
		w.And("article_collection_id", "=?", "fewe7cba797-3d83-c2fb-d931-43d98427e5a3").Or("title", "like ?", "%天气%")
	}).Select("id")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id string
		rows.Scan(&id)
		dogo.Dglog.Debugf("id :%s", id)
	}
	dogo.Dglog.Debugf("rows:%+v", rows)

}

func jsonFunc(ctx *dogo.Context) {
	ctx.W.Json(map[string]string{
		"name": "wuciyou",
	})
}

func htmlFunc(ctx *dogo.Context) {
	ctx.W.Assign("name", "iwuciyou")
	ctx.W.Display("./test_data/base.html")
}

func main() {
	dogo.Route.Get("/hello/wuciyou/abc", myFunc)
	dogo.Route.Get("/hello/[a-b]/abc", myFunc)
	dogo.Route.Any("/hello/html", htmlFunc)
	dogo.Route.Any("/hello/mysql", mysqlFunc)
	dogo.Route.Get("/hello/json", jsonFunc)
	dogo.Route.Post("/hello/wuciyou", func(ctx *dogo.Context) {
		_, file, _ := ctx.R.FormFile("myFile")
		dogo.Dglog.Debugf("file:%+v \n", file)
	})
	dogo.App().Run()
}
