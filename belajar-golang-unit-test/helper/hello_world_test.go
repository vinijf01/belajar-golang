package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name, request string
	}{
		{
			name:    "eko",
			request: "eko",
		},
		{
			name:    "vini",
			request: "vini",
		},
		{
			name:    "vini jumatul fitri",
			request: "vini jumatul fitri",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSubBench(b *testing.B) {
	b.Run("vini", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("vini")
		}
	})
	b.Run("jumatul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("jumatul")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("vini")
	}
}

func BenchmarkHelloWorldVini(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("jumatul")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "eko",
			request:  "eko",
			expected: "Hello eko",
		},
		{
			name:     "vini",
			request:  "vini",
			expected: "Hello vini",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "Result must be 'Hello eko'")
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("eko")
		assert.Equal(t, "Hello eko", result, "Result must be 'Hello eko'")
	})
	t.Run("Khanady", func(t *testing.T) {
		result := HelloWorld("khanady")
		assert.Equal(t, "Hello khanady", result, "Result must be 'Hello khanady'")
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before Unit test")
	m.Run()

	//after
	fmt.Println("After Unit test")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on windows")
	}
	result := HelloWorld("vini")
	assert.Equal(t, "Hello vini", result, "Result must be 'Hello vini'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("vini")
	assert.Equal(t, "Hello vini", result, "Result must be 'Hello vini'")
	fmt.Println("TestHelloWorld with assert Done")

}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("vini")
	require.Equal(t, "Hello vini", result, "Result must be 'Hello vini'")
	fmt.Println("TestHelloWorld with require Done")

}

func TestHelloWorldJumatul(t *testing.T) {
	result := HelloWorld("vini")
	//CARA I
	if result != "Hello vini" {
		//error
		// panic("Result is not 'Hello vini'")
		// t.Fail()
		t.Error("Harusnya 'Hello vini'")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldVini(t *testing.T) {
	result := HelloWorld("jumatul")
	if result != "Hello jumatul" {
		//error
		// panic("Result is not 'Hello jumatul'")
		// t.FailNow()
		t.Fatal("Harusnya 'Hello vini'")

	}
	fmt.Println("TestHelloWorld Vini Done")

}
