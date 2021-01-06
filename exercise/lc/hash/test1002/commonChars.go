package test1002

import "math"

func commonChars(A []string) (res []string) {
	table := [26]int{}

	for i := range table {
		table[i] = math.MaxInt64
	}
	for _, word := range A {
		p := [26]int{}
		for _, b := range word {
			p[b-'a']++
		}
		for i, f := range p[:] {
			if f < table[i] {
				table[i] = f
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < table[i]; j++ {
			res = append(res, string('a'+i))
		}
	}
	return
}
