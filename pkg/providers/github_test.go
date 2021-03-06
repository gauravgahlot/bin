package providers

import "testing"

func TestSanitizeName(t *testing.T) {
	cases := []struct {
		in  string
		v   string
		out string
	}{
		{"bin_amd64_linux", "v0.0.1", "bin"},
		{"bin_0.0.1_amd64_linux", "0.0.1", "bin"},
		{"bin_0.0.1_amd64_linux", "v0.0.1", "bin"},
	}

	for _, c := range cases {
		if n := sanitizeName(c.in, c.v); n != c.out {
			t.Fatalf("Error replacing %s: %s does not match %s", c.in, n, c.out)
		}
	}

}
