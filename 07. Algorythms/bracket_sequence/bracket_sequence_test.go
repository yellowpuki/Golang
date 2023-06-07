package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	name  string
	input string
	exp   bool
}{
	{
		name:  "Test1",
		input: "{}",
		exp:   true,
	},
	{
		name:  "Test2",
		input: "",
		exp:   true,
	},
	{
		name:  "Test3",
		input: "{",
		exp:   false,
	},
	{
		name:  "Test4",
		input: "{[(",
		exp:   false,
	},
	{
		name:  "Test5",
		input: "{[()",
		exp:   false,
	},
	{
		name:  "Test6",
		input: "{[]()",
		exp:   false,
	},
	{
		name:  "Test7",
		input: "{[]()}",
		exp:   true,
	},
	{
		name:  "Test8",
		input: "][",
		exp:   false,
	},
	{
		name:  "Test9",
		input: "}{",
		exp:   false,
	},
	{
		name:  "Test10",
		input: ")(",
		exp:   false,
	},
}

func TestCheckSequence(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.exp, CheckSequence(tc.input))
		})
	}
}
