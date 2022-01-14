package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0") // 监听一个未被占用的端口，并返回 Listener。
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)
	go http.Serve(ln, nil) // 调用 http.Serve(ln, nil) 启动 http 服务

	// 使用 http.Get 发起一个 Get 请求，检查返回值是否正确
	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handleError(t, err)

	// 尽量不对 http 和 net 库使用 mock，这样可以覆盖较为真实的场景。
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "hello world" {
		t.Fatal("expected hello world, but got", string(body))
	}
}

func TestConn1(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example/com/foo", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)

	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
}