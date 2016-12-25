package pipeline

import (
	"github.com/wuciyou/dogo/common"
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
)

type PipelineHandle interface {
	PipelineRun(ctx *context.Context) bool
}

type pipelineNode struct {
	name common.PipelineKey
	h    PipelineHandle
}

type pipeline struct {
	pipelineNodeSize int
	pipelineNode     []*pipelineNode
}

var commonpipeline = &pipeline{}

func (p *pipeline) addLast(name common.PipelineKey, handle PipelineHandle) {
	p.checkName(name)
	node := &pipelineNode{name: name, h: handle}
	p.pipelineNode = append(p.pipelineNode, node)
	dglog.Debugf("Add pipelineHandle to last for name[%s],pipelineNode:%+v", name, p.pipelineNode)
}

func (p *pipeline) addFirst(name common.PipelineKey, handle PipelineHandle) {
	p.checkName(name)
	node := &pipelineNode{name: name, h: handle}
	p.pipelineNode = append([]*pipelineNode{node}, p.pipelineNode...)
	dglog.Debugf("Add pipelineHandle to first for name[%s]", name)
}

func (p *pipeline) addAfter(afterName common.PipelineKey, name common.PipelineKey, handle PipelineHandle) {
	if tempPipelineHandle, i := p.getByName(afterName); tempPipelineHandle != nil {
		node := &pipelineNode{name: name, h: handle}

		pipelineNodeSize := len(p.pipelineNode)
		if i == pipelineNodeSize-1 {
			p.pipelineNode = append(p.pipelineNode, node)
		} else {
			p.pipelineNode = append(p.pipelineNode[:i+1], append([]*pipelineNode{node}, p.pipelineNode[i+1:]...)...)
		}
		dglog.Debugf("Add pipelineHandle to name[%s] after for name[%s]", afterName, name)
	} else {
		dglog.Errorf("Can't found name[%s] on the pipeline\n ", afterName)
	}
}

func (p *pipeline) addBefore(beforeName common.PipelineKey, name common.PipelineKey, handle PipelineHandle) {
	if tempPipelineHandle, i := p.getByName(beforeName); tempPipelineHandle != nil {
		node := &pipelineNode{name: name, h: handle}
		p.pipelineNode = append(p.pipelineNode[:i], append([]*pipelineNode{node}, p.pipelineNode[i:]...)...)
		dglog.Debugf("Add pipelineHandle to name[%s] before for name[%s]", beforeName, name)
	} else {
		dglog.Errorf("Can't found name[%s] on the pipeline\n ", beforeName)
	}
}
func (p *pipeline) replace(oldeName common.PipelineKey, name common.PipelineKey, handle PipelineHandle) {
	if tempPipelineHandle, i := p.getByName(oldeName); tempPipelineHandle != nil {
		node := &pipelineNode{name: name, h: handle}
		p.pipelineNode[i] = node
		dglog.Debugf("Replace pipelineHandle  name[%s] to old[%s]", name, oldeName)
	} else {
		dglog.Errorf("Can't found name[%s] on the pipeline\n ", oldeName)
	}
}

func (p *pipeline) checkName(name common.PipelineKey) {
	if node, _ := p.getByName(name); node != nil {
		dglog.Errorf("The same key[%s] already exists", name)
	}
}

func (p *pipeline) getByName(name common.PipelineKey) (PipelineHandle, int) {
	var tempPipelineHandle PipelineHandle
	var index int
	p.each(func(eachName common.PipelineKey, handle PipelineHandle) bool {
		if eachName == name {
			tempPipelineHandle = handle
			return false
		} else {
			index++
		}
		return true
	})
	return tempPipelineHandle, index
}

func (p *pipeline) each(f func(name common.PipelineKey, handle PipelineHandle) bool) bool {
	for i := 0; i < len(p.pipelineNode); i++ {
		if p.pipelineNode[i] == nil {
			continue
		}
		is_break := f(p.pipelineNode[i].name, p.pipelineNode[i].h)
		if !is_break {
			return false
		}
	}
	return true
}

func AddLast(name common.PipelineKey, handle PipelineHandle) {
	commonpipeline.addLast(name, handle)
}
func AddFirst(name common.PipelineKey, handle PipelineHandle) {
	commonpipeline.addFirst(name, handle)
}
func AddAfter(afterName common.PipelineKey, name common.PipelineKey, handle PipelineHandle) {
	commonpipeline.addAfter(afterName, name, handle)
}
func AddBefore(beforeName common.PipelineKey, name common.PipelineKey, handle PipelineHandle) {
	commonpipeline.addBefore(beforeName, name, handle)
}
func Replace(oldName common.PipelineKey, name common.PipelineKey, handle PipelineHandle) {
	commonpipeline.replace(oldName, name, handle)
}
func Each(f func(name common.PipelineKey, handle PipelineHandle) bool) bool {
	return commonpipeline.each(f)
}
