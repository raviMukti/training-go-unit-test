package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Before
	fmt.Println("Before Unit Test, ex do something")

	m.Run()

	// After
	fmt.Println("After Unit Test, ex do something")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ravi")
	if result != "Hello Ravi" {
		// Error
		// panic("Result UnMatch")
		// t.Fail()
		// t.FailNow()
		// t.Error("Result Must Be Hello Ravi")
		t.Fatal("Result Must Be Hello Ravi")
	}
	fmt.Println("TestHelloWorld")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ravi")
	assert.Equal(t, "Hello Ravi", result, "Harusnya Hello Ravi")
	fmt.Println("TestHelloWorld With Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ravi")
	require.Equal(t, "Hello Ravi", result, "Harusnya Hello Ravi")
	fmt.Println("TestHelloWorld With Require Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot Run In Linux")
	}

	result := HelloWorld("Ravi")
	require.Equal(t, "Hello Ravi", result, "Harusnya Hello Ravi")
}

func TestSubTest(t *testing.T) {
	t.Run("Ravi", func(t *testing.T) {
		result := HelloWorld("Ravi")
		require.Equal(t, "Hello Ravi", result, "Harusnya Hello Ravi")
	})
	t.Run("Mukti", func(t *testing.T) {
		result := HelloWorld("Mukti")
		require.Equal(t, "Hello Mukti", result, "Harusnya Hello Mukti")
	})
}
