package main

import (
	"reflect"

	"github.com/k0kubun/pp"
)

type BaseSt struct {
	FieldA string
}

type TypeA struct {
	BaseSt
	//FieldB int
	//FieldC TypeB
	FieldE MyInter
}

type TypeB struct {
	//FieldD string
}

func (TypeB) Run() string {
	return "IM TYPE B"
}

type TypeC struct {
	FieldD string
}

func (TypeC) Run() string {
	return "IM TYPE C"
}

type MyInter interface {
	Run() string
}

func main() {
	b := TypeB{}
	a := &TypeA{
		FieldE: &b,
	}

	c := &TypeA{}

	//pp.Println(a.FieldE.Run())
	elemVal := reflect.ValueOf(a).Elem()
	elemType := reflect.TypeOf(a).Elem()

	numOfFields := elemVal.NumField()
	for i := 0; i < numOfFields; i++ {
		fieldVal := elemVal.Field(i)
		fieldType := elemType.Field(i)

		if fieldType.Name == "BaseSt" {
			continue
		}

		//pp.Println(field.Elem().Interface())

		cElem := reflect.ValueOf(c).Elem()
		cField := cElem.FieldByName(fieldType.Name)
		cField.Set(fieldVal)

	}



}
