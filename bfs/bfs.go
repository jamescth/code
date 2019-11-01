package main

// https://cybernetist.com/2019/03/09/breadth-first-search-using-go-standard-library/
import (
	"container/list"
	"fmt"
)

type FIFO struct {
	queue []interface{}
}

func New() *FIFO {
	return &FIFO{
		queue: make([]interface{}, 0),
	}
}

// Push pushed node to the back of the queue
func (f *FIFO) Push(node interface{}) {
	f.queue = append(f.queue, node)
}

// Front takes a value from the front of the queue and returns it
func (f *FIFO) Front() interface{} {
	if len(f.queue) == 0 {
		return nil
	}

	node := f.queue[0]
	f.queue[0] = nil
	f.queue = f.queue[1:]

	return node
}

type node struct {
	Id      string
	friends map[string]*node
}

// Nodes returns a list of all graph nodes
func Nodes(n *node) []*node {
	// track the visited nodes
	visited := make(map[string]*node)

	// queue of the nodes to visit
	queue := list.New()
	queue.PushBack(n)

	// add the root node to the map of the visited nodes
	visited[n.Id] = n

	for qnode := queue.Front(); qnode != nil; qnode = qnode.Next() {
		queue.Remove(qnode)

		// iterate through all of its friends
		// mark the visited nodes; enqueue the non-visited
		for id, node := range qnode.Value.(*node).friends {
			if _, ok := visited[id]; !ok {
				visited[id] = node
				queue.PushBack(node)
			}
		}
	}

	nodes := make([]*node, 0)
	// collect all the nodes into slice
	for _, node := range visited {
		nodes = append(nodes, node)
	}

	return nodes
}

func main() {
	vals := []int{1, 2, 3, 4}
	fifo := New()

	for _, val := range vals {
		fifo.Push(val)
	}
	for {
		front := fifo.Front()
		if front == nil {
			break
		}
		fmt.Println(front)
	}

	// bfs
	node1 := &node{Id: "node1", friends: make(map[string]*node)}
	node2 := &node{Id: "node2", friends: make(map[string]*node)}
	node3 := &node{Id: "node3", friends: make(map[string]*node)}

	// node2 is directly connected to node1
	node2.friends[node1.Id] = node1
	// node 3 is directly connected to node2
	node3.friends[node2.Id] = node2
	// node 1 is directly connected to node3
	node1.friends[node3.Id] = node3

	// root node; this graph is actually not a tree, so it has no root node
	root := &node{Id: "root", friends: make(map[string]*node)}
	n := make(map[string]*node)

	n[root.Id] = root
	n[root.Id].friends[node1.Id] = node1
	n[root.Id].friends[node2.Id] = node2
	n[root.Id].friends[node3.Id] = node3

	nodes := Nodes(root)
	for _, node := range nodes {
		fmt.Printf("%+v\n", node.Id)
	}
}
