package core

import (
        //"container/heap"
)

type ANode struct {
        P *Point

        F, G, H, Step int

        Parent *ANode
}

type PriorityQueue []*ANode

func (pq PriorityQueue) Len() int {
        return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
        return pq[i].G + pq[i].H < pq[j].G + pq[j].H
}

func (pq PriorityQueue) Swap(i, j int) {
        pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
        item := x.(*ANode)
        *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
        old := *pq
        n := len(old)
        item := old[n-1]
        *pq = old[0 : n-1]
        return item
}

func (pq *PriorityQueue) update(item *ANode) {
}


