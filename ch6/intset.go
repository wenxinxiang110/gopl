package ch6

import (
	"bytes"
	"fmt"
	"gopl/ch2"
	"math"
)

//IntSet 模拟的集合操作
type IntSet struct {
	words []uint64
}

// 序列化
func (s *IntSet) String() string {
	var buf bytes.Buffer

	buf.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)

	// 扩容，这方法好蠢
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, t.words[i:]...)
			break
		}
	}
}

// 交集,元素在A集合B集合均出现 pass
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			break
		}

		s.words[i] &= tword
	}

	//	todo: 需要更加高效的方法
	for j := len(t.words); j < len(s.words); j++ {
		s.words[j] = 0
	}

}

// 差集：元素出现在s集合，未出现在t集合
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i >= len(t.words) {
			break
		}
		s.words[i] &= ^t.words[i]
	}
}

//并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			s.words = append(s.words, t.words[i:]...)
			break
		}

		s.words[i] ^= tword

	}
}

// return the number of elements
func (s *IntSet) Len() (count int) {
	for _, word := range s.words {
		if word != 0 {
			count += ch2.PopCount(word)
		}
	}
	return
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)

	if word < len(s.words) {
		s.words[word] &= math.MaxUint64 ^ (1 << bit)
	}

}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	//var another []uint64
	another := make([]uint64, len(s.words))
	for index, word := range s.words {
		another[index] = word
	}
	return &IntSet{words: another}
}

func (s *IntSet) AddAll(nums ...int) *IntSet {
	for _, num := range nums {
		s.Add(num)
	}

	return s
}

func (s *IntSet) Elems() (elems []int) {

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}

	return
}
