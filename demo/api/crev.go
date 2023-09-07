package handler
import (
  "net/http",
  "jcbhmr.me/crev"
)
func Handler(res http.ResponseWriter, req *http.Request) {
	crev.ServeHTTP(res, req)
}
