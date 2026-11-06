package utils

import (
	"os"
	"strings"
	"testing"
)

func TestGKEConfigValidation(t *testing.T) {
	tests := []struct {
		name      string
		config    GKEConfig
		wantError bool
	}{
		{
			name: "valid config",
			config: GKEConfig{
				ConfigPath: "/tmp/config",
				SAName:     "test",
				Namespace:  "default",
			},
			wantError: false,
		},
		{
			name: "invalid SAName with command injection semicolon",
			config: GKEConfig{
				ConfigPath: "/tmp/config",
				SAName:     "test; id",
				Namespace:  "default",
			},
			wantError: true,
		},
		{
			name: "invalid SAName with command substitution",
			config: GKEConfig{
				ConfigPath: "/tmp/config",
				SAName:     "test$(id)",
				Namespace:  "default",
			},
			wantError: true,
		},
		{
			name: "invalid Namespace with uppercase characters or injection",
			config: GKEConfig{
				ConfigPath: "/tmp/config",
				SAName:     "test-sa",
				Namespace:  "Default; reboot",
			},
			wantError: true,
		},
		{
			name: "invalid ConfigPath with shell metacharacters",
			config: GKEConfig{
				ConfigPath: "/tmp/config; rm -rf /",
				SAName:     "test-sa",
				Namespace:  "default",
			},
			wantError: true,
		},
		{
			name: "invalid ConfigPath with backticks",
			config: GKEConfig{
				ConfigPath: "/tmp/`whoami`",
				SAName:     "test-sa",
				Namespace:  "default",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.validate()
			if (err != nil) != tt.wantError {
				t.Errorf("validate() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestPrerequisiteChecking(t *testing.T) {
	// Temporarily disable stdout
	old := os.Stdout
	os.Stdout = nil
	defer func() { os.Stdout = old }()

	cfg := &GKEConfig{
		ConfigPath: "/tmp/config",
		SAName:     "test-sa",
		Namespace:  "default",
	}

	err := cfg.checkPrerequisites()
	if err != nil {
		t.Logf("Prerequisites check failed: %v", err)
	}
}

func TestScriptGeneration(t *testing.T) {
	cfg := &GKEConfig{
		ConfigPath: "/tmp/test-config",
		SAName:     "test-service-account",
		Namespace:  "test-namespace",
	}

	script := cfg.generateScript()
	if script == "" {
		t.Error("Script generation failed: empty script")
	}

	expectedConfigPath := "KUBECFG_FILE_NAME='/tmp/test-config'"
	expectedSAName := "SERVICE_ACCOUNT_NAME='test-service-account'"
	expectedNamespace := "NAMESPACE='test-namespace'"

	if !strings.Contains(script, expectedConfigPath) {
		t.Errorf("Expected script to contain %q, got script: %s", expectedConfigPath, script)
	}
	if !strings.Contains(script, expectedSAName) {
		t.Errorf("Expected script to contain %q, got script: %s", expectedSAName, script)
	}
	if !strings.Contains(script, expectedNamespace) {
		t.Errorf("Expected script to contain %q, got script: %s", expectedNamespace, script)
	}
}

func TestInvalidConfigurations(t *testing.T) {
	invalidCfg := &GKEConfig{
		ConfigPath: "",
		SAName:     "",
		Namespace:  "",
	}

	if err := invalidCfg.validate(); err == nil {
		t.Error("Expected validation to fail for empty config")
	}
}

func TestFullConfiguration(t *testing.T) {
	// Temporarily disable stdout
	old := os.Stdout
	os.Stdout = nil
	defer func() { os.Stdout = old }()

	err := GenerateConfigGKE("/tmp/test-config", "test-sa", "default")
	if err == nil {
		t.Error("Expected error in non-k8s environment")
	}
}
