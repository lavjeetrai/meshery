package helpers

import (
	"errors"
	"strings"
	"testing"

	meshkitErrors "github.com/meshery/meshkit/errors"
)

func TestErrorWrappers(t *testing.T) {
	mockErr := errors.New("mock error cause")

	tests := []struct {
		name          string
		err           error
		expectedCode  string
		expectedSubstr string
	}{
		{
			name:         "ErrNewDynamicClientGenerator",
			err:          ErrNewDynamicClientGenerator(mockErr),
			expectedCode: ErrErrNewDynamicClientGeneratorCode,
		},
		{
			name:         "ErrInvalidK8SConfig",
			err:          ErrInvalidK8SConfig(mockErr),
			expectedCode: ErrInvalidK8SConfigCode,
		},
		{
			name:         "ErrClientConfig",
			err:          ErrClientConfig(mockErr),
			expectedCode: ErrClientConfigCode,
		},
		{
			name:         "ErrFetchKubernetesNodes",
			err:          ErrFetchKubernetesNodes(mockErr),
			expectedCode: ErrFetchKubernetesNodesCode,
		},
		{
			name:         "ErrFetchNodes",
			err:          ErrFetchNodes(mockErr),
			expectedCode: ErrFetchNodesCode,
		},
		{
			name:         "ErrFetchKubernetesVersion",
			err:          ErrFetchKubernetesVersion(mockErr),
			expectedCode: ErrFetchKubernetesVersionCode,
		},
		{
			name:         "ErrScanKubernetes",
			err:          ErrScanKubernetes(mockErr),
			expectedCode: ErrScanKubernetesCode,
		},
		{
			name:         "ErrRetrievePodList",
			err:          ErrRetrievePodList(mockErr),
			expectedCode: ErrRetrievePodListCode,
		},
		{
			name:         "ErrDetectServiceForDeploymentImage",
			err:          ErrDetectServiceForDeploymentImage(mockErr),
			expectedCode: ErrDetectServiceForDeploymentImageCode,
		},
		{
			name:         "ErrRetrieveNamespacesList",
			err:          ErrRetrieveNamespacesList(mockErr),
			expectedCode: ErrRetrieveNamespacesListCode,
		},
		{
			name:           "ErrGetNamespaceDeployments",
			err:            ErrGetNamespaceDeployments(mockErr, "test-namespace"),
			expectedCode:   ErrGetNamespaceDeploymentsCode,
			expectedSubstr: "test-namespace",
		},
		{
			name:         "ErrDetectServiceWithName",
			err:          ErrDetectServiceWithName(mockErr),
			expectedCode: ErrDetectServiceWithNameCode,
		},
		{
			name:         "ErrGeneratingLoadTest",
			err:          ErrGeneratingLoadTest(mockErr),
			expectedCode: ErrGeneratingLoadTestCode,
		},
		{
			name:         "ErrRunningTest",
			err:          ErrRunningTest(mockErr),
			expectedCode: ErrRunningTestCode,
		},
		{
			name:         "ErrConvertingResultToMap",
			err:          ErrConvertingResultToMap(mockErr),
			expectedCode: ErrConvertingResultToMapCode,
		},
		{
			name:           "ErrGrpcSupport",
			err:            ErrGrpcSupport(mockErr, "my-grpc-service"),
			expectedCode:   ErrGrpcSupportCode,
			expectedSubstr: "my-grpc-service",
		},
		{
			name:         "ErrTransformingData",
			err:          ErrTransformingData(mockErr),
			expectedCode: ErrTransformingDataCode,
		},
		{
			name:         "ErrAddAndValidateExtraHeader",
			err:          ErrAddAndValidateExtraHeader(mockErr),
			expectedCode: ErrAddAndValidateExtraHeaderCode,
		},
		{
			name:         "ErrInClusterConfig",
			err:          ErrInClusterConfig(mockErr),
			expectedCode: ErrInClusterConfigCode,
		},
		{
			name:         "ErrNewKubeClientGenerator",
			err:          ErrNewKubeClientGenerator(mockErr),
			expectedCode: ErrNewKubeClientGeneratorCode,
		},
		{
			name:         "ErrRestConfigFromKubeConfig",
			err:          ErrRestConfigFromKubeConfig(mockErr),
			expectedCode: ErrRestConfigFromKubeConfigCode,
		},
		{
			name:         "ErrClientSet",
			err:          ErrClientSet(mockErr),
			expectedCode: ErrClientSetCode,
		},
		{
			name:         "ErrNewKubeClient",
			err:          ErrNewKubeClient(mockErr),
			expectedCode: ErrNewKubeClientCode,
		},
		{
			name:         "ErrDeployingAdapterInK8s",
			err:          ErrDeployingAdapterInK8s(mockErr),
			expectedCode: ErrDeployingAdapterInK8sEnvCode,
		},
		{
			name:         "ErrUnDeployingAdapterInK8s",
			err:          ErrUnDeployingAdapterInK8s(mockErr),
			expectedCode: ErrUnDeployingAdapterInK8sEnvCode,
		},
		{
			name:         "ErrDeployingAdapterInDocker",
			err:          ErrDeployingAdapterInDocker(mockErr),
			expectedCode: ErrDeployingAdapterInDockerEnvCode,
		},
		{
			name:         "ErrUnDeployingAdapterInDocker",
			err:          ErrUnDeployingAdapterInDocker(mockErr),
			expectedCode: ErrUnDeployingAdapterInDockerEnvCode,
		},
		{
			name:         "ErrDeployingAdapterInUnknownPlatform",
			err:          ErrDeployingAdapterInUnknownPlatform(mockErr),
			expectedCode: ErrDeployingAdapterCode,
		},
		{
			name:         "ErrUnDeployingAdapterInUnknownPlatform",
			err:          ErrUnDeployingAdapterInUnknownPlatform(mockErr),
			expectedCode: ErrUnDeployingAdapterCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				t.Fatalf("expected non-nil error")
			}

			code := meshkitErrors.GetCode(tt.err)
			if code != tt.expectedCode {
				t.Errorf("expected code %s, got %s", tt.expectedCode, code)
			}

			errMsg := tt.err.Error()
			if !strings.Contains(errMsg, mockErr.Error()) {
				t.Errorf("expected error message to contain cause %q, got %q", mockErr.Error(), errMsg)
			}

			if tt.expectedSubstr != "" && !strings.Contains(errMsg, tt.expectedSubstr) {
				t.Errorf("expected error message to contain %q, got %q", tt.expectedSubstr, errMsg)
			}
		})
	}
}
