package cmdutils

import (
	"strings"
	"testing"
)

// Acceptance / integration tests for the shared org resolver.
// These tests exercise the full chain: IsRobotAuth() + ResolveOrgID()
// together, validating spec acceptance scenarios end-to-end.
//
// References:
//   Spec: kitty-specs/001-cli-org-optionality-by-auth-type/spec.md
//   FR-001 .. FR-008, SC-001 .. SC-004, US-1 .. US-3

// clearRobotKeys ensures the test runs in standard-auth context by
// explicitly setting robot env vars to empty (via t.Setenv, which
// restores them after the test).
func clearRobotKeys(t *testing.T) {
	t.Helper()
	t.Setenv("TAIKUN_ACCESS_KEY", "")
	t.Setenv("TAIKUN_SECRET_KEY", "")
}

// setRobotKeys configures robot-auth env vars so IsRobotAuth() returns true.
func setRobotKeys(t *testing.T) {
	t.Helper()
	t.Setenv("TAIKUN_ACCESS_KEY", "test-access-key")
	t.Setenv("TAIKUN_SECRET_KEY", "test-secret-key")
}

// --------------------------------------------------------------------------
// T038 — Standard user with TAIKUN_ORGANIZATION_ID set (US1-1)
// --------------------------------------------------------------------------
// Spec ref: FR-002 precedence (env fallback), SC-001, US-1 acceptance 1
func TestAcceptance_T038_StandardUserWithEnvOrgID(t *testing.T) {
	clearRobotKeys(t)
	t.Setenv(EnvOrgID, "42")

	orgID, err := ResolveOrgID(0, IsRobotAuth())
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if orgID != 42 {
		t.Fatalf("expected orgID=42, got %d", orgID)
	}
}

// --------------------------------------------------------------------------
// T039 — Standard user without any org source (US1-2)
// --------------------------------------------------------------------------
// Spec ref: FR-003, FR-008, SC-002, US-1 acceptance 2
func TestAcceptance_T039_StandardUserWithoutOrg(t *testing.T) {
	clearRobotKeys(t)
	t.Setenv(EnvOrgID, "") // explicitly unset

	_, err := ResolveOrgID(0, IsRobotAuth())
	if err == nil {
		t.Fatal("expected error for standard user without org, got nil")
	}

	msg := err.Error()

	// FR-008: error must mention both remediation paths
	if !strings.Contains(msg, "--organization-id") {
		t.Errorf("error should mention --organization-id flag; got: %s", msg)
	}
	if !strings.Contains(msg, "TAIKUN_ORGANIZATION_ID") {
		t.Errorf("error should mention TAIKUN_ORGANIZATION_ID env var; got: %s", msg)
	}
}

// --------------------------------------------------------------------------
// T040 — Explicit flag overrides env var (US2-1)
// --------------------------------------------------------------------------
// Spec ref: FR-002 precedence (explicit flag first), FR-007, SC-001, US-2 acceptance 1
func TestAcceptance_T040_FlagOverridesEnv(t *testing.T) {
	// Auth type should not matter for this test; run for standard auth
	clearRobotKeys(t)
	t.Setenv(EnvOrgID, "100")

	orgID, err := ResolveOrgID(200, IsRobotAuth())
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if orgID != 200 {
		t.Fatalf("expected flag value 200 to win over env 100, got %d", orgID)
	}
}

// --------------------------------------------------------------------------
// T041 — Robot user without org, no local blocking (US3-1)
// --------------------------------------------------------------------------
// Spec ref: FR-004, SC-003, US-3 acceptance 1
func TestAcceptance_T041_RobotUserWithoutOrg(t *testing.T) {
	setRobotKeys(t)
	t.Setenv(EnvOrgID, "") // no org configured

	orgID, err := ResolveOrgID(0, IsRobotAuth())
	if err != nil {
		t.Fatalf("expected passthrough (no error) for robot user without org, got: %v", err)
	}
	if orgID != 0 {
		t.Fatalf("expected orgID=0 (passthrough) for robot user, got %d", orgID)
	}
}

// --------------------------------------------------------------------------
// T042 — Invalid TAIKUN_ORGANIZATION_ID values
// --------------------------------------------------------------------------
// Spec ref: Edge Cases (invalid env), FR-003
func TestAcceptance_T042_InvalidEnvOrgID(t *testing.T) {
	t.Run("non-numeric value", func(t *testing.T) {
		clearRobotKeys(t)
		t.Setenv(EnvOrgID, "abc")

		_, err := ResolveOrgID(0, IsRobotAuth())
		if err == nil {
			t.Fatal("expected error for non-numeric env value, got nil")
		}
		if !strings.Contains(err.Error(), "positive integer") {
			t.Errorf("error should mention 'positive integer'; got: %s", err.Error())
		}
	})

	t.Run("zero value", func(t *testing.T) {
		clearRobotKeys(t)
		t.Setenv(EnvOrgID, "0")

		_, err := ResolveOrgID(0, IsRobotAuth())
		if err == nil {
			t.Fatal("expected error for zero env value, got nil")
		}
		if !strings.Contains(err.Error(), "positive integer") {
			t.Errorf("error should mention 'positive integer'; got: %s", err.Error())
		}
	})

	t.Run("negative value", func(t *testing.T) {
		clearRobotKeys(t)
		t.Setenv(EnvOrgID, "-5")

		_, err := ResolveOrgID(0, IsRobotAuth())
		if err == nil {
			t.Fatal("expected error for negative env value, got nil")
		}
		if !strings.Contains(err.Error(), "positive integer") {
			t.Errorf("error should mention 'positive integer'; got: %s", err.Error())
		}
	})

	t.Run("empty string treated as unset", func(t *testing.T) {
		// Empty string → env is treated as unset → follows standard/robot path.
		// For standard user this means an error (no org source).
		// For robot user this means passthrough.
		t.Run("standard user", func(t *testing.T) {
			clearRobotKeys(t)
			t.Setenv(EnvOrgID, "")

			_, err := ResolveOrgID(0, IsRobotAuth())
			if err == nil {
				t.Fatal("expected error for standard user with empty env (treated as unset)")
			}
		})

		t.Run("robot user", func(t *testing.T) {
			setRobotKeys(t)
			t.Setenv(EnvOrgID, "")

			orgID, err := ResolveOrgID(0, IsRobotAuth())
			if err != nil {
				t.Fatalf("expected passthrough for robot user with empty env, got: %v", err)
			}
			if orgID != 0 {
				t.Fatalf("expected orgID=0 for robot passthrough, got %d", orgID)
			}
		})
	})
}

// --------------------------------------------------------------------------
// T043 — Backward compatibility
// --------------------------------------------------------------------------
// Spec ref: FR-007, SC-001
func TestAcceptance_T043_BackwardCompatibility(t *testing.T) {
	t.Run("flag overrides env for standard user", func(t *testing.T) {
		// FR-007: scripts passing explicit flags keep working
		clearRobotKeys(t)
		t.Setenv(EnvOrgID, "100")

		orgID, err := ResolveOrgID(200, IsRobotAuth())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if orgID != 200 {
			t.Fatalf("expected 200 (flag), got %d", orgID)
		}
	})

	t.Run("flag works without env for standard user", func(t *testing.T) {
		// Explicit flag alone is sufficient, no env needed
		clearRobotKeys(t)
		t.Setenv(EnvOrgID, "")

		orgID, err := ResolveOrgID(200, IsRobotAuth())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if orgID != 200 {
			t.Fatalf("expected 200 (flag), got %d", orgID)
		}
	})

	t.Run("flag works for robot user", func(t *testing.T) {
		// Robot user with explicit flag → flag value used
		setRobotKeys(t)
		t.Setenv(EnvOrgID, "")

		orgID, err := ResolveOrgID(300, IsRobotAuth())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if orgID != 300 {
			t.Fatalf("expected 300 (flag), got %d", orgID)
		}
	})

	t.Run("robot user with env and flag prefers flag", func(t *testing.T) {
		// FR-002: explicit flag always wins regardless of auth type
		setRobotKeys(t)
		t.Setenv(EnvOrgID, "50")

		orgID, err := ResolveOrgID(300, IsRobotAuth())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if orgID != 300 {
			t.Fatalf("expected 300 (flag over env), got %d", orgID)
		}
	})
}
