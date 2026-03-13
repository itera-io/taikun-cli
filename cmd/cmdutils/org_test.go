package cmdutils

import (
	"testing"
)

func TestResolveOrgID_FlagSet(t *testing.T) {
	id, err := ResolveOrgID(42, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 42 {
		t.Fatalf("expected 42, got %d", id)
	}
}

func TestResolveOrgID_FlagWinsOverEnv(t *testing.T) {
	t.Setenv(EnvOrgID, "99")
	id, err := ResolveOrgID(42, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 42 {
		t.Fatalf("expected 42, got %d", id)
	}
}

func TestResolveOrgID_EnvValid(t *testing.T) {
	t.Setenv(EnvOrgID, "55")
	id, err := ResolveOrgID(0, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 55 {
		t.Fatalf("expected 55, got %d", id)
	}
}

func TestResolveOrgID_EnvValidRobot(t *testing.T) {
	t.Setenv(EnvOrgID, "55")
	id, err := ResolveOrgID(0, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 55 {
		t.Fatalf("expected 55, got %d", id)
	}
}

func TestResolveOrgID_EnvInvalidNonNumeric(t *testing.T) {
	t.Setenv(EnvOrgID, "abc")
	_, err := ResolveOrgID(0, false)
	if err == nil {
		t.Fatal("expected error for non-numeric env var")
	}
}

func TestResolveOrgID_EnvInvalidZero(t *testing.T) {
	t.Setenv(EnvOrgID, "0")
	_, err := ResolveOrgID(0, false)
	if err == nil {
		t.Fatal("expected error for zero env var")
	}
}

func TestResolveOrgID_EnvInvalidNegative(t *testing.T) {
	t.Setenv(EnvOrgID, "-5")
	_, err := ResolveOrgID(0, false)
	if err == nil {
		t.Fatal("expected error for negative env var")
	}
}

func TestResolveOrgID_UnsetRobot(t *testing.T) {
	id, err := ResolveOrgID(0, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 0 {
		t.Fatalf("expected 0, got %d", id)
	}
}

func TestResolveOrgID_UnsetStandard(t *testing.T) {
	_, err := ResolveOrgID(0, false)
	if err == nil {
		t.Fatal("expected ErrMissingOrg for standard auth without org")
	}
}

func TestIsRobotAuth_BothSet(t *testing.T) {
	t.Setenv("TAIKUN_ACCESS_KEY", "ak")
	t.Setenv("TAIKUN_SECRET_KEY", "sk")
	if !IsRobotAuth() {
		t.Fatal("expected true when both keys are set")
	}
}

func TestIsRobotAuth_OnlyAccessKey(t *testing.T) {
	t.Setenv("TAIKUN_ACCESS_KEY", "ak")
	if IsRobotAuth() {
		t.Fatal("expected false when only access key is set")
	}
}

func TestIsRobotAuth_OnlySecretKey(t *testing.T) {
	t.Setenv("TAIKUN_SECRET_KEY", "sk")
	if IsRobotAuth() {
		t.Fatal("expected false when only secret key is set")
	}
}

func TestIsRobotAuth_NeitherSet(t *testing.T) {
	if IsRobotAuth() {
		t.Fatal("expected false when neither key is set")
	}
}
