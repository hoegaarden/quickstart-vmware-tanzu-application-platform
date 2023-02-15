package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
)

func ExecuteContext(ctx context.Context) error {
	var (
		cfg    aws.Config
		region string
	)

	var rootCmd = &cobra.Command{
		Use:          "helper", // TODO
		Short:        "helper for the VMware Tanzu Application Platform Quickstart / Cloudformation stack",
		SilenceUsage: true,

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			cfg, err = config.LoadDefaultConfig(cmd.Context())
			if err != nil {
				return err
			}
			if region != "" {
				cfg.Region = region
			}

			return nil
		},
	}

	rootCmd.PersistentFlags().StringVar(&region, "region", "", "the AWS region to run against")

	rootCmd.AddCommand(createQuotaCheckCmd(&cfg))

	return rootCmd.ExecuteContext(ctx)
}
