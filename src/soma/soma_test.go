package main

import (
  "testing"
)

func TestSoma(t *testing.T) {
  expected := 10
  actual := Soma(5,5)

  if actual != expected {
    t.Errorf("Soma incorreta %d!",actual)
  }
}