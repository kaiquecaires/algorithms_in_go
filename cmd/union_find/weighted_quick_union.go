package union_find

import "errors"

type WeightedQuickUnion struct {
	Id     []int
	Weight []int
}

func (wqu *WeightedQuickUnion) Union(p, q int) error {
	if p < 0 || p > len(wqu.Id)-1 || q < 0 || q > len(wqu.Id)-1 {
		return errors.New("invalid params")
	}

	pr := wqu.root(p)
	qr := wqu.root(q)

	if pr == qr {
		return nil
	}

	if wqu.Weight[pr] <= wqu.Weight[qr] {
		wqu.Id[pr] = wqu.Id[qr]
		wqu.Weight[qr] += wqu.Weight[pr]
	} else {
		wqu.Id[qr] = wqu.Id[pr]
		wqu.Weight[pr] += wqu.Weight[qr]
	}

	return nil
}

func (wqu *WeightedQuickUnion) Connected(p int, q int) (bool, error) {
	if q < 0 || q > len(wqu.Id)-1 || p < 0 || p > len(wqu.Id)-1 {
		return false, errors.New("invalid Params")
	}

	return wqu.root(p) == wqu.root(q), nil
}

func (wqu *WeightedQuickUnion) root(i int) int {
	for i != wqu.Id[i] {
		// path compression
		wqu.Id[i] = wqu.Id[wqu.Id[i]]
		i = wqu.Id[i]
	}
	return i
}

func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	wqu := &WeightedQuickUnion{
		Id:     make([]int, 10),
		Weight: make([]int, 10),
	}

	for i := 0; i < len(wqu.Id); i++ {
		wqu.Id[i] = i
		wqu.Weight[i] = 1
	}

	return wqu
}
