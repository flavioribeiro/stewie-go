package stewie

import (
  "net/http"
  "io/ioutil"
	"testing"
	"github.com/bmizerany/assert"
)

const STEWIE_HOST string = "http://localhost:3001"

func Test_get_token_should_return_token(t *testing.T) {
  resp, _ := http.Get(STEWIE_HOST + "/get_token")
  body, _ := ioutil.ReadAll(resp.Body)

  assert.Equal(t, 200, resp.StatusCode)
  assert.Equal(t, 10, len(string(body)))
}

func Test_get_token_should_return_unique_key(t *testing.T) {
//  keys = new List()
  var keys []string

  for i:=0; i < 100; i++ {
    resp, _ := http.Get(STEWIE_HOST + "/get_token")
    body, _ := ioutil.ReadAll(resp.Body)

    keys = append(keys, string(body))
  }
}
