package quota

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
)

const (
	pageSize = 100
)

//counterfeiter:generate . ServiceQuotaClient
type ServiceQuotaClient interface {
	ListServiceQuotas(ctx context.Context, params *servicequotas.ListServiceQuotasInput, optFns ...func(*servicequotas.Options)) (*servicequotas.ListServiceQuotasOutput, error)
}

func NewServiceQuotaCodes(context context.Context, client ServiceQuotaClient, serviceCode string) *ServiceQuotaCodes {
	return &ServiceQuotaCodes{
		ctx:         context,
		client:      client,
		serviceCode: serviceCode,
	}
}

type ServiceQuotaCodes struct {
	client      ServiceQuotaClient
	serviceCode string
	ctx         context.Context

	once       sync.Once
	onceErr    error
	quotaCodes map[string]string
}

var ErrNotFound = errors.New("Resource was not found")

func (s *ServiceQuotaCodes) ServiceCode() string {
	return s.serviceCode
}

func (s *ServiceQuotaCodes) Get(serviceQuotaName string) (string, error) {
	s.once.Do(s.loadAllCodes)

	if s.onceErr != nil {
		return "", s.onceErr
	}

	quotaCode, ok := s.quotaCodes[serviceQuotaName]
	if !ok {
		return "", fmt.Errorf("getting service quota code for service quota name %q: %w", serviceQuotaName, ErrNotFound)
	}
	return quotaCode, nil
}

func (s *ServiceQuotaCodes) loadAllCodes() {
	s.quotaCodes = map[string]string{}

	params := &servicequotas.ListServiceQuotasInput{
		ServiceCode: &s.serviceCode,
	}
	paginator := servicequotas.NewListServiceQuotasPaginator(s.client, params, func(o *servicequotas.ListServiceQuotasPaginatorOptions) {
		o.Limit = pageSize
	})

	for paginator.HasMorePages() {
		o, err := paginator.NextPage(s.ctx)
		if err != nil {
			s.onceErr = err
			return
		}
		for _, v := range o.Quotas {
			s.quotaCodes[*v.QuotaName] = *v.QuotaCode
		}
	}
}
