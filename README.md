# go-tucker

mux란?

  - multiplexer로 입력에 따라 다른 모듈로 분기해준다.

  - 웹 request에 따라 다른 핸들러로 연결해주는 역할

<br/>

mux를 정해주는 것과 http로 바로 연결하는 것에는 어떤차이?

  - http.HandleFunc()은 http 패키지의 함수를 직절 호출하는 것

  ```
    mux := http.NewServeMux()
    mux.HandleFunc() 

    는 mux 인스턴스를 만들고 그 인스턴스의 메서드를 콜한다.
  ```

  - 패키지의 함수를 콜하면 패키지 전역으로 핸들러가 등록되고 mux 인스턴스를 만들어서 메서드를 호출하면 전역이 아닌, mux 인스턴스 안에 등록된다. 

  - 따라서 특정 핸들러를 어떤 상황에서도 호출하고 싶을 경우, http 패키지 전역에 핸들러를 등록하면 되고, 각각의 상황에 따라 다른 핸들러를 호출하고 싶을 경우에는 mux 인스턴스를 여러개 만들어 이용하면 된다.

  - 예를 들어 테스트 환경에서는 특정 핸들러를 호출하고, 일반 상황에서는 일반 핸들러를 호출하려고 하는 경우처럼, 조건에 따라 다른 핸들러를 호출해줘야 할 때 여러 mux 인스턴스를 만들어서 사용할 수 있다.


* HTTP: Hyper Text Transfer Protocol// hyper text를 전달하는 규약

  - Request와 Response를 어떤 형식으로 할 것인지는 http 프로토콜에 정의되어 있다.

  - Server Rendering: 서버에서 html을 전부 그려 client로 전달
  
  - tucker의 go lang 강의 Repository
