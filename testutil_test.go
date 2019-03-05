package schema_test

import "testing"

// noErr returns a sub-test that fails if there is an error.
func noErr(err error) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()

		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}
}

// errMatch returns a sub-test that fails if err.Error() does not match
// the expected value.
func errMatch(err error, expect string) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()

		if err == nil || err.Error() != expect {
			t.Errorf("\n got: %v\nwant: %s", err, expect)
		}
	}
}

// isNil returns a sub-test that fails if v is not nil.
func isNill(v interface{}) func(t *testing.T) {
	return func(t *testing.T) {
		t.Helper()

		if v != nil {
			t.Errorf("\n got: %v\nwant: <nil>", v)
		}
	}
}
