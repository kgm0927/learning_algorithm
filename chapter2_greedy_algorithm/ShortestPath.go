package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Item struct {
	node     int
	distance int
	index    int
}

// 우선순위 큐 (최소 힙) 구조체
type MinHeap []*Item

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}
func (h *MinHeap) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*h)
	*h = append(*h, item)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
func dijkstra(graph [][][2]int, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[start] = 0

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, &Item{node: start, distance: 0})

	for minHeap.Len() > 0 {
		item := heap.Pop(minHeap).(*Item)
		u := item.node

		if item.distance > dist[u] {
			continue
		}

		// 노드 u의 인접 노드들에 대해서 거리 갱신
		for _, edge := range graph[u] {
			v := edge[0]      // 인접 노드 v
			weight := edge[1] // u에서 v로 가는 간선의 가중치

			// 출발점에서 v까지의 새로운 거리
			newDist := dist[u] + weight
			if newDist < dist[v] {
				dist[v] = newDist // v까지의 거리가 더 짧으면 갱신
				// 새로운 거리가 갱신된 v를 힙에 추가
				heap.Push(minHeap, &Item{node: v, distance: dist[v]})

			}

		}

	}
	return dist
}
func main() {
	// 그래프 정의: graph[u]=[(v,weight),(v2,weight2), ...]
	// u-> v로 가는 간선의 가중치
	graph := [][][2]int{
		{{1, 10}, {2, 5}},
		{{3, 1}},
		{{3, 2}},
		{},
	}

	start := 0
	distances := dijkstra(graph, start)

	fmt.Println("출발점", start, "에서 각 노드까지 최단 거리:")
	for i, d := range distances {
		fmt.Printf("노드 %d:%d\n", i, d)
	}
}
