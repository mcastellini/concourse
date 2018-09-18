// Code generated by counterfeiter. DO NOT EDIT.
package factoryfakes

import (
	"sync"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/scheduler/factory"
)

type FakeBuildFactory struct {
	CreateStub        func(atc.JobConfig, atc.ResourceConfigs, atc.VersionedResourceTypes, []db.BuildInput) (atc.Plan, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 atc.JobConfig
		arg2 atc.ResourceConfigs
		arg3 atc.VersionedResourceTypes
		arg4 []db.BuildInput
	}
	createReturns struct {
		result1 atc.Plan
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 atc.Plan
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildFactory) Create(arg1 atc.JobConfig, arg2 atc.ResourceConfigs, arg3 atc.VersionedResourceTypes, arg4 []db.BuildInput) (atc.Plan, error) {
	var arg4Copy []db.BuildInput
	if arg4 != nil {
		arg4Copy = make([]db.BuildInput, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 atc.JobConfig
		arg2 atc.ResourceConfigs
		arg3 atc.VersionedResourceTypes
		arg4 []db.BuildInput
	}{arg1, arg2, arg3, arg4Copy})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4Copy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *FakeBuildFactory) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeBuildFactory) CreateArgsForCall(i int) (atc.JobConfig, atc.ResourceConfigs, atc.VersionedResourceTypes, []db.BuildInput) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1, fake.createArgsForCall[i].arg2, fake.createArgsForCall[i].arg3, fake.createArgsForCall[i].arg4
}

func (fake *FakeBuildFactory) CreateReturns(result1 atc.Plan, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 atc.Plan
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildFactory) CreateReturnsOnCall(i int, result1 atc.Plan, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 atc.Plan
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 atc.Plan
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildFactory) recordInvocation(key string, args []interface{}) {
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

var _ factory.BuildFactory = new(FakeBuildFactory)
