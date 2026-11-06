package helpers

import (
	"errors"
	"strings"
	"testing"

	meshkiterrors "github.com/meshery/meshkit/errors"
)

func TestErrInvalidK8SConfig(t *testing.T) {
	underlyingErr := errors.New("invalid k8s configuration file")
	err := ErrInvalidK8SConfig(underlyingErr)

	if err == nil {
		t.Fatal("expected non-nil error")
	}

	if code := meshkiterrors.GetCode(err); code != ErrInvalidK8SConfigCode {
		t.Errorf("expected error code %q, got %q", ErrInvalidK8SConfigCode, code)
	}

	if !strings.Contains(err.Error(), underlyingErr.Error()) {
		t.Errorf("expected error message to contain %q, got %q", underlyingErr.Error(), err.Error())
	}
}

func TestErrorWrappers(t *testing.T) {
	underlyingErr := errors.New("underlying sample error")

	tests := []struct {
		name     string
		err      error
		wantCode string
	}{
		{
			name:     "ErrNewDynamicClientGenerator",
			err:      ErrNewDynamicClientGenerator(underlyingErr),
			wantCode: ErrErrNewDynamicClientGeneratorCode,
		},
		{
			name:     "ErrInvalidK8SConfig",
			err:      ErrInvalidK8SConfig(underlyingErr),
			wantCode: ErrInvalidK8SConfigCode,
		},
		{
			name:     "ErrClientConfig",
			err:      ErrClientConfig(underlyingErr),
			wantCode: ErrClientConfigCode,
		},
		{
			name:     "ErrFetchKubernetesNodes",
			err:      ErrFetchKubernetesNodes(underlyingErr),
			wantCode: ErrFetchKubernetesNodesCode,
		},
		{
			name:     "ErrFetchNodes",
			err:      ErrFetchNodes(underlyingErr),
			wantCode: ErrFetchNodesCode,
		},
		{
			name:     "ErrFetchKubernetesVersion",
			err:      ErrFetchKubernetesVersion(underlyingErr),
			wantCode: ErrFetchKubernetesVersionCode,
		},
		{
			name:     "ErrScanKubernetes",
			err:      ErrScanKubernetes(underlyingErr),
			wantCode: ErrScanKubernetesCode,
		},
		{
			name:     "ErrRetrievePodList",
			err:      ErrRetrievePodList(underlyingErr),
			wantCode: ErrRetrievePodListCode,
		},
		{
			name:     "ErrDetectServiceForDeploymentImage",
			err:      ErrDetectServiceForDeploymentImage(underlyingErr),
			wantCode: ErrDetectServiceForDeploymentImageCode,
		},
		{
			name:     "ErrRetrieveNamespacesList",
			err:      ErrRetrieveNamespacesList(underlyingErr),
			wantCode: ErrRetrieveNamespacesListCode,
		},
		{
			name:     "ErrGetNamespaceDeployments",
			err:      ErrGetNamespaceDeployments(underlyingErr, "default"),
			wantCode: ErrGetNamespaceDeploymentsCode,
		},
		{
			name:     "ErrDetectServiceWithName",
			err:      ErrDetectServiceWithName(underlyingErr),
			wantCode: ErrDetectServiceWithNameCode,
		},
		{
			name:     "ErrGeneratingLoadTest",
			err:      ErrGeneratingLoadTest(underlyingErr),
			wantCode: ErrGeneratingLoadTestCode,
		},
		{
			name:     "ErrRunningTest",
			err:      ErrRunningTest(underlyingErr),
			wantCode: ErrRunningTestCode,
		},
		{
			name:     "ErrConvertingResultToMap",
			err:      ErrConvertingResultToMap(underlyingErr),
			wantCode: ErrConvertingResultToMapCode,
		},
		{
			name:     "ErrGrpcSupport",
			err:      ErrGrpcSupport(underlyingErr, "service"),
			wantCode: ErrGrpcSupportCode,
		},
		{
			name:     "ErrTransformingData",
			err:      ErrTransformingData(underlyingErr),
			wantCode: ErrTransformingDataCode,
		},
		{
			name:     "ErrAddAndValidateExtraHeader",
			err:      ErrAddAndValidateExtraHeader(underlyingErr),
			wantCode: ErrAddAndValidateExtraHeaderCode,
		},
		{
			name:     "ErrInClusterConfig",
			err:      ErrInClusterConfig(underlyingErr),
			wantCode: ErrInClusterConfigCode,
		},
		{
			name:     "ErrNewKubeClientGenerator",
			err:      ErrNewKubeClientGenerator(underlyingErr),
			wantCode: ErrNewKubeClientGeneratorCode,
		},
		{
			name:     "ErrRestConfigFromKubeConfig",
			err:      ErrRestConfigFromKubeConfig(underlyingErr),
			wantCode: ErrRestConfigFromKubeConfigCode,
		},
		{
			name:     "ErrClientSet",
			err:      ErrClientSet(underlyingErr),
			wantCode: ErrClientSetCode,
		},
		{
			name:     "ErrNewKubeClient",
			err:      ErrNewKubeClient(underlyingErr),
			wantCode: ErrNewKubeClientCode,
		},
		{
			name:     "ErrDeployingAdapterInK8s",
			err:      ErrDeployingAdapterInK8s(underlyingErr),
			wantCode: ErrDeployingAdapterInK8sEnvCode,
		},
		{
			name:     "ErrUnDeployingAdapterInK8s",
			err:      ErrUnDeployingAdapterInK8s(underlyingErr),
			wantCode: ErrUnDeployingAdapterInK8sEnvCode,
		},
		{
			name:     "ErrDeployingAdapterInDocker",
			err:      ErrDeployingAdapterInDocker(underlyingErr),
			wantCode: ErrDeployingAdapterInDockerEnvCode,
		},
		{
			name:     "ErrUnDeployingAdapterInDocker",
			err:      ErrUnDeployingAdapterInDocker(underlyingErr),
			wantCode: ErrUnDeployingAdapterInDockerEnvCode,
		},
		{
			name:     "ErrDeployingAdapterInUnknownPlatform",
			err:      ErrDeployingAdapterInUnknownPlatform(underlyingErr),
			wantCode: ErrDeployingAdapterCode,
		},
		{
			name:     "ErrUnDeployingAdapterInUnknownPlatform",
			err:      ErrUnDeployingAdapterInUnknownPlatform(underlyingErr),
			wantCode: ErrUnDeployingAdapterCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				t.Fatalf("%s returned nil, expected error", tt.name)
			}
			if code := meshkiterrors.GetCode(tt.err); code != tt.wantCode {
				t.Errorf("%s code = %q, want %q", tt.name, code, tt.wantCode)
			}
			if !strings.Contains(tt.err.Error(), underlyingErr.Error()) {
				t.Errorf("%s error message = %q, want to contain %q", tt.name, tt.err.Error(), underlyingErr.Error())
			}
		})
	}
}
