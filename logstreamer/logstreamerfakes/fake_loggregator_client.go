// Code generated by counterfeiter. DO NOT EDIT.
package logstreamerfakes

import (
	context "context"
	sync "sync"

	logstreamer "code.cloudfoundry.org/cpu-entitlement-plugin/logstreamer"
	loggregator "code.cloudfoundry.org/go-loggregator"
	loggregator_v2 "code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
)

type FakeLoggregatorClient struct {
	StreamStub        func(context.Context, *loggregator_v2.EgressBatchRequest) loggregator.EnvelopeStream
	streamMutex       sync.RWMutex
	streamArgsForCall []struct {
		arg1 context.Context
		arg2 *loggregator_v2.EgressBatchRequest
	}
	streamReturns struct {
		result1 loggregator.EnvelopeStream
	}
	streamReturnsOnCall map[int]struct {
		result1 loggregator.EnvelopeStream
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLoggregatorClient) Stream(arg1 context.Context, arg2 *loggregator_v2.EgressBatchRequest) loggregator.EnvelopeStream {
	fake.streamMutex.Lock()
	ret, specificReturn := fake.streamReturnsOnCall[len(fake.streamArgsForCall)]
	fake.streamArgsForCall = append(fake.streamArgsForCall, struct {
		arg1 context.Context
		arg2 *loggregator_v2.EgressBatchRequest
	}{arg1, arg2})
	fake.recordInvocation("Stream", []interface{}{arg1, arg2})
	fake.streamMutex.Unlock()
	if fake.StreamStub != nil {
		return fake.StreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.streamReturns
	return fakeReturns.result1
}

func (fake *FakeLoggregatorClient) StreamCallCount() int {
	fake.streamMutex.RLock()
	defer fake.streamMutex.RUnlock()
	return len(fake.streamArgsForCall)
}

func (fake *FakeLoggregatorClient) StreamCalls(stub func(context.Context, *loggregator_v2.EgressBatchRequest) loggregator.EnvelopeStream) {
	fake.streamMutex.Lock()
	defer fake.streamMutex.Unlock()
	fake.StreamStub = stub
}

func (fake *FakeLoggregatorClient) StreamArgsForCall(i int) (context.Context, *loggregator_v2.EgressBatchRequest) {
	fake.streamMutex.RLock()
	defer fake.streamMutex.RUnlock()
	argsForCall := fake.streamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLoggregatorClient) StreamReturns(result1 loggregator.EnvelopeStream) {
	fake.streamMutex.Lock()
	defer fake.streamMutex.Unlock()
	fake.StreamStub = nil
	fake.streamReturns = struct {
		result1 loggregator.EnvelopeStream
	}{result1}
}

func (fake *FakeLoggregatorClient) StreamReturnsOnCall(i int, result1 loggregator.EnvelopeStream) {
	fake.streamMutex.Lock()
	defer fake.streamMutex.Unlock()
	fake.StreamStub = nil
	if fake.streamReturnsOnCall == nil {
		fake.streamReturnsOnCall = make(map[int]struct {
			result1 loggregator.EnvelopeStream
		})
	}
	fake.streamReturnsOnCall[i] = struct {
		result1 loggregator.EnvelopeStream
	}{result1}
}

func (fake *FakeLoggregatorClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.streamMutex.RLock()
	defer fake.streamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLoggregatorClient) recordInvocation(key string, args []interface{}) {
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

var _ logstreamer.LoggregatorClient = new(FakeLoggregatorClient)