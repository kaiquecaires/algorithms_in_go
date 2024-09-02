package union_find

import "errors"

type WeightedQuickUnion struct {
	id     []int
	weight []int
}

func (wqu *WeightedQuickUnion) Union(p, q int) error {
	if p < 0 || p > len(wqu.id)-1 || q < 0 || q > len(wqu.id)-1 {
		return errors.New("invalid params")
	}

	pr := wqu.root(p)
	qr := wqu.root(q)

	if wqu.weight[pr] <= wqu.weight[qr] {
		wqu.id[pr] = wqu.id[qr]
		wqu.weight[qr] += wqu.weight[pr]
	} else {
		wqu.id[qr] = wqu.id[pr]
		wqu.weight[pr] += wqu.weight[qr]
	}

	return nil
}

func (wqu *WeightedQuickUnion) Connected(p int, q int) (bool, error) {
	if q < 0 || q > len(wqu.id)-1 || p < 0 || p > len(wqu.id)-1 {
		return false, errors.New("invalid Params")
	}

	pr := wqu.root(p)
	qr := wqu.root(q)

	return pr == qr, nil
}

func (wqu *WeightedQuickUnion) root(i int) int {
	for i != wqu.id[i] {
		i = wqu.id[i]
	}
	return i
}

func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	wqu := &WeightedQuickUnion{
		id:     make([]int, 10),
		weight: make([]int, 10),
	}

	for i := 0; i < len(wqu.id); i++ {
		wqu.id[i] = i
		wqu.weight[i] = 1
	}

	return wqu
}
