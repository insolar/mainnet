package testutils

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	mm_insolar "github.com/insolar/insolar/insolar"
)

// CertificateManagerMock implements insolar.CertificateManager
type CertificateManagerMock struct {
	t minimock.Tester

	funcGetCertificate          func() (c1 mm_insolar.Certificate)
	inspectFuncGetCertificate   func()
	afterGetCertificateCounter  uint64
	beforeGetCertificateCounter uint64
	GetCertificateMock          mCertificateManagerMockGetCertificate
}

// NewCertificateManagerMock returns a mock for insolar.CertificateManager
func NewCertificateManagerMock(t minimock.Tester) *CertificateManagerMock {
	m := &CertificateManagerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetCertificateMock = mCertificateManagerMockGetCertificate{mock: m}

	return m
}

type mCertificateManagerMockGetCertificate struct {
	mock               *CertificateManagerMock
	defaultExpectation *CertificateManagerMockGetCertificateExpectation
	expectations       []*CertificateManagerMockGetCertificateExpectation
}

// CertificateManagerMockGetCertificateExpectation specifies expectation struct of the CertificateManager.GetCertificate
type CertificateManagerMockGetCertificateExpectation struct {
	mock *CertificateManagerMock

	results *CertificateManagerMockGetCertificateResults
	Counter uint64
}

// CertificateManagerMockGetCertificateResults contains results of the CertificateManager.GetCertificate
type CertificateManagerMockGetCertificateResults struct {
	c1 mm_insolar.Certificate
}

// Expect sets up expected params for CertificateManager.GetCertificate
func (mmGetCertificate *mCertificateManagerMockGetCertificate) Expect() *mCertificateManagerMockGetCertificate {
	if mmGetCertificate.mock.funcGetCertificate != nil {
		mmGetCertificate.mock.t.Fatalf("CertificateManagerMock.GetCertificate mock is already set by Set")
	}

	if mmGetCertificate.defaultExpectation == nil {
		mmGetCertificate.defaultExpectation = &CertificateManagerMockGetCertificateExpectation{}
	}

	return mmGetCertificate
}

// Inspect accepts an inspector function that has same arguments as the CertificateManager.GetCertificate
func (mmGetCertificate *mCertificateManagerMockGetCertificate) Inspect(f func()) *mCertificateManagerMockGetCertificate {
	if mmGetCertificate.mock.inspectFuncGetCertificate != nil {
		mmGetCertificate.mock.t.Fatalf("Inspect function is already set for CertificateManagerMock.GetCertificate")
	}

	mmGetCertificate.mock.inspectFuncGetCertificate = f

	return mmGetCertificate
}

// Return sets up results that will be returned by CertificateManager.GetCertificate
func (mmGetCertificate *mCertificateManagerMockGetCertificate) Return(c1 mm_insolar.Certificate) *CertificateManagerMock {
	if mmGetCertificate.mock.funcGetCertificate != nil {
		mmGetCertificate.mock.t.Fatalf("CertificateManagerMock.GetCertificate mock is already set by Set")
	}

	if mmGetCertificate.defaultExpectation == nil {
		mmGetCertificate.defaultExpectation = &CertificateManagerMockGetCertificateExpectation{mock: mmGetCertificate.mock}
	}
	mmGetCertificate.defaultExpectation.results = &CertificateManagerMockGetCertificateResults{c1}
	return mmGetCertificate.mock
}

//Set uses given function f to mock the CertificateManager.GetCertificate method
func (mmGetCertificate *mCertificateManagerMockGetCertificate) Set(f func() (c1 mm_insolar.Certificate)) *CertificateManagerMock {
	if mmGetCertificate.defaultExpectation != nil {
		mmGetCertificate.mock.t.Fatalf("Default expectation is already set for the CertificateManager.GetCertificate method")
	}

	if len(mmGetCertificate.expectations) > 0 {
		mmGetCertificate.mock.t.Fatalf("Some expectations are already set for the CertificateManager.GetCertificate method")
	}

	mmGetCertificate.mock.funcGetCertificate = f
	return mmGetCertificate.mock
}

// GetCertificate implements insolar.CertificateManager
func (mmGetCertificate *CertificateManagerMock) GetCertificate() (c1 mm_insolar.Certificate) {
	mm_atomic.AddUint64(&mmGetCertificate.beforeGetCertificateCounter, 1)
	defer mm_atomic.AddUint64(&mmGetCertificate.afterGetCertificateCounter, 1)

	if mmGetCertificate.inspectFuncGetCertificate != nil {
		mmGetCertificate.inspectFuncGetCertificate()
	}

	if mmGetCertificate.GetCertificateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetCertificate.GetCertificateMock.defaultExpectation.Counter, 1)

		mm_results := mmGetCertificate.GetCertificateMock.defaultExpectation.results
		if mm_results == nil {
			mmGetCertificate.t.Fatal("No results are set for the CertificateManagerMock.GetCertificate")
		}
		return (*mm_results).c1
	}
	if mmGetCertificate.funcGetCertificate != nil {
		return mmGetCertificate.funcGetCertificate()
	}
	mmGetCertificate.t.Fatalf("Unexpected call to CertificateManagerMock.GetCertificate.")
	return
}

// GetCertificateAfterCounter returns a count of finished CertificateManagerMock.GetCertificate invocations
func (mmGetCertificate *CertificateManagerMock) GetCertificateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCertificate.afterGetCertificateCounter)
}

// GetCertificateBeforeCounter returns a count of CertificateManagerMock.GetCertificate invocations
func (mmGetCertificate *CertificateManagerMock) GetCertificateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCertificate.beforeGetCertificateCounter)
}

// MinimockGetCertificateDone returns true if the count of the GetCertificate invocations corresponds
// the number of defined expectations
func (m *CertificateManagerMock) MinimockGetCertificateDone() bool {
	for _, e := range m.GetCertificateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCertificateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCertificateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCertificate != nil && mm_atomic.LoadUint64(&m.afterGetCertificateCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetCertificateInspect logs each unmet expectation
func (m *CertificateManagerMock) MinimockGetCertificateInspect() {
	for _, e := range m.GetCertificateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CertificateManagerMock.GetCertificate")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCertificateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCertificateCounter) < 1 {
		m.t.Error("Expected call to CertificateManagerMock.GetCertificate")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCertificate != nil && mm_atomic.LoadUint64(&m.afterGetCertificateCounter) < 1 {
		m.t.Error("Expected call to CertificateManagerMock.GetCertificate")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CertificateManagerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetCertificateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CertificateManagerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CertificateManagerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetCertificateDone()
}