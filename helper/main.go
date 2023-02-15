package main

import (
	"context"
	"os"

	"github.com/vmware-tap-on-public-cloud/quickstart-vmware-tanzu-application-platform/helper/cmd"
)

const (
	region = "us-east-1"

	serviceCodeEC2 = "ec2"
)

func main() {
	ctx := context.TODO()
	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
