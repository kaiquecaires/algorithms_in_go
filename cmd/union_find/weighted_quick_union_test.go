package union_find_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kaiquecaires/algorithms_in_go/cmd/union_find"
)

func TestWeightedQuickUnion(t *testing.T) {
	wqu := union_find.NewWeightedQuickUnion(10)

	wqu.Union(1, 3)
	wqu.Union(3, 5)
	wqu.Union(3, 9)
	wqu.Union(0, 2)
	wqu.Union(1, 2)

	expectedID := []int{2, 3, 3, 3, 4, 3, 6, 7, 8, 3}
	if !reflect.DeepEqual(wqu.Id, expectedID) {
		t.Fatalf("expected id array %v, but got %v", expectedID, wqu.Id)
	}

	fmt.Println(wqu)
}
