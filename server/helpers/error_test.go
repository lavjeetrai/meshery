package helpers

import (
	"errors"
	"strings"
	"testing"

	meshkiterrors "github.com/meshery/meshkit/errors"
)

func TestErrorWrappers(t *testing.T) {
	dummyErr := errors.New("underlying sample error")
	testObj := "default-namespace"

	tests := []struct {
		name         string
		errFunc      func() error
		wantCode     string
		checkMsgText string
	}{
		{
			name:     "ErrNewDynamicClientGenerator",
			errFunc:  func() error { return ErrNewDynamicClientGenerator(dummyErr) },
			wantCode: ErrErrNewDynamicClientGeneratorCode,
		},
		{
			name:     "ErrInvalidK8SConfig",
			errFunc:  func() error { return ErrInvalidK8SConfig(dummyErr) },
			wantCode: ErrInvalidK8SConfigCode,
		},
		{
			name:     "ErrClientConfig",
			errFunc:  func() error { return ErrClientConfig(dummyErr) },
			wantCode: ErrClientConfigCode,
		},
		{
			name:     "ErrFetchKubernetesNodes",
			errFunc:  func() error { return ErrFetchKubernetesNodes(dummyErr) },
			wantCode: ErrFetchKubernetesNodesCode,
		},
		{
			name:     "ErrFetchNodes",
			errFunc:  func() error { return ErrFetchNodes(dummyErr) },
			wantCode: ErrFetchNodesCode,
		},
		{
			name:     "ErrFetchKubernetesVersion",
			errFunc:  func() error { return ErrFetchKubernetesVersion(dummyErr) },
			wantCode: ErrFetchKubernetesVersionCode,
		},
		{
			name:     "ErrScanKubernetes",
			errFunc:  func() error { return ErrScanKubernetes(dummyErr) },
			wantCode: ErrScanKubernetesCode,
		},
		{
			name:     "ErrRetrievePodList",
			errFunc:  func() error { return ErrRetrievePodList(dummyErr) },
			wantCode: ErrRetrievePodListCode,
		},
		{
			name:     "ErrDetectServiceForDeploymentImage",
			errFunc:  func() error { return ErrDetectServiceForDeploymentImage(dummyErr) },
			wantCode: ErrDetectServiceForDeploymentImageCode,
		},
		{
			name:     "ErrRetrieveNamespacesList",
			errFunc:  func() error { return ErrRetrieveNamespacesList(dummyErr) },
			wantCode: ErrRetrieveNamespacesListCode,
		},
		{
			name:         "ErrGetNamespaceDeployments",
			errFunc:      func() error { return ErrGetNamespaceDeployments(dummyErr, testObj) },
			wantCode:     ErrGetNamespaceDeploymentsCode,
			checkMsgText: testObj,
		},
		{
			name:     "ErrDetectServiceWithName",
			errFunc:  func() error { return ErrDetectServiceWithName(dummyErr) },
			wantCode: ErrDetectServiceWithNameCode,
		},
		{
			name:     "ErrGeneratingLoadTest",
			errFunc:  func() error { return ErrGeneratingLoadTest(dummyErr) },
			wantCode: ErrGeneratingLoadTestCode,
		},
		{
			name:     "ErrRunningTest",
			errFunc:  func() error { return ErrRunningTest(dummyErr) },
			wantCode: ErrRunningTestCode,
		},
		{
			name:     "ErrConvertingResultToMap",
			errFunc:  func() error { return ErrConvertingResultToMap(dummyErr) },
			wantCode: ErrConvertingResultToMapCode,
		},
		{
			name:         "ErrGrpcSupport",
			errFunc:      func() error { return ErrGrpcSupport(dummyErr, testObj) },
			wantCode:     ErrGrpcSupportCode,
			checkMsgText: testObj,
		},
		{
			name:     "ErrTransformingData",
			errFunc:  func() error { return ErrTransformingData(dummyErr) },
			wantCode: ErrTransformingDataCode,
		},
		{
			name:     "ErrAddAndValidateExtraHeader",
			errFunc:  func() error { return ErrAddAndValidateExtraHeader(dummyErr) },
			wantCode: ErrAddAndValidateExtraHeaderCode,
		},
		{
			name:     "ErrInClusterConfig",
			errFunc:  func() error { return ErrInClusterConfig(dummyErr) },
			wantCode: ErrInClusterConfigCode,
		},
		{
			name:     "ErrNewKubeClientGenerator",
			errFunc:  func() error { return ErrNewKubeClientGenerator(dummyErr) },
			wantCode: ErrNewKubeClientGeneratorCode,
		},
		{
			name:     "ErrRestConfigFromKubeConfig",
			errFunc:  func() error { return ErrRestConfigFromKubeConfig(dummyErr) },
			wantCode: ErrRestConfigFromKubeConfigCode,
		},
		{
			name:     "ErrClientSet",
			errFunc:  func() error { return ErrClientSet(dummyErr) },
			wantCode: ErrClientSetCode,
		},
		{
			name:     "ErrNewKubeClient",
			errFunc:  func() error { return ErrNewKubeClient(dummyErr) },
			wantCode: ErrNewKubeClientCode,
		},
		{
			name:     "ErrDeployingAdapterInK8s",
			errFunc:  func() error { return ErrDeployingAdapterInK8s(dummyErr) },
			wantCode: ErrDeployingAdapterInK8sEnvCode,
		},
		{
			name:     "ErrUnDeployingAdapterInK8s",
			errFunc:  func() error { return ErrUnDeployingAdapterInK8s(dummyErr) },
			wantCode: ErrUnDeployingAdapterInK8sEnvCode,
		},
		{
			name:     "ErrDeployingAdapterInDocker",
			errFunc:  func() error { return ErrDeployingAdapterInDocker(dummyErr) },
			wantCode: ErrDeployingAdapterInDockerEnvCode,
		},
		{
			name:     "ErrUnDeployingAdapterInDocker",
			errFunc:  func() error { return ErrUnDeployingAdapterInDocker(dummyErr) },
			wantCode: ErrUnDeployingAdapterInDockerEnvCode,
		},
		{
			name:     "ErrDeployingAdapterInUnknownPlatform",
			errFunc:  func() error { return ErrDeployingAdapterInUnknownPlatform(dummyErr) },
			wantCode: ErrDeployingAdapterCode,
		},
		{
			name:     "ErrUnDeployingAdapterInUnknownPlatform",
			errFunc:  func() error { return ErrUnDeployingAdapterInUnknownPlatform(dummyErr) },
			wantCode: ErrUnDeployingAdapterCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.errFunc()
			if err == nil {
				t.Fatalf("%s returned nil, expected error", tt.name)
			}

			if gotCode := meshkiterrors.GetCode(err); gotCode != tt.wantCode {
				t.Errorf("expected error code %q, got %q", tt.wantCode, gotCode)
			}

			errMsg := err.Error()
			if errMsg == "" {
				t.Errorf("expected non-empty error message for %s", tt.name)
			}

			if tt.checkMsgText != "" && !strings.Contains(errMsg, tt.checkMsgText) {
				t.Errorf("expected error message to contain %q, got %q", tt.checkMsgText, errMsg)
			}
		})
	}
}
