package semver

import "testing"

// Regression: prerelease comparison must give higher precedence
// to the version with MORE segments when all shared segments are equal.
// Per SemVer 2.0.0 spec section 11.4.4.

func TestComparePreSegmentCount(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		// Fewer segments = lower precedence
		{"1.0.0-alpha", "1.0.0-alpha.1", -1},
		{"1.0.0-alpha.1", "1.0.0-alpha", 1},
		// Same segments = equal
		{"1.0.0-alpha.1", "1.0.0-alpha.1", 0},
		// Real-world chain from SemVer spec
		{"1.0.0-beta", "1.0.0-beta.2", -1},
		{"1.0.0-beta.2", "1.0.0-beta.11", -1},
		// Multiple extra segments
		{"1.0.0-a.b", "1.0.0-a.b.c", -1},
	}
	for _, c := range cases {
		a, err := Parse(c.a)
		if err != nil {
			t.Fatalf("Parse(%q): %v", c.a, err)
		}
		b, err := Parse(c.b)
		if err != nil {
			t.Fatalf("Parse(%q): %v", c.b, err)
		}
		if got := Compare(a, b); got != c.want {
			t.Errorf("Compare(%q, %q) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
