package integrationTest

import (
	"testing"

	"github.com/gavv/httpexpect"
)

type testClient struct {
	*httpexpect.Expect
}

func Client(t *testing.T) *testClient {
	t.Helper()
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  TestServer.GetURL(),
		Reporter: httpexpect.NewAssertReporter(t),
	})
	return &testClient{e}
}
