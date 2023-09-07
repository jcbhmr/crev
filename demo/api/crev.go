package handler
import (
  "net/http"
  crev "jcbhmr.me/crev/pkg"
)
func Handler(w http.ResponseWriter, r *http.Request) {
  crev.ServeHTTP(w, r)
}
