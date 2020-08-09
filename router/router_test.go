package router

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/jcbritobr/gorillaapi/model"

	"github.com/gorilla/mux"
)

// setup adds some data to test
func setup() {
	fmt.Println("Setting ...")
	model.Insert(model.NewToDoItem(1, "A"))
	model.Insert(model.NewToDoItem(2, "B"))
}

// tearDown cleans all data from tests
func tearDown() {
	model.Delete(1)
	model.Delete(2)
	data := model.FindAll()
	fmt.Println("Teared down ...", data)
}

func TestFindItem(t *testing.T) {
	setup()
	tt := []struct {
		name   string
		path   string
		result string
	}{
		{"FindItem test response body A", "/api/items/1", `{"id":1,"description":"A"}`},
		{"FindItem test response body B", "/api/items/2", `{"id":2,"description":"B"}`},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			SetupRouter(router)
			router.ServeHTTP(rr, req)

			if rr.Body.String() != tc.result {
				t.Errorf("Unexpected body: got %v want %v status:%v\n", rr.Body.String(), tc.result, rr.Code)
				return
			}
		})
	}
	tearDown()
}

func TestCheckBadRequest(t *testing.T) {
	t.Run("checkBadRequest A", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/test/br", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			err := fmt.Errorf("testing bad request")
			checkBadRequest(err, w)
		})

		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Invalid http response got %v want %v", rr.Code, http.StatusBadRequest)
		}
	})
}

func TestCreate(t *testing.T) {
	type args struct {
		id          int
		description string
	}
	tt := []struct {
		name string
		args args
		want string
	}{
		{"create A", args{id: 1, description: "A"}, `{"id":1,"description":"A"}`},
		{"create B", args{id: 2, description: "B"}, `{"id":2,"description":"B"}`},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			formData := url.Values{}
			formData.Set("id", strconv.Itoa(tc.args.id))
			formData.Set("description", tc.args.description)
			req, err := http.NewRequest(http.MethodPost, "/api/items", bytes.NewBufferString(formData.Encode()))
			if err != nil {
				t.Errorf("Can't create a request object")
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

			rr := httptest.NewRecorder()
			m := mux.NewRouter()
			SetupRouter(m)
			m.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK && rr.Body.String() != tc.want {
				t.Errorf("Wrong response body. Got %v want %v", rr.Body.String(), tc.want)
			}
		})
	}
	tearDown()
}

func TestRouteFindAll(t *testing.T) {
	setup()
	expected := `[{"id":1,"description":"A"},{"id":2,"description":"B"}]`
	t.Run("route findall A", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/items", nil)
		if err != nil {
			t.Errorf("Can't create a request object")
		}
		rr := httptest.NewRecorder()
		m := mux.NewRouter()
		SetupRouter(m)
		m.ServeHTTP(rr, req)

		str := rr.Body.String()

		if rr.Code != http.StatusOK || str != expected {
			t.Errorf("Unexpected response body. Got %v want %v", str, expected)
		}
	})
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
		{"delete A", args{id: 1}, true},
		{"delete B", args{id: 2}, true},
		{"delete C", args{id: 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/items/%d", tt.args.id), nil)
			if err != nil {
				t.Errorf("Can't create a new request")
			}
			rr := httptest.NewRecorder()
			m := mux.NewRouter()
			SetupRouter(m)
			m.ServeHTTP(rr, req)

			result, err := strconv.ParseBool(rr.Body.String())
			if err != nil {
				t.Errorf("Unexpected response body")
			}

			if rr.Code != http.StatusOK && result != tt.want {
				t.Errorf("Test failed with got %v want %v", result, tt.want)
			}
		})
	}
	tearDown()
}
