package helpers

import (
	"errors"
	"strings"
	"testing"

	mesheryerrors "github.com/meshery/meshkit/errors"
)

func TestErrNewDynamicClientGenerator(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrNewDynamicClientGenerator(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrErrNewDynamicClientGeneratorCode {
		t.Errorf("Expected error code %s, got %s", ErrErrNewDynamicClientGeneratorCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrInvalidK8SConfig(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrInvalidK8SConfig(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrInvalidK8SConfigCode {
		t.Errorf("Expected error code %s, got %s", ErrInvalidK8SConfigCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrClientConfig(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrClientConfig(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrClientConfigCode {
		t.Errorf("Expected error code %s, got %s", ErrClientConfigCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrFetchKubernetesNodes(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrFetchKubernetesNodes(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrFetchKubernetesNodesCode {
		t.Errorf("Expected error code %s, got %s", ErrFetchKubernetesNodesCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrFetchNodes(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrFetchNodes(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrFetchNodesCode {
		t.Errorf("Expected error code %s, got %s", ErrFetchNodesCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrFetchKubernetesVersion(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrFetchKubernetesVersion(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrFetchKubernetesVersionCode {
		t.Errorf("Expected error code %s, got %s", ErrFetchKubernetesVersionCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrScanKubernetes(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrScanKubernetes(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrScanKubernetesCode {
		t.Errorf("Expected error code %s, got %s", ErrScanKubernetesCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrRetrievePodList(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrRetrievePodList(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrRetrievePodListCode {
		t.Errorf("Expected error code %s, got %s", ErrRetrievePodListCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrDetectServiceForDeploymentImage(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrDetectServiceForDeploymentImage(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrDetectServiceForDeploymentImageCode {
		t.Errorf("Expected error code %s, got %s", ErrDetectServiceForDeploymentImageCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrRetrieveNamespacesList(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrRetrieveNamespacesList(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrRetrieveNamespacesListCode {
		t.Errorf("Expected error code %s, got %s", ErrRetrieveNamespacesListCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrGetNamespaceDeployments(t *testing.T) {
	err := errors.New("test error")
	obj := "test_obj"
	resultErr := ErrGetNamespaceDeployments(err, obj)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrGetNamespaceDeploymentsCode {
		t.Errorf("Expected error code %s, got %s", ErrGetNamespaceDeploymentsCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrDetectServiceWithName(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrDetectServiceWithName(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrDetectServiceWithNameCode {
		t.Errorf("Expected error code %s, got %s", ErrDetectServiceWithNameCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrGeneratingLoadTest(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrGeneratingLoadTest(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrGeneratingLoadTestCode {
		t.Errorf("Expected error code %s, got %s", ErrGeneratingLoadTestCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrRunningTest(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrRunningTest(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrRunningTestCode {
		t.Errorf("Expected error code %s, got %s", ErrRunningTestCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrConvertingResultToMap(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrConvertingResultToMap(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrConvertingResultToMapCode {
		t.Errorf("Expected error code %s, got %s", ErrConvertingResultToMapCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrGrpcSupport(t *testing.T) {
	err := errors.New("test error")
	obj := "test_obj"
	resultErr := ErrGrpcSupport(err, obj)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrGrpcSupportCode {
		t.Errorf("Expected error code %s, got %s", ErrGrpcSupportCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrTransformingData(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrTransformingData(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrTransformingDataCode {
		t.Errorf("Expected error code %s, got %s", ErrTransformingDataCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrAddAndValidateExtraHeader(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrAddAndValidateExtraHeader(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrAddAndValidateExtraHeaderCode {
		t.Errorf("Expected error code %s, got %s", ErrAddAndValidateExtraHeaderCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrInClusterConfig(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrInClusterConfig(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrInClusterConfigCode {
		t.Errorf("Expected error code %s, got %s", ErrInClusterConfigCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrNewKubeClientGenerator(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrNewKubeClientGenerator(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrNewKubeClientGeneratorCode {
		t.Errorf("Expected error code %s, got %s", ErrNewKubeClientGeneratorCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrRestConfigFromKubeConfig(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrRestConfigFromKubeConfig(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrRestConfigFromKubeConfigCode {
		t.Errorf("Expected error code %s, got %s", ErrRestConfigFromKubeConfigCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrClientSet(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrClientSet(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrClientSetCode {
		t.Errorf("Expected error code %s, got %s", ErrClientSetCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrNewKubeClient(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrNewKubeClient(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrNewKubeClientCode {
		t.Errorf("Expected error code %s, got %s", ErrNewKubeClientCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrDeployingAdapterInK8s(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrDeployingAdapterInK8s(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrDeployingAdapterInK8sEnvCode {
		t.Errorf("Expected error code %s, got %s", ErrDeployingAdapterInK8sEnvCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrUnDeployingAdapterInK8s(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrUnDeployingAdapterInK8s(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrUnDeployingAdapterInK8sEnvCode {
		t.Errorf("Expected error code %s, got %s", ErrUnDeployingAdapterInK8sEnvCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrDeployingAdapterInDocker(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrDeployingAdapterInDocker(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrDeployingAdapterInDockerEnvCode {
		t.Errorf("Expected error code %s, got %s", ErrDeployingAdapterInDockerEnvCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrUnDeployingAdapterInDocker(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrUnDeployingAdapterInDocker(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrUnDeployingAdapterInDockerEnvCode {
		t.Errorf("Expected error code %s, got %s", ErrUnDeployingAdapterInDockerEnvCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrDeployingAdapterInUnknownPlatform(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrDeployingAdapterInUnknownPlatform(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrDeployingAdapterCode {
		t.Errorf("Expected error code %s, got %s", ErrDeployingAdapterCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}

func TestErrUnDeployingAdapterInUnknownPlatform(t *testing.T) {
	err := errors.New("test error")
	resultErr := ErrUnDeployingAdapterInUnknownPlatform(err)

	if resultErr == nil {
		t.Errorf("Expected error, got nil")
	}
	if mesheryerrors.GetCode(resultErr) != ErrUnDeployingAdapterCode {
		t.Errorf("Expected error code %s, got %s", ErrUnDeployingAdapterCode, mesheryerrors.GetCode(resultErr))
	}
	if !strings.Contains(resultErr.Error(), err.Error()) {
		t.Errorf("Expected error message to contain '%s', got '%s'", err.Error(), resultErr.Error())
	}
}
