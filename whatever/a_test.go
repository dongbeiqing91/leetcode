package whatever

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	m := map[string]interface{}{}
	m["a"] = struct{}{}
}

func TestB(t *testing.T) {
	var a int
	a = 3 / 2
	print(a)
}

func TestC(t *testing.T) {
	m1 := []string{"a", "b"}
	m1 = m1[0 : len(m1)-1]
	fmt.Print(m1)
	m2 := m1[0:0]
	fmt.Print(m2)
}
