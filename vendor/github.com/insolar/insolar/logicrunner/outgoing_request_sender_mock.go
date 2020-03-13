package logicrunner

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
)

// OutgoingRequestSenderMock implements OutgoingRequestSender
type OutgoingRequestSenderMock struct {
	t minimock.Tester

	funcSendAbandonedOutgoingRequest          func(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest)
	inspectFuncSendAbandonedOutgoingRequest   func(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest)
	afterSendAbandonedOutgoingRequestCounter  uint64
	beforeSendAbandonedOutgoingRequestCounter uint64
	SendAbandonedOutgoingRequestMock          mOutgoingRequestSenderMockSendAbandonedOutgoingRequest

	funcSendOutgoingRequest          func(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest) (a1 insolar.Arguments, ip1 *record.IncomingRequest, err error)
	inspectFuncSendOutgoingRequest   func(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest)
	afterSendOutgoingRequestCounter  uint64
	beforeSendOutgoingRequestCounter uint64
	SendOutgoingRequestMock          mOutgoingRequestSenderMockSendOutgoingRequest

	funcStop          func(ctx context.Context)
	inspectFuncStop   func(ctx context.Context)
	afterStopCounter  uint64
	beforeStopCounter uint64
	StopMock          mOutgoingRequestSenderMockStop
}

// NewOutgoingRequestSenderMock returns a mock for OutgoingRequestSender
func NewOutgoingRequestSenderMock(t minimock.Tester) *OutgoingRequestSenderMock {
	m := &OutgoingRequestSenderMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SendAbandonedOutgoingRequestMock = mOutgoingRequestSenderMockSendAbandonedOutgoingRequest{mock: m}
	m.SendAbandonedOutgoingRequestMock.callArgs = []*OutgoingRequestSenderMockSendAbandonedOutgoingRequestParams{}

	m.SendOutgoingRequestMock = mOutgoingRequestSenderMockSendOutgoingRequest{mock: m}
	m.SendOutgoingRequestMock.callArgs = []*OutgoingRequestSenderMockSendOutgoingRequestParams{}

	m.StopMock = mOutgoingRequestSenderMockStop{mock: m}
	m.StopMock.callArgs = []*OutgoingRequestSenderMockStopParams{}

	return m
}

type mOutgoingRequestSenderMockSendAbandonedOutgoingRequest struct {
	mock               *OutgoingRequestSenderMock
	defaultExpectation *OutgoingRequestSenderMockSendAbandonedOutgoingRequestExpectation
	expectations       []*OutgoingRequestSenderMockSendAbandonedOutgoingRequestExpectation

	callArgs []*OutgoingRequestSenderMockSendAbandonedOutgoingRequestParams
	mutex    sync.RWMutex
}

// OutgoingRequestSenderMockSendAbandonedOutgoingRequestExpectation specifies expectation struct of the OutgoingRequestSender.SendAbandonedOutgoingRequest
type OutgoingRequestSenderMockSendAbandonedOutgoingRequestExpectation struct {
	mock   *OutgoingRequestSenderMock
	params *OutgoingRequestSenderMockSendAbandonedOutgoingRequestParams

	Counter uint64
}

// OutgoingRequestSenderMockSendAbandonedOutgoingRequestParams contains parameters of the OutgoingRequestSender.SendAbandonedOutgoingRequest
type OutgoingRequestSenderMockSendAbandonedOutgoingRequestParams struct {
	ctx    context.Context
	reqRef insolar.Reference
	req    *record.OutgoingRequest
}

// Expect sets up expected params for OutgoingRequestSender.SendAbandonedOutgoingRequest
func (mmSendAbandonedOutgoingRequest *mOutgoingRequestSenderMockSendAbandonedOutgoingRequest) Expect(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest) *mOutgoingRequestSenderMockSendAbandonedOutgoingRequest {
	if mmSendAbandonedOutgoingRequest.mock.funcSendAbandonedOutgoingRequest != nil {
		mmSendAbandonedOutgoingRequest.mock.t.Fatalf("OutgoingRequestSenderMock.SendAbandonedOutgoingRequest mock is already set by Set")
	}

	if mmSendAbandonedOutgoingRequest.defaultExpectation == nil {
		mmSendAbandonedOutgoingRequest.defaultExpectation = &OutgoingRequestSenderMockSendAbandonedOutgoingRequestExpectation{}
	}

	mmSendAbandonedOutgoingRequest.defaultExpectation.params = &OutgoingRequestSenderMockSendAbandonedOutgoingRequestParams{ctx, reqRef, req}
	for _, e := range mmSendAbandonedOutgoingRequest.expectations {
		if minimock.Equal(e.params, mmSendAbandonedOutgoingRequest.defaultExpectation.params) {
			mmSendAbandonedOutgoingRequest.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendAbandonedOutgoingRequest.defaultExpectation.params)
		}
	}

	return mmSendAbandonedOutgoingRequest
}

// Inspect accepts an inspector function that has same arguments as the OutgoingRequestSender.SendAbandonedOutgoingRequest
func (mmSendAbandonedOutgoingRequest *mOutgoingRequestSenderMockSendAbandonedOutgoingRequest) Inspect(f func(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest)) *mOutgoingRequestSenderMockSendAbandonedOutgoingRequest {
	if mmSendAbandonedOutgoingRequest.mock.inspectFuncSendAbandonedOutgoingRequest != nil {
		mmSendAbandonedOutgoingRequest.mock.t.Fatalf("Inspect function is already set for OutgoingRequestSenderMock.SendAbandonedOutgoingRequest")
	}

	mmSendAbandonedOutgoingRequest.mock.inspectFuncSendAbandonedOutgoingRequest = f

	return mmSendAbandonedOutgoingRequest
}

// Return sets up results that will be returned by OutgoingRequestSender.SendAbandonedOutgoingRequest
func (mmSendAbandonedOutgoingRequest *mOutgoingRequestSenderMockSendAbandonedOutgoingRequest) Return() *OutgoingRequestSenderMock {
	if mmSendAbandonedOutgoingRequest.mock.funcSendAbandonedOutgoingRequest != nil {
		mmSendAbandonedOutgoingRequest.mock.t.Fatalf("OutgoingRequestSenderMock.SendAbandonedOutgoingRequest mock is already set by Set")
	}

	if mmSendAbandonedOutgoingRequest.defaultExpectation == nil {
		mmSendAbandonedOutgoingRequest.defaultExpectation = &OutgoingRequestSenderMockSendAbandonedOutgoingRequestExpectation{mock: mmSendAbandonedOutgoingRequest.mock}
	}

	return mmSendAbandonedOutgoingRequest.mock
}

//Set uses given function f to mock the OutgoingRequestSender.SendAbandonedOutgoingRequest method
func (mmSendAbandonedOutgoingRequest *mOutgoingRequestSenderMockSendAbandonedOutgoingRequest) Set(f func(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest)) *OutgoingRequestSenderMock {
	if mmSendAbandonedOutgoingRequest.defaultExpectation != nil {
		mmSendAbandonedOutgoingRequest.mock.t.Fatalf("Default expectation is already set for the OutgoingRequestSender.SendAbandonedOutgoingRequest method")
	}

	if len(mmSendAbandonedOutgoingRequest.expectations) > 0 {
		mmSendAbandonedOutgoingRequest.mock.t.Fatalf("Some expectations are already set for the OutgoingRequestSender.SendAbandonedOutgoingRequest method")
	}

	mmSendAbandonedOutgoingRequest.mock.funcSendAbandonedOutgoingRequest = f
	return mmSendAbandonedOutgoingRequest.mock
}

// SendAbandonedOutgoingRequest implements OutgoingRequestSender
func (mmSendAbandonedOutgoingRequest *OutgoingRequestSenderMock) SendAbandonedOutgoingRequest(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest) {
	mm_atomic.AddUint64(&mmSendAbandonedOutgoingRequest.beforeSendAbandonedOutgoingRequestCounter, 1)
	defer mm_atomic.AddUint64(&mmSendAbandonedOutgoingRequest.afterSendAbandonedOutgoingRequestCounter, 1)

	if mmSendAbandonedOutgoingRequest.inspectFuncSendAbandonedOutgoingRequest != nil {
		mmSendAbandonedOutgoingRequest.inspectFuncSendAbandonedOutgoingRequest(ctx, reqRef, req)
	}

	mm_params := &OutgoingRequestSenderMockSendAbandonedOutgoingRequestParams{ctx, reqRef, req}

	// Record call args
	mmSendAbandonedOutgoingRequest.SendAbandonedOutgoingRequestMock.mutex.Lock()
	mmSendAbandonedOutgoingRequest.SendAbandonedOutgoingRequestMock.callArgs = append(mmSendAbandonedOutgoingRequest.SendAbandonedOutgoingRequestMock.callArgs, mm_params)
	mmSendAbandonedOutgoingRequest.SendAbandonedOutgoingRequestMock.mutex.Unlock()

	for _, e := range mmSendAbandonedOutgoingRequest.SendAbandonedOutgoingRequestMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmSendAbandonedOutgoingRequest.SendAbandonedOutgoingRequestMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendAbandonedOutgoingRequest.SendAbandonedOutgoingRequestMock.defaultExpectation.Counter, 1)
		mm_want := mmSendAbandonedOutgoingRequest.SendAbandonedOutgoingRequestMock.defaultExpectation.params
		mm_got := OutgoingRequestSenderMockSendAbandonedOutgoingRequestParams{ctx, reqRef, req}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendAbandonedOutgoingRequest.t.Errorf("OutgoingRequestSenderMock.SendAbandonedOutgoingRequest got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmSendAbandonedOutgoingRequest.funcSendAbandonedOutgoingRequest != nil {
		mmSendAbandonedOutgoingRequest.funcSendAbandonedOutgoingRequest(ctx, reqRef, req)
		return
	}
	mmSendAbandonedOutgoingRequest.t.Fatalf("Unexpected call to OutgoingRequestSenderMock.SendAbandonedOutgoingRequest. %v %v %v", ctx, reqRef, req)

}

// SendAbandonedOutgoingRequestAfterCounter returns a count of finished OutgoingRequestSenderMock.SendAbandonedOutgoingRequest invocations
func (mmSendAbandonedOutgoingRequest *OutgoingRequestSenderMock) SendAbandonedOutgoingRequestAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendAbandonedOutgoingRequest.afterSendAbandonedOutgoingRequestCounter)
}

// SendAbandonedOutgoingRequestBeforeCounter returns a count of OutgoingRequestSenderMock.SendAbandonedOutgoingRequest invocations
func (mmSendAbandonedOutgoingRequest *OutgoingRequestSenderMock) SendAbandonedOutgoingRequestBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendAbandonedOutgoingRequest.beforeSendAbandonedOutgoingRequestCounter)
}

// Calls returns a list of arguments used in each call to OutgoingRequestSenderMock.SendAbandonedOutgoingRequest.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendAbandonedOutgoingRequest *mOutgoingRequestSenderMockSendAbandonedOutgoingRequest) Calls() []*OutgoingRequestSenderMockSendAbandonedOutgoingRequestParams {
	mmSendAbandonedOutgoingRequest.mutex.RLock()

	argCopy := make([]*OutgoingRequestSenderMockSendAbandonedOutgoingRequestParams, len(mmSendAbandonedOutgoingRequest.callArgs))
	copy(argCopy, mmSendAbandonedOutgoingRequest.callArgs)

	mmSendAbandonedOutgoingRequest.mutex.RUnlock()

	return argCopy
}

// MinimockSendAbandonedOutgoingRequestDone returns true if the count of the SendAbandonedOutgoingRequest invocations corresponds
// the number of defined expectations
func (m *OutgoingRequestSenderMock) MinimockSendAbandonedOutgoingRequestDone() bool {
	for _, e := range m.SendAbandonedOutgoingRequestMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendAbandonedOutgoingRequestMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendAbandonedOutgoingRequestCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendAbandonedOutgoingRequest != nil && mm_atomic.LoadUint64(&m.afterSendAbandonedOutgoingRequestCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendAbandonedOutgoingRequestInspect logs each unmet expectation
func (m *OutgoingRequestSenderMock) MinimockSendAbandonedOutgoingRequestInspect() {
	for _, e := range m.SendAbandonedOutgoingRequestMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OutgoingRequestSenderMock.SendAbandonedOutgoingRequest with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendAbandonedOutgoingRequestMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendAbandonedOutgoingRequestCounter) < 1 {
		if m.SendAbandonedOutgoingRequestMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to OutgoingRequestSenderMock.SendAbandonedOutgoingRequest")
		} else {
			m.t.Errorf("Expected call to OutgoingRequestSenderMock.SendAbandonedOutgoingRequest with params: %#v", *m.SendAbandonedOutgoingRequestMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendAbandonedOutgoingRequest != nil && mm_atomic.LoadUint64(&m.afterSendAbandonedOutgoingRequestCounter) < 1 {
		m.t.Error("Expected call to OutgoingRequestSenderMock.SendAbandonedOutgoingRequest")
	}
}

type mOutgoingRequestSenderMockSendOutgoingRequest struct {
	mock               *OutgoingRequestSenderMock
	defaultExpectation *OutgoingRequestSenderMockSendOutgoingRequestExpectation
	expectations       []*OutgoingRequestSenderMockSendOutgoingRequestExpectation

	callArgs []*OutgoingRequestSenderMockSendOutgoingRequestParams
	mutex    sync.RWMutex
}

// OutgoingRequestSenderMockSendOutgoingRequestExpectation specifies expectation struct of the OutgoingRequestSender.SendOutgoingRequest
type OutgoingRequestSenderMockSendOutgoingRequestExpectation struct {
	mock    *OutgoingRequestSenderMock
	params  *OutgoingRequestSenderMockSendOutgoingRequestParams
	results *OutgoingRequestSenderMockSendOutgoingRequestResults
	Counter uint64
}

// OutgoingRequestSenderMockSendOutgoingRequestParams contains parameters of the OutgoingRequestSender.SendOutgoingRequest
type OutgoingRequestSenderMockSendOutgoingRequestParams struct {
	ctx    context.Context
	reqRef insolar.Reference
	req    *record.OutgoingRequest
}

// OutgoingRequestSenderMockSendOutgoingRequestResults contains results of the OutgoingRequestSender.SendOutgoingRequest
type OutgoingRequestSenderMockSendOutgoingRequestResults struct {
	a1  insolar.Arguments
	ip1 *record.IncomingRequest
	err error
}

// Expect sets up expected params for OutgoingRequestSender.SendOutgoingRequest
func (mmSendOutgoingRequest *mOutgoingRequestSenderMockSendOutgoingRequest) Expect(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest) *mOutgoingRequestSenderMockSendOutgoingRequest {
	if mmSendOutgoingRequest.mock.funcSendOutgoingRequest != nil {
		mmSendOutgoingRequest.mock.t.Fatalf("OutgoingRequestSenderMock.SendOutgoingRequest mock is already set by Set")
	}

	if mmSendOutgoingRequest.defaultExpectation == nil {
		mmSendOutgoingRequest.defaultExpectation = &OutgoingRequestSenderMockSendOutgoingRequestExpectation{}
	}

	mmSendOutgoingRequest.defaultExpectation.params = &OutgoingRequestSenderMockSendOutgoingRequestParams{ctx, reqRef, req}
	for _, e := range mmSendOutgoingRequest.expectations {
		if minimock.Equal(e.params, mmSendOutgoingRequest.defaultExpectation.params) {
			mmSendOutgoingRequest.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendOutgoingRequest.defaultExpectation.params)
		}
	}

	return mmSendOutgoingRequest
}

// Inspect accepts an inspector function that has same arguments as the OutgoingRequestSender.SendOutgoingRequest
func (mmSendOutgoingRequest *mOutgoingRequestSenderMockSendOutgoingRequest) Inspect(f func(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest)) *mOutgoingRequestSenderMockSendOutgoingRequest {
	if mmSendOutgoingRequest.mock.inspectFuncSendOutgoingRequest != nil {
		mmSendOutgoingRequest.mock.t.Fatalf("Inspect function is already set for OutgoingRequestSenderMock.SendOutgoingRequest")
	}

	mmSendOutgoingRequest.mock.inspectFuncSendOutgoingRequest = f

	return mmSendOutgoingRequest
}

// Return sets up results that will be returned by OutgoingRequestSender.SendOutgoingRequest
func (mmSendOutgoingRequest *mOutgoingRequestSenderMockSendOutgoingRequest) Return(a1 insolar.Arguments, ip1 *record.IncomingRequest, err error) *OutgoingRequestSenderMock {
	if mmSendOutgoingRequest.mock.funcSendOutgoingRequest != nil {
		mmSendOutgoingRequest.mock.t.Fatalf("OutgoingRequestSenderMock.SendOutgoingRequest mock is already set by Set")
	}

	if mmSendOutgoingRequest.defaultExpectation == nil {
		mmSendOutgoingRequest.defaultExpectation = &OutgoingRequestSenderMockSendOutgoingRequestExpectation{mock: mmSendOutgoingRequest.mock}
	}
	mmSendOutgoingRequest.defaultExpectation.results = &OutgoingRequestSenderMockSendOutgoingRequestResults{a1, ip1, err}
	return mmSendOutgoingRequest.mock
}

//Set uses given function f to mock the OutgoingRequestSender.SendOutgoingRequest method
func (mmSendOutgoingRequest *mOutgoingRequestSenderMockSendOutgoingRequest) Set(f func(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest) (a1 insolar.Arguments, ip1 *record.IncomingRequest, err error)) *OutgoingRequestSenderMock {
	if mmSendOutgoingRequest.defaultExpectation != nil {
		mmSendOutgoingRequest.mock.t.Fatalf("Default expectation is already set for the OutgoingRequestSender.SendOutgoingRequest method")
	}

	if len(mmSendOutgoingRequest.expectations) > 0 {
		mmSendOutgoingRequest.mock.t.Fatalf("Some expectations are already set for the OutgoingRequestSender.SendOutgoingRequest method")
	}

	mmSendOutgoingRequest.mock.funcSendOutgoingRequest = f
	return mmSendOutgoingRequest.mock
}

// When sets expectation for the OutgoingRequestSender.SendOutgoingRequest which will trigger the result defined by the following
// Then helper
func (mmSendOutgoingRequest *mOutgoingRequestSenderMockSendOutgoingRequest) When(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest) *OutgoingRequestSenderMockSendOutgoingRequestExpectation {
	if mmSendOutgoingRequest.mock.funcSendOutgoingRequest != nil {
		mmSendOutgoingRequest.mock.t.Fatalf("OutgoingRequestSenderMock.SendOutgoingRequest mock is already set by Set")
	}

	expectation := &OutgoingRequestSenderMockSendOutgoingRequestExpectation{
		mock:   mmSendOutgoingRequest.mock,
		params: &OutgoingRequestSenderMockSendOutgoingRequestParams{ctx, reqRef, req},
	}
	mmSendOutgoingRequest.expectations = append(mmSendOutgoingRequest.expectations, expectation)
	return expectation
}

// Then sets up OutgoingRequestSender.SendOutgoingRequest return parameters for the expectation previously defined by the When method
func (e *OutgoingRequestSenderMockSendOutgoingRequestExpectation) Then(a1 insolar.Arguments, ip1 *record.IncomingRequest, err error) *OutgoingRequestSenderMock {
	e.results = &OutgoingRequestSenderMockSendOutgoingRequestResults{a1, ip1, err}
	return e.mock
}

// SendOutgoingRequest implements OutgoingRequestSender
func (mmSendOutgoingRequest *OutgoingRequestSenderMock) SendOutgoingRequest(ctx context.Context, reqRef insolar.Reference, req *record.OutgoingRequest) (a1 insolar.Arguments, ip1 *record.IncomingRequest, err error) {
	mm_atomic.AddUint64(&mmSendOutgoingRequest.beforeSendOutgoingRequestCounter, 1)
	defer mm_atomic.AddUint64(&mmSendOutgoingRequest.afterSendOutgoingRequestCounter, 1)

	if mmSendOutgoingRequest.inspectFuncSendOutgoingRequest != nil {
		mmSendOutgoingRequest.inspectFuncSendOutgoingRequest(ctx, reqRef, req)
	}

	mm_params := &OutgoingRequestSenderMockSendOutgoingRequestParams{ctx, reqRef, req}

	// Record call args
	mmSendOutgoingRequest.SendOutgoingRequestMock.mutex.Lock()
	mmSendOutgoingRequest.SendOutgoingRequestMock.callArgs = append(mmSendOutgoingRequest.SendOutgoingRequestMock.callArgs, mm_params)
	mmSendOutgoingRequest.SendOutgoingRequestMock.mutex.Unlock()

	for _, e := range mmSendOutgoingRequest.SendOutgoingRequestMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.a1, e.results.ip1, e.results.err
		}
	}

	if mmSendOutgoingRequest.SendOutgoingRequestMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendOutgoingRequest.SendOutgoingRequestMock.defaultExpectation.Counter, 1)
		mm_want := mmSendOutgoingRequest.SendOutgoingRequestMock.defaultExpectation.params
		mm_got := OutgoingRequestSenderMockSendOutgoingRequestParams{ctx, reqRef, req}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendOutgoingRequest.t.Errorf("OutgoingRequestSenderMock.SendOutgoingRequest got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendOutgoingRequest.SendOutgoingRequestMock.defaultExpectation.results
		if mm_results == nil {
			mmSendOutgoingRequest.t.Fatal("No results are set for the OutgoingRequestSenderMock.SendOutgoingRequest")
		}
		return (*mm_results).a1, (*mm_results).ip1, (*mm_results).err
	}
	if mmSendOutgoingRequest.funcSendOutgoingRequest != nil {
		return mmSendOutgoingRequest.funcSendOutgoingRequest(ctx, reqRef, req)
	}
	mmSendOutgoingRequest.t.Fatalf("Unexpected call to OutgoingRequestSenderMock.SendOutgoingRequest. %v %v %v", ctx, reqRef, req)
	return
}

// SendOutgoingRequestAfterCounter returns a count of finished OutgoingRequestSenderMock.SendOutgoingRequest invocations
func (mmSendOutgoingRequest *OutgoingRequestSenderMock) SendOutgoingRequestAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendOutgoingRequest.afterSendOutgoingRequestCounter)
}

// SendOutgoingRequestBeforeCounter returns a count of OutgoingRequestSenderMock.SendOutgoingRequest invocations
func (mmSendOutgoingRequest *OutgoingRequestSenderMock) SendOutgoingRequestBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendOutgoingRequest.beforeSendOutgoingRequestCounter)
}

// Calls returns a list of arguments used in each call to OutgoingRequestSenderMock.SendOutgoingRequest.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendOutgoingRequest *mOutgoingRequestSenderMockSendOutgoingRequest) Calls() []*OutgoingRequestSenderMockSendOutgoingRequestParams {
	mmSendOutgoingRequest.mutex.RLock()

	argCopy := make([]*OutgoingRequestSenderMockSendOutgoingRequestParams, len(mmSendOutgoingRequest.callArgs))
	copy(argCopy, mmSendOutgoingRequest.callArgs)

	mmSendOutgoingRequest.mutex.RUnlock()

	return argCopy
}

// MinimockSendOutgoingRequestDone returns true if the count of the SendOutgoingRequest invocations corresponds
// the number of defined expectations
func (m *OutgoingRequestSenderMock) MinimockSendOutgoingRequestDone() bool {
	for _, e := range m.SendOutgoingRequestMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendOutgoingRequestMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendOutgoingRequestCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendOutgoingRequest != nil && mm_atomic.LoadUint64(&m.afterSendOutgoingRequestCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendOutgoingRequestInspect logs each unmet expectation
func (m *OutgoingRequestSenderMock) MinimockSendOutgoingRequestInspect() {
	for _, e := range m.SendOutgoingRequestMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OutgoingRequestSenderMock.SendOutgoingRequest with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendOutgoingRequestMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendOutgoingRequestCounter) < 1 {
		if m.SendOutgoingRequestMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to OutgoingRequestSenderMock.SendOutgoingRequest")
		} else {
			m.t.Errorf("Expected call to OutgoingRequestSenderMock.SendOutgoingRequest with params: %#v", *m.SendOutgoingRequestMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendOutgoingRequest != nil && mm_atomic.LoadUint64(&m.afterSendOutgoingRequestCounter) < 1 {
		m.t.Error("Expected call to OutgoingRequestSenderMock.SendOutgoingRequest")
	}
}

type mOutgoingRequestSenderMockStop struct {
	mock               *OutgoingRequestSenderMock
	defaultExpectation *OutgoingRequestSenderMockStopExpectation
	expectations       []*OutgoingRequestSenderMockStopExpectation

	callArgs []*OutgoingRequestSenderMockStopParams
	mutex    sync.RWMutex
}

// OutgoingRequestSenderMockStopExpectation specifies expectation struct of the OutgoingRequestSender.Stop
type OutgoingRequestSenderMockStopExpectation struct {
	mock   *OutgoingRequestSenderMock
	params *OutgoingRequestSenderMockStopParams

	Counter uint64
}

// OutgoingRequestSenderMockStopParams contains parameters of the OutgoingRequestSender.Stop
type OutgoingRequestSenderMockStopParams struct {
	ctx context.Context
}

// Expect sets up expected params for OutgoingRequestSender.Stop
func (mmStop *mOutgoingRequestSenderMockStop) Expect(ctx context.Context) *mOutgoingRequestSenderMockStop {
	if mmStop.mock.funcStop != nil {
		mmStop.mock.t.Fatalf("OutgoingRequestSenderMock.Stop mock is already set by Set")
	}

	if mmStop.defaultExpectation == nil {
		mmStop.defaultExpectation = &OutgoingRequestSenderMockStopExpectation{}
	}

	mmStop.defaultExpectation.params = &OutgoingRequestSenderMockStopParams{ctx}
	for _, e := range mmStop.expectations {
		if minimock.Equal(e.params, mmStop.defaultExpectation.params) {
			mmStop.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmStop.defaultExpectation.params)
		}
	}

	return mmStop
}

// Inspect accepts an inspector function that has same arguments as the OutgoingRequestSender.Stop
func (mmStop *mOutgoingRequestSenderMockStop) Inspect(f func(ctx context.Context)) *mOutgoingRequestSenderMockStop {
	if mmStop.mock.inspectFuncStop != nil {
		mmStop.mock.t.Fatalf("Inspect function is already set for OutgoingRequestSenderMock.Stop")
	}

	mmStop.mock.inspectFuncStop = f

	return mmStop
}

// Return sets up results that will be returned by OutgoingRequestSender.Stop
func (mmStop *mOutgoingRequestSenderMockStop) Return() *OutgoingRequestSenderMock {
	if mmStop.mock.funcStop != nil {
		mmStop.mock.t.Fatalf("OutgoingRequestSenderMock.Stop mock is already set by Set")
	}

	if mmStop.defaultExpectation == nil {
		mmStop.defaultExpectation = &OutgoingRequestSenderMockStopExpectation{mock: mmStop.mock}
	}

	return mmStop.mock
}

//Set uses given function f to mock the OutgoingRequestSender.Stop method
func (mmStop *mOutgoingRequestSenderMockStop) Set(f func(ctx context.Context)) *OutgoingRequestSenderMock {
	if mmStop.defaultExpectation != nil {
		mmStop.mock.t.Fatalf("Default expectation is already set for the OutgoingRequestSender.Stop method")
	}

	if len(mmStop.expectations) > 0 {
		mmStop.mock.t.Fatalf("Some expectations are already set for the OutgoingRequestSender.Stop method")
	}

	mmStop.mock.funcStop = f
	return mmStop.mock
}

// Stop implements OutgoingRequestSender
func (mmStop *OutgoingRequestSenderMock) Stop(ctx context.Context) {
	mm_atomic.AddUint64(&mmStop.beforeStopCounter, 1)
	defer mm_atomic.AddUint64(&mmStop.afterStopCounter, 1)

	if mmStop.inspectFuncStop != nil {
		mmStop.inspectFuncStop(ctx)
	}

	mm_params := &OutgoingRequestSenderMockStopParams{ctx}

	// Record call args
	mmStop.StopMock.mutex.Lock()
	mmStop.StopMock.callArgs = append(mmStop.StopMock.callArgs, mm_params)
	mmStop.StopMock.mutex.Unlock()

	for _, e := range mmStop.StopMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmStop.StopMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmStop.StopMock.defaultExpectation.Counter, 1)
		mm_want := mmStop.StopMock.defaultExpectation.params
		mm_got := OutgoingRequestSenderMockStopParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmStop.t.Errorf("OutgoingRequestSenderMock.Stop got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmStop.funcStop != nil {
		mmStop.funcStop(ctx)
		return
	}
	mmStop.t.Fatalf("Unexpected call to OutgoingRequestSenderMock.Stop. %v", ctx)

}

// StopAfterCounter returns a count of finished OutgoingRequestSenderMock.Stop invocations
func (mmStop *OutgoingRequestSenderMock) StopAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStop.afterStopCounter)
}

// StopBeforeCounter returns a count of OutgoingRequestSenderMock.Stop invocations
func (mmStop *OutgoingRequestSenderMock) StopBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStop.beforeStopCounter)
}

// Calls returns a list of arguments used in each call to OutgoingRequestSenderMock.Stop.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmStop *mOutgoingRequestSenderMockStop) Calls() []*OutgoingRequestSenderMockStopParams {
	mmStop.mutex.RLock()

	argCopy := make([]*OutgoingRequestSenderMockStopParams, len(mmStop.callArgs))
	copy(argCopy, mmStop.callArgs)

	mmStop.mutex.RUnlock()

	return argCopy
}

// MinimockStopDone returns true if the count of the Stop invocations corresponds
// the number of defined expectations
func (m *OutgoingRequestSenderMock) MinimockStopDone() bool {
	for _, e := range m.StopMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StopMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStopCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStop != nil && mm_atomic.LoadUint64(&m.afterStopCounter) < 1 {
		return false
	}
	return true
}

// MinimockStopInspect logs each unmet expectation
func (m *OutgoingRequestSenderMock) MinimockStopInspect() {
	for _, e := range m.StopMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OutgoingRequestSenderMock.Stop with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StopMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStopCounter) < 1 {
		if m.StopMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to OutgoingRequestSenderMock.Stop")
		} else {
			m.t.Errorf("Expected call to OutgoingRequestSenderMock.Stop with params: %#v", *m.StopMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStop != nil && mm_atomic.LoadUint64(&m.afterStopCounter) < 1 {
		m.t.Error("Expected call to OutgoingRequestSenderMock.Stop")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *OutgoingRequestSenderMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockSendAbandonedOutgoingRequestInspect()

		m.MinimockSendOutgoingRequestInspect()

		m.MinimockStopInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *OutgoingRequestSenderMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *OutgoingRequestSenderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSendAbandonedOutgoingRequestDone() &&
		m.MinimockSendOutgoingRequestDone() &&
		m.MinimockStopDone()
}
