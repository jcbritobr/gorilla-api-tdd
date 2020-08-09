package model

// ToDoItem is a card item
type ToDoItem struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// toDoBuffer is the fake database in memory
var toDoBuffer []ToDoItem = []ToDoItem{}

// NewToDoItem creates a todo item
func NewToDoItem(id int, description string) *ToDoItem {
	return &ToDoItem{ID: id, Description: description}
}

// Insert inserts an item in buffer
func Insert(t *ToDoItem) {
	toDoBuffer = append(toDoBuffer, *t)
}

// FindToDoItem returns an item
func FindToDoItem(id int) *ToDoItem {
	for _, item := range toDoBuffer {
		if item.ID == id {
			return &item
		}
	}
	return &ToDoItem{}
}

// FindAll returns all items from buffer
func FindAll() []ToDoItem {
	return toDoBuffer
}

// Delete removes an item from buffer
func Delete(id int) bool {
	for i, item := range toDoBuffer {
		if item.ID == id {
			toDoBuffer = append(toDoBuffer[:i], toDoBuffer[i+1:]...)
			return true
		}
	}
	return false
}
