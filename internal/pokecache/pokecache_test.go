package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 100)

	if cache.cache == nil {
		t.Error("expected cache to be initialized")
	}
}

func TestAddAndGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 100)

	cases := []struct {
		key   string
		value []byte
	}{
		{"key1", []byte("value1")},
		{"key2", []byte("value2")},
		{"key3", []byte("value3")},
	}

	for _, tt := range cases {
		cache.Add(tt.key, tt.value)
		value, exists := cache.Get(tt.key)

		if !exists {
			t.Errorf("%s expected key to exist in cache", tt.key)
		}

		if string(value) != string(tt.value) {
			t.Errorf("expected value to be 'value1', got %s", string(value))
		}
	}
}

func TestReapCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 50)

	cases := []struct {
		key   string
		value []byte
	}{
		{"key1", []byte("value1")},
		{"key2", []byte("value2")},
		{"key3", []byte("value3")},
	}

	for _, tt := range cases {
		cache.Add(tt.key, tt.value)
	}

	if len(cache.cache) != 3 {
		t.Error("expected cache to have 3 entries")
	}

	time.Sleep(time.Millisecond * 100)

	if len(cache.cache) != 0 {
		t.Error("expected cache to be empty")
	}
}
