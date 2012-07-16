package stewie

import (
  "net/http"
  "io/ioutil"
	"testing"
	"github.com/bmizerany/assert"
)

func Test_hello_world(t *testing.T) {
  resp, _ := http.Get("http://localhost:3000/hello")
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)

  assert.Equal(t, 200, resp.StatusCode)
  assert.Equal(t, "hello", string(body))
}
