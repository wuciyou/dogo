package dogo

import (
	"net/http"
)

type PipelineHandle interface {
	PipelineRun(http.ResponseWriter, *http.Request) bool
}

type pipelineNode struct {
	name PipelineKey
	h    PipelineHandle
}

type pipeline struct {
	pipelineNodeSize int
	pipelineNode     []*pipelineNode
}

var Commonpipeline = &pipeline{}

func (p *pipeline) AddLast(name PipelineKey, handle PipelineHandle) {
	p.checkName(name)
	node := &pipelineNode{name: name, h: handle}
	p.pipelineNode = append(p.pipelineNode, node)
	DogoLog.Debugf("Add pipelineHandle to last for name[%s],pipelineNode:%+v", name, p.pipelineNode)
}

func (p *pipeline) AddFirst(name PipelineKey, handle PipelineHandle) {
	p.checkName(name)
	node := &pipelineNode{name: name, h: handle}
	p.pipelineNode = append([]*pipelineNode{node}, p.pipelineNode...)
	DogoLog.Debugf("Add pipelineHandle to first for name[%s]", name)
}

func (p *pipeline) AddAfter(afterName PipelineKey, name PipelineKey, handle PipelineHandle) {
	if tempPipelineNode, i := p.getByName(afterName); tempPipelineNode != nil {
		node := &pipelineNode{name: name, h: handle}

		pipelineNodeSize := len(p.pipelineNode)
		if i == pipelineNodeSize-1 {
			p.pipelineNode = append(p.pipelineNode, node)
		} else {
			p.pipelineNode = append(p.pipelineNode[:i+1], append([]*pipelineNode{node}, p.pipelineNode[i+1:]...)...)
		}
		DogoLog.Debugf("Add pipelineHandle to name[%s] after for name[%s]", afterName, name)
	} else {
		DogoLog.Errorf("Can't found name[%s] on the pipeline\n ", afterName)
	}
}

func (p *pipeline) AddBefore(beforeName PipelineKey, name PipelineKey, handle PipelineHandle) {
	if tempPipelineNode, i := p.getByName(beforeName); tempPipelineNode != nil {
		node := &pipelineNode{name: name, h: handle}
		p.pipelineNode = append(p.pipelineNode[:i], append([]*pipelineNode{node}, p.pipelineNode[i:]...)...)
		DogoLog.Debugf("Add pipelineHandle to name[%s] before for name[%s]", beforeName, name)
	} else {
		DogoLog.Errorf("Can't found name[%s] on the pipeline\n ", beforeName)
	}
}

func (p *pipeline) checkName(name PipelineKey) {
	if node, _ := p.getByName(name); node != nil {
		DogoLog.Errorf("The same key[%s] already exists", name)
	}
}

func (p *pipeline) getByName(name PipelineKey) (*pipelineNode, int) {
	var tempPipeline *pipelineNode
	var index int
	p.each(func(n *pipelineNode) bool {
		if n.name == name {
			tempPipeline = n
			return false
		} else {
			index++
		}
		return true
	})
	return tempPipeline, index
}

func (p *pipeline) each(f func(*pipelineNode) bool) bool {
	for i := 0; i < len(p.pipelineNode); i++ {
		if p.pipelineNode[i] == nil {
			continue
		}
		is_break := f(p.pipelineNode[i])
		if !is_break {
			return false
		}
	}
	return true
}
