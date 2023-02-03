package envs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	COMMENT_SYMBOL   = "#"
	SEPARATOR_SYMBOL = "="
	ERROR_LINE       = "file is not in correct format"
)

// LoadEnvFromFile loads environment variables from a file
func LoadEnvFromFile(filenames ...string) error {
	for _, filename := range filenames {
		if err := loadEnvs(filename); err != nil {
			return err
		}
	}
	return nil
}

// loadEnvs loads environment variables from a file
func loadEnvs(filename string) error {
	if strings.TrimSpace(filename) == "" {
		return fmt.Errorf("filename is empty")
	}

	// Read environment variables from envlist.env
	fs, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fs.Close()

	// Iterate over each line in the file
	scanner := bufio.NewScanner(fs)
	for scanner.Scan() {
		// Exclude the comments
		if strings.HasPrefix(scanner.Text(), COMMENT_SYMBOL) {
			continue
		}

		// Skip line if it is empty
		if len(strings.TrimSpace(scanner.Text())) == 0 {
			continue
		}

		// Split the line on the = character
		line := strings.Split(scanner.Text(), SEPARATOR_SYMBOL)
		countIndex := len(line)

		// Get the environment variable name and value
		envVarName := strings.TrimSpace(line[0])
		envVarValue := strings.TrimSpace(line[1])

		// Check if countIndex >= 2, concatenate the value if countIndex > 2
		if countIndex >= 2 {
			for i := 2; i < countIndex; i++ {
				envVarValue = envVarValue + SEPARATOR_SYMBOL + strings.TrimSpace(line[i])
			}
		}

		// Validate line[0] to block empty environment variable name
		if envVarName == "" {
			return fmt.Errorf(filename + " file is not in correct format, variable name is empty")
		}

		// Validate line[0] to block invalid environment variable name
		if regexp.MustCompile("^[a-zA-Z_]{1,}[a-zA-Z0-9_]{0,}$").MatchString(line[0]) == false {
			return fmt.Errorf(filename + " file is not in correct format, variable name is invalid")
		}

		// Remove comments from the envVarValue starts with #
		envVarValue = strings.Split(envVarValue, COMMENT_SYMBOL)[0]

		// Set the environment variable
		os.Setenv(envVarName, envVarValue)
	}
	return nil
}
