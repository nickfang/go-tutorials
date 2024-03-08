package main

import "fmt"

type Item struct {
	Value string
	Next *Item
	Prev *Item
}

type LinkedList struct {
	Head *Item
	Tail *Item
}

func (i *Item) NewItem(value string) *Item {
	i.Value = value
	i.Next = nil
	i.Prev = nil
	return i
}

func (i *Item) InsertItem(newItem *Item) {
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

func showListItem(item *Item) {
	if item == nil { return }
	fmt.Println(item.Value)
	showListItem(item.Next)
}




func (ll *LinkedList) CheckHeadAndTail() error{
	if (ll.Head == nil && ll.Tail != nil) || (ll.Head != nil && ll.Tail == nil) {
		error := fmt.Errorf("the head (%p) and tail (%p) must both have a value or both be nil", ll.Head, ll.Tail)
		return error
	}
	return nil
}

func (ll *LinkedList) AddToBeginning(newItem *Item) {
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

func (ll *LinkedList) AddToEnd(newItem * Item) {
	error := ll.CheckHeadAndTail()
	if (error != nil) {
		return
	}
	if ll.Tail == nil {
		ll.Head = newItem
		ll.Tail = newItem
	}
	newItem.Prev = ll.Tail
	ll.Tail.Next = newItem
	ll.Tail = newItem
}

func (ll *LinkedList) GetLength() int {
	length := 0
	if ll.Head == nil {
		return length
	}
	index := ll.Head
	for index != nil {
		length = length + 1
		index = index.Next
	}
	return length
}

func (ll *LinkedList) GetIndexValue(index int) string{

}

func (ll *LinkedList) GetValueIndex(value string) int {

}



func (ll *LinkedList) showList() {
	item := ll.Head
	for item != nil {
		fmt.Println(item.Value)
		item = item.Next
	}
}

func main() {
	list := LinkedList{}
	item1 := Item{Value: "1"}
	item2 := Item{Value: "2"}
	item3 := Item{Value: "3"}
	list.AddToBeginning(&item2)
	list.AddToBeginning(&item1)
	item2.InsertItem(&item3)

	// showListItem(list.Head)
	// showListItem(&item1)
	// showListItem(&item2)
	list.showList()
	fmt.Printf("length %d\n", list.GetLength())
}