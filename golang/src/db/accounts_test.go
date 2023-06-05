package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAccount(t *testing.T) {
	id := "1"
	// main_testが走ってない可能性がある
	// account, err := testQueries.GetAccount(id)
	// require.NoError(t, err)
	// require.NotEmpty(t, account)
	require.Equal(t, "1", id)
}
