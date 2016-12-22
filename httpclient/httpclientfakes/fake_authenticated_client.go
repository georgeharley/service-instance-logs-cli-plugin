// This file was generated by counterfeiter
package httpclientfakes

import (
	"bytes"
	"sync"

	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/httpclient"
)

type FakeAuthenticatedClient struct {
	DoAuthenticatedGetStub        func(url string, accessToken string) (*bytes.Buffer, error)
	doAuthenticatedGetMutex       sync.RWMutex
	doAuthenticatedGetArgsForCall []struct {
		url         string
		accessToken string
	}
	doAuthenticatedGetReturns struct {
		result1 *bytes.Buffer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedGet(url string, accessToken string) (*bytes.Buffer, error) {
	fake.doAuthenticatedGetMutex.Lock()
	fake.doAuthenticatedGetArgsForCall = append(fake.doAuthenticatedGetArgsForCall, struct {
		url         string
		accessToken string
	}{url, accessToken})
	fake.recordInvocation("DoAuthenticatedGet", []interface{}{url, accessToken})
	fake.doAuthenticatedGetMutex.Unlock()
	if fake.DoAuthenticatedGetStub != nil {
		return fake.DoAuthenticatedGetStub(url, accessToken)
	} else {
		return fake.doAuthenticatedGetReturns.result1, fake.doAuthenticatedGetReturns.result2
	}
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedGetCallCount() int {
	fake.doAuthenticatedGetMutex.RLock()
	defer fake.doAuthenticatedGetMutex.RUnlock()
	return len(fake.doAuthenticatedGetArgsForCall)
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedGetArgsForCall(i int) (string, string) {
	fake.doAuthenticatedGetMutex.RLock()
	defer fake.doAuthenticatedGetMutex.RUnlock()
	return fake.doAuthenticatedGetArgsForCall[i].url, fake.doAuthenticatedGetArgsForCall[i].accessToken
}

func (fake *FakeAuthenticatedClient) DoAuthenticatedGetReturns(result1 *bytes.Buffer, result2 error) {
	fake.DoAuthenticatedGetStub = nil
	fake.doAuthenticatedGetReturns = struct {
		result1 *bytes.Buffer
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthenticatedClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doAuthenticatedGetMutex.RLock()
	defer fake.doAuthenticatedGetMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAuthenticatedClient) recordInvocation(key string, args []interface{}) {
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

var _ httpclient.AuthenticatedClient = new(FakeAuthenticatedClient)