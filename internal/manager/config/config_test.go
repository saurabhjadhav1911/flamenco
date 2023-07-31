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

	assert.Equal(t, `C:\Downloads\tab\newline.exe`, vars["single-backslash-common-escapechar"]["blender"])
	assert.Equal(t, `C:\Downloads\blender-1.0\`, vars["single-backslash-trailing"]["blender"])
	assert.Equal(t, `F:\`, vars["single-backslash-drive-only"]["blender"])
}

func TestReplaceTwowayVariables(t *testing.T) {
	c := DefaultConfig(func(c *Conf) {
		c.Variables["shared"] = Variable{
			IsTwoWay: true,
			Values: []VariableValue{
				{Value: "/shared/flamenco", Platform: VariablePlatformLinux},
				{Value: `Y:\shared\flamenco`, Platform: VariablePlatformWindows},
			},
		}
	})

	feeder := make(chan string, 2)
	receiver := make(chan string, 2)

	feeder <- `Y:\shared\flamenco\shot\file.blend`

	// This is the real reason for this test: forward slashes in the path should
	// still be matched to the backslashes in the variable value.
	feeder <- `Y:/shared/flamenco/shot/file.blend`
	close(feeder)

	c.ConvertTwoWayVariables(feeder, receiver, VariableAudienceUsers, VariablePlatformWindows)

	assert.Equal(t, `{shared}\shot\file.blend`, <-receiver)
	assert.Equal(t, `{shared}/shot/file.blend`, <-receiver)
}
