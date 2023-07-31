package config

import (
	"fmt"
	"strings"

	"git.blender.org/flamenco/pkg/crosspath"
	"github.com/rs/zerolog/log"
)

type ValueToVariableReplacer struct {
	twoWayVars map[string]string // Mapping from variable name to value.
}

// NewVariableToValueConverter returns a ValueToVariableReplacer for the given audience & platform.
func (c *Conf) NewVariableToValueConverter(audience VariableAudience, platform VariablePlatform) *ValueToVariableReplacer {
	// Get the variables for the given audience & platform.
	twoWayVars := c.GetTwoWayVariables(audience, platform)

	if len(twoWayVars) == 0 {
		log.Debug().
			Str("audience", string(audience)).
			Str("platform", string(platform)).
			Msg("no two-way variables defined for this platform given this audience")
	}

	return &ValueToVariableReplacer{
		twoWayVars: twoWayVars,
	}
}

// ValueToVariableReplacer replaces any variable values it recognises in
// valueToConvert to the actual variable. For example, `/path/to/file.blend` can
// be changed to `{my_storage}/file.blend`.
func (vvc *ValueToVariableReplacer) Replace(valueToConvert string) string {
	result := valueToConvert

	for varName, varValue := range vvc.twoWayVars {
		if !isValueMatch(result, varValue) {
			continue
		}
		result = vvc.join(varName, result[len(varValue):])
	}

	log.Debug().
		Str("from", valueToConvert).
		Str("to", result).
		Msg("first step of two-way variable replacement")

	return result
}

func (vvc *ValueToVariableReplacer) join(varName, value string) string {
	return fmt.Sprintf("{%s}%s", varName, value)
}

// isValueMatch returns whether `valueToMatch` starts with `variableValue`.
// When `variableValue` is a Windows path (with backslash separators), it is
// also tested with forward slashes against `valueToMatch`.
func isValueMatch(valueToMatch, variableValue string) bool {
	if strings.HasPrefix(valueToMatch, variableValue) {
		return true
	}

	// If the variable value has a backslash, assume it is a Windows path.
	// Convert it to slash notation just to see if that would provide a
	// match.
	if strings.ContainsRune(variableValue, '\\') {
		slashedValue := crosspath.ToSlash(variableValue)
		return strings.HasPrefix(valueToMatch, slashedValue)
	}

	return false
}
