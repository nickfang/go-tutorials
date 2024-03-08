package main

import "fmt"

type Item struct {
	Value string
	Next *Item
	Prev *Item
}

type LinkedList struct {
	Head *Item
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

func (ll *LinkedList) InsertAtBeginning(newItem *Item) {
	newItem.Next = ll.Head
	if ll.Head != nil {
		ll.Head.Prev = newItem
	}
	ll.Head = newItem
}

func showList(item *Item) {
	if item == nil { return }
	fmt.Println(item.Value)
	showList(item.Next)
}

func (ll *LinkedList) showList() {
	item := ll.Head
	for item != nil {
		fmt.Println(item.Value)
		item = item.Next
	}
}

func showListItem(item *Item) {
	fmt.Println(item.Value, item.Next, item.Prev)
}

func main() {
	list := LinkedList{}
	item1 := Item{Value: "1"}
	item2 := Item{Value: "2"}
	item3 := Item{Value: "3"}
	list.InsertAtBeginning(&item2)
	list.InsertAtBeginning(&item1)
	item2.InsertItem(&item3)

	showListItem(list.Head)
	showListItem(&item1)
	showListItem(&item2)
	list.showList()
}