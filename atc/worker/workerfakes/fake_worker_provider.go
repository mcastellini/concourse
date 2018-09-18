// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"sync"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/worker"
)

type FakeWorkerProvider struct {
	RunningWorkersStub        func(lager.Logger) ([]worker.Worker, error)
	runningWorkersMutex       sync.RWMutex
	runningWorkersArgsForCall []struct {
		arg1 lager.Logger
	}
	runningWorkersReturns struct {
		result1 []worker.Worker
		result2 error
	}
	runningWorkersReturnsOnCall map[int]struct {
		result1 []worker.Worker
		result2 error
	}
	FindWorkerForContainerStub        func(logger lager.Logger, teamID int, handle string) (worker.Worker, bool, error)
	findWorkerForContainerMutex       sync.RWMutex
	findWorkerForContainerArgsForCall []struct {
		logger lager.Logger
		teamID int
		handle string
	}
	findWorkerForContainerReturns struct {
		result1 worker.Worker
		result2 bool
		result3 error
	}
	findWorkerForContainerReturnsOnCall map[int]struct {
		result1 worker.Worker
		result2 bool
		result3 error
	}
	FindWorkerForContainerByOwnerStub        func(logger lager.Logger, teamID int, owner db.ContainerOwner) (worker.Worker, bool, error)
	findWorkerForContainerByOwnerMutex       sync.RWMutex
	findWorkerForContainerByOwnerArgsForCall []struct {
		logger lager.Logger
		teamID int
		owner  db.ContainerOwner
	}
	findWorkerForContainerByOwnerReturns struct {
		result1 worker.Worker
		result2 bool
		result3 error
	}
	findWorkerForContainerByOwnerReturnsOnCall map[int]struct {
		result1 worker.Worker
		result2 bool
		result3 error
	}
	NewGardenWorkerStub        func(logger lager.Logger, tikTok clock.Clock, savedWorker db.Worker) worker.Worker
	newGardenWorkerMutex       sync.RWMutex
	newGardenWorkerArgsForCall []struct {
		logger      lager.Logger
		tikTok      clock.Clock
		savedWorker db.Worker
	}
	newGardenWorkerReturns struct {
		result1 worker.Worker
	}
	newGardenWorkerReturnsOnCall map[int]struct {
		result1 worker.Worker
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorkerProvider) RunningWorkers(arg1 lager.Logger) ([]worker.Worker, error) {
	fake.runningWorkersMutex.Lock()
	ret, specificReturn := fake.runningWorkersReturnsOnCall[len(fake.runningWorkersArgsForCall)]
	fake.runningWorkersArgsForCall = append(fake.runningWorkersArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("RunningWorkers", []interface{}{arg1})
	fake.runningWorkersMutex.Unlock()
	if fake.RunningWorkersStub != nil {
		return fake.RunningWorkersStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runningWorkersReturns.result1, fake.runningWorkersReturns.result2
}

func (fake *FakeWorkerProvider) RunningWorkersCallCount() int {
	fake.runningWorkersMutex.RLock()
	defer fake.runningWorkersMutex.RUnlock()
	return len(fake.runningWorkersArgsForCall)
}

func (fake *FakeWorkerProvider) RunningWorkersArgsForCall(i int) lager.Logger {
	fake.runningWorkersMutex.RLock()
	defer fake.runningWorkersMutex.RUnlock()
	return fake.runningWorkersArgsForCall[i].arg1
}

func (fake *FakeWorkerProvider) RunningWorkersReturns(result1 []worker.Worker, result2 error) {
	fake.RunningWorkersStub = nil
	fake.runningWorkersReturns = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerProvider) RunningWorkersReturnsOnCall(i int, result1 []worker.Worker, result2 error) {
	fake.RunningWorkersStub = nil
	if fake.runningWorkersReturnsOnCall == nil {
		fake.runningWorkersReturnsOnCall = make(map[int]struct {
			result1 []worker.Worker
			result2 error
		})
	}
	fake.runningWorkersReturnsOnCall[i] = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerProvider) FindWorkerForContainer(logger lager.Logger, teamID int, handle string) (worker.Worker, bool, error) {
	fake.findWorkerForContainerMutex.Lock()
	ret, specificReturn := fake.findWorkerForContainerReturnsOnCall[len(fake.findWorkerForContainerArgsForCall)]
	fake.findWorkerForContainerArgsForCall = append(fake.findWorkerForContainerArgsForCall, struct {
		logger lager.Logger
		teamID int
		handle string
	}{logger, teamID, handle})
	fake.recordInvocation("FindWorkerForContainer", []interface{}{logger, teamID, handle})
	fake.findWorkerForContainerMutex.Unlock()
	if fake.FindWorkerForContainerStub != nil {
		return fake.FindWorkerForContainerStub(logger, teamID, handle)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findWorkerForContainerReturns.result1, fake.findWorkerForContainerReturns.result2, fake.findWorkerForContainerReturns.result3
}

func (fake *FakeWorkerProvider) FindWorkerForContainerCallCount() int {
	fake.findWorkerForContainerMutex.RLock()
	defer fake.findWorkerForContainerMutex.RUnlock()
	return len(fake.findWorkerForContainerArgsForCall)
}

func (fake *FakeWorkerProvider) FindWorkerForContainerArgsForCall(i int) (lager.Logger, int, string) {
	fake.findWorkerForContainerMutex.RLock()
	defer fake.findWorkerForContainerMutex.RUnlock()
	return fake.findWorkerForContainerArgsForCall[i].logger, fake.findWorkerForContainerArgsForCall[i].teamID, fake.findWorkerForContainerArgsForCall[i].handle
}

func (fake *FakeWorkerProvider) FindWorkerForContainerReturns(result1 worker.Worker, result2 bool, result3 error) {
	fake.FindWorkerForContainerStub = nil
	fake.findWorkerForContainerReturns = struct {
		result1 worker.Worker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerProvider) FindWorkerForContainerReturnsOnCall(i int, result1 worker.Worker, result2 bool, result3 error) {
	fake.FindWorkerForContainerStub = nil
	if fake.findWorkerForContainerReturnsOnCall == nil {
		fake.findWorkerForContainerReturnsOnCall = make(map[int]struct {
			result1 worker.Worker
			result2 bool
			result3 error
		})
	}
	fake.findWorkerForContainerReturnsOnCall[i] = struct {
		result1 worker.Worker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerProvider) FindWorkerForContainerByOwner(logger lager.Logger, teamID int, owner db.ContainerOwner) (worker.Worker, bool, error) {
	fake.findWorkerForContainerByOwnerMutex.Lock()
	ret, specificReturn := fake.findWorkerForContainerByOwnerReturnsOnCall[len(fake.findWorkerForContainerByOwnerArgsForCall)]
	fake.findWorkerForContainerByOwnerArgsForCall = append(fake.findWorkerForContainerByOwnerArgsForCall, struct {
		logger lager.Logger
		teamID int
		owner  db.ContainerOwner
	}{logger, teamID, owner})
	fake.recordInvocation("FindWorkerForContainerByOwner", []interface{}{logger, teamID, owner})
	fake.findWorkerForContainerByOwnerMutex.Unlock()
	if fake.FindWorkerForContainerByOwnerStub != nil {
		return fake.FindWorkerForContainerByOwnerStub(logger, teamID, owner)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findWorkerForContainerByOwnerReturns.result1, fake.findWorkerForContainerByOwnerReturns.result2, fake.findWorkerForContainerByOwnerReturns.result3
}

func (fake *FakeWorkerProvider) FindWorkerForContainerByOwnerCallCount() int {
	fake.findWorkerForContainerByOwnerMutex.RLock()
	defer fake.findWorkerForContainerByOwnerMutex.RUnlock()
	return len(fake.findWorkerForContainerByOwnerArgsForCall)
}

func (fake *FakeWorkerProvider) FindWorkerForContainerByOwnerArgsForCall(i int) (lager.Logger, int, db.ContainerOwner) {
	fake.findWorkerForContainerByOwnerMutex.RLock()
	defer fake.findWorkerForContainerByOwnerMutex.RUnlock()
	return fake.findWorkerForContainerByOwnerArgsForCall[i].logger, fake.findWorkerForContainerByOwnerArgsForCall[i].teamID, fake.findWorkerForContainerByOwnerArgsForCall[i].owner
}

func (fake *FakeWorkerProvider) FindWorkerForContainerByOwnerReturns(result1 worker.Worker, result2 bool, result3 error) {
	fake.FindWorkerForContainerByOwnerStub = nil
	fake.findWorkerForContainerByOwnerReturns = struct {
		result1 worker.Worker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerProvider) FindWorkerForContainerByOwnerReturnsOnCall(i int, result1 worker.Worker, result2 bool, result3 error) {
	fake.FindWorkerForContainerByOwnerStub = nil
	if fake.findWorkerForContainerByOwnerReturnsOnCall == nil {
		fake.findWorkerForContainerByOwnerReturnsOnCall = make(map[int]struct {
			result1 worker.Worker
			result2 bool
			result3 error
		})
	}
	fake.findWorkerForContainerByOwnerReturnsOnCall[i] = struct {
		result1 worker.Worker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerProvider) NewGardenWorker(logger lager.Logger, tikTok clock.Clock, savedWorker db.Worker) worker.Worker {
	fake.newGardenWorkerMutex.Lock()
	ret, specificReturn := fake.newGardenWorkerReturnsOnCall[len(fake.newGardenWorkerArgsForCall)]
	fake.newGardenWorkerArgsForCall = append(fake.newGardenWorkerArgsForCall, struct {
		logger      lager.Logger
		tikTok      clock.Clock
		savedWorker db.Worker
	}{logger, tikTok, savedWorker})
	fake.recordInvocation("NewGardenWorker", []interface{}{logger, tikTok, savedWorker})
	fake.newGardenWorkerMutex.Unlock()
	if fake.NewGardenWorkerStub != nil {
		return fake.NewGardenWorkerStub(logger, tikTok, savedWorker)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newGardenWorkerReturns.result1
}

func (fake *FakeWorkerProvider) NewGardenWorkerCallCount() int {
	fake.newGardenWorkerMutex.RLock()
	defer fake.newGardenWorkerMutex.RUnlock()
	return len(fake.newGardenWorkerArgsForCall)
}

func (fake *FakeWorkerProvider) NewGardenWorkerArgsForCall(i int) (lager.Logger, clock.Clock, db.Worker) {
	fake.newGardenWorkerMutex.RLock()
	defer fake.newGardenWorkerMutex.RUnlock()
	return fake.newGardenWorkerArgsForCall[i].logger, fake.newGardenWorkerArgsForCall[i].tikTok, fake.newGardenWorkerArgsForCall[i].savedWorker
}

func (fake *FakeWorkerProvider) NewGardenWorkerReturns(result1 worker.Worker) {
	fake.NewGardenWorkerStub = nil
	fake.newGardenWorkerReturns = struct {
		result1 worker.Worker
	}{result1}
}

func (fake *FakeWorkerProvider) NewGardenWorkerReturnsOnCall(i int, result1 worker.Worker) {
	fake.NewGardenWorkerStub = nil
	if fake.newGardenWorkerReturnsOnCall == nil {
		fake.newGardenWorkerReturnsOnCall = make(map[int]struct {
			result1 worker.Worker
		})
	}
	fake.newGardenWorkerReturnsOnCall[i] = struct {
		result1 worker.Worker
	}{result1}
}

func (fake *FakeWorkerProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runningWorkersMutex.RLock()
	defer fake.runningWorkersMutex.RUnlock()
	fake.findWorkerForContainerMutex.RLock()
	defer fake.findWorkerForContainerMutex.RUnlock()
	fake.findWorkerForContainerByOwnerMutex.RLock()
	defer fake.findWorkerForContainerByOwnerMutex.RUnlock()
	fake.newGardenWorkerMutex.RLock()
	defer fake.newGardenWorkerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorkerProvider) recordInvocation(key string, args []interface{}) {
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

var _ worker.WorkerProvider = new(FakeWorkerProvider)
