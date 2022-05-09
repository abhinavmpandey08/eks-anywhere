// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/eks-anywhere/pkg/providers/tinkerbell (interfaces: ProviderDockerClient,ProviderKubectlClient,ProviderTinkClient,ProviderPbnjClient,SSHAuthKeyGenerator)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	executables "github.com/aws/eks-anywhere/pkg/executables"
	filewriter "github.com/aws/eks-anywhere/pkg/filewriter"
	hardware "github.com/aws/eks-anywhere/pkg/providers/tinkerbell/hardware"
	pbnj "github.com/aws/eks-anywhere/pkg/providers/tinkerbell/pbnj"
	types "github.com/aws/eks-anywhere/pkg/types"
	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/mrajashree/etcdadm-controller/api/v1beta1"
	v1alpha10 "github.com/tinkerbell/cluster-api-provider-tinkerbell/tink/api/v1alpha1"
	v1 "github.com/tinkerbell/pbnj/api/v1"
	hardware0 "github.com/tinkerbell/tink/protos/hardware"
	workflow "github.com/tinkerbell/tink/protos/workflow"
	v10 "k8s.io/api/core/v1"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
	v1beta11 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1beta1"
)

// MockProviderDockerClient is a mock of ProviderDockerClient interface.
type MockProviderDockerClient struct {
	ctrl     *gomock.Controller
	recorder *MockProviderDockerClientMockRecorder
}

// MockProviderDockerClientMockRecorder is the mock recorder for MockProviderDockerClient.
type MockProviderDockerClientMockRecorder struct {
	mock *MockProviderDockerClient
}

// NewMockProviderDockerClient creates a new mock instance.
func NewMockProviderDockerClient(ctrl *gomock.Controller) *MockProviderDockerClient {
	mock := &MockProviderDockerClient{ctrl: ctrl}
	mock.recorder = &MockProviderDockerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderDockerClient) EXPECT() *MockProviderDockerClientMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockProviderDockerClient) Run(arg0 context.Context, arg1, arg2 string, arg3 []string, arg4 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockProviderDockerClientMockRecorder) Run(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockProviderDockerClient)(nil).Run), varargs...)
}

// MockProviderKubectlClient is a mock of ProviderKubectlClient interface.
type MockProviderKubectlClient struct {
	ctrl     *gomock.Controller
	recorder *MockProviderKubectlClientMockRecorder
}

// MockProviderKubectlClientMockRecorder is the mock recorder for MockProviderKubectlClient.
type MockProviderKubectlClientMockRecorder struct {
	mock *MockProviderKubectlClient
}

// NewMockProviderKubectlClient creates a new mock instance.
func NewMockProviderKubectlClient(ctrl *gomock.Controller) *MockProviderKubectlClient {
	mock := &MockProviderKubectlClient{ctrl: ctrl}
	mock.recorder = &MockProviderKubectlClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderKubectlClient) EXPECT() *MockProviderKubectlClientMockRecorder {
	return m.recorder
}

// ApplyKubeSpec mocks base method.
func (m *MockProviderKubectlClient) ApplyKubeSpec(arg0 context.Context, arg1 *types.Cluster, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyKubeSpec", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyKubeSpec indicates an expected call of ApplyKubeSpec.
func (mr *MockProviderKubectlClientMockRecorder) ApplyKubeSpec(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyKubeSpec", reflect.TypeOf((*MockProviderKubectlClient)(nil).ApplyKubeSpec), arg0, arg1, arg2)
}

// ApplyKubeSpecFromBytesForce mocks base method.
func (m *MockProviderKubectlClient) ApplyKubeSpecFromBytesForce(arg0 context.Context, arg1 *types.Cluster, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyKubeSpecFromBytesForce", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyKubeSpecFromBytesForce indicates an expected call of ApplyKubeSpecFromBytesForce.
func (mr *MockProviderKubectlClientMockRecorder) ApplyKubeSpecFromBytesForce(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyKubeSpecFromBytesForce", reflect.TypeOf((*MockProviderKubectlClient)(nil).ApplyKubeSpecFromBytesForce), arg0, arg1, arg2)
}

// DeleteEksaDatacenterConfig mocks base method.
func (m *MockProviderKubectlClient) DeleteEksaDatacenterConfig(arg0 context.Context, arg1, arg2, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEksaDatacenterConfig", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEksaDatacenterConfig indicates an expected call of DeleteEksaDatacenterConfig.
func (mr *MockProviderKubectlClientMockRecorder) DeleteEksaDatacenterConfig(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEksaDatacenterConfig", reflect.TypeOf((*MockProviderKubectlClient)(nil).DeleteEksaDatacenterConfig), arg0, arg1, arg2, arg3, arg4)
}

// DeleteEksaMachineConfig mocks base method.
func (m *MockProviderKubectlClient) DeleteEksaMachineConfig(arg0 context.Context, arg1, arg2, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEksaMachineConfig", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEksaMachineConfig indicates an expected call of DeleteEksaMachineConfig.
func (mr *MockProviderKubectlClientMockRecorder) DeleteEksaMachineConfig(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEksaMachineConfig", reflect.TypeOf((*MockProviderKubectlClient)(nil).DeleteEksaMachineConfig), arg0, arg1, arg2, arg3, arg4)
}

// GetBmcsPowerState mocks base method.
func (m *MockProviderKubectlClient) GetBmcsPowerState(arg0 context.Context, arg1 []string, arg2, arg3 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBmcsPowerState", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBmcsPowerState indicates an expected call of GetBmcsPowerState.
func (mr *MockProviderKubectlClientMockRecorder) GetBmcsPowerState(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBmcsPowerState", reflect.TypeOf((*MockProviderKubectlClient)(nil).GetBmcsPowerState), arg0, arg1, arg2, arg3)
}

// GetEksaCluster mocks base method.
func (m *MockProviderKubectlClient) GetEksaCluster(arg0 context.Context, arg1 *types.Cluster, arg2 string) (*v1alpha1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaCluster indicates an expected call of GetEksaCluster.
func (mr *MockProviderKubectlClientMockRecorder) GetEksaCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaCluster", reflect.TypeOf((*MockProviderKubectlClient)(nil).GetEksaCluster), arg0, arg1, arg2)
}

// GetEksaTinkerbellDatacenterConfig mocks base method.
func (m *MockProviderKubectlClient) GetEksaTinkerbellDatacenterConfig(arg0 context.Context, arg1, arg2, arg3 string) (*v1alpha1.TinkerbellDatacenterConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaTinkerbellDatacenterConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.TinkerbellDatacenterConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaTinkerbellDatacenterConfig indicates an expected call of GetEksaTinkerbellDatacenterConfig.
func (mr *MockProviderKubectlClientMockRecorder) GetEksaTinkerbellDatacenterConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaTinkerbellDatacenterConfig", reflect.TypeOf((*MockProviderKubectlClient)(nil).GetEksaTinkerbellDatacenterConfig), arg0, arg1, arg2, arg3)
}

// GetEksaTinkerbellMachineConfig mocks base method.
func (m *MockProviderKubectlClient) GetEksaTinkerbellMachineConfig(arg0 context.Context, arg1, arg2, arg3 string) (*v1alpha1.TinkerbellMachineConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaTinkerbellMachineConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.TinkerbellMachineConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaTinkerbellMachineConfig indicates an expected call of GetEksaTinkerbellMachineConfig.
func (mr *MockProviderKubectlClientMockRecorder) GetEksaTinkerbellMachineConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaTinkerbellMachineConfig", reflect.TypeOf((*MockProviderKubectlClient)(nil).GetEksaTinkerbellMachineConfig), arg0, arg1, arg2, arg3)
}

// GetEtcdadmCluster mocks base method.
func (m *MockProviderKubectlClient) GetEtcdadmCluster(arg0 context.Context, arg1 *types.Cluster, arg2 string, arg3 ...executables.KubectlOpt) (*v1beta1.EtcdadmCluster, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEtcdadmCluster", varargs...)
	ret0, _ := ret[0].(*v1beta1.EtcdadmCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEtcdadmCluster indicates an expected call of GetEtcdadmCluster.
func (mr *MockProviderKubectlClientMockRecorder) GetEtcdadmCluster(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEtcdadmCluster", reflect.TypeOf((*MockProviderKubectlClient)(nil).GetEtcdadmCluster), varargs...)
}

// GetHardwareWithLabel mocks base method.
func (m *MockProviderKubectlClient) GetHardwareWithLabel(arg0 context.Context, arg1, arg2, arg3 string) ([]v1alpha10.Hardware, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHardwareWithLabel", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]v1alpha10.Hardware)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHardwareWithLabel indicates an expected call of GetHardwareWithLabel.
func (mr *MockProviderKubectlClientMockRecorder) GetHardwareWithLabel(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHardwareWithLabel", reflect.TypeOf((*MockProviderKubectlClient)(nil).GetHardwareWithLabel), arg0, arg1, arg2, arg3)
}

// GetKubeadmControlPlane mocks base method.
func (m *MockProviderKubectlClient) GetKubeadmControlPlane(arg0 context.Context, arg1 *types.Cluster, arg2 string, arg3 ...executables.KubectlOpt) (*v1beta11.KubeadmControlPlane, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKubeadmControlPlane", varargs...)
	ret0, _ := ret[0].(*v1beta11.KubeadmControlPlane)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKubeadmControlPlane indicates an expected call of GetKubeadmControlPlane.
func (mr *MockProviderKubectlClientMockRecorder) GetKubeadmControlPlane(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubeadmControlPlane", reflect.TypeOf((*MockProviderKubectlClient)(nil).GetKubeadmControlPlane), varargs...)
}

// GetMachineDeployment mocks base method.
func (m *MockProviderKubectlClient) GetMachineDeployment(arg0 context.Context, arg1 string, arg2 ...executables.KubectlOpt) (*v1beta10.MachineDeployment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMachineDeployment", varargs...)
	ret0, _ := ret[0].(*v1beta10.MachineDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineDeployment indicates an expected call of GetMachineDeployment.
func (mr *MockProviderKubectlClientMockRecorder) GetMachineDeployment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineDeployment", reflect.TypeOf((*MockProviderKubectlClient)(nil).GetMachineDeployment), varargs...)
}

// GetSecret mocks base method.
func (m *MockProviderKubectlClient) GetSecret(arg0 context.Context, arg1 string, arg2 ...executables.KubectlOpt) (*v10.Secret, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSecret", varargs...)
	ret0, _ := ret[0].(*v10.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockProviderKubectlClientMockRecorder) GetSecret(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockProviderKubectlClient)(nil).GetSecret), varargs...)
}

// UpdateAnnotation mocks base method.
func (m *MockProviderKubectlClient) UpdateAnnotation(arg0 context.Context, arg1, arg2 string, arg3 map[string]string, arg4 ...executables.KubectlOpt) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAnnotation", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAnnotation indicates an expected call of UpdateAnnotation.
func (mr *MockProviderKubectlClientMockRecorder) UpdateAnnotation(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnnotation", reflect.TypeOf((*MockProviderKubectlClient)(nil).UpdateAnnotation), varargs...)
}

// WaitForDeployment mocks base method.
func (m *MockProviderKubectlClient) WaitForDeployment(arg0 context.Context, arg1 *types.Cluster, arg2, arg3, arg4, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForDeployment", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForDeployment indicates an expected call of WaitForDeployment.
func (mr *MockProviderKubectlClientMockRecorder) WaitForDeployment(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForDeployment", reflect.TypeOf((*MockProviderKubectlClient)(nil).WaitForDeployment), arg0, arg1, arg2, arg3, arg4, arg5)
}

// MockProviderTinkClient is a mock of ProviderTinkClient interface.
type MockProviderTinkClient struct {
	ctrl     *gomock.Controller
	recorder *MockProviderTinkClientMockRecorder
}

// MockProviderTinkClientMockRecorder is the mock recorder for MockProviderTinkClient.
type MockProviderTinkClientMockRecorder struct {
	mock *MockProviderTinkClient
}

// NewMockProviderTinkClient creates a new mock instance.
func NewMockProviderTinkClient(ctrl *gomock.Controller) *MockProviderTinkClient {
	mock := &MockProviderTinkClient{ctrl: ctrl}
	mock.recorder = &MockProviderTinkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderTinkClient) EXPECT() *MockProviderTinkClientMockRecorder {
	return m.recorder
}

// DeleteWorkflow mocks base method.
func (m *MockProviderTinkClient) DeleteWorkflow(arg0 context.Context, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWorkflow", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkflow indicates an expected call of DeleteWorkflow.
func (mr *MockProviderTinkClientMockRecorder) DeleteWorkflow(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflow", reflect.TypeOf((*MockProviderTinkClient)(nil).DeleteWorkflow), varargs...)
}

// GetHardware mocks base method.
func (m *MockProviderTinkClient) GetHardware(arg0 context.Context) ([]*hardware0.Hardware, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHardware", arg0)
	ret0, _ := ret[0].([]*hardware0.Hardware)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHardware indicates an expected call of GetHardware.
func (mr *MockProviderTinkClientMockRecorder) GetHardware(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHardware", reflect.TypeOf((*MockProviderTinkClient)(nil).GetHardware), arg0)
}

// GetHardwareByUuid mocks base method.
func (m *MockProviderTinkClient) GetHardwareByUuid(arg0 context.Context, arg1 string) (*hardware.Hardware, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHardwareByUuid", arg0, arg1)
	ret0, _ := ret[0].(*hardware.Hardware)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHardwareByUuid indicates an expected call of GetHardwareByUuid.
func (mr *MockProviderTinkClientMockRecorder) GetHardwareByUuid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHardwareByUuid", reflect.TypeOf((*MockProviderTinkClient)(nil).GetHardwareByUuid), arg0, arg1)
}

// GetWorkflow mocks base method.
func (m *MockProviderTinkClient) GetWorkflow(arg0 context.Context) ([]*workflow.Workflow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkflow", arg0)
	ret0, _ := ret[0].([]*workflow.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflow indicates an expected call of GetWorkflow.
func (mr *MockProviderTinkClientMockRecorder) GetWorkflow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflow", reflect.TypeOf((*MockProviderTinkClient)(nil).GetWorkflow), arg0)
}

// PushHardware mocks base method.
func (m *MockProviderTinkClient) PushHardware(arg0 context.Context, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushHardware", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushHardware indicates an expected call of PushHardware.
func (mr *MockProviderTinkClientMockRecorder) PushHardware(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushHardware", reflect.TypeOf((*MockProviderTinkClient)(nil).PushHardware), arg0, arg1)
}

// MockProviderPbnjClient is a mock of ProviderPbnjClient interface.
type MockProviderPbnjClient struct {
	ctrl     *gomock.Controller
	recorder *MockProviderPbnjClientMockRecorder
}

// MockProviderPbnjClientMockRecorder is the mock recorder for MockProviderPbnjClient.
type MockProviderPbnjClientMockRecorder struct {
	mock *MockProviderPbnjClient
}

// NewMockProviderPbnjClient creates a new mock instance.
func NewMockProviderPbnjClient(ctrl *gomock.Controller) *MockProviderPbnjClient {
	mock := &MockProviderPbnjClient{ctrl: ctrl}
	mock.recorder = &MockProviderPbnjClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderPbnjClient) EXPECT() *MockProviderPbnjClientMockRecorder {
	return m.recorder
}

// GetPowerState mocks base method.
func (m *MockProviderPbnjClient) GetPowerState(arg0 context.Context, arg1 pbnj.BmcSecretConfig) (pbnj.PowerState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPowerState", arg0, arg1)
	ret0, _ := ret[0].(pbnj.PowerState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPowerState indicates an expected call of GetPowerState.
func (mr *MockProviderPbnjClientMockRecorder) GetPowerState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPowerState", reflect.TypeOf((*MockProviderPbnjClient)(nil).GetPowerState), arg0, arg1)
}

// PowerOff mocks base method.
func (m *MockProviderPbnjClient) PowerOff(arg0 context.Context, arg1 pbnj.BmcSecretConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOff", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PowerOff indicates an expected call of PowerOff.
func (mr *MockProviderPbnjClientMockRecorder) PowerOff(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOff", reflect.TypeOf((*MockProviderPbnjClient)(nil).PowerOff), arg0, arg1)
}

// PowerOn mocks base method.
func (m *MockProviderPbnjClient) PowerOn(arg0 context.Context, arg1 pbnj.BmcSecretConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PowerOn indicates an expected call of PowerOn.
func (mr *MockProviderPbnjClientMockRecorder) PowerOn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOn", reflect.TypeOf((*MockProviderPbnjClient)(nil).PowerOn), arg0, arg1)
}

// SetBootDevice mocks base method.
func (m *MockProviderPbnjClient) SetBootDevice(arg0 context.Context, arg1 pbnj.BmcSecretConfig, arg2 v1.BootDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBootDevice", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBootDevice indicates an expected call of SetBootDevice.
func (mr *MockProviderPbnjClientMockRecorder) SetBootDevice(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBootDevice", reflect.TypeOf((*MockProviderPbnjClient)(nil).SetBootDevice), arg0, arg1, arg2)
}

// MockSSHAuthKeyGenerator is a mock of SSHAuthKeyGenerator interface.
type MockSSHAuthKeyGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockSSHAuthKeyGeneratorMockRecorder
}

// MockSSHAuthKeyGeneratorMockRecorder is the mock recorder for MockSSHAuthKeyGenerator.
type MockSSHAuthKeyGeneratorMockRecorder struct {
	mock *MockSSHAuthKeyGenerator
}

// NewMockSSHAuthKeyGenerator creates a new mock instance.
func NewMockSSHAuthKeyGenerator(ctrl *gomock.Controller) *MockSSHAuthKeyGenerator {
	mock := &MockSSHAuthKeyGenerator{ctrl: ctrl}
	mock.recorder = &MockSSHAuthKeyGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSSHAuthKeyGenerator) EXPECT() *MockSSHAuthKeyGeneratorMockRecorder {
	return m.recorder
}

// GenerateSSHAuthKey mocks base method.
func (m *MockSSHAuthKeyGenerator) GenerateSSHAuthKey(arg0 filewriter.FileWriter) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSSHAuthKey", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateSSHAuthKey indicates an expected call of GenerateSSHAuthKey.
func (mr *MockSSHAuthKeyGeneratorMockRecorder) GenerateSSHAuthKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSSHAuthKey", reflect.TypeOf((*MockSSHAuthKeyGenerator)(nil).GenerateSSHAuthKey), arg0)
}
