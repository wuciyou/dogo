package router

import (
	"fmt"
	"github.com/wuciyou/dogo/context"
	"github.com/wuciyou/dogo/dglog"
	"strings"
)

type nodeType byte

const (
	routerName nodeType = iota
	paramName
	prefixParamName
)

type node struct {
	pathName      string
	pathOtherMeta string
	ctxHander     context.ContextHandle
	nType         nodeType
	childrenNode  []*node
}

func (n *node) addRouter(pathNode []string, ctxHander context.ContextHandle) {
	var isChildrenNode bool
	var newNode *node
	pathNodeLen := len(pathNode)
	if pathNodeLen <= 0 {
		return
	}

	var newPathNode []string

	for _, tempPath := range pathNode {
		tempPath := strings.TrimSpace(tempPath)
		if tempPath != "" {
			newPathNode = append(newPathNode, tempPath)

		}
	}

	dglog.Debugf("pathNode:%+v", pathNode)
	if newNode == nil {
		return
	}

	for _, cNode := range n.childrenNode {
		if cNode.pathName == newNode.pathName && cNode.nType == newNode.nType {
			isChildrenNode = true
			if pathNodeLen > 1 {
				cNode.addRouter(pathNode[1:], ctxHander)
			}
			break
		}
	}

	if !isChildrenNode {
		n.childrenNode = append(n.childrenNode, newNode)
	}
}

func (n *node) Print() {
	fmt.Printf("%s/", n.pathName)
	for _, childern := range n.childrenNode {
		fmt.Printf("%s/", childern.pathName)
		if len(childern.childrenNode) > 0 {
			childern.Print()
		}
		fmt.Println("")
	}
}

func getNode(path string, ctxHander context.ContextHandle) *node {
	var nType nodeType
	var pathName string
	var pathOtherMeta string

	pn := strings.TrimSpace(path)

	pnLen := len(pn)
	dglog.Debugf("pnLen:`%d`", pnLen)
	if pnLen <= 0 {
		return nil
	}

	dglog.Debugf("pn v:`%s`", pn)

	switch pn[0] {

	case ':':
		if pnLen < 2 {
			return nil
		}
		dglog.Debugf("this %s route is a paramName node", pn)
		nType = paramName
		pathName = pn[1:]

	default:
		paramNameIndex := strings.IndexAny(pn, ":")
		if paramNameIndex < 0 {
			nType = routerName
			pathName = pn
		} else {
			nType = prefixParamName
			pathName = pn[:paramNameIndex]
			pathOtherMeta = pn[paramNameIndex:]
		}
	}

	return &node{
		pathName:      pathName,
		pathOtherMeta: pathOtherMeta,
		ctxHander:     ctxHander,
		nType:         nType,
	}
}
