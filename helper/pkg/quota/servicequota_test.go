package quota_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/stretchr/testify/assert"
	"github.com/vmware-tap-on-public-cloud/quickstart-vmware-tanzu-application-platform/helper/pkg/quota"

	"github.com/vmware-tap-on-public-cloud/quickstart-vmware-tanzu-application-platform/helper/pkg/quota/quotafakes"
)

type testClient = *quotafakes.FakeServiceQuotaClient

func p(s string) *string { return &s }

func TestSeriveQuotaCodes(t *testing.T) {
	type testCase struct {
		prepareClient func(testClient)
		serviceName   string
		check         func(t *testing.T, code string, err error, client testClient)
	}

	testCases := map[string]testCase{
		"client returns an error": {
			prepareClient: func(c testClient) {
				c.ListServiceQuotasReturns(nil, fmt.Errorf("some error"))
			},
			check: func(t *testing.T, _ string, err error, client testClient) {
				assert.EqualError(t, err, "some error")
				assert.Equal(t, client.ListServiceQuotasCallCount(), 1)
			},
		},
		"code cannot be found": {
			prepareClient: func(c testClient) {
				c.ListServiceQuotasReturns(&servicequotas.ListServiceQuotasOutput{}, nil)
			},
			serviceName: "some name",
			check: func(t *testing.T, _ string, err error, client testClient) {
				assert.ErrorIs(t, err, quota.ErrNotFound)
				assert.ErrorContains(t, err, "some name")
				assert.Equal(t, client.ListServiceQuotasCallCount(), 1)
			},
		},
		"code can be found": {
			prepareClient: func(c testClient) {
				c.ListServiceQuotasReturns(&servicequotas.ListServiceQuotasOutput{
					Quotas: []types.ServiceQuota{
						{QuotaName: p("some name"), QuotaCode: p("some code")},
					},
				}, nil)
			},
			serviceName: "some name",
			check: func(t *testing.T, actualCode string, err error, client testClient) {
				assert.NoError(t, err)
				assert.Equal(t, actualCode, "some code")
				assert.Equal(t, client.ListServiceQuotasCallCount(), 1)
			},
		},
		"code can be found, multi page": {
			prepareClient: func(c testClient) {
				c.ListServiceQuotasReturnsOnCall(0, &servicequotas.ListServiceQuotasOutput{
					NextToken: p("gimme next page"),
					Quotas: []types.ServiceQuota{
						{QuotaName: p("some other name"), QuotaCode: p("some other code")},
					},
				}, nil)
				c.ListServiceQuotasReturnsOnCall(1, &servicequotas.ListServiceQuotasOutput{
					Quotas: []types.ServiceQuota{
						{QuotaName: p("some name"), QuotaCode: p("some code")},
					},
				}, nil)
			},
			serviceName: "some name",
			check: func(t *testing.T, actualCode string, err error, client testClient) {
				assert.NoError(t, err)
				assert.Equal(t, actualCode, "some code")
				assert.Equal(t, client.ListServiceQuotasCallCount(), 2)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			client := &quotafakes.FakeServiceQuotaClient{}
			if tc.prepareClient != nil {
				tc.prepareClient(client)
			}

			codes := quota.NewServiceQuotaCodes(context.TODO(), client, "ec2")
			code, err := codes.Get(tc.serviceName)
			tc.check(t, code, err, client)
		})
	}
}
