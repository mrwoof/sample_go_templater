package templater

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadTemplates(t *testing.T) {
	_, err := LoadTemplates("../../templates")
	require.NoError(t, err)
}
