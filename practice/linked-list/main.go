package main

import (
	"fmt"
	// "time"
)

type Item[T string|int] struct {
	Value T
	Next *Item[T]
	Prev *Item[T]
}

type LinkedList[T string|int] struct {
	Head *Item[T]
	Tail *Item[T]
}

func (i *Item[T]) NewItem(value T) *Item[T] {
	i.Value = value
	i.Next = nil
	i.Prev = nil
	return i
}

func (i *Item[T]) InsertItem(newItem *Item[T]) {
	if (i == nil || newItem == nil) {
		return
	}
	newItem.Next = i.Next
	newItem.Prev = i
	if i.Next != nil {
		i.Next.Prev = newItem
	}
	i.Next = newItem
}

func (i *Item[T]) showListItem() {
	if i == nil { return }
	fmt.Println(i.Value)
	i.Next.showListItem()
}


func (ll *LinkedList[T]) CheckHeadAndTail() error{
	if (ll.Head == nil && ll.Tail != nil) || (ll.Head != nil && ll.Tail == nil) {
		error := fmt.Errorf("the head (%p) and tail (%p) must both have a value or both be nil", ll.Head, ll.Tail)
		return error
	}
	return nil
}

func (ll *LinkedList[T]) IsEmpty() bool{
	return ll.Head == nil && ll.Tail == nil
}


func (ll *LinkedList[T]) AddToBeginning(newItem *Item[T]) {
	error := ll.CheckHeadAndTail()
	if (error != nil) {
		return
	}
	newItem.Next = ll.Head
	if ll.Head != nil {
		ll.Head.Prev = newItem
	} else {
		ll.Tail = newItem
	}
	ll.Head = newItem
}

func (ll *LinkedList[T]) AddToEnd(newItem *Item[T]) {
	error := ll.CheckHeadAndTail()
	if error != nil {
		return
	}
	if ll.IsEmpty() {
		ll.Head = newItem
		ll.Tail = newItem
		return
	}
	newItem.Prev = ll.Tail
	ll.Tail.Next = newItem
	ll.Tail = newItem
}

func (ll *LinkedList[T]) GetLength() int {
	error := ll.CheckHeadAndTail()
	if error != nil {
		return 0
	}
	length := 0
	if ll.IsEmpty() {
		return length
	}
	index := ll.Head
	for index != nil {
		length = length + 1
		index = index.Next
	}
	return length
}

func (ll *LinkedList[T]) GetValueOfIndex(index int) T {
	var defaultValue T
	error := ll.CheckHeadAndTail()
	if error != nil {
		return defaultValue
	}
	if ll.IsEmpty() {
		fmt.Printf("Empty list.\n")
		return defaultValue
	}
	i := 0
	item := ll.Head
	for item != nil {
		if i == index {
			return item.Value
		}
		i = i+1
		item = item.Next
	}
	fmt.Printf("Index is outside of the list.")
	return defaultValue
}

func (ll *LinkedList[T]) GetIndexOfValue(value T) int {
	var defaultValue int
	error := ll.CheckHeadAndTail()
	if error != nil {
		return defaultValue
	}
	if ll.IsEmpty() {
		fmt.Printf("Empty list.\n")
		return defaultValue
	}
	i := 0
	item := ll.Head
	for item.Value != value {
		i = i + 1
		item = item.Next
	}
	return i
}

func (ll *LinkedList[T]) showList() {
	item := ll.Head
	for item != nil {
		fmt.Println(item.Value)
		item = item.Next
	}
}

func main() {
	list := LinkedList[int]{}
	for i := 0; i < 100000; i++ {
		item := Item[int]{Value: i}
		list.AddToEnd(&item)
	}

	fmt.Printf("length %d\n", list.GetLength())
	resultChan := make(chan string)
	go func() {
		result := list.GetValueOfIndex(99998)
		message := fmt.Sprintf("Index 99998: %d\n", result)
		// time.Sleep(1 * time.Second)
		resultChan <- message
	}()
	go func() {
		result := list.GetValueOfIndex(1)
		message := fmt.Sprintf("Index 1: %d\n", result)
		// time.Sleep(2 * time.Second)
		resultChan <- message
	}()
	go func() {
		result := list.GetValueOfIndex(50000)
		message := fmt.Sprintf("Index 50000: %d\n", result)
		// time.Sleep(3 * time.Second)
		resultChan <- message
	}()
	fmt.Printf(<- resultChan)
	fmt.Printf(<- resultChan)
	fmt.Printf(<- resultChan)
}