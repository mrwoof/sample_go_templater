package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHandlers(t *testing.T) {
	_, err := NewHandlers("../../templates")
	require.NoError(t, err)
}

func TestGetSampleOne(t *testing.T) {
	// create a new handler
	handler, err := NewHandlers("../../templates")
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/sample_one", nil)

	// handle the request
	handler.GetSampleOne(w, req)

	// test the code and status text of the response
	require.Equal(t, http.StatusOK, w.Code)
}
