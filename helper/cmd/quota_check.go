package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/spf13/cobra"
	"github.com/vmware-tap-on-public-cloud/quickstart-vmware-tanzu-application-platform/helper/pkg/quota"
)

func createQuotaCheckCmd(awsConfig *aws.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "quota-check",
		Aliases: []string{"quota", "qc"},
		Short:   "check AWS quotas",

		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			client := servicequotas.NewFromConfig(*awsConfig)

			quotas := quota.NewGetter(ctx, client, quota.NewServiceQuotaCodes(ctx, client, "ec2"))

			val, err := quotas.Get("EC2-VPC Elastic IPs")
			if err != nil {
				return err
			}

			expected := 100
			if val <= expected {
				return ExitCodeError{
					Message: fmt.Sprintf("Expected quota '%s' to be greater or equal to %d, but got %d", "EC2-VPC Elastic IPs", expected, val),
					Code:    42,
				}
			}

			return nil
		},
	}
}

func p(s string) *string { return &s }
