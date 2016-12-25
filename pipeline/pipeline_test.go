package pipeline

import (
	"testing"

	"github.com/wuciyou/dogo/context"

	"github.com/wuciyou/dogo/common"
	"github.com/wuciyou/dogo/dglog"
)

var pipelineArr = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "n", "m", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

type logTest struct {
}

func (l *logTest) PipelineRun(ctx *context.Context) bool {

	dglog.Infof("request:%+v \n ", ctx.Request)
	return true
}

func TestAdd(t *testing.T) {
	log := &logTest{}
	AddLast("d", log)
	AddLast("e", log)
	AddLast("f", log)
	AddLast("g", log)
	AddFirst("c", log)
	AddLast("h", log)
	AddFirst("b", log)
	AddFirst("a", log)
	AddLast("k", log)
	AddBefore("k", "l", log)
	AddAfter("l", "j", log)
	Replace("l", "i", log)
	var count int
	Each(func(name common.PipelineKey, handle PipelineHandle) bool {

		if string(name) != pipelineArr[count] {
			t.Errorf("Commonpipeline index:%d , name[%s] not equal pipelineArr[%s] \n", count, name, pipelineArr[count])
		}
		count++
		return true
	})
}
