package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(w http.ResponseWriter, r *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(w http.ResponseWriter, r *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(w http.ResponseWriter, r *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(w http.ResponseWriter, r *http.Request) error {
	return os.ErrPermission
}

func errUnknown(w http.ResponseWriter, r *http.Request) error {
	return errors.New("unknown error")
}

func noError(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "no error")
	return nil
}

// 表格驱动测试
var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com",
			nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)

		verifyResponse(response, tt.code, tt.message, t)
	}
}

func verifyResponse(response *http.Response, expectedCode int, expectMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expectedCode || body != expectMsg {
		t.Errorf("expect (%d, %s); got (%d, %s)", expectedCode, expectMsg, response.StatusCode, body)
	}
}
