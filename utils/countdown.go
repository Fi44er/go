package utils

import "iter"

func Countdown(v []string) iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for index := 0; index < len(v); index++ {
			if !yield(v[index], index) {
				return
			}
		}
	}
}
