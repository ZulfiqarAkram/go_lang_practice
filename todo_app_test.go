package main

import "testing"

func TestAddItem(t *testing.T) {
	var todoItems []string
	todoItemsAfterAddedItem := addItem(todoItems, "") //should return empty list
	if len(todoItemsAfterAddedItem) != len(todoItems) {
		t.Errorf("FAILED: Empty text is added into todo list")
	} else {
		t.Logf("PASS: Empty text not added in the list.")
	}

	var todoItems2 []string
	todoItemsAfterAddedItem2 := addItem(todoItems2, "Tomorrow, I'll go to market and buy some fruits.") //should return list with one item
	if len(todoItemsAfterAddedItem2) == len(todoItems2) {
		t.Errorf("FAILED: Valid Item is not added into todo list")
	} else {
		t.Logf("PASS: valid item is added into list.")
	}

}

func TestRemoveItem(t *testing.T) {
	var todoItems = []string{"Hello world", "There", "Come here", "Test"}

	//should return same list
	listAfterRemove := removeItem(todoItems, -1)
	if len(listAfterRemove) != len(todoItems) {
		t.Errorf("FAILED: negative index number not allowed.")
	} else {
		t.Logf("PASS: Negative index is not removed item from list.")
	}

	//should not removed item from list
	var todoItems2 = []string{"Hello world", "There", "Come here", "Test"}
	result := removeItem(todoItems2, 7)
	if len(result) != len(todoItems2) {
		t.Errorf("FAILED: Index no is out of range.")
	} else {
		t.Logf("PASS: item is not removed against out of range index.")
	}

	//should return list where len=3
	var todoItems3 = []string{"Hello world", "There", "Come here", "Test"}
	result3 := removeItem(todoItems3, 2)
	if len(result3) == len(todoItems3) {
		t.Errorf("FAILED: item is not removed against valid index.")
	} else {
		t.Logf("PASS: item has been removed against valid index.")
	}

}
