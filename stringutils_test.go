package main

import "github.com/voidjuice/stringutils"
import "testing"


func TestStringUtil(t *testing.T) {
test_string := "sTrInG"
if test_string.Upper != "STRING" {
t.Error("expected test_string to be uppercase")
}
if test_string.Lower != "string" {
t.Error("expected test_string to be lowercase")
}
