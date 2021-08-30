package geminirouter

import (
	"net/url"
	"testing"

	"github.com/kulak/gemini"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	url, err := url.ParseRequestURI("GEMINI://localhost/")
	require.NoError(t, err)
	require.Equal(t, gemini.SchemaGemini, url.Scheme)
	t.Logf("scheme: %s", url.Scheme)
}
