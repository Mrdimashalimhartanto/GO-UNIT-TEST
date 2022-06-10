package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// todo banchmark testing befungsi untuk test kecepatan kodingan kita
// todo * kecepatan untuk benchmark 0.127s *
func BenchmarkHelloWorld1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWord("hello")
	}
}

// function test main golang
func TestMain(m *testing.M) {
	// sebelum unit test
	fmt.Println("SEBELUM UNIT TEST")
	m.Run()

	// setelah
	fmt.Println("SETELAH UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("hello", func(t *testing.T) {
		result := HelloWord("hello")
		require.Equal(t, "hello dimas", result, "hasil di temukan")
	})
}

func TestSubTestAlamat(t *testing.T) {
	t.Run("hello", func(t *testing.T) {
		result := HelloWord("hello")
		require.Equal(t, "hello nasabah", result, "NASABAH BERHASIL DI TAMPILKAN")

	})
}

//func TestSubTestBerhasil(t *testing.T) {
//	t.Run("hello", func(t *testing.T) {
//		result := HelloWord("hello")
//		require.Equal(t, "hello", result, "selamat datang")
//
//	})
//}

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
	result := HelloWord("halim")
	assert.Equal(t, "hello halim", result, "halim berhasil ketampil")
	//assert.Equal(t, "hello Halim", result)
	fmt.Println("Tidak di eksekusi")
}

//func TestSkip(t *testing.T) {
//	if runtime.GOOS == "darwin" {
//		t.Skip("gabisa jalan di server dev")
//	}
//
//	result := HelloWord("dimas")
//	require.Equal(t, "hello dimas", result, "TEST CMS SENIN DI MULAI")
//	fmt.Println("data dimas berhasil di tampilkan")
//}

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

func TestDataNasabah(t *testing.T) {
	result := HelloWord("Dimas")

	if result != "hello dimas" {
		t.Error("hasil tidak hello dimas")
	}
	fmt.Println("TestDataNasabah done")
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
