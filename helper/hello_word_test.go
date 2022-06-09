package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// function test main
func TestMain(m *testing.M) {
	// sebelum unit test
	fmt.Println("SEBELUM UNIT TEST")
	m.Run()

	// setelah
	fmt.Println("SETELAH UNIT TEST")
}

//todo *panic jangan di pake untuk test unit*

// todo assert unit test function
func TestHelloWorldAssert(t *testing.T) {
	// menggunakan assertion
	result := HelloWord("Dimas")
	//required.Equal(t, "Hello Dimas", result, "yeay berhasil nampilin data")
	require.Equal(t, "hello Dimas", result, "Hasil nya harus 'Hallo halim'")
	fmt.Println("data nasabah belum bisa tampil")
}

// todo menggunakan assert require
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWord("Halim")
	assert.Equal(t, "hello Halim", result)
	fmt.Println("Tidak di eksekusi")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("gabisa jalan di server dev")
	}

	result := HelloWord("CMS-EPAY")
	require.Equal(t, "hello CMS", result, "TEST CMS SENIN DI MULAI")
}

// testing nasabah
func TestNasabah(t *testing.T) {
	result := HelloNasabah("dimas")

	assert.Equal(t, "hai dimas", result, "data ada")
	fmt.Println("data berhasil tampil")
}

func TestHelloWorldDimasHalim(t *testing.T) {
	result := HelloWord("Dimas")

	if result != "hello Dimas" {
		// error
		//panic("Result is not 'hello eko'")
		//t.Fail()
		t.Error("hasil tidak hello dimas")
	}
	fmt.Println("TestHelloWorldDimasHalim Done")
}

func TestHelloWorldDimas(t *testing.T) {
	result := HelloWord("Dimas")

	if result != "hello Dimas" {
		//panic("result is not 'hello dimas'")
		//t.FailNow()
		t.Fatal("Hasil bukan hello dimas")
	}

	fmt.Println("TestHelloWorldDimas")
}
