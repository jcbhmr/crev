package handler
import (
  "net/http",
  "jcbhmr.me/crev"
)
func Handler(w http.ResponseWriter, r *http.Request) {
  crev.ServeHTTP(w, r)
}
