package asynccache

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAsyncCacheSet(t *testing.T) {
	ac := InitAsyncCache(5*time.Minute, 10*time.Minute)

	ctxBase := context.Background()
	ctx, _ := context.WithTimeout(ctxBase, time.Nanosecond/10)

	err := ac.Add(ctx, "guest", "id: 5456")
	if err != ErrTimeout {
		t.Error("Expected timeout error")
	} else {
		t.Log("timeout ok")
	}

	ctx, _ = context.WithTimeout(ctxBase, time.Millisecond*2)

	err = ac.Add(ctx, "guest", "id: 5456")
	if err != nil {
		t.Errorf("Expected success: %v", err.Error())
	}

}

func TestAsyncCacheGet(t *testing.T) {
	ac := InitAsyncCache(5*time.Minute, 10*time.Minute)

	ctxBase := context.Background()
	ctx, _ := context.WithTimeout(ctxBase, time.Nanosecond/10)

	_ = ac.Add(ctxBase, "guest", "v")
	_, err := ac.Get(ctx, "guest")
	if err != ErrTimeout {
		t.Error("Expected timeout error")
	} else {
		t.Log("timeout ok")
	}

	ctx, _ = context.WithTimeout(ctxBase, time.Millisecond*2)

	v, err := ac.Get(ctx, "guest")
	if err != nil {
		t.Error("Expected success")
	}

	assert.Equal(t, "v", v)
}

func TestAsyncCacheDelete(t *testing.T) {
	ac := InitAsyncCache(5*time.Minute, 10*time.Minute)

	ctxBase := context.Background()
	ctx, _ := context.WithTimeout(ctxBase, time.Nanosecond/10)

	_ = ac.Add(ctxBase, "guest", "v")

	err := ac.Delete(ctx, "guest")
	if err != ErrTimeout {
		t.Error("Expected timeout error")
	} else {
		t.Log("timeout ok")
	}

	err = ac.Delete(ctxBase, "guest")
	if err != nil {
		t.Error("Expected success")
	}
}
