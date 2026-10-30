package helpers

import (
	"errors"
	"testing"

	meshkiterrors "github.com/meshery/meshkit/errors"
)

func TestErrorWrappers(t *testing.T) {
	testErr := errors.New("underlying test error")
	testObj := "test-object"

	tests := []struct {
		name     string
		err      error
		wantCode string
	}{
		{
			name:     "ErrNewDynamicClientGenerator",
			err:      ErrNewDynamicClientGenerator(testErr),
			wantCode: ErrErrNewDynamicClientGeneratorCode,
		},
		{
			name:     "ErrInvalidK8SConfig",
			err:      ErrInvalidK8SConfig(testErr),
			wantCode: ErrInvalidK8SConfigCode,
		},
		{
			name:     "ErrClientConfig",
			err:      ErrClientConfig(testErr),
			wantCode: ErrClientConfigCode,
		},
		{
			name:     "ErrFetchKubernetesNodes",
			err:      ErrFetchKubernetesNodes(testErr),
			wantCode: ErrFetchKubernetesNodesCode,
		},
		{
			name:     "ErrFetchNodes",
			err:      ErrFetchNodes(testErr),
			wantCode: ErrFetchNodesCode,
		},
		{
			name:     "ErrFetchKubernetesVersion",
			err:      ErrFetchKubernetesVersion(testErr),
			wantCode: ErrFetchKubernetesVersionCode,
		},
		{
			name:     "ErrScanKubernetes",
			err:      ErrScanKubernetes(testErr),
			wantCode: ErrScanKubernetesCode,
		},
		{
			name:     "ErrRetrievePodList",
			err:      ErrRetrievePodList(testErr),
			wantCode: ErrRetrievePodListCode,
		},
		{
			name:     "ErrDetectServiceForDeploymentImage",
			err:      ErrDetectServiceForDeploymentImage(testErr),
			wantCode: ErrDetectServiceForDeploymentImageCode,
		},
		{
			name:     "ErrRetrieveNamespacesList",
			err:      ErrRetrieveNamespacesList(testErr),
			wantCode: ErrRetrieveNamespacesListCode,
		},
		{
			name:     "ErrGetNamespaceDeployments",
			err:      ErrGetNamespaceDeployments(testErr, testObj),
			wantCode: ErrGetNamespaceDeploymentsCode,
		},
		{
			name:     "ErrDetectServiceWithName",
			err:      ErrDetectServiceWithName(testErr),
			wantCode: ErrDetectServiceWithNameCode,
		},
		{
			name:     "ErrGeneratingLoadTest",
			err:      ErrGeneratingLoadTest(testErr),
			wantCode: ErrGeneratingLoadTestCode,
		},
		{
			name:     "ErrRunningTest",
			err:      ErrRunningTest(testErr),
			wantCode: ErrRunningTestCode,
		},
		{
			name:     "ErrConvertingResultToMap",
			err:      ErrConvertingResultToMap(testErr),
			wantCode: ErrConvertingResultToMapCode,
		},
		{
			name:     "ErrGrpcSupport",
			err:      ErrGrpcSupport(testErr, testObj),
			wantCode: ErrGrpcSupportCode,
		},
		{
			name:     "ErrTransformingData",
			err:      ErrTransformingData(testErr),
			wantCode: ErrTransformingDataCode,
		},
		{
			name:     "ErrAddAndValidateExtraHeader",
			err:      ErrAddAndValidateExtraHeader(testErr),
			wantCode: ErrAddAndValidateExtraHeaderCode,
		},
		{
			name:     "ErrInClusterConfig",
			err:      ErrInClusterConfig(testErr),
			wantCode: ErrInClusterConfigCode,
		},
		{
			name:     "ErrNewKubeClientGenerator",
			err:      ErrNewKubeClientGenerator(testErr),
			wantCode: ErrNewKubeClientGeneratorCode,
		},
		{
			name:     "ErrRestConfigFromKubeConfig",
			err:      ErrRestConfigFromKubeConfig(testErr),
			wantCode: ErrRestConfigFromKubeConfigCode,
		},
		{
			name:     "ErrClientSet",
			err:      ErrClientSet(testErr),
			wantCode: ErrClientSetCode,
		},
		{
			name:     "ErrNewKubeClient",
			err:      ErrNewKubeClient(testErr),
			wantCode: ErrNewKubeClientCode,
		},
		{
			name:     "ErrDeployingAdapterInK8s",
			err:      ErrDeployingAdapterInK8s(testErr),
			wantCode: ErrDeployingAdapterInK8sEnvCode,
		},
		{
			name:     "ErrUnDeployingAdapterInK8s",
			err:      ErrUnDeployingAdapterInK8s(testErr),
			wantCode: ErrUnDeployingAdapterInK8sEnvCode,
		},
		{
			name:     "ErrDeployingAdapterInDocker",
			err:      ErrDeployingAdapterInDocker(testErr),
			wantCode: ErrDeployingAdapterInDockerEnvCode,
		},
		{
			name:     "ErrUnDeployingAdapterInDocker",
			err:      ErrUnDeployingAdapterInDocker(testErr),
			wantCode: ErrUnDeployingAdapterInDockerEnvCode,
		},
		{
			name:     "ErrDeployingAdapterInUnknownPlatform",
			err:      ErrDeployingAdapterInUnknownPlatform(testErr),
			wantCode: ErrDeployingAdapterCode,
		},
		{
			name:     "ErrUnDeployingAdapterInUnknownPlatform",
			err:      ErrUnDeployingAdapterInUnknownPlatform(testErr),
			wantCode: ErrUnDeployingAdapterCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				t.Fatalf("%s returned nil, expected error", tt.name)
			}

			code := meshkiterrors.GetCode(tt.err)
			if code != tt.wantCode {
				t.Errorf("%s code = %q, want %q", tt.name, code, tt.wantCode)
			}

			if gotMsg := tt.err.Error(); gotMsg == "" {
				t.Errorf("%s returned empty error message", tt.name)
			}
		})
	}
}
