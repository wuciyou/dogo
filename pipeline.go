package thinkgo

import (
	"log"
	"net/http"
)

type PipelineHandle interface {
	PipelineRun(http.ResponseWriter, *http.Request) bool
}

type pipeline struct {
	pipelineNodeSize int
	node             *pipelineNode
	firsetNode       *pipelineNode
	lastNode         *pipeline
}

var Commonpipeline *pipeline

func (p *pipeline) AddLast(name string) {

}

func (p *pipeline) AddFirst(name string) {

}

func (p *pipeline) AddAfter() {

}

func (p *pipeline) AddBefore() {

}

func (p *pipeline) checkName(name string) {
	if pi := p.getByName(name); pi != nil {
		log.Panicf("The [%s] key already exists \n", name)
	}
}

func (p *pipeline) getByName(name string) *pipeline {
	var newPipeline *pipeline
	p.each(func(p1 *pipeline) {
		if p1.name == name {
			newPipeline = p1
		}
	})
	return newPipeline
}

func (p *pipeline) each(f func(*pipeline)) {
	each := p
	for {
		if next := each.next; next != nil {
			f(each)
			each = next
			if next.isFirst {
				return
			}
		} else {
			break
		}
	}

}

func init() {
	Commonpipeline = &pipeline{isFirst: true}
}
