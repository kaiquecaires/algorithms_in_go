package union_find_test

import (
	"fmt"
	"testing"

	"github.com/kaiquecaires/algorithms_in_go/cmd/union_find"
)

func TestWeightedQuickUnion(t *testing.T) {
	wqf := union_find.NewWeightedQuickUnion(10)
	wqf.Union(1, 3)
	wqf.Union(3, 5)
	wqf.Union(3, 9)
	wqf.Union(0, 2)
	wqf.Union(1, 2)

	if connected, _ := wqf.Connected(1, 5); !connected {
		t.Fatalf("1 and 5 should be connected")
	}

	if connected, _ := wqf.Connected(1, 9); !connected {
		t.Fatalf("1 and 5 should be connected")
	}

	if connected, _ := wqf.Connected(1, 8); connected {
		t.Fatalf("1 and 5 should not be connected")
	}

	if _, error := wqf.Connected(-1, 6); error == nil {
		t.Fatalf("error should not be nil")
	}

	if _, error := wqf.Connected(1, 10); error == nil {
		t.Fatalf("error should not be nil")
	}

	fmt.Println(wqf)
}
