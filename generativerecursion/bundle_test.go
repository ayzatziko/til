package generativerecursion

import (
	"reflect"
	"testing"
)

func TestBundles(t *testing.T) {
	type in struct {
		chars []string
		n     int
	}
	tests := []struct {
		in  in
		out []string
	}{
		{in{[]string{"a", "b", "c", "d"}, 3}, []string{"abc", "d"}},
		{in{[]string{"a", "b", "c", "d"}, 5}, []string{"abcd"}},
	}

	bundleFuncs := map[string]func([]string, int) []string{
		"main": Bundle,
		"v1":   bundleV1,
		"v2":   bundleV2,
		"v3":   bundleV3,
	}

	for i, tt := range tests {
		for ver, bundle := range bundleFuncs {
			o := bundle(tt.in.chars, tt.in.n)
			if !reflect.DeepEqual(tt.out, o) {
				t.Errorf("test: %d, version %s: want %v, got %v\n", i, ver, tt.out, o)
			}
		}
	}
}
