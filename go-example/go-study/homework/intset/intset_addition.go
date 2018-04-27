package intset

//import (
//)

func (s *IntSet) Len() int {
    l := 0;
    for _, word := range s.words {
        if word == 0 {
            continue
        }
        for j := 0; j < 64; j++ {
            if word&(1<<uint(j)) != 0 {
                l++
            }
        }
    }
    return l
}

func (s *IntSet) Remove(x int) {
    word, bit := x/64, uint(x%64)
    if word >= len(s.words) {
        return
    }
    s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
    for i, _ := range s.words {
        s.words[i] = 0
    }
}

func (s *IntSet) Copy() *IntSet {
    t := new(IntSet)
    for _, word := range s.words {
        t.words  = append(t.words, word)
    }
    return t
}
