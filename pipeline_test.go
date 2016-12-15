package thinkgo

import (
	"testing"
)

func TestAddLast(t *testing.T) {
	Commonpipeline.AddLast("a")
	Commonpipeline.AddLast("b")
	Commonpipeline.AddLast("c")
	Commonpipeline.AddLast("d")
	Commonpipeline.AddLast("d1")
	Commonpipeline.each(func(p *pipeline) {
		t.Logf("pipeline:%+v", p)

	})
	Commonpipeline.each(func(p *pipeline) {
		t.Logf("pipeline each:%+v", p)

	})

}
