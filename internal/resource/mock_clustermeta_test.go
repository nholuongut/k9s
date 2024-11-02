// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/nholuongut/k9s/internal/resource (interfaces: ClusterMeta)

package resource_test

import (
	k8s "github.com/nholuongut/k9s/internal/k8s"
	pegomock "github.com/petergtz/pegomock"
	v1 "k8s.io/api/core/v1"
	version "k8s.io/apimachinery/pkg/version"
	dynamic "k8s.io/client-go/dynamic"
	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	versioned "k8s.io/metrics/pkg/client/clientset/versioned"
	"reflect"
	"time"
)

type MockClusterMeta struct {
	fail func(message string, callerSkip ...int)
}

func NewMockClusterMeta(options ...pegomock.Option) *MockClusterMeta {
	mock := &MockClusterMeta{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockClusterMeta) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockClusterMeta) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockClusterMeta) CanIAccess(_param0 string, _param1 string, _param2 []string) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CanIAccess", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterMeta) CheckListNSAccess() error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CheckListNSAccess", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) CheckNSAccess(_param0 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CheckNSAccess", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) ClusterName() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ClusterName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) Config() *k8s.Config {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Config", params, []reflect.Type{reflect.TypeOf((**k8s.Config)(nil)).Elem()})
	var ret0 *k8s.Config
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*k8s.Config)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) ContextName() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ContextName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) CurrentNamespaceName() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentNamespaceName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterMeta) DialOrDie() kubernetes.Interface {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DialOrDie", params, []reflect.Type{reflect.TypeOf((*kubernetes.Interface)(nil)).Elem()})
	var ret0 kubernetes.Interface
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(kubernetes.Interface)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) DynDialOrDie() dynamic.Interface {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DynDialOrDie", params, []reflect.Type{reflect.TypeOf((*dynamic.Interface)(nil)).Elem()})
	var ret0 dynamic.Interface
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(dynamic.Interface)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) FetchNodes() (*v1.NodeList, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("FetchNodes", params, []reflect.Type{reflect.TypeOf((**v1.NodeList)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *v1.NodeList
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*v1.NodeList)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterMeta) GetNodes() (*v1.NodeList, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetNodes", params, []reflect.Type{reflect.TypeOf((**v1.NodeList)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *v1.NodeList
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*v1.NodeList)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterMeta) HasMetrics() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("HasMetrics", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) IsNamespaced(_param0 string) bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("IsNamespaced", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) MXDial() (*versioned.Clientset, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("MXDial", params, []reflect.Type{reflect.TypeOf((**versioned.Clientset)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *versioned.Clientset
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*versioned.Clientset)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterMeta) NSDialOrDie() dynamic.NamespaceableResourceInterface {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("NSDialOrDie", params, []reflect.Type{reflect.TypeOf((*dynamic.NamespaceableResourceInterface)(nil)).Elem()})
	var ret0 dynamic.NamespaceableResourceInterface
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(dynamic.NamespaceableResourceInterface)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) NodePods(_param0 string) (*v1.PodList, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("NodePods", params, []reflect.Type{reflect.TypeOf((**v1.PodList)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *v1.PodList
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*v1.PodList)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterMeta) RestConfigOrDie() *rest.Config {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("RestConfigOrDie", params, []reflect.Type{reflect.TypeOf((**rest.Config)(nil)).Elem()})
	var ret0 *rest.Config
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*rest.Config)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) ServerVersion() (*version.Info, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ServerVersion", params, []reflect.Type{reflect.TypeOf((**version.Info)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *version.Info
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*version.Info)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterMeta) SupportsRes(_param0 string, _param1 []string) (string, bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("SupportsRes", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 bool
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(bool)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockClusterMeta) SupportsResource(_param0 string) bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("SupportsResource", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) SwitchContextOrDie(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SwitchContextOrDie", params, []reflect.Type{})
}

func (mock *MockClusterMeta) UserName() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UserName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) ValidNamespaces() ([]v1.Namespace, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ValidNamespaces", params, []reflect.Type{reflect.TypeOf((*[]v1.Namespace)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []v1.Namespace
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]v1.Namespace)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterMeta) Version() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Version", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterMeta) VerifyWasCalledOnce() *VerifierMockClusterMeta {
	return &VerifierMockClusterMeta{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockClusterMeta) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockClusterMeta {
	return &VerifierMockClusterMeta{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockClusterMeta) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockClusterMeta {
	return &VerifierMockClusterMeta{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockClusterMeta) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockClusterMeta {
	return &VerifierMockClusterMeta{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockClusterMeta struct {
	mock                   *MockClusterMeta
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockClusterMeta) CanIAccess(_param0 string, _param1 string, _param2 []string) *MockClusterMeta_CanIAccess_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CanIAccess", params, verifier.timeout)
	return &MockClusterMeta_CanIAccess_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_CanIAccess_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_CanIAccess_OngoingVerification) GetCapturedArguments() (string, string, []string) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockClusterMeta_CanIAccess_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 [][]string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([][]string, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.([]string)
		}
	}
	return
}

func (verifier *VerifierMockClusterMeta) CheckListNSAccess() *MockClusterMeta_CheckListNSAccess_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CheckListNSAccess", params, verifier.timeout)
	return &MockClusterMeta_CheckListNSAccess_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_CheckListNSAccess_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_CheckListNSAccess_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_CheckListNSAccess_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) CheckNSAccess(_param0 string) *MockClusterMeta_CheckNSAccess_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CheckNSAccess", params, verifier.timeout)
	return &MockClusterMeta_CheckNSAccess_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_CheckNSAccess_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_CheckNSAccess_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockClusterMeta_CheckNSAccess_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClusterMeta) ClusterName() *MockClusterMeta_ClusterName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ClusterName", params, verifier.timeout)
	return &MockClusterMeta_ClusterName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_ClusterName_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_ClusterName_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_ClusterName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) Config() *MockClusterMeta_Config_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Config", params, verifier.timeout)
	return &MockClusterMeta_Config_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_Config_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_Config_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_Config_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) ContextName() *MockClusterMeta_ContextName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ContextName", params, verifier.timeout)
	return &MockClusterMeta_ContextName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_ContextName_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_ContextName_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_ContextName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) CurrentNamespaceName() *MockClusterMeta_CurrentNamespaceName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentNamespaceName", params, verifier.timeout)
	return &MockClusterMeta_CurrentNamespaceName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_CurrentNamespaceName_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_CurrentNamespaceName_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_CurrentNamespaceName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) DialOrDie() *MockClusterMeta_DialOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DialOrDie", params, verifier.timeout)
	return &MockClusterMeta_DialOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_DialOrDie_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_DialOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_DialOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) DynDialOrDie() *MockClusterMeta_DynDialOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DynDialOrDie", params, verifier.timeout)
	return &MockClusterMeta_DynDialOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_DynDialOrDie_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_DynDialOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_DynDialOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) FetchNodes() *MockClusterMeta_FetchNodes_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "FetchNodes", params, verifier.timeout)
	return &MockClusterMeta_FetchNodes_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_FetchNodes_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_FetchNodes_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_FetchNodes_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) GetNodes() *MockClusterMeta_GetNodes_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetNodes", params, verifier.timeout)
	return &MockClusterMeta_GetNodes_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_GetNodes_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_GetNodes_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_GetNodes_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) HasMetrics() *MockClusterMeta_HasMetrics_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "HasMetrics", params, verifier.timeout)
	return &MockClusterMeta_HasMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_HasMetrics_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_HasMetrics_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_HasMetrics_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) IsNamespaced(_param0 string) *MockClusterMeta_IsNamespaced_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "IsNamespaced", params, verifier.timeout)
	return &MockClusterMeta_IsNamespaced_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_IsNamespaced_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_IsNamespaced_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockClusterMeta_IsNamespaced_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClusterMeta) MXDial() *MockClusterMeta_MXDial_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "MXDial", params, verifier.timeout)
	return &MockClusterMeta_MXDial_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_MXDial_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_MXDial_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_MXDial_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) NSDialOrDie() *MockClusterMeta_NSDialOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NSDialOrDie", params, verifier.timeout)
	return &MockClusterMeta_NSDialOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_NSDialOrDie_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_NSDialOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_NSDialOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) NodePods(_param0 string) *MockClusterMeta_NodePods_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NodePods", params, verifier.timeout)
	return &MockClusterMeta_NodePods_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_NodePods_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_NodePods_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockClusterMeta_NodePods_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClusterMeta) RestConfigOrDie() *MockClusterMeta_RestConfigOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RestConfigOrDie", params, verifier.timeout)
	return &MockClusterMeta_RestConfigOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_RestConfigOrDie_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_RestConfigOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_RestConfigOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) ServerVersion() *MockClusterMeta_ServerVersion_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ServerVersion", params, verifier.timeout)
	return &MockClusterMeta_ServerVersion_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_ServerVersion_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_ServerVersion_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_ServerVersion_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) SupportsRes(_param0 string, _param1 []string) *MockClusterMeta_SupportsRes_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SupportsRes", params, verifier.timeout)
	return &MockClusterMeta_SupportsRes_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_SupportsRes_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_SupportsRes_OngoingVerification) GetCapturedArguments() (string, []string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockClusterMeta_SupportsRes_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 [][]string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([][]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.([]string)
		}
	}
	return
}

func (verifier *VerifierMockClusterMeta) SupportsResource(_param0 string) *MockClusterMeta_SupportsResource_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SupportsResource", params, verifier.timeout)
	return &MockClusterMeta_SupportsResource_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_SupportsResource_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_SupportsResource_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockClusterMeta_SupportsResource_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClusterMeta) SwitchContextOrDie(_param0 string) *MockClusterMeta_SwitchContextOrDie_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SwitchContextOrDie", params, verifier.timeout)
	return &MockClusterMeta_SwitchContextOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_SwitchContextOrDie_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_SwitchContextOrDie_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockClusterMeta_SwitchContextOrDie_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClusterMeta) UserName() *MockClusterMeta_UserName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UserName", params, verifier.timeout)
	return &MockClusterMeta_UserName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_UserName_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_UserName_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_UserName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) ValidNamespaces() *MockClusterMeta_ValidNamespaces_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ValidNamespaces", params, verifier.timeout)
	return &MockClusterMeta_ValidNamespaces_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_ValidNamespaces_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_ValidNamespaces_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_ValidNamespaces_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockClusterMeta) Version() *MockClusterMeta_Version_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Version", params, verifier.timeout)
	return &MockClusterMeta_Version_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClusterMeta_Version_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClusterMeta_Version_OngoingVerification) GetCapturedArguments() {
}

func (c *MockClusterMeta_Version_OngoingVerification) GetAllCapturedArguments() {
}
