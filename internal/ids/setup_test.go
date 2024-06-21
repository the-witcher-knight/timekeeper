package ids

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetup(t *testing.T) {
	err := Setup()
	require.NoError(t, err)

	require.NotNil(t, BlockChainAttendanceRecord)
	require.NotNil(t, User)
	require.NotNil(t, Attendance)
}
