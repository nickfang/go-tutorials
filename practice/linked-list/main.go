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

func createListItem(value string) Item {
	item := Item{value, nil, nil}
	return item
}

func insertListItem(item *Item, location *Item) {
	if (item == nil || location == nil) {
		return
	}
	temp := location.Next
	location.Next = item
	item.Next = temp
}



func showListItem(item *Item) {
	fmt.Println(item.Value, item.Next, item.Prev)
}

func main() {
	head := createListItem("head")
	item1 := createListItem("1")
	insertListItem(&item1, &head)
	item2 := createListItem("2")
	insertListItem(&item2, &item1)

	showListItem(&head)
	showListItem(&item1)
	showListItem(&item2)
}