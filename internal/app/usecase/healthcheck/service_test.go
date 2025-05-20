package healthcheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService_Ping(t *testing.T) {
	s := NewService()

	require.Equal(t, map[string]string{"status": "ok"}, s.Ping())
}
