package union_find

import "errors"

type QuickFind []int

func (qf *QuickFind) Union(p int, q int) error {
	if q < 0 || q > len(*qf)-1 || p < 0 || p > len(*qf)-1 {
		return errors.New("Invalid Params")
	}

	qid := (*qf)[q]
	pid := (*qf)[p]

	for i, num := range *qf {
		if num == pid {
			(*qf)[i] = qid
		}
	}

	return nil
}

func (qf *QuickFind) Connected(p int, q int) (bool, error) {
	if q < 0 || q > len(*qf)-1 || p < 0 || p > len(*qf)-1 {
		return false, errors.New("Invalid Params")
	}

	qid := (*qf)[q]
	pid := (*qf)[p]

	return qid == pid, nil
}

func NewQuickFind(n int) *QuickFind {
	qf := make(QuickFind, n)
	for i := 0; i < n; i++ {
		qf[i] = i
	}
	return &qf
}
