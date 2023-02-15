package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/spf13/cobra"
	"github.com/vmware-tap-on-public-cloud/quickstart-vmware-tanzu-application-platform/helper/pkg/quota"
)

func createQuotaCheckCmd(cfg *aws.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "quota-check",
		Aliases: []string{"quota", "qc"},
		Short:   "check AWS quotas",

		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			client := servicequotas.NewFromConfig(*cfg)

			quotaCodesEC2 := quota.NewServiceQuotaCodes(ctx, client, "ec2")

			vpcVipCode, err := quotaCodesEC2.Get("EC2-VPC Elastic IPs")
			if err != nil {
				return err
			}

			vpcVipQuota, err := client.GetServiceQuota(ctx, &servicequotas.GetServiceQuotaInput{
				ServiceCode: p("ec2"),
				QuotaCode:   p(vpcVipCode),
			})

			if val, expected := int(*vpcVipQuota.Quota.Value), 100; val < expected {
				return fmt.Errorf("Expected quota '%s' to be greater or equal to %d, but got %d", "EC2-VPC Elastic IPs", expected, val)
			}

			return nil
		},
	}
}

func p(s string) *string { return &s }
