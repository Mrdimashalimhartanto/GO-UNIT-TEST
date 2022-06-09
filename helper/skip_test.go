package helper

import (
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

/**
todo *SKIP TEST di gunakan untuk melawati sebuah test jika terjadi sebuah kegagalan*
*/

func TestSkip2(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("CMS EPAY ON")
	}

	result := HelloWord("Dimas")
	require.Equal(t, "hello dimas", result, "Test server belum bisa")
}
