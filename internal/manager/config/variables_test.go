package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	replacer := c.NewVariableToValueConverter(VariableAudienceUsers, VariablePlatformWindows)

	// This is the real reason for this test: forward slashes in the path should
	// still be matched to the backslashes in the variable value.
	assert.Equal(t, `{shared}\shot\file.blend`, replacer.Replace(`Y:\shared\flamenco\shot\file.blend`))
	assert.Equal(t, `{shared}/shot/file.blend`, replacer.Replace(`Y:/shared/flamenco/shot/file.blend`))
}
