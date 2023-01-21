package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	Print()
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	tail *ListItem
	head *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.head == nil {
		l.head = newItem
		l.tail = newItem
	} else {
		newItem.Next = l.head
		l.head.Prev = newItem
		l.head = newItem
	}
	l.len++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.head == nil {
		l.head = newItem
		l.tail = newItem
	} else {
		currentItem := l.head
		for currentItem.Next != nil {
			currentItem = currentItem.Next
		}
		newItem.Prev = currentItem
		currentItem.Next = newItem
		l.tail = newItem
	}
	l.len++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i.Next != nil {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		i.Next = nil // avoid memory leaks
		i.Prev = nil // avoid memory leaks
		i.Value = nil
		l.len--
	} else {
		i.Prev.Next = nil
		i.Next = nil
		l.tail = i.Prev
		l.len--
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Value == l.Front().Value {
		return
	}

	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}

func (l *list) Print() {
	fmt.Println("Print:", l.head.Value, l.tail.Value)
}
