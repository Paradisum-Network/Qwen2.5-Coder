package variableUnknown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymflowerVariableUnknown(t *testing.T) {
	type testCase struct {
		Name string

		X int

		ExpectedInt int
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualInt := variableUnknown(tc.X)

			assert.Equal(t, tc.ExpectedInt, actualInt)
		})
	}

	validate(t, &testCase{
		X: -1,

		ExpectedInt: 1,
	})
	validate(t, &testCase{
		X: 0,

		ExpectedInt: 0,
	})
	validate(t, &testCase{
		X: 1,

		ExpectedInt: 1,
	})
}