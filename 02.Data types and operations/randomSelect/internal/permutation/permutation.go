package permutation

import (
	"fmt"
	"math/rand"
	"time"
)

type Permutation struct {
	a []int
	n int
	m int
}

func NewPermutation(n, m int) *Permutation {
	a := make([]int, n)

	return &Permutation{
		a: a,
		n: n,
		m: m,
	}
}

func (p *Permutation) Generate() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < p.n; i++ {
		p.a[i] = i
	}

	for i := 0; i < p.m; i++ {
		rnd := r.Intn(p.n-i) + 1
		p.a[rnd], p.a[i] = p.a[i], p.a[rnd]
	}
}

func (p *Permutation) SetSelectSize(m int) error {
	if m < 0 || m > p.n {
		return fmt.Errorf("select size must be in 0 .. %d", p.n)
	}

	p.m = m
	return nil
}

func (p *Permutation) Print() {
	for i := 0; i < p.m; i++ {
		fmt.Printf("%d ", p.a[i])
	}
	fmt.Println()
}
