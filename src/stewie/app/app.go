package main

import (
  "mango" // Point this to mango
)

func Hello(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
  env.Logger().Println("Got a", env.Request().Method, "request")
  return 200, mango.Headers{}, mango.Body("Hello World!")
}

func main() {
  stack := new(mango.Stack)
  stack.Address = ":3000"
  stack.Run(Hello)
}
