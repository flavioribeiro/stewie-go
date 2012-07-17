package main

import (
  "mango" // Point this to mango
)

func Index(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
  env.Logger().Println("Got a", env.Request().Method, "request")
  return 200, mango.Headers{}, mango.Body("Index here. untested :P")
}

func GetToken(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
  return 200, mango.Headers{}, mango.Body("1234567890")
}

func main() {
  routes := make(map[string]mango.App)
  routes["/"] = new(mango.Stack).Compile(Index)
  routes["/get_token"] = new(mango.Stack).Compile(GetToken)

  stack := new(mango.Stack)
  stack.Middleware(mango.ShowErrors("<html><body>{Error|html}</body></html>"), mango.Routing(routes))
  stack.Address = ":3001"
  stack.Run(Index)
}
