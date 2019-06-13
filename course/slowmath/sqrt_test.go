package slowmath

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestSqrtMany(t *testing.T) {
	var testCases = []struct {
		val       float64
		expected  float64
		shouldErr bool
	}{
		{2.0, 1.4142, false},
		{0, 0, false},
		{-1, 0, true},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%v", tc.val)
		t.Run(name, func(t *testing.T) {
			req := require.New(t)
			val, err := Sqrt(tc.val)
			hasErr := err != nil
			req.Equal(tc.shouldErr, hasErr, "Error")
			if !tc.shouldErr {
				req.InDelta(tc.expected, val, 0.0001, tc.val)
			}
		})
	}

}

func TestWithJson(t *testing.T) {
	type testCase struct {
		Value    float64 `json:"value"`
		Expected float64 `json:"expected"`
		Error    bool    `json:"error"`
	}

	jsonFile, err := os.Open("cases.json")
	if err != nil {
		t.Fatal("cant open file", err)
	}

	var testCases []testCase
	dec := json.NewDecoder(jsonFile)
	require.Nil(t, dec.Decode(&testCases))

	for _, tc := range testCases {
		name := fmt.Sprintf("%v", tc.Value)
		t.Run(name, func(t *testing.T) {
			req := require.New(t)
			val, err := Sqrt(tc.Value)
			hasErr := err != nil
			req.Equal(tc.Error, hasErr, "Error")
			if !tc.Error {
				req.InDelta(tc.Expected, val, 0.0001, tc.Value)
			}
		})
	}

}
