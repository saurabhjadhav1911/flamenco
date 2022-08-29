package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVariablesWithBackslashes(t *testing.T) {
	c, err := loadConf("test_files/config_with_backslashes.yaml")
	require.NoError(t, err)

	vars := c.VariablesLookup[VariableAudienceUsers]
	expectSingle := `C:\Downloads\blender-1.0\blender.exe`
	expectDouble := `C:\\Downloads\\blender-1.0\\blender.exe`
	assert.Equal(t, expectSingle, vars["single-backslash"]["blender"])
	assert.Equal(t, expectDouble, vars["double-backslash"]["blender"])
	assert.Equal(t, expectSingle, vars["quoted-double-backslash"]["blender"])
}
