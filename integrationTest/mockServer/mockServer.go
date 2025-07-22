package mockServer

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"tester/domain"
)

type MockServer struct {
	server *httptest.Server
}

func NewMockServer() *MockServer {
	return &MockServer{}
}

func (ms *MockServer) Init() *MockServer {
	ms.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/posts" {
			ms.handlePosts(w)
		} else {
			http.NotFound(w, r)
		}
	}))
	return ms
}

func (ms *MockServer) handlePosts(w http.ResponseWriter) {
	posts := domain.Posts{
		{
			UserId: 1,
			Title:  "Test Title",
			Body:   "Test Body",
		},
	}

	bytes, _ := json.Marshal(posts)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func (ms *MockServer) GetURL() string {
	if ms.server == nil {
		log.Fatalln("Mock server not initialized")
	}
	return ms.server.URL
}

func (ms *MockServer) Close() {
	if ms.server != nil {
		ms.server.Close()
	} else {
		log.Println("Mock server is not initialized, nothing to close")
	}
}
