package server

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/sourcegraph/jsonrpc2"
)

func TestLSPHandler(t *testing.T) {
	wg := sync.WaitGroup{}
	received := make(chan string, 20)
	mockHandler := &MockHandler{
		handler: func(ctx context.Context, conn *jsonrpc2.Conn, request *jsonrpc2.Request) {
			t.Logf("Received request: %s", request.Method)
			time.Sleep(5 * time.Millisecond)
			received <- request.Method
			wg.Done()
		},
	}

	ah := newLSPHandler(mockHandler)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		ah.Handle(context.Background(), &jsonrpc2.Conn{}, &jsonrpc2.Request{
			Method: fmt.Sprintf("call-%d", i),
		})
	}

	wg.Wait()
	close(received)
	t.Log("heard all")

	var ordered []string
	for v := range received {
		ordered = append(ordered, v)
	}
	for i := 0; i < 10; i++ {
		if ordered[i] != fmt.Sprintf("call-%d", i) {
			t.Errorf("Expected call-%d but got %v", i, ordered)
		}
	}
}

func TestLSPHandler_Cancel(t *testing.T) {
	done := make(chan struct{})
	firstCallCtx, cancel := context.WithCancel(context.Background())

	mockHandler := &MockHandler{
		handler: func(ctx context.Context, conn *jsonrpc2.Conn, request *jsonrpc2.Request) {
			switch request.Method {
			case "$/cancelRequest":
				cancel()
			case "call":
				<-ctx.Done()
				close(done)
			}
		},
	}

	ah := newLSPHandler(mockHandler)
	ah.Handle(firstCallCtx, &jsonrpc2.Conn{}, &jsonrpc2.Request{
		Method: "call",
	})

	ah.Handle(context.Background(), &jsonrpc2.Conn{}, &jsonrpc2.Request{
		Method: "$/cancelRequest",
	})

	select {
	case <-time.After(50 * time.Millisecond):
		t.Errorf("expected request to be cancelled")
	case <-done:
	}
}

type MockHandler struct {
	handler func(ctx context.Context, conn *jsonrpc2.Conn, request *jsonrpc2.Request)
}

func (m *MockHandler) Handle(ctx context.Context, conn *jsonrpc2.Conn, request *jsonrpc2.Request) {
	m.handler(ctx, conn, request)
}
