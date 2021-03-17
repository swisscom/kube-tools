package kubetools

import (
	"fmt"
	"os"
)

func Errorf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args)
}
