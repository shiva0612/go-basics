package calctest_test

import (
	"fmt"
	"shiva/testing-framework/calc"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	testName string
	ip1      int
	ip2      int
	op       int
}

func TestAdd5(t *testing.T) {
	assert.Equal(t, 3, calc.Add(1, 2))
}

func BenchmarkAdd(b *testing.B) {

	for i := 0; i < b.N; i++ {
		calc.Add(1, 2)
	}

}

func TestAdd4(t *testing.T) {
	test_inputs := []Input{
		Input{
			testName: "test1",
			ip1:      1,
			ip2:      2,
			op:       3,
		},
		Input{
			testName: "test2",
			ip1:      2,
			ip2:      3,
			op:       5,
		},
	}

	for _, test_input := range test_inputs {
		t.Run(test_input.testName, func(t *testing.T) {
			t.Logf("running test with input %v", test_input)
			got := calc.Add(test_input.ip1, test_input.ip2)

			if got != test_input.op {
				t.Fail()
			}
		})
	}
}
func TestAdd1(t *testing.T) {

	//create input and expected output
	a, b := 1, 1
	expected := 2

	//call
	got := calc.Add(a, b)

	//assert output
	if expected != got {
		t.Fail()
	}

}

func TestAdd2(t *testing.T) {
	t.Run("positive test", func(t *testing.T) {
		//create input and expected output
		a, b := 1, 1
		expected := 2

		//call
		got := calc.Add(a, b)

		//assert output
		if expected != got {
			t.Fail()
		}
	})
	t.Run("negative test", func(t *testing.T) {
		//create input and expected output
		a, b := 1, 1
		expected := 1

		//call
		got := calc.Add(a, b)

		//assert output
		if expected != got {
			t.Fail()
		}
	})
}

func TestAdd3(t *testing.T) {
	test_inputs := [][]int{{1, 2}, {3, 4}, {5, 6}}
	expected_outputs := []int{3, 7, 11}

	for i, test_input := range test_inputs {
		testName := fmt.Sprintf("test-%d", i)
		t.Run(testName, func(t *testing.T) {
			got := calc.Add(test_input[0], test_input[1])

			if got != expected_outputs[i] {
				t.Fail()
			}
		})
	}
}
