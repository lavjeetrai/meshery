package helpers

import (
	"errors"
	"testing"

	meshkiterrors "github.com/meshery/meshkit/errors"
)

func TestErrors(t *testing.T) {
	dummyErr := errors.New("sample error")

	tests := []struct {
		name         string
		fn           func() error
		expectedCode string
	}{
		{
			name:         "ErrNewDynamicClientGenerator",
			fn:           func() error { return ErrNewDynamicClientGenerator(dummyErr) },
			expectedCode: ErrErrNewDynamicClientGeneratorCode,
		},
		{
			name:         "ErrInvalidK8SConfig",
			fn:           func() error { return ErrInvalidK8SConfig(dummyErr) },
			expectedCode: ErrInvalidK8SConfigCode,
		},
		{
			name:         "ErrClientConfig",
			fn:           func() error { return ErrClientConfig(dummyErr) },
			expectedCode: ErrClientConfigCode,
		},
		{
			name:         "ErrFetchKubernetesNodes",
			fn:           func() error { return ErrFetchKubernetesNodes(dummyErr) },
			expectedCode: ErrFetchKubernetesNodesCode,
		},
		{
			name:         "ErrFetchNodes",
			fn:           func() error { return ErrFetchNodes(dummyErr) },
			expectedCode: ErrFetchNodesCode,
		},
		{
			name:         "ErrFetchKubernetesVersion",
			fn:           func() error { return ErrFetchKubernetesVersion(dummyErr) },
			expectedCode: ErrFetchKubernetesVersionCode,
		},
		{
			name:         "ErrScanKubernetes",
			fn:           func() error { return ErrScanKubernetes(dummyErr) },
			expectedCode: ErrScanKubernetesCode,
		},
		{
			name:         "ErrRetrievePodList",
			fn:           func() error { return ErrRetrievePodList(dummyErr) },
			expectedCode: ErrRetrievePodListCode,
		},
		{
			name:         "ErrDetectServiceForDeploymentImage",
			fn:           func() error { return ErrDetectServiceForDeploymentImage(dummyErr) },
			expectedCode: ErrDetectServiceForDeploymentImageCode,
		},
		{
			name:         "ErrRetrieveNamespacesList",
			fn:           func() error { return ErrRetrieveNamespacesList(dummyErr) },
			expectedCode: ErrRetrieveNamespacesListCode,
		},
		{
			name:         "ErrGetNamespaceDeployments",
			fn:           func() error { return ErrGetNamespaceDeployments(dummyErr, "default") },
			expectedCode: ErrGetNamespaceDeploymentsCode,
		},
		{
			name:         "ErrDetectServiceWithName",
			fn:           func() error { return ErrDetectServiceWithName(dummyErr) },
			expectedCode: ErrDetectServiceWithNameCode,
		},
		{
			name:         "ErrGeneratingLoadTest",
			fn:           func() error { return ErrGeneratingLoadTest(dummyErr) },
			expectedCode: ErrGeneratingLoadTestCode,
		},
		{
			name:         "ErrRunningTest",
			fn:           func() error { return ErrRunningTest(dummyErr) },
			expectedCode: ErrRunningTestCode,
		},
		{
			name:         "ErrConvertingResultToMap",
			fn:           func() error { return ErrConvertingResultToMap(dummyErr) },
			expectedCode: ErrConvertingResultToMapCode,
		},
		{
			name:         "ErrGrpcSupport",
			fn:           func() error { return ErrGrpcSupport(dummyErr, "service") },
			expectedCode: ErrGrpcSupportCode,
		},
		{
			name:         "ErrTransformingData",
			fn:           func() error { return ErrTransformingData(dummyErr) },
			expectedCode: ErrTransformingDataCode,
		},
		{
			name:         "ErrAddAndValidateExtraHeader",
			fn:           func() error { return ErrAddAndValidateExtraHeader(dummyErr) },
			expectedCode: ErrAddAndValidateExtraHeaderCode,
		},
		{
			name:         "ErrInClusterConfig",
			fn:           func() error { return ErrInClusterConfig(dummyErr) },
			expectedCode: ErrInClusterConfigCode,
		},
		{
			name:         "ErrNewKubeClientGenerator",
			fn:           func() error { return ErrNewKubeClientGenerator(dummyErr) },
			expectedCode: ErrNewKubeClientGeneratorCode,
		},
		{
			name:         "ErrRestConfigFromKubeConfig",
			fn:           func() error { return ErrRestConfigFromKubeConfig(dummyErr) },
			expectedCode: ErrRestConfigFromKubeConfigCode,
		},
		{
			name:         "ErrClientSet",
			fn:           func() error { return ErrClientSet(dummyErr) },
			expectedCode: ErrClientSetCode,
		},
		{
			name:         "ErrNewKubeClient",
			fn:           func() error { return ErrNewKubeClient(dummyErr) },
			expectedCode: ErrNewKubeClientCode,
		},
		{
			name:         "ErrDeployingAdapterInK8s",
			fn:           func() error { return ErrDeployingAdapterInK8s(dummyErr) },
			expectedCode: ErrDeployingAdapterInK8sEnvCode,
		},
		{
			name:         "ErrUnDeployingAdapterInK8s",
			fn:           func() error { return ErrUnDeployingAdapterInK8s(dummyErr) },
			expectedCode: ErrUnDeployingAdapterInK8sEnvCode,
		},
		{
			name:         "ErrDeployingAdapterInDocker",
			fn:           func() error { return ErrDeployingAdapterInDocker(dummyErr) },
			expectedCode: ErrDeployingAdapterInDockerEnvCode,
		},
		{
			name:         "ErrUnDeployingAdapterInDocker",
			fn:           func() error { return ErrUnDeployingAdapterInDocker(dummyErr) },
			expectedCode: ErrUnDeployingAdapterInDockerEnvCode,
		},
		{
			name:         "ErrDeployingAdapterInUnknownPlatform",
			fn:           func() error { return ErrDeployingAdapterInUnknownPlatform(dummyErr) },
			expectedCode: ErrDeployingAdapterCode,
		},
		{
			name:         "ErrUnDeployingAdapterInUnknownPlatform",
			fn:           func() error { return ErrUnDeployingAdapterInUnknownPlatform(dummyErr) },
			expectedCode: ErrUnDeployingAdapterCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			if err == nil {
				t.Fatalf("%s returned nil, expected error", tt.name)
			}
			code := meshkiterrors.GetCode(err)
			if code != tt.expectedCode {
				t.Errorf("expected error code %s, got %s", tt.expectedCode, code)
			}
			if err.Error() == "" {
				t.Errorf("expected non-empty error message")
			}
		})
	}
}
