// Code generated by counterfeiter. DO NOT EDIT.
package pluginfakes

import (
	"sync"

	"code.cloudfoundry.org/cpu-entitlement-plugin/metadata"
	"code.cloudfoundry.org/cpu-entitlement-plugin/plugin"
	"code.cloudfoundry.org/cpu-entitlement-plugin/reporter"
)

type FakeMetricsRenderer struct {
	ShowInstanceReportsStub        func(metadata.CFAppInfo, []reporter.InstanceReport) error
	showInstanceReportsMutex       sync.RWMutex
	showInstanceReportsArgsForCall []struct {
		arg1 metadata.CFAppInfo
		arg2 []reporter.InstanceReport
	}
	showInstanceReportsReturns struct {
		result1 error
	}
	showInstanceReportsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMetricsRenderer) ShowInstanceReports(arg1 metadata.CFAppInfo, arg2 []reporter.InstanceReport) error {
	var arg2Copy []reporter.InstanceReport
	if arg2 != nil {
		arg2Copy = make([]reporter.InstanceReport, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.showInstanceReportsMutex.Lock()
	ret, specificReturn := fake.showInstanceReportsReturnsOnCall[len(fake.showInstanceReportsArgsForCall)]
	fake.showInstanceReportsArgsForCall = append(fake.showInstanceReportsArgsForCall, struct {
		arg1 metadata.CFAppInfo
		arg2 []reporter.InstanceReport
	}{arg1, arg2Copy})
	fake.recordInvocation("ShowInstanceReports", []interface{}{arg1, arg2Copy})
	fake.showInstanceReportsMutex.Unlock()
	if fake.ShowInstanceReportsStub != nil {
		return fake.ShowInstanceReportsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.showInstanceReportsReturns
	return fakeReturns.result1
}

func (fake *FakeMetricsRenderer) ShowInstanceReportsCallCount() int {
	fake.showInstanceReportsMutex.RLock()
	defer fake.showInstanceReportsMutex.RUnlock()
	return len(fake.showInstanceReportsArgsForCall)
}

func (fake *FakeMetricsRenderer) ShowInstanceReportsCalls(stub func(metadata.CFAppInfo, []reporter.InstanceReport) error) {
	fake.showInstanceReportsMutex.Lock()
	defer fake.showInstanceReportsMutex.Unlock()
	fake.ShowInstanceReportsStub = stub
}

func (fake *FakeMetricsRenderer) ShowInstanceReportsArgsForCall(i int) (metadata.CFAppInfo, []reporter.InstanceReport) {
	fake.showInstanceReportsMutex.RLock()
	defer fake.showInstanceReportsMutex.RUnlock()
	argsForCall := fake.showInstanceReportsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMetricsRenderer) ShowInstanceReportsReturns(result1 error) {
	fake.showInstanceReportsMutex.Lock()
	defer fake.showInstanceReportsMutex.Unlock()
	fake.ShowInstanceReportsStub = nil
	fake.showInstanceReportsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMetricsRenderer) ShowInstanceReportsReturnsOnCall(i int, result1 error) {
	fake.showInstanceReportsMutex.Lock()
	defer fake.showInstanceReportsMutex.Unlock()
	fake.ShowInstanceReportsStub = nil
	if fake.showInstanceReportsReturnsOnCall == nil {
		fake.showInstanceReportsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.showInstanceReportsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMetricsRenderer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.showInstanceReportsMutex.RLock()
	defer fake.showInstanceReportsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMetricsRenderer) recordInvocation(key string, args []interface{}) {
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

var _ plugin.MetricsRenderer = new(FakeMetricsRenderer)
