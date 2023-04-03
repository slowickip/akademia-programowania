package testProject

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	want := "Hello World"
	got := HelloWorld()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

type PriceTest struct {
	name      string
	price     int64
	tax       int64
	wantNet   float64
	wantGross float64
}

func TestCalculatePrice(t *testing.T) {
	testSlice := []PriceTest{
		{"19.99 at 8%", 1999, 8, 18.51, 19.99},
		{"123 at 23%", 12300, 23, 100.00, 123.00},
	}

	for _, test := range testSlice {
		t.Run(test.name, func(t *testing.T) {
			gotNet, gotGross := CalculatePrice(test.price, test.tax)
			assert.Equal(t, test.wantNet, gotNet)
			assert.Equal(t, test.wantGross, gotGross)
		})
	}
}

func BenchmarkBufferString(b *testing.B) {
	tests := []struct {
		name     string
		testText string
	}{
		{"a", "a"},
		{"Lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas facilisis odio nec erat ornare aliquam. Aenean vitae faucibus lacus. Maecenas suscipit elit eu sem vestibulum, eget egestas ante pulvinar. Nam non metus commodo, porta ligula id, volutpat velit. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc quis venenatis nisl. Maecenas suscipit mollis malesuada. Maecenas congue leo at mattis placerat."},
	}
	for _, tt := range tests {
		var buffer bytes.Buffer
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buffer.WriteString(tt.testText)
			}
		})
	}
}

func BenchmarkStringBuilderString(b *testing.B) {
	tests := []struct {
		name     string
		testText string
	}{
		{"a", "a"},
		{"Lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas facilisis odio nec erat ornare aliquam. Aenean vitae faucibus lacus. Maecenas suscipit elit eu sem vestibulum, eget egestas ante pulvinar. Nam non metus commodo, porta ligula id, volutpat velit. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc quis venenatis nisl. Maecenas suscipit mollis malesuada. Maecenas congue leo at mattis placerat."},
	}
	for _, tt := range tests {
		var builder strings.Builder
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				builder.WriteString(tt.testText)
			}
		})
	}
}
