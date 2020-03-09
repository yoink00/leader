package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetQuorum(t *testing.T) {
	var peers []string
	expected := []int{1, 2, 2, 3, 3, 4, 4, 5, 5, 6}
	for i := 0; i < 10; i++ {
		peers = append(peers, fmt.Sprintf("FOO%d", i))
		t.Logf("Expected: %d, Actual: %d, Index: %d, Len: %d", expected[i], getQuorum(peers), i, len(peers))
		assert.Equal(t, expected[i], getQuorum(peers))
	}
}
