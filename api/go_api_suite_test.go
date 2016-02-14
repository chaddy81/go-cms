package main

import (
  "github.com/gin-gonic/gin"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "io"
  "net/http"
  "net/http/httptest"
  "testing"
)

var (
  response *httptest.ResponseRecorder
)

func TestGoApi(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "GoApi Suite")
}

func Request(method string, route string, handler gin.HandlerFunc) {
  m := gin.Default()
  m.GET(route, handler)
  request, _ := http.NewRequest(method, route, nil)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}

func PostRequest(method string, route string, handler gin.HandlerFunc, body io.Reader) {
  m := gin.Default()
  m.POST(route, handler)
  request, _ := http.NewRequest(method, route, body)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
}
