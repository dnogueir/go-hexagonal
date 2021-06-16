package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	expected := []byte(`{"message":"Hello Json"}`)
	require.Equal(t, expected, result)
}
