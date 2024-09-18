package jsonstring

import (
	"github.com/LucianoPerez314/json/visitor/jsondata"
)

type jsonstringVisitor struct {
}

func New() jsonstringVisitor {
	return jsonstringVisitor{}
}

// Convert json value to string.
func Map(map[string]jsondata.JSONValue) (string, error) {

}

func Slice([]jsondata.JSONValue) (string, error) {

}

func Bool(bool) (string, error) {

}

func Float64(float64) (string, error) {

}

func String(string) (string, error) {

}

func Null() (string, error) {

}
