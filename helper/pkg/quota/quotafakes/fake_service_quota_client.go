// Code generated by counterfeiter. DO NOT EDIT.
package quotafakes

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/vmware-tap-on-public-cloud/quickstart-vmware-tanzu-application-platform/helper/pkg/quota"
)

type FakeServiceQuotaClient struct {
	ListServiceQuotasStub        func(context.Context, *servicequotas.ListServiceQuotasInput, ...func(*servicequotas.Options)) (*servicequotas.ListServiceQuotasOutput, error)
	listServiceQuotasMutex       sync.RWMutex
	listServiceQuotasArgsForCall []struct {
		arg1 context.Context
		arg2 *servicequotas.ListServiceQuotasInput
		arg3 []func(*servicequotas.Options)
	}
	listServiceQuotasReturns struct {
		result1 *servicequotas.ListServiceQuotasOutput
		result2 error
	}
	listServiceQuotasReturnsOnCall map[int]struct {
		result1 *servicequotas.ListServiceQuotasOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceQuotaClient) ListServiceQuotas(arg1 context.Context, arg2 *servicequotas.ListServiceQuotasInput, arg3 ...func(*servicequotas.Options)) (*servicequotas.ListServiceQuotasOutput, error) {
	fake.listServiceQuotasMutex.Lock()
	ret, specificReturn := fake.listServiceQuotasReturnsOnCall[len(fake.listServiceQuotasArgsForCall)]
	fake.listServiceQuotasArgsForCall = append(fake.listServiceQuotasArgsForCall, struct {
		arg1 context.Context
		arg2 *servicequotas.ListServiceQuotasInput
		arg3 []func(*servicequotas.Options)
	}{arg1, arg2, arg3})
	stub := fake.ListServiceQuotasStub
	fakeReturns := fake.listServiceQuotasReturns
	fake.recordInvocation("ListServiceQuotas", []interface{}{arg1, arg2, arg3})
	fake.listServiceQuotasMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceQuotaClient) ListServiceQuotasCallCount() int {
	fake.listServiceQuotasMutex.RLock()
	defer fake.listServiceQuotasMutex.RUnlock()
	return len(fake.listServiceQuotasArgsForCall)
}

func (fake *FakeServiceQuotaClient) ListServiceQuotasCalls(stub func(context.Context, *servicequotas.ListServiceQuotasInput, ...func(*servicequotas.Options)) (*servicequotas.ListServiceQuotasOutput, error)) {
	fake.listServiceQuotasMutex.Lock()
	defer fake.listServiceQuotasMutex.Unlock()
	fake.ListServiceQuotasStub = stub
}

func (fake *FakeServiceQuotaClient) ListServiceQuotasArgsForCall(i int) (context.Context, *servicequotas.ListServiceQuotasInput, []func(*servicequotas.Options)) {
	fake.listServiceQuotasMutex.RLock()
	defer fake.listServiceQuotasMutex.RUnlock()
	argsForCall := fake.listServiceQuotasArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeServiceQuotaClient) ListServiceQuotasReturns(result1 *servicequotas.ListServiceQuotasOutput, result2 error) {
	fake.listServiceQuotasMutex.Lock()
	defer fake.listServiceQuotasMutex.Unlock()
	fake.ListServiceQuotasStub = nil
	fake.listServiceQuotasReturns = struct {
		result1 *servicequotas.ListServiceQuotasOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceQuotaClient) ListServiceQuotasReturnsOnCall(i int, result1 *servicequotas.ListServiceQuotasOutput, result2 error) {
	fake.listServiceQuotasMutex.Lock()
	defer fake.listServiceQuotasMutex.Unlock()
	fake.ListServiceQuotasStub = nil
	if fake.listServiceQuotasReturnsOnCall == nil {
		fake.listServiceQuotasReturnsOnCall = make(map[int]struct {
			result1 *servicequotas.ListServiceQuotasOutput
			result2 error
		})
	}
	fake.listServiceQuotasReturnsOnCall[i] = struct {
		result1 *servicequotas.ListServiceQuotasOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceQuotaClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listServiceQuotasMutex.RLock()
	defer fake.listServiceQuotasMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceQuotaClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ quota.ServiceQuotaClient = new(FakeServiceQuotaClient)
