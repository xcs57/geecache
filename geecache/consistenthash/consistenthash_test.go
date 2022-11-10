package consistenthash

import (
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	hash := New(3, func(data []byte) uint32 {
		i, err := strconv.Atoi(string(data))
		if err != nil {
			t.Fatalf("inner error")
		}
		return uint32(i)
	})

	hash.Add("6", "4", "2")

	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Ask for %v,should be %v, but is %v", k, v, hash.Get(k))
		}

	}

	// 27 should now map to 8.
	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

}
