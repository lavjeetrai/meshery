package helpers

import (
	"errors"
	"strings"
	"testing"

	meshkiterrors "github.com/meshery/meshkit/errors"
)

func TestErrNewDynamicClientGenerator(t *testing.T) {
	underlyingErr := errors.New("failed to generate dynamic client")

	err := ErrNewDynamicClientGenerator(underlyingErr)
	if err == nil {
		t.Fatal("expected non-nil error, got nil")
	}

	code := meshkiterrors.GetCode(err)
	if code != ErrErrNewDynamicClientGeneratorCode {
		t.Errorf("expected error code %q, got %q", ErrErrNewDynamicClientGeneratorCode, code)
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, underlyingErr.Error()) {
		t.Errorf("expected error message to contain %q, got %q", underlyingErr.Error(), errMsg)
	}
}
