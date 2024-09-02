package union_find

import (
	"errors"
	"fmt"
)

type QuickUnion []int

func (qu *QuickUnion) Union(p int, q int) error {
	if q < 0 || q > len(*qu)-1 || p < 0 || p > len(*qu)-1 {
		return errors.New("invalid Params")
	}

	i, err := qu.Root(p)
	if err != nil {
		return err
	}

	j, err := qu.Root(q)
	if err != nil {
		return err
	}

	fmt.Println(i, j)

	(*qu)[i] = j
	return nil
}

func (qu *QuickUnion) Connected(p int, q int) (bool, error) {
	if q < 0 || q > len(*qu)-1 || p < 0 || p > len(*qu)-1 {
		return false, errors.New("invalid Params")
	}

	pRoot, err := qu.Root(p)
	if err != nil {
		return false, err
	}

	qRoot, err := qu.Root(q)
	if err != nil {
		return false, err
	}

	return pRoot == qRoot, nil
}

func (qu *QuickUnion) Root(n int) (int, error) {
	if n < 0 || n > len(*qu)-1 {
		return 0, errors.New("invalid Param")
	}
	i := n
	for i != (*qu)[i] {
		i = (*qu)[i]
	}
	return i, nil
}

func NewQuickUnion(n int) *QuickUnion {
	qu := make(QuickUnion, n)
	for i := 0; i < n; i++ {
		qu[i] = i
	}
	return &qu
}
