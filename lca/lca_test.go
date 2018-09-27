package lca

import "testing"

func Test_Hello(t *testing.T) {
  var v string
  v = Hello()
  if v != "Hello" {
    t.Error("Expected \"Hello\", got ", v)
  }
}
