package collection

import (
	"fmt"
	"strconv"
	"testing"
)

func TestList(t *testing.T) {
	list := new(List[string])
	list.Add("1")
	list.Add("3")
	list.Add("2")
	fmt.Println(list.Values())
	list.Sort(func(a string, b string) int {
		n1, _ := strconv.Atoi(a)
		n2, _ := strconv.Atoi(b)
		return n1 - n2
	})
	fmt.Println(list.Values())
}
