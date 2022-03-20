package parser

import (
	"encoding/json"
	"fmt"
)

// Format takes in multiple configurations input and formats the configuration
// to be more human readable. The key of each configuration should be its filepath.
func Format(configurations map[string]interface{}) (string, error) {
	var output string
	for file, config := range configurations {
		output += file + "\n"

		current, err := format(config)
		if err != nil {
			return "", fmt.Errorf("marshal output to json: %w", err)
		}

		output += current
	}

	return output, nil
}

// FormatJSON takes in multiple configurations and formats them as a JSON list
// of objects with a filePath and configuration key for each.
func FormatJSON(configurations map[string]interface{}) (string, error) {
	type configWithFilePath struct {
		FilePath      string
		Configuration interface{}
	}
	var configs []configWithFilePath
	for fp, cfg := range configurations {
		configs = append(configs, configWithFilePath{FilePath: fp, Configuration: cfg})
	}
	marshalled, err := json.MarshalIndent(configs, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal configs: %w", err)
	}

	return string(marshalled), nil
}

// FormatCombined takes in multiple configurations, combines them, and formats the
// configuration to be more human readable. The key of each configuration should be
// its filepath.
func FormatCombined(configurations map[string]interface{}) (string, error) {
	combinedConfigurations := CombineConfigurations(configurations)

	formattedConfigs, err := format(combinedConfigurations["Combined"])
	if err != nil {
		return "", fmt.Errorf("formatting configs: %w", err)
	}

	return formattedConfigs, nil
}

func format(configs interface{}) (string, error) {
	out, err := json.MarshalIndent(configs, "", "\t")
	if err != nil {
		return "", fmt.Errorf("marshal output to json: %w", err)
	}

	return string(out) + "\n", nil
}
