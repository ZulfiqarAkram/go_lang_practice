package main

import (
	"bufio"
	"fmt"
	"os"
)

// Add an item to list/array
// Display all items
//remove item from _array

func addItem(todoItems []string, text string) []string {
	todoItems = append(todoItems, text)
	println("Item Added.\n")
	return todoItems
}

func removeItem(todoItems []string, indexNo int) []string {
	todoItems = append(todoItems[:indexNo], todoItems[(indexNo+1):]...) // slice left and right array then concatenate the array
	println("Item Removed.\n")
	return todoItems
}

func displayItems(todoItems []string) {
	println("------------------------")
	for index, item := range todoItems {
		println(index, item)
	}
	println("------------------------")
}

func main() {
	var todoItems []string
	scanner := bufio.NewScanner(os.Stdin)

	var askAgain = true
	for askAgain == true {
		fmt.Println("Please select\n1-Add new item\n2-Remove item\n3-Display items")
		scanner.Scan()
		var input string
		input = scanner.Text()

		if input == "1" {
			fmt.Print("Please enter text: ")
			scanner.Scan()
			var text string
			text = scanner.Text()
			todoItems = addItem(todoItems, text)

		} else if input == "2" {
			fmt.Println("Please select index no of item that will be removed.")
			var indexItemToRemove int
			fmt.Scanln(&indexItemToRemove)
			todoItems = removeItem(todoItems, indexItemToRemove)

		} else if input == "3" {
			displayItems(todoItems)

		} else {
			println("App has been stopped.")
			askAgain = false
		}
	}

}
