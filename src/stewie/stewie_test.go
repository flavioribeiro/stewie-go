package stewie

import (
  "net/http"
	"testing"
	"github.com/bmizerany/assert"
)

func test_hello_world(t *testing.T) {
  resp, _ := http.Get("http://localhost:3000/hello")
  assert.Equal(t, "hello", resp)
}
