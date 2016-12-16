package dogo

import (
	"net/http"
)

type PipelineHandle interface {
	PipelineRun(http.ResponseWriter, *http.Request) bool
}

type pipelineNode struct {
	name string
	h    PipelineHandle
	prev *pipelineNode
	next *pipelineNode
}

type pipeline struct {
	pipelineNodeSize int
	firsetNode       *pipelineNode
}

var Commonpipeline *pipeline

func (p *pipeline) AddLast(name string, handle PipelineHandle) {
	if p.firsetNode == nil {
		node := &pipelineNode{name: name, h: handle}
		p.firsetNode = node
	} else {
		p.checkName(name)
		node := &pipelineNode{name: name, h: handle}

		if p.firsetNode.prev != nil {
			node.prev = p.firsetNode.prev
			p.firsetNode.prev.next = node
		}
		if p.firsetNode.next == nil {
			node.prev = p.firsetNode
			p.firsetNode.next = node
		}

		p.firsetNode.prev = node
		node.next = p.firsetNode

	}
}

func (p *pipeline) AddFirst(name string, handle PipelineHandle) {
	if p.firsetNode == nil {
		node := &pipelineNode{name: name, h: handle}
		p.firsetNode = node
	} else {
		p.checkName(name)
		node := &pipelineNode{name: name, h: handle}

		if p.firsetNode.prev != nil {
			p.firsetNode.prev.next = node
			node.prev = p.firsetNode.prev
		}

		if p.firsetNode.next == nil {
			p.firsetNode.next = node
		}

		node.next = p.firsetNode
		p.firsetNode.prev = node
		p.firsetNode = node
	}
}

func (p *pipeline) AddAfter(afterName string, name string, handle PipelineHandle) {
	if tempPipelineNode := p.getByName(afterName); tempPipelineNode != nil {
		node := &pipelineNode{name: name, h: handle}

		node.next = tempPipelineNode.next
		node.prev = tempPipelineNode

		tempPipelineNode.next.prev = node
		tempPipelineNode.next = node
	} else {
		DogoLog.Panicf("Can't found name[%s] on the pipeline\n ", afterName)
	}
}

func (p *pipeline) AddBefore(beforeName string, name string, handle PipelineHandle) {
	if tempPipelineNode := p.getByName(beforeName); tempPipelineNode != nil {
		node := &pipelineNode{name: name, h: handle}

		node.prev = tempPipelineNode.prev
		node.next = tempPipelineNode

		tempPipelineNode.prev.next = node
		tempPipelineNode.prev = node

	} else {
		DogoLog.Panicf("Can't found name[%s] on the pipeline\n ", beforeName)
	}
}

func (p *pipeline) checkName(name string) {
	if node := p.getByName(name); node != nil {
		DogoLog.Panicf("The same key[%s] already exists", name)
	}
}

func (p *pipeline) getByName(name string) *pipelineNode {
	var tempPipeline *pipelineNode
	p.each(func(n *pipelineNode) bool {
		if n.name == name {
			tempPipeline = n
		}
		return true

	})
	return tempPipeline
}

func (p *pipeline) each(f func(*pipelineNode) bool) bool {
	if !f(p.firsetNode) {
		return false
	}
	each := p.firsetNode
	for {
		if next := each.next; next != nil {

			each = next
			if next == p.firsetNode {
				return true
			}

			is_break := f(each)
			if !is_break {
				return false
			}

		} else {
			break
		}
	}

	return true

}

func init() {
	Commonpipeline = &pipeline{}
}
