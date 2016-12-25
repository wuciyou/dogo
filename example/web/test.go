package main

import (
	"log"
	"net/http"
	"path"
	"path/filepath"
	"sync"
	"time"
)

var w *sync.WaitGroup

func main() {
	// f0()
	f0()
}

func f0() {
	w = &sync.WaitGroup{}

	for i := 0; i <= 10000; i++ {
		w.Add(1)
		go f1(i)
	}

	w.Wait()
}

func f1(index int) {
	for i := 0; i <= 100; i++ {
		f2(index)
	}
	w.Done()
}

func f2(index int) {
	var beginTime, endTime int64
	beginTime = time.Now().UnixNano()
	http.Get("http://127.0.0.1:8089")
	endTime = time.Now().UnixNano()
	log.Printf("当前：%d, 请求时间：%d us \n", index, (endTime-beginTime)/1000)

}

func f3() {
	p := "/static/dogo.png"
	d := string(http.Dir("."))
	p = filepath.Join(d, filepath.FromSlash(path.Clean("/"+p)))
	log.Printf("path:%s\n", p)
}
