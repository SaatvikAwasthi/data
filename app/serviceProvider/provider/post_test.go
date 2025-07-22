package provider_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"

	"tester/app/serviceProvider/config"
	"tester/app/serviceProvider/provider"
	httpConfig "tester/crosscutting/http/config"
	"tester/crosscutting/util"
	"tester/domain"
)

func TestPostsServiceProvider_Fetch(t *testing.T) {
	// Define test cases
	ctx := context.Background()
	posts := domain.Posts{
		{
			UserId: 1,
			Title:  "Test Title",
			Body:   "Test Body",
		},
	}

	type expectations struct {
		resp domain.Posts
		err  error
	}

	testCases := []struct {
		name       string
		setupMock  http.HandlerFunc
		expected   expectations
		assertions func(expected, got expectations)
	}{
		{
			name: "success fetch posts",
			setupMock: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				bytes, _ := json.Marshal(posts)
				_, _ = w.Write(bytes)
			},
			expected: expectations{
				resp: posts,
				err:  nil,
			},
			assertions: func(expected, got expectations) {
				assert.NilError(t, got.err)
				assert.DeepEqual(t, expected.resp, got.resp)
			},
		},
		{
			name: "error fetch posts - non-200 status code",
			setupMock: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("Internal Server Error"))
			},
			expected: expectations{
				resp: nil,
				err:  util.NewError("failed to fetch data %d", http.StatusInternalServerError),
			},
			assertions: func(expected, got expectations) {
				assert.ErrorContains(t, got.err, expected.err.Error())
				assert.DeepEqual(t, got.resp, expected.resp)
			},
		},
		{
			name: "error fetch posts - empty response",
			setupMock: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("[]")) // Empty JSON array
			},
			expected: expectations{
				resp: nil,
				err:  util.NewError("no data found"),
			},
			assertions: func(expected, got expectations) {
				assert.ErrorContains(t, got.err, expected.err.Error())
				assert.DeepEqual(t, got.resp, expected.resp)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			testServer := httptest.NewServer(tc.setupMock)
			cfg := config.Provider{BaseUrl: testServer.URL}
			httpCfg := httpConfig.NewHTTPConfig().Default()
			p := provider.New(cfg, httpCfg)

			// Act
			result, err := p.Fetch(ctx)

			// Assert
			tc.assertions(tc.expected, expectations{resp: result, err: err})
		})
	}

}
