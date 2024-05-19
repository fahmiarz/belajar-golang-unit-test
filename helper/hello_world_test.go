package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Fahmi",
			request: "Fahmi",
		},
		{
			name:    "Arzalega",
			request: "Arzalega",
		},
		{
			name:    "FahmiArzalega",
			request: "Fahmi Arzalega Seazz",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Fahmi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fahmi")
		}
	})
	b.Run("Kurniawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kurniawan")
		}
	})

}

func BenchmarkHelloWorldFahmi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fahmi")
	}
}
func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

func TestTableHelloWorld(t *testing.T) { //table test
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fahmi",
			request:  "Fahmi",
			expected: "Hello Fahmi",
		},
		{
			name:     "Arzalega",
			request:  "Arzalega",
			expected: "Hello Arzalega",
		},
		{
			name:     "Odetta",
			request:  "Odetta",
			expected: "Hello Odetta",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) { //sub test
	t.Run("Fahmi", func(t *testing.T) {
		result := HelloWorld("Fahmi")
		require.Equal(t, "Hello Fahmi", result, "Result must be 'Hello Fahmi'")
	})
	t.Run("Arzalega", func(t *testing.T) {
		result := HelloWorld("Arzalega")
		require.Equal(t, "Hello Arzalega", result, "Result must be 'Hello Arzalega'")
	})
}

func TestMain(m *testing.M) { //test main
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST)")
}

func TestSkip(t *testing.T) { //test skip
	if runtime.GOOS == "darwin" {
		t.Skip("Cant not run on Mac OS")

	}
	result := HelloWorld("Fahmi")
	require.Equal(t, "Hello Fahmi", result, "Result must be 'Hello Fahmi'")

}

func TestHelloWorldRequire(t *testing.T) { //test require
	result := HelloWorld("Fahmi")
	require.Equal(t, "Hello Fahmi", result, "Result must be 'Hello Fahmi'")
	fmt.Println("TestHelloWorldAssert With Require Done")
}

func TestHelloWorldAssert(t *testing.T) { //test assert
	result := HelloWorld("Fahmi")
	assert.Equal(t, "Hello Fahmi", result, "Result must be 'Hello Fahmi'")
	fmt.Println("TestHelloWorldAssert With Assert Done")
}

func TestHelloWorldFahmi(t *testing.T) { //contoh unit test 1
	result := HelloWorld("Fahmi")

	if result != "Hello Fahmi" {
		//error
		t.Error("Result should be 'Hello Fahmi'")
	}
	fmt.Println("TestHelloWorldFahmi Done")
}
func TestHelloArzalega(t *testing.T) { //contoh unit test 2
	result := HelloWorld("Arzalega")

	if result != "Hello Arzalega" {
		//error
		t.Fatal("Result should be 'Hello Arzalega'")
	}
	fmt.Println("TestHelloWorldArzalega Done")
}
