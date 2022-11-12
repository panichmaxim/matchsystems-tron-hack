package tron

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i gitlab.com/rubin-dev/api/pkg/neo4jstore.CategoryRisk -o ./category_risk_mock_test.go -n CategoryRiskMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"gitlab.com/rubin-dev/api/pkg/models"
)

// CategoryRiskMock implements neo4jstore.CategoryRisk
type CategoryRiskMock struct {
	t minimock.Tester

	funcCategoryFindByName          func(ctx context.Context, name string) (cp1 *models.Category, err error)
	inspectFuncCategoryFindByName   func(ctx context.Context, name string)
	afterCategoryFindByNameCounter  uint64
	beforeCategoryFindByNameCounter uint64
	CategoryFindByNameMock          mCategoryRiskMockCategoryFindByName

	funcCategoryFindByNumber          func(ctx context.Context, id int) (cp1 *models.Category, err error)
	inspectFuncCategoryFindByNumber   func(ctx context.Context, id int)
	afterCategoryFindByNumberCounter  uint64
	beforeCategoryFindByNumberCounter uint64
	CategoryFindByNumberMock          mCategoryRiskMockCategoryFindByNumber
}

// NewCategoryRiskMock returns a mock for neo4jstore.CategoryRisk
func NewCategoryRiskMock(t minimock.Tester) *CategoryRiskMock {
	m := &CategoryRiskMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CategoryFindByNameMock = mCategoryRiskMockCategoryFindByName{mock: m}
	m.CategoryFindByNameMock.callArgs = []*CategoryRiskMockCategoryFindByNameParams{}

	m.CategoryFindByNumberMock = mCategoryRiskMockCategoryFindByNumber{mock: m}
	m.CategoryFindByNumberMock.callArgs = []*CategoryRiskMockCategoryFindByNumberParams{}

	return m
}

type mCategoryRiskMockCategoryFindByName struct {
	mock               *CategoryRiskMock
	defaultExpectation *CategoryRiskMockCategoryFindByNameExpectation
	expectations       []*CategoryRiskMockCategoryFindByNameExpectation

	callArgs []*CategoryRiskMockCategoryFindByNameParams
	mutex    sync.RWMutex
}

// CategoryRiskMockCategoryFindByNameExpectation specifies expectation struct of the CategoryRisk.CategoryFindByName
type CategoryRiskMockCategoryFindByNameExpectation struct {
	mock    *CategoryRiskMock
	params  *CategoryRiskMockCategoryFindByNameParams
	results *CategoryRiskMockCategoryFindByNameResults
	Counter uint64
}

// CategoryRiskMockCategoryFindByNameParams contains parameters of the CategoryRisk.CategoryFindByName
type CategoryRiskMockCategoryFindByNameParams struct {
	ctx  context.Context
	name string
}

// CategoryRiskMockCategoryFindByNameResults contains results of the CategoryRisk.CategoryFindByName
type CategoryRiskMockCategoryFindByNameResults struct {
	cp1 *models.Category
	err error
}

// Expect sets up expected params for CategoryRisk.CategoryFindByName
func (mmCategoryFindByName *mCategoryRiskMockCategoryFindByName) Expect(ctx context.Context, name string) *mCategoryRiskMockCategoryFindByName {
	if mmCategoryFindByName.mock.funcCategoryFindByName != nil {
		mmCategoryFindByName.mock.t.Fatalf("CategoryRiskMock.CategoryFindByName mock is already set by Set")
	}

	if mmCategoryFindByName.defaultExpectation == nil {
		mmCategoryFindByName.defaultExpectation = &CategoryRiskMockCategoryFindByNameExpectation{}
	}

	mmCategoryFindByName.defaultExpectation.params = &CategoryRiskMockCategoryFindByNameParams{ctx, name}
	for _, e := range mmCategoryFindByName.expectations {
		if minimock.Equal(e.params, mmCategoryFindByName.defaultExpectation.params) {
			mmCategoryFindByName.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCategoryFindByName.defaultExpectation.params)
		}
	}

	return mmCategoryFindByName
}

// Inspect accepts an inspector function that has same arguments as the CategoryRisk.CategoryFindByName
func (mmCategoryFindByName *mCategoryRiskMockCategoryFindByName) Inspect(f func(ctx context.Context, name string)) *mCategoryRiskMockCategoryFindByName {
	if mmCategoryFindByName.mock.inspectFuncCategoryFindByName != nil {
		mmCategoryFindByName.mock.t.Fatalf("Inspect function is already set for CategoryRiskMock.CategoryFindByName")
	}

	mmCategoryFindByName.mock.inspectFuncCategoryFindByName = f

	return mmCategoryFindByName
}

// Return sets up results that will be returned by CategoryRisk.CategoryFindByName
func (mmCategoryFindByName *mCategoryRiskMockCategoryFindByName) Return(cp1 *models.Category, err error) *CategoryRiskMock {
	if mmCategoryFindByName.mock.funcCategoryFindByName != nil {
		mmCategoryFindByName.mock.t.Fatalf("CategoryRiskMock.CategoryFindByName mock is already set by Set")
	}

	if mmCategoryFindByName.defaultExpectation == nil {
		mmCategoryFindByName.defaultExpectation = &CategoryRiskMockCategoryFindByNameExpectation{mock: mmCategoryFindByName.mock}
	}
	mmCategoryFindByName.defaultExpectation.results = &CategoryRiskMockCategoryFindByNameResults{cp1, err}
	return mmCategoryFindByName.mock
}

// Set uses given function f to mock the CategoryRisk.CategoryFindByName method
func (mmCategoryFindByName *mCategoryRiskMockCategoryFindByName) Set(f func(ctx context.Context, name string) (cp1 *models.Category, err error)) *CategoryRiskMock {
	if mmCategoryFindByName.defaultExpectation != nil {
		mmCategoryFindByName.mock.t.Fatalf("Default expectation is already set for the CategoryRisk.CategoryFindByName method")
	}

	if len(mmCategoryFindByName.expectations) > 0 {
		mmCategoryFindByName.mock.t.Fatalf("Some expectations are already set for the CategoryRisk.CategoryFindByName method")
	}

	mmCategoryFindByName.mock.funcCategoryFindByName = f
	return mmCategoryFindByName.mock
}

// When sets expectation for the CategoryRisk.CategoryFindByName which will trigger the result defined by the following
// Then helper
func (mmCategoryFindByName *mCategoryRiskMockCategoryFindByName) When(ctx context.Context, name string) *CategoryRiskMockCategoryFindByNameExpectation {
	if mmCategoryFindByName.mock.funcCategoryFindByName != nil {
		mmCategoryFindByName.mock.t.Fatalf("CategoryRiskMock.CategoryFindByName mock is already set by Set")
	}

	expectation := &CategoryRiskMockCategoryFindByNameExpectation{
		mock:   mmCategoryFindByName.mock,
		params: &CategoryRiskMockCategoryFindByNameParams{ctx, name},
	}
	mmCategoryFindByName.expectations = append(mmCategoryFindByName.expectations, expectation)
	return expectation
}

// Then sets up CategoryRisk.CategoryFindByName return parameters for the expectation previously defined by the When method
func (e *CategoryRiskMockCategoryFindByNameExpectation) Then(cp1 *models.Category, err error) *CategoryRiskMock {
	e.results = &CategoryRiskMockCategoryFindByNameResults{cp1, err}
	return e.mock
}

// CategoryFindByName implements neo4jstore.CategoryRisk
func (mmCategoryFindByName *CategoryRiskMock) CategoryFindByName(ctx context.Context, name string) (cp1 *models.Category, err error) {
	mm_atomic.AddUint64(&mmCategoryFindByName.beforeCategoryFindByNameCounter, 1)
	defer mm_atomic.AddUint64(&mmCategoryFindByName.afterCategoryFindByNameCounter, 1)

	if mmCategoryFindByName.inspectFuncCategoryFindByName != nil {
		mmCategoryFindByName.inspectFuncCategoryFindByName(ctx, name)
	}

	mm_params := &CategoryRiskMockCategoryFindByNameParams{ctx, name}

	// Record call args
	mmCategoryFindByName.CategoryFindByNameMock.mutex.Lock()
	mmCategoryFindByName.CategoryFindByNameMock.callArgs = append(mmCategoryFindByName.CategoryFindByNameMock.callArgs, mm_params)
	mmCategoryFindByName.CategoryFindByNameMock.mutex.Unlock()

	for _, e := range mmCategoryFindByName.CategoryFindByNameMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.cp1, e.results.err
		}
	}

	if mmCategoryFindByName.CategoryFindByNameMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCategoryFindByName.CategoryFindByNameMock.defaultExpectation.Counter, 1)
		mm_want := mmCategoryFindByName.CategoryFindByNameMock.defaultExpectation.params
		mm_got := CategoryRiskMockCategoryFindByNameParams{ctx, name}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCategoryFindByName.t.Errorf("CategoryRiskMock.CategoryFindByName got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCategoryFindByName.CategoryFindByNameMock.defaultExpectation.results
		if mm_results == nil {
			mmCategoryFindByName.t.Fatal("No results are set for the CategoryRiskMock.CategoryFindByName")
		}
		return (*mm_results).cp1, (*mm_results).err
	}
	if mmCategoryFindByName.funcCategoryFindByName != nil {
		return mmCategoryFindByName.funcCategoryFindByName(ctx, name)
	}
	mmCategoryFindByName.t.Fatalf("Unexpected call to CategoryRiskMock.CategoryFindByName. %v %v", ctx, name)
	return
}

// CategoryFindByNameAfterCounter returns a count of finished CategoryRiskMock.CategoryFindByName invocations
func (mmCategoryFindByName *CategoryRiskMock) CategoryFindByNameAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCategoryFindByName.afterCategoryFindByNameCounter)
}

// CategoryFindByNameBeforeCounter returns a count of CategoryRiskMock.CategoryFindByName invocations
func (mmCategoryFindByName *CategoryRiskMock) CategoryFindByNameBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCategoryFindByName.beforeCategoryFindByNameCounter)
}

// Calls returns a list of arguments used in each call to CategoryRiskMock.CategoryFindByName.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCategoryFindByName *mCategoryRiskMockCategoryFindByName) Calls() []*CategoryRiskMockCategoryFindByNameParams {
	mmCategoryFindByName.mutex.RLock()

	argCopy := make([]*CategoryRiskMockCategoryFindByNameParams, len(mmCategoryFindByName.callArgs))
	copy(argCopy, mmCategoryFindByName.callArgs)

	mmCategoryFindByName.mutex.RUnlock()

	return argCopy
}

// MinimockCategoryFindByNameDone returns true if the count of the CategoryFindByName invocations corresponds
// the number of defined expectations
func (m *CategoryRiskMock) MinimockCategoryFindByNameDone() bool {
	for _, e := range m.CategoryFindByNameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CategoryFindByNameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCategoryFindByNameCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCategoryFindByName != nil && mm_atomic.LoadUint64(&m.afterCategoryFindByNameCounter) < 1 {
		return false
	}
	return true
}

// MinimockCategoryFindByNameInspect logs each unmet expectation
func (m *CategoryRiskMock) MinimockCategoryFindByNameInspect() {
	for _, e := range m.CategoryFindByNameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CategoryRiskMock.CategoryFindByName with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CategoryFindByNameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCategoryFindByNameCounter) < 1 {
		if m.CategoryFindByNameMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CategoryRiskMock.CategoryFindByName")
		} else {
			m.t.Errorf("Expected call to CategoryRiskMock.CategoryFindByName with params: %#v", *m.CategoryFindByNameMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCategoryFindByName != nil && mm_atomic.LoadUint64(&m.afterCategoryFindByNameCounter) < 1 {
		m.t.Error("Expected call to CategoryRiskMock.CategoryFindByName")
	}
}

type mCategoryRiskMockCategoryFindByNumber struct {
	mock               *CategoryRiskMock
	defaultExpectation *CategoryRiskMockCategoryFindByNumberExpectation
	expectations       []*CategoryRiskMockCategoryFindByNumberExpectation

	callArgs []*CategoryRiskMockCategoryFindByNumberParams
	mutex    sync.RWMutex
}

// CategoryRiskMockCategoryFindByNumberExpectation specifies expectation struct of the CategoryRisk.CategoryFindByNumber
type CategoryRiskMockCategoryFindByNumberExpectation struct {
	mock    *CategoryRiskMock
	params  *CategoryRiskMockCategoryFindByNumberParams
	results *CategoryRiskMockCategoryFindByNumberResults
	Counter uint64
}

// CategoryRiskMockCategoryFindByNumberParams contains parameters of the CategoryRisk.CategoryFindByNumber
type CategoryRiskMockCategoryFindByNumberParams struct {
	ctx context.Context
	id  int
}

// CategoryRiskMockCategoryFindByNumberResults contains results of the CategoryRisk.CategoryFindByNumber
type CategoryRiskMockCategoryFindByNumberResults struct {
	cp1 *models.Category
	err error
}

// Expect sets up expected params for CategoryRisk.CategoryFindByNumber
func (mmCategoryFindByNumber *mCategoryRiskMockCategoryFindByNumber) Expect(ctx context.Context, id int) *mCategoryRiskMockCategoryFindByNumber {
	if mmCategoryFindByNumber.mock.funcCategoryFindByNumber != nil {
		mmCategoryFindByNumber.mock.t.Fatalf("CategoryRiskMock.CategoryFindByNumber mock is already set by Set")
	}

	if mmCategoryFindByNumber.defaultExpectation == nil {
		mmCategoryFindByNumber.defaultExpectation = &CategoryRiskMockCategoryFindByNumberExpectation{}
	}

	mmCategoryFindByNumber.defaultExpectation.params = &CategoryRiskMockCategoryFindByNumberParams{ctx, id}
	for _, e := range mmCategoryFindByNumber.expectations {
		if minimock.Equal(e.params, mmCategoryFindByNumber.defaultExpectation.params) {
			mmCategoryFindByNumber.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCategoryFindByNumber.defaultExpectation.params)
		}
	}

	return mmCategoryFindByNumber
}

// Inspect accepts an inspector function that has same arguments as the CategoryRisk.CategoryFindByNumber
func (mmCategoryFindByNumber *mCategoryRiskMockCategoryFindByNumber) Inspect(f func(ctx context.Context, id int)) *mCategoryRiskMockCategoryFindByNumber {
	if mmCategoryFindByNumber.mock.inspectFuncCategoryFindByNumber != nil {
		mmCategoryFindByNumber.mock.t.Fatalf("Inspect function is already set for CategoryRiskMock.CategoryFindByNumber")
	}

	mmCategoryFindByNumber.mock.inspectFuncCategoryFindByNumber = f

	return mmCategoryFindByNumber
}

// Return sets up results that will be returned by CategoryRisk.CategoryFindByNumber
func (mmCategoryFindByNumber *mCategoryRiskMockCategoryFindByNumber) Return(cp1 *models.Category, err error) *CategoryRiskMock {
	if mmCategoryFindByNumber.mock.funcCategoryFindByNumber != nil {
		mmCategoryFindByNumber.mock.t.Fatalf("CategoryRiskMock.CategoryFindByNumber mock is already set by Set")
	}

	if mmCategoryFindByNumber.defaultExpectation == nil {
		mmCategoryFindByNumber.defaultExpectation = &CategoryRiskMockCategoryFindByNumberExpectation{mock: mmCategoryFindByNumber.mock}
	}
	mmCategoryFindByNumber.defaultExpectation.results = &CategoryRiskMockCategoryFindByNumberResults{cp1, err}
	return mmCategoryFindByNumber.mock
}

// Set uses given function f to mock the CategoryRisk.CategoryFindByNumber method
func (mmCategoryFindByNumber *mCategoryRiskMockCategoryFindByNumber) Set(f func(ctx context.Context, id int) (cp1 *models.Category, err error)) *CategoryRiskMock {
	if mmCategoryFindByNumber.defaultExpectation != nil {
		mmCategoryFindByNumber.mock.t.Fatalf("Default expectation is already set for the CategoryRisk.CategoryFindByNumber method")
	}

	if len(mmCategoryFindByNumber.expectations) > 0 {
		mmCategoryFindByNumber.mock.t.Fatalf("Some expectations are already set for the CategoryRisk.CategoryFindByNumber method")
	}

	mmCategoryFindByNumber.mock.funcCategoryFindByNumber = f
	return mmCategoryFindByNumber.mock
}

// When sets expectation for the CategoryRisk.CategoryFindByNumber which will trigger the result defined by the following
// Then helper
func (mmCategoryFindByNumber *mCategoryRiskMockCategoryFindByNumber) When(ctx context.Context, id int) *CategoryRiskMockCategoryFindByNumberExpectation {
	if mmCategoryFindByNumber.mock.funcCategoryFindByNumber != nil {
		mmCategoryFindByNumber.mock.t.Fatalf("CategoryRiskMock.CategoryFindByNumber mock is already set by Set")
	}

	expectation := &CategoryRiskMockCategoryFindByNumberExpectation{
		mock:   mmCategoryFindByNumber.mock,
		params: &CategoryRiskMockCategoryFindByNumberParams{ctx, id},
	}
	mmCategoryFindByNumber.expectations = append(mmCategoryFindByNumber.expectations, expectation)
	return expectation
}

// Then sets up CategoryRisk.CategoryFindByNumber return parameters for the expectation previously defined by the When method
func (e *CategoryRiskMockCategoryFindByNumberExpectation) Then(cp1 *models.Category, err error) *CategoryRiskMock {
	e.results = &CategoryRiskMockCategoryFindByNumberResults{cp1, err}
	return e.mock
}

// CategoryFindByNumber implements neo4jstore.CategoryRisk
func (mmCategoryFindByNumber *CategoryRiskMock) CategoryFindByNumber(ctx context.Context, id int) (cp1 *models.Category, err error) {
	mm_atomic.AddUint64(&mmCategoryFindByNumber.beforeCategoryFindByNumberCounter, 1)
	defer mm_atomic.AddUint64(&mmCategoryFindByNumber.afterCategoryFindByNumberCounter, 1)

	if mmCategoryFindByNumber.inspectFuncCategoryFindByNumber != nil {
		mmCategoryFindByNumber.inspectFuncCategoryFindByNumber(ctx, id)
	}

	mm_params := &CategoryRiskMockCategoryFindByNumberParams{ctx, id}

	// Record call args
	mmCategoryFindByNumber.CategoryFindByNumberMock.mutex.Lock()
	mmCategoryFindByNumber.CategoryFindByNumberMock.callArgs = append(mmCategoryFindByNumber.CategoryFindByNumberMock.callArgs, mm_params)
	mmCategoryFindByNumber.CategoryFindByNumberMock.mutex.Unlock()

	for _, e := range mmCategoryFindByNumber.CategoryFindByNumberMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.cp1, e.results.err
		}
	}

	if mmCategoryFindByNumber.CategoryFindByNumberMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCategoryFindByNumber.CategoryFindByNumberMock.defaultExpectation.Counter, 1)
		mm_want := mmCategoryFindByNumber.CategoryFindByNumberMock.defaultExpectation.params
		mm_got := CategoryRiskMockCategoryFindByNumberParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCategoryFindByNumber.t.Errorf("CategoryRiskMock.CategoryFindByNumber got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCategoryFindByNumber.CategoryFindByNumberMock.defaultExpectation.results
		if mm_results == nil {
			mmCategoryFindByNumber.t.Fatal("No results are set for the CategoryRiskMock.CategoryFindByNumber")
		}
		return (*mm_results).cp1, (*mm_results).err
	}
	if mmCategoryFindByNumber.funcCategoryFindByNumber != nil {
		return mmCategoryFindByNumber.funcCategoryFindByNumber(ctx, id)
	}
	mmCategoryFindByNumber.t.Fatalf("Unexpected call to CategoryRiskMock.CategoryFindByNumber. %v %v", ctx, id)
	return
}

// CategoryFindByNumberAfterCounter returns a count of finished CategoryRiskMock.CategoryFindByNumber invocations
func (mmCategoryFindByNumber *CategoryRiskMock) CategoryFindByNumberAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCategoryFindByNumber.afterCategoryFindByNumberCounter)
}

// CategoryFindByNumberBeforeCounter returns a count of CategoryRiskMock.CategoryFindByNumber invocations
func (mmCategoryFindByNumber *CategoryRiskMock) CategoryFindByNumberBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCategoryFindByNumber.beforeCategoryFindByNumberCounter)
}

// Calls returns a list of arguments used in each call to CategoryRiskMock.CategoryFindByNumber.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCategoryFindByNumber *mCategoryRiskMockCategoryFindByNumber) Calls() []*CategoryRiskMockCategoryFindByNumberParams {
	mmCategoryFindByNumber.mutex.RLock()

	argCopy := make([]*CategoryRiskMockCategoryFindByNumberParams, len(mmCategoryFindByNumber.callArgs))
	copy(argCopy, mmCategoryFindByNumber.callArgs)

	mmCategoryFindByNumber.mutex.RUnlock()

	return argCopy
}

// MinimockCategoryFindByNumberDone returns true if the count of the CategoryFindByNumber invocations corresponds
// the number of defined expectations
func (m *CategoryRiskMock) MinimockCategoryFindByNumberDone() bool {
	for _, e := range m.CategoryFindByNumberMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CategoryFindByNumberMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCategoryFindByNumberCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCategoryFindByNumber != nil && mm_atomic.LoadUint64(&m.afterCategoryFindByNumberCounter) < 1 {
		return false
	}
	return true
}

// MinimockCategoryFindByNumberInspect logs each unmet expectation
func (m *CategoryRiskMock) MinimockCategoryFindByNumberInspect() {
	for _, e := range m.CategoryFindByNumberMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CategoryRiskMock.CategoryFindByNumber with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CategoryFindByNumberMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCategoryFindByNumberCounter) < 1 {
		if m.CategoryFindByNumberMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CategoryRiskMock.CategoryFindByNumber")
		} else {
			m.t.Errorf("Expected call to CategoryRiskMock.CategoryFindByNumber with params: %#v", *m.CategoryFindByNumberMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCategoryFindByNumber != nil && mm_atomic.LoadUint64(&m.afterCategoryFindByNumberCounter) < 1 {
		m.t.Error("Expected call to CategoryRiskMock.CategoryFindByNumber")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CategoryRiskMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCategoryFindByNameInspect()

		m.MinimockCategoryFindByNumberInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CategoryRiskMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CategoryRiskMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCategoryFindByNameDone() &&
		m.MinimockCategoryFindByNumberDone()
}
