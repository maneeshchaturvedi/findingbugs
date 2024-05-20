package main

import "fmt"

// Entry represents a node in the linked list
type Entry struct {
    key   int
    data  int
    name  string
    next  *Entry
}

// insertLinkedList inserts a new element into the linked list
func insertLinkedList(currentHead *Entry, newElement *Entry) *Entry {
    var currentElement *Entry

    // If the list is empty, then just return newElement
    if currentHead == nil {
        newElement.next = nil
        return newElement
    }

    // If newElement should be the first on the list,
    // attach the current list to newElement and return
    // newElement as the new head of the list.
    if newElement.key < currentHead.key {
        newElement.next = currentHead
        return newElement
    }

    // Now walk the list, comparing key values. Exit the
    // while loop when currentElement points to the
    // value after which we should insert newElement.
    currentElement = currentHead
    for currentElement.next != nil {
        if newElement.key < currentElement.next.key {
            break
        }
        currentElement = currentElement.next
    }

    // Insert newElement after currentElement and return
    newElement.next = currentElement.next
    currentElement.next = newElement
    return currentHead
}
