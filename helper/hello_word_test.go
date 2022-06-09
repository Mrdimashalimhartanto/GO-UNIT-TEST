package helper

import (
	"fmt"
	"testing"
)

//panic jangan di pake untuk test unit

func TestHelloWorldAssert(t *testing.T) {
	//result := HelloWord("Dimas")
	//assert.Equal(t)
}

func TestHelloWorldDimasHalim(t *testing.T) {
	result := HelloWord("Dimas")

	if result != "Hello Dimas" {
		// error
		//panic("Result is not 'hello eko'")
		//t.Fail()
		t.Error("hasil tidak hello dimas")
	}
	fmt.Println("TestHelloWorldDimasHalim Done")
}

func TestHelloWorldDimas(t *testing.T) {
	result := HelloWord("Dimas")

	if result != "Hello Dimas" {
		//panic("result is not 'hello dimas'")
		//t.FailNow()
		t.Fatal("Hasil bukan hello dimas")
	}

	fmt.Println("TestHelloWorldDimas")
}
