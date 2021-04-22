package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetCIDR(t *testing.T) {

	t.Run("passing valid CIDR block argument returns cidr", func(t *testing.T) {
		_, err := GetCIDR([]string{"10.1.2.0/24"})
		require.NoError(t, err)
	})

	t.Run("passing no arguments returns error", func(t *testing.T) {
		_, err := GetCIDR([]string{})
		require.Error(t, err)
	})

	t.Run("passing more than 3 arguments returns error", func(t *testing.T) {
		_, err := GetCIDR([]string{"10.1.2.0/24", "4", "15", "9"})
		require.Error(t, err)
	})
}
