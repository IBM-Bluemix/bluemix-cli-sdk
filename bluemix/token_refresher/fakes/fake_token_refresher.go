// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.ibm.com/Bluemix/bluemix-cli-sdk/bluemix/token_refresher"
)

type FakeTokenRefresher struct {
	RefreshStub        func(oldToken string) (newToken string, refreshToken string, err error)
	refreshMutex       sync.RWMutex
	refreshArgsForCall []struct {
		oldToken string
	}
	refreshReturns struct {
		result1 string
		result2 string
		result3 error
	}
}

func (fake *FakeTokenRefresher) Refresh(oldToken string) (newToken string, refreshToken string, err error) {
	fake.refreshMutex.Lock()
	fake.refreshArgsForCall = append(fake.refreshArgsForCall, struct {
		oldToken string
	}{oldToken})
	fake.refreshMutex.Unlock()
	if fake.RefreshStub != nil {
		return fake.RefreshStub(oldToken)
	} else {
		return fake.refreshReturns.result1, fake.refreshReturns.result2, fake.refreshReturns.result3
	}
}

func (fake *FakeTokenRefresher) RefreshCallCount() int {
	fake.refreshMutex.RLock()
	defer fake.refreshMutex.RUnlock()
	return len(fake.refreshArgsForCall)
}

func (fake *FakeTokenRefresher) RefreshArgsForCall(i int) string {
	fake.refreshMutex.RLock()
	defer fake.refreshMutex.RUnlock()
	return fake.refreshArgsForCall[i].oldToken
}

func (fake *FakeTokenRefresher) RefreshReturns(result1 string, result2 string, result3 error) {
	fake.RefreshStub = nil
	fake.refreshReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

var _ token_refresher.TokenRefresher = new(FakeTokenRefresher)
