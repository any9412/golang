package myapp

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
)

type fooHandler struct {}

type User struct {
	FirstName string	`json:"first_name"`
	LastName string		`json:"last_name"`
	Email string		`json:"email"`
	CreatedAt time.Time	`json:"created_at"`
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Reuest: ", err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func barHanlder(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // "http://localhost:3000/bar?name=name" 이 형식의 url을 통해 name값을 받아온다
	if name == "" {
		name = "World"
	}
	fmt.Fprint(w, "Hello " + name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux() // mux instance(router)를 만들어서 사용
	// Handler를 function 형태로 직접 등록. "/" 경로 = index 경로, 첫번째 page 경로에 대한 request가 왔을때 어떤 function을 할지
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", barHanlder)

	// Handler instance를 등록. ServeHTTP interface를 구현한 뒤에 이를 통해 handler instance를 생성
	mux.Handle("/foo", &fooHandler{})
	return mux
}