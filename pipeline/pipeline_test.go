package pipeline

import (
	"testing"
)

var pipelineArr = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "n", "m", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func TestAdd(t *testing.T) {
	log := &LogPipeline{}
	Commonpipeline.AddLast("d", log)
	Commonpipeline.AddLast("e", log)
	Commonpipeline.AddLast("f", log)
	Commonpipeline.AddLast("g", log)
	Commonpipeline.AddFirst("c", log)
	Commonpipeline.AddLast("h", log)
	Commonpipeline.AddFirst("b", log)
	Commonpipeline.AddFirst("a", log)
	Commonpipeline.AddLast("k", log)
	Commonpipeline.AddBefore("k", "i", log)
	Commonpipeline.AddAfter("i", "j", log)

	var count int
	Commonpipeline.each(func(p *pipelineNode) bool {
		if p.name != pipelineArr[count] {
			t.Errorf("Commonpipeline index:%d , name[%s] not equal pipelineArr[%s] \n", count, p.name, pipelineArr[count])
		}
		count++
		return true
	})
}
