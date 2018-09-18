// Code generated by counterfeiter. DO NOT EDIT.
package encryptionfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/db/encryption"
)

type FakeStrategy struct {
	EncryptStub        func([]byte) (string, *string, error)
	encryptMutex       sync.RWMutex
	encryptArgsForCall []struct {
		arg1 []byte
	}
	encryptReturns struct {
		result1 string
		result2 *string
		result3 error
	}
	encryptReturnsOnCall map[int]struct {
		result1 string
		result2 *string
		result3 error
	}
	DecryptStub        func(string, *string) ([]byte, error)
	decryptMutex       sync.RWMutex
	decryptArgsForCall []struct {
		arg1 string
		arg2 *string
	}
	decryptReturns struct {
		result1 []byte
		result2 error
	}
	decryptReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStrategy) Encrypt(arg1 []byte) (string, *string, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.encryptMutex.Lock()
	ret, specificReturn := fake.encryptReturnsOnCall[len(fake.encryptArgsForCall)]
	fake.encryptArgsForCall = append(fake.encryptArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Encrypt", []interface{}{arg1Copy})
	fake.encryptMutex.Unlock()
	if fake.EncryptStub != nil {
		return fake.EncryptStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.encryptReturns.result1, fake.encryptReturns.result2, fake.encryptReturns.result3
}

func (fake *FakeStrategy) EncryptCallCount() int {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	return len(fake.encryptArgsForCall)
}

func (fake *FakeStrategy) EncryptArgsForCall(i int) []byte {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	return fake.encryptArgsForCall[i].arg1
}

func (fake *FakeStrategy) EncryptReturns(result1 string, result2 *string, result3 error) {
	fake.EncryptStub = nil
	fake.encryptReturns = struct {
		result1 string
		result2 *string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStrategy) EncryptReturnsOnCall(i int, result1 string, result2 *string, result3 error) {
	fake.EncryptStub = nil
	if fake.encryptReturnsOnCall == nil {
		fake.encryptReturnsOnCall = make(map[int]struct {
			result1 string
			result2 *string
			result3 error
		})
	}
	fake.encryptReturnsOnCall[i] = struct {
		result1 string
		result2 *string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStrategy) Decrypt(arg1 string, arg2 *string) ([]byte, error) {
	fake.decryptMutex.Lock()
	ret, specificReturn := fake.decryptReturnsOnCall[len(fake.decryptArgsForCall)]
	fake.decryptArgsForCall = append(fake.decryptArgsForCall, struct {
		arg1 string
		arg2 *string
	}{arg1, arg2})
	fake.recordInvocation("Decrypt", []interface{}{arg1, arg2})
	fake.decryptMutex.Unlock()
	if fake.DecryptStub != nil {
		return fake.DecryptStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.decryptReturns.result1, fake.decryptReturns.result2
}

func (fake *FakeStrategy) DecryptCallCount() int {
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	return len(fake.decryptArgsForCall)
}

func (fake *FakeStrategy) DecryptArgsForCall(i int) (string, *string) {
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	return fake.decryptArgsForCall[i].arg1, fake.decryptArgsForCall[i].arg2
}

func (fake *FakeStrategy) DecryptReturns(result1 []byte, result2 error) {
	fake.DecryptStub = nil
	fake.decryptReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeStrategy) DecryptReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.DecryptStub = nil
	if fake.decryptReturnsOnCall == nil {
		fake.decryptReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.decryptReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeStrategy) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStrategy) recordInvocation(key string, args []interface{}) {
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

var _ encryption.Strategy = new(FakeStrategy)
