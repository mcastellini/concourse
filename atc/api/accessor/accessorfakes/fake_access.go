// Code generated by counterfeiter. DO NOT EDIT.
package accessorfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/api/accessor"
)

type FakeAccess struct {
	IsAuthenticatedStub        func() bool
	isAuthenticatedMutex       sync.RWMutex
	isAuthenticatedArgsForCall []struct{}
	isAuthenticatedReturns     struct {
		result1 bool
	}
	isAuthenticatedReturnsOnCall map[int]struct {
		result1 bool
	}
	IsAuthorizedStub        func(string) bool
	isAuthorizedMutex       sync.RWMutex
	isAuthorizedArgsForCall []struct {
		arg1 string
	}
	isAuthorizedReturns struct {
		result1 bool
	}
	isAuthorizedReturnsOnCall map[int]struct {
		result1 bool
	}
	IsAdminStub        func() bool
	isAdminMutex       sync.RWMutex
	isAdminArgsForCall []struct{}
	isAdminReturns     struct {
		result1 bool
	}
	isAdminReturnsOnCall map[int]struct {
		result1 bool
	}
	IsSystemStub        func() bool
	isSystemMutex       sync.RWMutex
	isSystemArgsForCall []struct{}
	isSystemReturns     struct {
		result1 bool
	}
	isSystemReturnsOnCall map[int]struct {
		result1 bool
	}
	TeamNamesStub        func() []string
	teamNamesMutex       sync.RWMutex
	teamNamesArgsForCall []struct{}
	teamNamesReturns     struct {
		result1 []string
	}
	teamNamesReturnsOnCall map[int]struct {
		result1 []string
	}
	CSRFTokenStub        func() string
	cSRFTokenMutex       sync.RWMutex
	cSRFTokenArgsForCall []struct{}
	cSRFTokenReturns     struct {
		result1 string
	}
	cSRFTokenReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAccess) IsAuthenticated() bool {
	fake.isAuthenticatedMutex.Lock()
	ret, specificReturn := fake.isAuthenticatedReturnsOnCall[len(fake.isAuthenticatedArgsForCall)]
	fake.isAuthenticatedArgsForCall = append(fake.isAuthenticatedArgsForCall, struct{}{})
	fake.recordInvocation("IsAuthenticated", []interface{}{})
	fake.isAuthenticatedMutex.Unlock()
	if fake.IsAuthenticatedStub != nil {
		return fake.IsAuthenticatedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isAuthenticatedReturns.result1
}

func (fake *FakeAccess) IsAuthenticatedCallCount() int {
	fake.isAuthenticatedMutex.RLock()
	defer fake.isAuthenticatedMutex.RUnlock()
	return len(fake.isAuthenticatedArgsForCall)
}

func (fake *FakeAccess) IsAuthenticatedReturns(result1 bool) {
	fake.IsAuthenticatedStub = nil
	fake.isAuthenticatedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAccess) IsAuthenticatedReturnsOnCall(i int, result1 bool) {
	fake.IsAuthenticatedStub = nil
	if fake.isAuthenticatedReturnsOnCall == nil {
		fake.isAuthenticatedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isAuthenticatedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAccess) IsAuthorized(arg1 string) bool {
	fake.isAuthorizedMutex.Lock()
	ret, specificReturn := fake.isAuthorizedReturnsOnCall[len(fake.isAuthorizedArgsForCall)]
	fake.isAuthorizedArgsForCall = append(fake.isAuthorizedArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("IsAuthorized", []interface{}{arg1})
	fake.isAuthorizedMutex.Unlock()
	if fake.IsAuthorizedStub != nil {
		return fake.IsAuthorizedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isAuthorizedReturns.result1
}

func (fake *FakeAccess) IsAuthorizedCallCount() int {
	fake.isAuthorizedMutex.RLock()
	defer fake.isAuthorizedMutex.RUnlock()
	return len(fake.isAuthorizedArgsForCall)
}

func (fake *FakeAccess) IsAuthorizedArgsForCall(i int) string {
	fake.isAuthorizedMutex.RLock()
	defer fake.isAuthorizedMutex.RUnlock()
	return fake.isAuthorizedArgsForCall[i].arg1
}

func (fake *FakeAccess) IsAuthorizedReturns(result1 bool) {
	fake.IsAuthorizedStub = nil
	fake.isAuthorizedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAccess) IsAuthorizedReturnsOnCall(i int, result1 bool) {
	fake.IsAuthorizedStub = nil
	if fake.isAuthorizedReturnsOnCall == nil {
		fake.isAuthorizedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isAuthorizedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAccess) IsAdmin() bool {
	fake.isAdminMutex.Lock()
	ret, specificReturn := fake.isAdminReturnsOnCall[len(fake.isAdminArgsForCall)]
	fake.isAdminArgsForCall = append(fake.isAdminArgsForCall, struct{}{})
	fake.recordInvocation("IsAdmin", []interface{}{})
	fake.isAdminMutex.Unlock()
	if fake.IsAdminStub != nil {
		return fake.IsAdminStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isAdminReturns.result1
}

func (fake *FakeAccess) IsAdminCallCount() int {
	fake.isAdminMutex.RLock()
	defer fake.isAdminMutex.RUnlock()
	return len(fake.isAdminArgsForCall)
}

func (fake *FakeAccess) IsAdminReturns(result1 bool) {
	fake.IsAdminStub = nil
	fake.isAdminReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAccess) IsAdminReturnsOnCall(i int, result1 bool) {
	fake.IsAdminStub = nil
	if fake.isAdminReturnsOnCall == nil {
		fake.isAdminReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isAdminReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAccess) IsSystem() bool {
	fake.isSystemMutex.Lock()
	ret, specificReturn := fake.isSystemReturnsOnCall[len(fake.isSystemArgsForCall)]
	fake.isSystemArgsForCall = append(fake.isSystemArgsForCall, struct{}{})
	fake.recordInvocation("IsSystem", []interface{}{})
	fake.isSystemMutex.Unlock()
	if fake.IsSystemStub != nil {
		return fake.IsSystemStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isSystemReturns.result1
}

func (fake *FakeAccess) IsSystemCallCount() int {
	fake.isSystemMutex.RLock()
	defer fake.isSystemMutex.RUnlock()
	return len(fake.isSystemArgsForCall)
}

func (fake *FakeAccess) IsSystemReturns(result1 bool) {
	fake.IsSystemStub = nil
	fake.isSystemReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAccess) IsSystemReturnsOnCall(i int, result1 bool) {
	fake.IsSystemStub = nil
	if fake.isSystemReturnsOnCall == nil {
		fake.isSystemReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isSystemReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAccess) TeamNames() []string {
	fake.teamNamesMutex.Lock()
	ret, specificReturn := fake.teamNamesReturnsOnCall[len(fake.teamNamesArgsForCall)]
	fake.teamNamesArgsForCall = append(fake.teamNamesArgsForCall, struct{}{})
	fake.recordInvocation("TeamNames", []interface{}{})
	fake.teamNamesMutex.Unlock()
	if fake.TeamNamesStub != nil {
		return fake.TeamNamesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.teamNamesReturns.result1
}

func (fake *FakeAccess) TeamNamesCallCount() int {
	fake.teamNamesMutex.RLock()
	defer fake.teamNamesMutex.RUnlock()
	return len(fake.teamNamesArgsForCall)
}

func (fake *FakeAccess) TeamNamesReturns(result1 []string) {
	fake.TeamNamesStub = nil
	fake.teamNamesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeAccess) TeamNamesReturnsOnCall(i int, result1 []string) {
	fake.TeamNamesStub = nil
	if fake.teamNamesReturnsOnCall == nil {
		fake.teamNamesReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.teamNamesReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeAccess) CSRFToken() string {
	fake.cSRFTokenMutex.Lock()
	ret, specificReturn := fake.cSRFTokenReturnsOnCall[len(fake.cSRFTokenArgsForCall)]
	fake.cSRFTokenArgsForCall = append(fake.cSRFTokenArgsForCall, struct{}{})
	fake.recordInvocation("CSRFToken", []interface{}{})
	fake.cSRFTokenMutex.Unlock()
	if fake.CSRFTokenStub != nil {
		return fake.CSRFTokenStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cSRFTokenReturns.result1
}

func (fake *FakeAccess) CSRFTokenCallCount() int {
	fake.cSRFTokenMutex.RLock()
	defer fake.cSRFTokenMutex.RUnlock()
	return len(fake.cSRFTokenArgsForCall)
}

func (fake *FakeAccess) CSRFTokenReturns(result1 string) {
	fake.CSRFTokenStub = nil
	fake.cSRFTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAccess) CSRFTokenReturnsOnCall(i int, result1 string) {
	fake.CSRFTokenStub = nil
	if fake.cSRFTokenReturnsOnCall == nil {
		fake.cSRFTokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cSRFTokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeAccess) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isAuthenticatedMutex.RLock()
	defer fake.isAuthenticatedMutex.RUnlock()
	fake.isAuthorizedMutex.RLock()
	defer fake.isAuthorizedMutex.RUnlock()
	fake.isAdminMutex.RLock()
	defer fake.isAdminMutex.RUnlock()
	fake.isSystemMutex.RLock()
	defer fake.isSystemMutex.RUnlock()
	fake.teamNamesMutex.RLock()
	defer fake.teamNamesMutex.RUnlock()
	fake.cSRFTokenMutex.RLock()
	defer fake.cSRFTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAccess) recordInvocation(key string, args []interface{}) {
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

var _ accessor.Access = new(FakeAccess)
