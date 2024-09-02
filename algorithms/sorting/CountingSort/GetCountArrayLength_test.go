package CountingSort

import (
    "testing"
)

// TestGetCountArrayLength tests the behavior of getCountArrayLength function with various scenarios.
func TestGetCountArrayLength(t *testing.T) {
    testCases := []struct {
        name     string
        input    []int
        expected int
    }{
        {
            name:     "Empty Array",
            input:    []int{},
            expected: 1,
        },
        {
            name:     "Single Element Array",
            input:    []int{5},
            expected: 6,
        },
        {
            name:     "Array with Multiple Elements",
            input:    []int{1, 3, 2, 5},
            expected: 6,
        },
        {
            name:     "Array with Negative Elements",
            input:    []int{-1, -3, 2, 5},
            expected: 6,
        },
        {
            name:     "Array with All Identical Elements",
            input:    []int{4, 4, 4, 4},
            expected: 5,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Act
            result := getCountArrayLength(tc.input)

            // Assert
            if result != tc.expected {
                t.Errorf("Failed %s: expected %d, got %d", tc.name, tc.expected, result)
            } else {
                t.Logf("Success %s: expected %d, got %d", tc.name, tc.expected, result)
            }
        })
    }
}
