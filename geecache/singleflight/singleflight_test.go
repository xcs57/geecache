package singleflight

import (
	"fmt"
	"testing"
)

func TestGroup_Do(t *testing.T) {
	var myFunc func() (interface{}, error)
	myFunc = func() (key interface{}, err error) {
		fmt.Printf("call func\n")
		return "a", nil
	}
	g := &Group{}
	for i := 0; i < 100; i++ {
		go g.Do("122", myFunc)
	}
}
