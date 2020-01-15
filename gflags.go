package gflags

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func FlagOrEnvString(flagPtr *string, env string) error {
	var err error = nil
	if flagPtr == nil {
		err = errors.New("flagPtr is nil")
	} else if *flagPtr == "" {
		*flagPtr = os.Getenv(env)
		if *flagPtr == "" {
			err = fmt.Errorf("expected flag or environment variable for %s expected", env)
		}
	}
	return err
}
