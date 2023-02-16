package quota

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
)

//counterfeiter:generate . ServiceQuotaGetterClient
type ServiceQuotaGetterClient interface {
	GetServiceQuota(context.Context, *servicequotas.GetServiceQuotaInput, ...func(*servicequotas.Options)) (*servicequotas.GetServiceQuotaOutput, error)
}

//counterfeiter:generate . ServicQuotaCodeGetter
type ServiceQuotaCodeGetter interface {
	Get(string) (string, error)
	ServiceCode() string
}

func NewGetter(ctx context.Context, client ServiceQuotaGetterClient, quotaCodeGetter ServiceQuotaCodeGetter) *Quotas {
	return &Quotas{
		ctx:    ctx,
		client: client,

		quotaCodeGetter: quotaCodeGetter,
	}
}

type Quotas struct {
	ctx    context.Context
	client ServiceQuotaGetterClient

	quotaCodeGetter ServiceQuotaCodeGetter
}

func (q *Quotas) Get(quotaName string) (int, error) {
	quotaCode, err := q.quotaCodeGetter.Get(quotaName)
	if err != nil {
		return 0, fmt.Errorf("getting service quota code: %w", err)
	}

	serviceCode := q.quotaCodeGetter.ServiceCode()

	quota, err := q.client.GetServiceQuota(q.ctx, &servicequotas.GetServiceQuotaInput{
		ServiceCode: &serviceCode,
		QuotaCode:   &quotaCode,
	})
	if err != nil {
		return 0, fmt.Errorf("getting quota: %w", err)
	}

	return int(*quota.Quota.Value), nil
}
