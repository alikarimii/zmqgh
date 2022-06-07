package bytegen_test

import (
	"testing"

	"github.com/alikarimii/zmqph/publisher/zero/infrastructure/adapter/bytegen"
)

func TestGenerateRandomByte(t *testing.T) {

}

func TestGenerateRandomBetween(t *testing.T) {
	inputs := make(map[int]int) // map[min]max
	inputs[0] = 5
	inputs[50] = 8192
	inputs[100] = 20
	for bt1, bt2 := range inputs {
		iterCount := 10
		for ; iterCount > 0; iterCount-- {
			if k := bytegen.GenerateRandomBetween(bt1, bt2); !((k >= bt1 && k <= bt2) || (k <= bt1 && k >= bt2)) {
				t.Errorf("expected: k in range %d-%d, got: %d", bt1, bt2, k)
			}
		}
	}
}
