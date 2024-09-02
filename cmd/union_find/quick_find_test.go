package union_find_test

import (
	"testing"

	"github.com/kaiquecaires/algorithms_in_go/cmd/union_find"
)

func TestQuickFind(t *testing.T) {
	qf := union_find.NewQuickFind(10)
	qf.Union(1, 3)
	qf.Union(3, 5)
	qf.Union(3, 9)

	if connected, _ := qf.Connected(1, 5); !connected {
		t.Fatalf("1 and 5 should be connected")
	}

	if connected, _ := qf.Connected(1, 9); !connected {
		t.Fatalf("1 and 5 should be connected")
	}

	if connected, _ := qf.Connected(1, 8); connected {
		t.Fatalf("1 and 5 should not be connected")
	}

	if _, error := qf.Connected(-1, 6); error == nil {
		t.Fatalf("error should not be nil")
	}

	if _, error := qf.Connected(1, 10); error == nil {
		t.Fatalf("error should not be nil")
	}
}
