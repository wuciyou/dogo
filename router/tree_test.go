package router

import (
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
	"strings"
	"testing"
)

func aTestGetNodeType(t *testing.T) {

	url_path := "/wuciyou/:age/:name(\\w)"

	url_path_split := strings.Split(url_path, "/")

	dglog.Debugf("url_path:%v", url_path_split)
}

func TestAddRouter(t *testing.T) {
	url_path := "/wuciyou//wuciyou:age/:name(\\w)"
	// url_path = "/*"
	n := &node{}
	pathNode := strings.Split(url_path, "/")

	n.addRouter(pathNode, func(ctx *context.Context) {

	})
	n.Print()
}
