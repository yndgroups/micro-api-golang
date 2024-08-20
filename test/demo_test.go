package test

import (
	"fmt"
	"math"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	fmt.Println("Hello World")
}

func TestYouAge(t *testing.T) {
	fmt.Println("请问您的年龄是：", getAge())
}

func getAge() float64 {
	return math.Round(100)
}
