package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRefNo(t *testing.T) {
	var listOfRandomNumber []string
	// Invoke 100 random Number
	testIteration := 100
	for index := 0; index < testIteration; index++ {
		listOfRandomNumber = append(listOfRandomNumber, GenerateRefNo())
	}
	// Temp map for test2
	test2 := map[string]int{}
	for k, v := range listOfRandomNumber {
		// Checking len of unique number
		test1 := len(v) >= 16
		expecttest1 := true
		assert.Equal(t, expecttest1, test1, fmt.Sprintf("length of %v is %v", v, len(v)))

		test2[v] = k
	}
	// After flip the value (unique number) to be key, if found any duplicate number it's should be replace the current key and decrease length of data that has been created before
	assert.Equal(t, testIteration, len(test2), fmt.Sprintf("took %v iteration but only %v unique number created", testIteration, len(test2)))
}
