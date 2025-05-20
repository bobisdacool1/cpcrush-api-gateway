package handler_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bobisdacool1/cpcrush/api-gateway/test/testutils"
)

func TestGetHealthcheck_Integration(t *testing.T) {
	app := testutils.NewTestApplication(t)

	req := httptest.NewRequest("GET", "/healthcheck", nil)
	resp, err := app.Server.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	err = resp.Body.Close()
	require.NoError(t, err)

	assert.JSONEq(t, `{"status":"ok"}`, string(body))
}
