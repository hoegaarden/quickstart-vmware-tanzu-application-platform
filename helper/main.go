package main

import (
	"context"
	"errors"
	"os"

	"github.com/vmware-tap-on-public-cloud/quickstart-vmware-tanzu-application-platform/helper/cmd"
)

func main() {
	ctx := context.TODO()
	if err := cmd.ExecuteContext(ctx); err != nil {
		var cmdErr cmd.ExitCodeError

		if errors.As(err, &cmdErr) {
			os.Exit(cmdErr.Code)
		}

		os.Exit(1)
	}
}
