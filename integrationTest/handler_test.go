package integrationTest

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func Test_ShouldCreateData(t *testing.T) {
	Client(t).
		POST("/data").
		Expect().
		Status(http.StatusCreated).
		JSON().Equal(JSONResponseFile(t, "addData.json"))
}

func JSONResponseFile(t *testing.T, fileName string) (expected interface{}) {
	t.Helper()
	jsonFile, _ := os.Open(filepath.Join(".", "response", fileName))
	bytes, _ := io.ReadAll(jsonFile)
	_ = json.Unmarshal(bytes, &expected)
	return
}
