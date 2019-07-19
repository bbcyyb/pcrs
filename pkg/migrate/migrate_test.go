package migrate

import (
	"testing"
)

func TestNumDownFromArgs(t *testing.T) {
	cases := []struct {
		name                string
		arg                 string
		expectedNum         int
		expectedNeedConfirm bool
		expectedErrStr      string
	}{
		{"no args", "", 0, false, "can't read limit argument N, only accept num or 'all'"},
		{"down all", "all", -1, true, ""},
		{"down 5", "5", 5, false, ""},
		{"down N", "N", 0, false, "can't read limit argument N, only accept num or 'all'"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			num, needsConfirm, err := numDownMigrationsFromArgs(c.arg)
			if needsConfirm != c.expectedNeedConfirm {
				t.Errorf("Incorrect needsConfirm was: %v wanted %v", needsConfirm, c.expectedNeedConfirm)
			}

			if num != c.expectedNum {
				t.Errorf("Incorrect num was: %v wanted %v", num, c.expectedNum)
			}

			if err != nil {
				if err.Error() != c.expectedErrStr {
					t.Error("Incorrect error: " + err.Error() + " != " + c.expectedErrStr)
				}
			} else if c.expectedErrStr != "" {
				t.Error("Expected error: " + c.expectedErrStr + " but got nil instead")
			}
		})
	}
}
