// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/worker"
)

type FakeContainerPlacementStrategy struct {
	ChooseStub        func([]worker.Worker, worker.ContainerSpec) (worker.Worker, error)
	chooseMutex       sync.RWMutex
	chooseArgsForCall []struct {
		arg1 []worker.Worker
		arg2 worker.ContainerSpec
	}
	chooseReturns struct {
		result1 worker.Worker
		result2 error
	}
	chooseReturnsOnCall map[int]struct {
		result1 worker.Worker
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerPlacementStrategy) Choose(arg1 []worker.Worker, arg2 worker.ContainerSpec) (worker.Worker, error) {
	var arg1Copy []worker.Worker
	if arg1 != nil {
		arg1Copy = make([]worker.Worker, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.chooseMutex.Lock()
	ret, specificReturn := fake.chooseReturnsOnCall[len(fake.chooseArgsForCall)]
	fake.chooseArgsForCall = append(fake.chooseArgsForCall, struct {
		arg1 []worker.Worker
		arg2 worker.ContainerSpec
	}{arg1Copy, arg2})
	fake.recordInvocation("Choose", []interface{}{arg1Copy, arg2})
	fake.chooseMutex.Unlock()
	if fake.ChooseStub != nil {
		return fake.ChooseStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.chooseReturns.result1, fake.chooseReturns.result2
}

func (fake *FakeContainerPlacementStrategy) ChooseCallCount() int {
	fake.chooseMutex.RLock()
	defer fake.chooseMutex.RUnlock()
	return len(fake.chooseArgsForCall)
}

func (fake *FakeContainerPlacementStrategy) ChooseArgsForCall(i int) ([]worker.Worker, worker.ContainerSpec) {
	fake.chooseMutex.RLock()
	defer fake.chooseMutex.RUnlock()
	return fake.chooseArgsForCall[i].arg1, fake.chooseArgsForCall[i].arg2
}

func (fake *FakeContainerPlacementStrategy) ChooseReturns(result1 worker.Worker, result2 error) {
	fake.ChooseStub = nil
	fake.chooseReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerPlacementStrategy) ChooseReturnsOnCall(i int, result1 worker.Worker, result2 error) {
	fake.ChooseStub = nil
	if fake.chooseReturnsOnCall == nil {
		fake.chooseReturnsOnCall = make(map[int]struct {
			result1 worker.Worker
			result2 error
		})
	}
	fake.chooseReturnsOnCall[i] = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerPlacementStrategy) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chooseMutex.RLock()
	defer fake.chooseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerPlacementStrategy) recordInvocation(key string, args []interface{}) {
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

var _ worker.ContainerPlacementStrategy = new(FakeContainerPlacementStrategy)
