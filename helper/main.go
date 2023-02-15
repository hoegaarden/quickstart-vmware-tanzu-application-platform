package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/vmware-tap-on-public-cloud/quickstart-vmware-tanzu-application-platform/helper/pkg/quota"
)

const (
	region = "us-east-1"

	serviceCodeEC2 = "ec2"
)

func main() {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		panic(err)
	}

	client := servicequotas.NewFromConfig(cfg)

	// Service quota names seem to be stable across regions, but service quota
	// codes are not. Thus we neeed something that can translate from one to the
	// other.
	quotaCodesEC2 := quota.NewServiceQuotaCodes(ctx, client, serviceCodeEC2)

	value, err := getCurrentQuota(ctx, client, serviceCodeEC2, quotaCodesEC2.MustGet("EC2-VPC Elastic IPs"))
	if err != nil {
		panic(err)
	}

	if value < 100.0 {
		log.Fatalf("EC2-VPC Elastic IPs quota needs to be at least at 100, got: %d", int(value))
	}
}

func getCurrentQuota(ctx context.Context, client *servicequotas.Client, serviceCode, quotaCode string) (float64, error) {
	quota, err := client.GetServiceQuota(ctx, &servicequotas.GetServiceQuotaInput{
		ServiceCode: &serviceCode,
		QuotaCode:   &quotaCode,
	})
	if err != nil {
		return 0, err
	}

	return *quota.Quota.Value, nil
}
