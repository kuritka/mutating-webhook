package flags

import (
	"fmt"
	"os"

	"mutating-webhook/common/log"
)

// GetStringFlagFromEnv attemps to lookup a flag with name 'flagName' from the environment
func GetStringFlagFromEnv(flagName string) (string, error) {
	value, lookup := os.LookupEnv(flagName)
	if !lookup {
		return "", fmt.Errorf("failed to get flag %s from env", flagName)
	}

	return value, nil
}

// MustGetStringFlagFromEnv returns the value of the flag 'flagName' or panics if no value is set
func MustGetStringFlagFromEnv(flagName string) string {
	value, err := GetStringFlagFromEnv(flagName)
	if err != nil {
		log.Logger().Panic().Err(err).Msg("failed to load string flag from env")
	}
	return value
}
