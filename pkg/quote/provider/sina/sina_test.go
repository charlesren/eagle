package sina

import (
	"fmt"
	"math"
	"testing"
)

func TestGet(t *testing.T) {
	data := Get("sz000027")
	fmt.Println("Highest value of today is:", data.High)
	fmt.Println("Lowest value of today is:", data.Low)
	high := math.Max(data.High, data.Low)
	if data.High != high {
		t.Error("Rong data")
	}
}
