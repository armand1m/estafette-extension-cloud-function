package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	trueValue   = true
	falseValue  = false
	validParams = Params{
		Runtime:        "go111",
		Memory:         "256MB",
		Source:         ".",
		TimeoutSeconds: 60,
	}
	validCredential = GKECredentials{
		Name: "gke-production",
	}
)

func TestSetDefaults(t *testing.T) {
	t.Run("DefaultsAppToGitNameIfAppParamIsEmptyAndAppLabelIsEmpty", func(t *testing.T) {

		params := Params{
			App: "",
		}
		gitName := "mygitrepo"
		appLabel := ""

		// act
		params.SetDefaults(gitName, appLabel, "", "", "", map[string]string{})

		assert.Equal(t, "mygitrepo", params.App)
	})

	t.Run("DefaultsAppToAppLabelIfEmpty", func(t *testing.T) {

		params := Params{
			App: "",
		}
		appLabel := "myapp"

		// act
		params.SetDefaults("", appLabel, "", "", "", map[string]string{})

		assert.Equal(t, "myapp", params.App)
	})

	t.Run("KeepsAppIfNotEmpty", func(t *testing.T) {

		params := Params{
			App: "yourapp",
		}
		appLabel := "myapp"

		// act
		params.SetDefaults("", appLabel, "", "", "", map[string]string{})

		assert.Equal(t, "yourapp", params.App)
	})

	t.Run("DefaultsMemoryTo256MB", func(t *testing.T) {

		params := Params{
			Memory: "256MB",
		}

		// act
		params.SetDefaults("", "", "", "", "", map[string]string{})

		assert.Equal(t, "256MB", params.Memory)
	})

	t.Run("KeepsMemoryIfNotEmpty", func(t *testing.T) {

		params := Params{
			Memory: "128MB",
		}

		// act
		params.SetDefaults("", "", "", "", "", map[string]string{})

		assert.Equal(t, "128MB", params.Memory)
	})

	t.Run("DefaultsSourceToCurrentDirectory", func(t *testing.T) {

		params := Params{
			Source: "",
		}

		// act
		params.SetDefaults("", "", "", "", "", map[string]string{})

		assert.Equal(t, ".", params.Source)
	})

	t.Run("KeepsSourceIfNotEmpty", func(t *testing.T) {

		params := Params{
			Source: "otherpath/",
		}

		// act
		params.SetDefaults("", "", "", "", "", map[string]string{})

		assert.Equal(t, "otherpath/", params.Source)
	})

	t.Run("DefaultsTimeoutTo60Seconds", func(t *testing.T) {

		params := Params{
			TimeoutSeconds: 0,
		}

		// act
		params.SetDefaults("", "", "", "", "", map[string]string{})

		assert.Equal(t, 60, params.TimeoutSeconds)
	})

	t.Run("KeepsTimeoutIfLargerThanZero", func(t *testing.T) {

		params := Params{
			TimeoutSeconds: 30,
		}

		// act
		params.SetDefaults("", "", "", "", "", map[string]string{})

		assert.Equal(t, 30, params.TimeoutSeconds)
	})
}

func TestValidateRequiredProperties(t *testing.T) {

	t.Run("ReturnsFalseIfRuntimeIsNotSupported", func(t *testing.T) {

		params := validParams
		params.Runtime = "nodejs6"

		// act
		valid, errors, _ := params.ValidateRequiredProperties()

		assert.False(t, valid)
		assert.True(t, len(errors) > 0)
	})

	t.Run("ReturnsTrueIfRuntimeIsSupported", func(t *testing.T) {

		params := validParams
		params.Runtime = "go111"

		// act
		valid, errors, _ := params.ValidateRequiredProperties()

		assert.True(t, valid)
		assert.True(t, len(errors) == 0)
	})

	t.Run("ReturnsFalseIfMemoryIsNotSupported", func(t *testing.T) {

		params := validParams
		params.Memory = "64MB"

		// act
		valid, errors, _ := params.ValidateRequiredProperties()

		assert.False(t, valid)
		assert.True(t, len(errors) > 0)
	})

	t.Run("ReturnsTrueIfRuntimeIsSupported", func(t *testing.T) {

		params := validParams
		params.Memory = "512MB"

		// act
		valid, errors, _ := params.ValidateRequiredProperties()

		assert.True(t, valid)
		assert.True(t, len(errors) == 0)
	})

	t.Run("ReturnsFalseIfTimeoutSecondsIsLargerThan540Seconds", func(t *testing.T) {

		params := validParams
		params.TimeoutSeconds = 541

		// act
		valid, errors, _ := params.ValidateRequiredProperties()

		assert.False(t, valid)
		assert.True(t, len(errors) > 0)
	})

	t.Run("ReturnsTrueIfTimeoutSecondsIsLessThan540Seconds", func(t *testing.T) {

		params := validParams
		params.TimeoutSeconds = 540

		// act
		valid, errors, _ := params.ValidateRequiredProperties()

		assert.True(t, valid)
		assert.True(t, len(errors) == 0)
	})
}
