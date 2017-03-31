package core

import (
        "container/heap"
        "strconv"
        "math"
        "fmt"
)

const (
        BARRIER = 1
        PATH = 2

        DIRECT_VALUE = 10
        OBLIQUE_VALUE = 14
)

var openNodes PriorityQueue
var closeNodes map[string] *ANode
func init() {

}

func FindPath(mapInfo *MapInfo) {
        var count = 0;

        openNodes = make(PriorityQueue, 0)
        heap.Init(&openNodes)

        closeNodes = make(map[string] *ANode)

        openNodes.Push(mapInfo.Start)
        MoveNodes(mapInfo, &count)

        fmt.Printf("loop count: %d  ", count)
}

func MoveNodes(mapInfo *MapInfo, count *int) {
        endPn := mapInfo.End.P
        for {
                if openNodes.Len() <= 0 {
                        return
                }

                if IsInClose(endPn.X, endPn.Y) {
                        buildPath(mapInfo, mapInfo.End)
                        return
                }

                current := heap.Pop(&openNodes).(*ANode)
                closeNodes[GetKey(current.P)] = current

                AddSurroundNode(mapInfo, current)

                *count++
        }
}

//draw path
func buildPath(mapInfo *MapInfo, node *ANode) {
        for {
                if node == nil {
                        return
                }

                pn := node.P
                mapInfo.Points[pn.Y][pn.X] = PATH
                node = node.Parent
        }
}

func AddSurroundNode(mapInfo *MapInfo, current *ANode) {
        x := current.P.X
        y := current.P.Y

        AddSurroundNode0(mapInfo, current, x - 1, y, DIRECT_VALUE)
        AddSurroundNode0(mapInfo, current, x, y - 1, DIRECT_VALUE)
        AddSurroundNode0(mapInfo, current, x + 1, y, DIRECT_VALUE)
        AddSurroundNode0(mapInfo, current, x, y + 1, DIRECT_VALUE)

        AddSurroundNode0(mapInfo, current, x - 1, y - 1, OBLIQUE_VALUE)
        AddSurroundNode0(mapInfo, current, x + 1, y - 1, OBLIQUE_VALUE)
        AddSurroundNode0(mapInfo, current, x - 1, y + 1, OBLIQUE_VALUE)
        AddSurroundNode0(mapInfo, current, x + 1, y + 1, OBLIQUE_VALUE)
}

func AddSurroundNode0(mapInfo *MapInfo, current *ANode, x, y, value int) {
        if CanAddToOpen(mapInfo, x, y) {
                end := mapInfo.End

                pn := &Point{x, y}
                g := current.G + value

                child := FindNodeInOpen(pn)
                if child == nil {
                        h := CalcH(end.P, pn)
                        if IsEnd(end.P, pn) {
                                child = end
                                child.Parent = current
                                child.G = g
                                child.H = h
                        } else {
                                child = &ANode{P: pn, Parent: current, G: g, H: h }
                        }
                        openNodes.Push(child)
                } else if child.G > g {
                        child.G = g
                        child.Parent = current
                        openNodes.Push(child)
                }
        }
}

// find node from open list
func FindNodeInOpen(point *Point) *ANode {
        if point == nil || openNodes.Len() <= 0 {
                return  nil
        }

        for _, node := range openNodes {
                if node.P.X == point.X && node.P.Y == point.Y {
                        return node;
                }
        }

        return nil
}

func CalcH(end *Point, point *Point) int {
        x := end.X - point.X
        y := end.Y - point.Y

        return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func IsEnd(end *Point, point *Point) bool {
        return point != nil && end.X == point.X && end.Y == point.Y
}

func CanAddToOpen(mapInfo *MapInfo, x, y int) bool {
        if x < 0 || x >= mapInfo.Width || y < 0 || y >= mapInfo.Height {
                return  false
        }

        if mapInfo.Points[y][x] == BARRIER {
                return false
        }

        if IsInClose(x, y) {
                return false
        }

        return true
}


func IsInClose(x ,y int) bool {
        key := strconv.Itoa(x) + "-" + strconv.Itoa(y)

        _, ok := closeNodes[key]
        return ok
}
