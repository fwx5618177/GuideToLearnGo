package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name string
	LifeEx int
}

func (b *Bird) Fly() {
	fmt.Println("Flying...")
}

func main() {
	sparrow := &Bird{"Sparrow", 3}
	
	s := reflect.Valueof(sparrow).Elem()

	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		fmt.Println("%d-%s-%s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

