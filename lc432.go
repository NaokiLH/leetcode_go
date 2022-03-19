package main

// type AllOne struct {
// 	lens []*Lnode
// 	mp   map[string]*Node
// 	max  *Lnode
// 	min  *Lnode
// }
// type Node struct {
// 	name string
// 	val  int
// 	prev *Node
// 	next *Node
// }
// type Lnode struct {
// 	prev *Lnode
// 	next *Lnode
// 	node *Node
// }

// func Constructor() AllOne {

// 	lens := make([]*Lnode, 5e4+5)
// 	mp := make(map[string]*Node, 0)

// 	return AllOne{
// 		lens,
// 		mp,
// 		nil,
// 		nil,
// 	}

// }

// func (this *AllOne) Inc(key string) {
// 	if this.mp[key] == nil {
// 		node := Node{
// 			key,
// 			1,
// 			nil,
// 			nil,
// 		}
// 		if this.lens[node.val] == nil {
// 			lnode := Lnode{
// 				nil,
// 				nil,
// 				nil,
// 			}
// 			h := this.min
// 			lnode.next = h
// 			h.prev = &lnode
// 			this.min = &lnode
// 			this.lens[node.val] = &lnode
// 		}
// 		lnode := this.lens[node.val]
// 		nnode := lnode.node
// 		lnode.node = &node
// 		node.next = nnode
// 		if nnode != nil {
// 			nnode.prev = &node
// 		}
// 	} else {
// 		node := this.mp[key]

// 		if this.lens[node.val+1] == nil {
// 			lnode := Lnode{
// 				nil,
// 				nil,
// 				nil,
// 			}
// 			h := this.lens[node.val]
// 			n := h.next
// 			h.next = &lnode
// 			lnode.prev = h
// 			lnode.next = n
// 			n.prev = &lnode
// 			this.lens[node.val+1] = &lnode
// 		}

// 		lnode := this.lens[node.val+1]
// 		node.val += 1

// 		nnode := lnode.node
// 		lnode.node = node
// 		node.next = nnode
// 		if nnode != nil {
// 			node.prev = node
// 		}
// 	}
// }

// func (this *AllOne) Dec(key string) {
// 	node := this.mp[key]

// 	if node.val == 1 {
// 		if this.lens[1].node == node {
// 			this.lens[1].node = node.next
// 		}

// 	}
// }

// func (this *AllOne) GetMaxKey() string {
// 	return this.max.node.name
// }

// func (this *AllOne) GetMinKey() string {
// 	return this.min.node.name
// }

// /**
//  * Your AllOne object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.Inc(key);
//  * obj.Dec(key);
//  * param_3 := obj.GetMaxKey();
//  * param_4 := obj.GetMinKey();
//  */
