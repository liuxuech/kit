package redis

import (
	"fmt"
	"testing"
)

func TestNewCache(t *testing.T) {
	newCache := NewCache()

	if err := newCache.Set("name", "刘学超"); err != nil {
		t.Error(err)
	}

	if v, err := newCache.Get("name"); err != nil {
		t.Error(err)
	} else {
		fmt.Println(v)
	}
}
