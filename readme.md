## A gorilla to-do rest api using TDD and Go tests best pratices
This is a simple rest api that implements a todo application. It uses [gorilla-mux](https://github.com/gorilla/mux) to create the routes, and **TDD** and **Table Tests**. **Gorilla-mux** and http standard library from Go are good combination to build microservices. They are small, and produces efficient executables.

* **Gorilla-mux** \
[Gorilla-mux](https://github.com/gorilla/mux) is a library used to implement requests and dispatcher, for matching incoming resquests.
```go
subroute.HandleFunc("/item/{id:[0-9]+}", findItem).Methods("GET")
```


* **Golang table tests pattern** \
This pattern allow us to build tests that are reusable.We may test many scenarios in the same test.
```go
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
```

* **Design**

<p align="center">
    <img src="image/rest-api.png">
</p>