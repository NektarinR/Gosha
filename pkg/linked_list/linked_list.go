package linked_list

import (
	"fmt"
	"io"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head  *Node
	trail *Node
	len   int
}

func CreateList() *LinkedList {
	return &LinkedList{}
}

func (p *LinkedList) Add(v int) {
	if p.head == nil {
		p.head = &Node{
			value: v,
			next:  nil,
		}
		p.trail = p.head
		p.len++
	} else {
		p.trail.next = &Node{
			value: v,
			next:  nil,
		}
		p.trail = p.trail.next
		p.len++
	}
}

func (p *LinkedList) Print(w io.Writer) {
	tmp := p.head
	for tmp != nil {
		fmt.Fprintf(w, "%d ", tmp.value)
		tmp = tmp.next
	}
}

func (p *LinkedList) Delete(v int) {
	if p.head == nil {
		return
	}
	if p.head.value == v {
		p.head = p.head.next
	} else {
		tmp := p.head
		for tmp.next != nil {
			if tmp.next.value == v {
				if tmp.next == p.trail {
					p.trail = tmp
				}
				s := tmp.next
				tmp.next = s.next
				s = nil
			} else {
				tmp = tmp.next
			}
		}
	}
}
