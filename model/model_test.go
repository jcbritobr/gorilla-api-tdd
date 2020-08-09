package model

import (
	"fmt"
	"reflect"
	"testing"
)

// setup adds some data to tests
func setup() {
	fmt.Println("Setting ...")
	Insert(NewToDoItem(1, "A"))
	Insert(NewToDoItem(2, "B"))
}

// tearDown cleans all data from tests
func tearDown() {
	Delete(1)
	Delete(2)
	fmt.Println("Teared down ...", toDoBuffer)
}

func TestNewToDoItem(t *testing.T) {
	type args struct {
		id          int
		description string
	}
	tests := []struct {
		name string
		args args
		want *ToDoItem
	}{
		{"NewToDoItem A", args{1, "simple test"}, &ToDoItem{ID: 1, Description: "simple test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewToDoItem(tt.args.id, tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewToDoItem() = %v, want %v", got, tt.want)
			}
		})
	}

	tearDown()
}

func TestFindToDoItem(t *testing.T) {
	setup()
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want *ToDoItem
	}{
		{"FindToDoItem A", args{id: 1}, NewToDoItem(1, "A")},
		{"FindToDoItem B", args{id: 0}, &ToDoItem{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindToDoItem(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindToDoItem() = %v, want %v", got, tt.want)
			}
		})
	}
	tearDown()
}

func TestFindAll(t *testing.T) {
	setup()
	tests := []struct {
		name string
		want []ToDoItem
	}{
		{"FindAll A", []ToDoItem{{ID: 1, Description: "A"}, {ID: 2, Description: "B"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
	tearDown()
}

func TestDelete(t *testing.T) {
	setup()
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Delete A", args{id: 1}, true},
		{"Delete B", args{id: 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(tt.args.id); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
	tearDown()
}

func TestInsert(t *testing.T) {
	type args struct {
		t *ToDoItem
	}
	tests := []struct {
		name string
		args args
	}{
		{"Insert A", args{t: NewToDoItem(1, "A")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insert(tt.args.t)
		})
	}
	tearDown()
}
