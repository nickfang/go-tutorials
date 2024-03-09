package main

import "fmt"

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
	list := LinkedList[string]{}
	item1 := Item[string]{Value: "1"}
	item2 := Item[string]{Value: "2"}
	item3 := Item[string]{Value: "3"}
	emptyValue := list.GetValueOfIndex(1)
	fmt.Printf("%s", emptyValue)
	list.AddToBeginning(&item2)
	list.AddToBeginning(&item1)
	item2.InsertItem(&item3)

	// showListItem(list.Head)
	// showListItem(&item1)
	// showListItem(&item2)
	list.showList()
	fmt.Printf("length %d\n", list.GetLength())
	value4 := list.GetValueOfIndex(4)
	fmt.Printf("%s\n", value4)
	value1 := list.GetValueOfIndex(1)
	fmt.Printf("%s\n", value1)
	index1 := list.GetIndexOfValue("1")
	fmt.Printf("Index of '1': %d\n", index1)
	index2 := list.GetIndexOfValue("2")
	fmt.Printf("Index of '2': %d\n", index2)
	index3 := list.GetIndexOfValue("3")
	fmt.Printf("Index of '3': %d\n", index3)
}