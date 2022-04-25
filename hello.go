package main

import (
	"fmt"
	"net/http"
)

// 인스턴스 생성
type fooHandler struct{}

// fooHandler 인스턴스의 메서드를 만듬
// 인스턴스의 ServeHTTP 인터페이스를 구현
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Foo")
}

func main() {
	// HandleFunc는 handler를 등록하는 것이다.
	// HandleFunc의 두번째 인자는 어떤 경로로 들어왔을 때 어떤 것을 실행시킬지에 대한 것.
	// 두번째 인자에 있는 함수의 형태는 정해져 있는 것
	// response를 write할 수 있는 ResponseWriter, 사용자 요청한 Request 정보를 가지고 있는 Request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Fprint는 w(writer)에 해당하는 것의 문자열을 출력하라.
		fmt.Fprint(w, "hello world")
	})

	// HandleFunc: handler를 function 형태로 직접 등록할 때 쓰는 함수
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello bar~")
	})

	// Handle: 인스턴스 형태로 등록할 때 쓰는 함수
	http.Handle("/foo", &fooHandler{})

	// web server를 3000번 포트에서 구동
	http.ListenAndServe(":3000", nil)
}
