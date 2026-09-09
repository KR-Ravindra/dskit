package flagext

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.yaml.in/yaml/v3"
)

func TestDayValueYAML(t *testing.T) {
	t.Run("embedding DayValue", func(t *testing.T) {
		type TestStruct struct {
			Day DayValue `yaml:"day"`
		}

		var testStruct TestStruct
		require.NoError(t, testStruct.Day.Set("1985-06-02"))
		expected := []byte(`day: "1985-06-02"
`)

		actual, err := yaml.Marshal(testStruct)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)

		var actualStruct TestStruct
		err = yaml.Unmarshal(expected, &actualStruct)
		require.NoError(t, err)
		assert.Equal(t, testStruct, actualStruct)
	})

	t.Run("pointer of DayValue", func(t *testing.T) {
		type TestStruct struct {
			Day *DayValue `yaml:"day"`
		}

		var testStruct TestStruct
		testStruct.Day = &DayValue{}
		require.NoError(t, testStruct.Day.Set("1985-06-02"))
		expected := []byte(`day: "1985-06-02"
`)

		actual, err := yaml.Marshal(testStruct)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)

		var actualStruct TestStruct
		err = yaml.Unmarshal(expected, &actualStruct)
		require.NoError(t, err)
		assert.Equal(t, testStruct, actualStruct)
	})
}

func Test_DayValue_SetInvalid(t *testing.T) {
	for _, tcase := range []struct {
		input string
	}{
		{"invalid"},
		{"abc"},
		{"2023-13-01"}, // Invalid month
		{"2023-00-01"}, // Invalid month
		{"2023-01-32"}, // Invalid day
		{"2023-01-00"}, // Invalid day
		{"2023-02-29"}, // Not a leap year (2023 is not leap)
		{""}, // Empty string
	} {
		var d DayValue
		err := d.Set(tcase.input)
		assert.Error(t, err, "input %q should fail", tcase.input)
	}
}
