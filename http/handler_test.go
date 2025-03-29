
// handler_test.go
package main

import (
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/hello?name=Alice", nil)
    rr := httptest.NewRecorder()
    HelloHandler(rr, req)

    expected := "Hello, Alice!"
    if rr.Body.String() != expected {
        t.Errorf("Handler returned: %s; want %s", rr.Body.String(), expected)
    }
}

func TestHelloHandler_NoName(t *testing.T) {
    req := httptest.NewRequest("GET", "/hello", nil)
    rr := httptest.NewRecorder()
    HelloHandler(rr, req)

    expected := "Hello, Guest!"
    if rr.Body.String() != expected {
        t.Errorf("Handler returned: %s; want %s", rr.Body.String(), expected)
    }
}
