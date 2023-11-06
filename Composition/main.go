package main

import (
	"fmt"
	"sort"
)

// The fields of an embedded structs get promoted to the level of embedding structure!

// func main() {

// 	type Pair struct {
// 		Hash string
// 		Path string
// 	}

// 	type PairWithLength struct {
// 		Pair
// 		Length int
// 	}

// 	p1 := PairWithLength{Pair{"Hello", "End"}, 21}
// 	fmt.Println(p1.Path, p1.Length) // Not p1.x.Path

// }

// type Pair struct {
// 	Path string
// 	Hash string
// }

// type PairWithLength struct {
// 	Pair
// 	Length int
// }

// type FileNamer interface {
// 	FileName() string
// }

// func (p Pair) String() string {
// 	return fmt.Sprintf("Path is :%s ||||| Hash is %s", p.Path, p.Hash)
// }

// func (p PairWithLength) String() string {
// 	return fmt.Sprintf("Path is :%s ||||| Hash is %s, Length: %d ", p.Path, p.Hash, p.Length)
// }

// func (p Pair) FileName() string {
// 	return filepath.Base(p.Path)
// }

// func main() {
// 	p := Pair{"/usr", "0xcdd"}
// 	k1 := PairWithLength{Pair{"/usr/time1/xyz/", "0xcddd23"}, 23}

// 	var fn FileNamer = PairWithLength{Pair{"/usr/time1/xyz/", "0xcddd23"}, 23}
// 	fmt.Println(p)
// 	fmt.Println(k1)
// 	//	fmt.Println(FileName(k1.Pair)) // In go there is no Inheritence-----> you can't do fmt.Println(FileName(k1)) despite
// 	// PairWithLength having Pair in it.
// 	fmt.Println(fn)
// 	fmt.Println(p.FileName())
// }

// COMPOSITION WITH POINTER TYPES!!!!!!!!

// type Fizgig struct {
// 	*PairWithLength
// 	Broken bool
// }

// type Pair struct {
// 	Path string
// 	Hash string
// }

// type PairWithLength struct {
// 	Pair
// 	Length int
// }

// func (p PairWithLength) String() string {
// 	return fmt.Sprintf("Path is :%s ||||| Hash is %s, Length: %d", p.Path, p.Hash, p.Length)
// }
// func (f Fizgig) String() string {
// 	return fmt.Sprintf("%s, Broken: %v", f.PairWithLength.String(), f.Broken)
// }

// func main() {
// 	fg := Fizgig{&PairWithLength{Pair{"/usr", "Oxfede"}, 121}, false}
// 	fmt.Println(fg)
// }

// SORTABLE INTERFACE

// func main() {
// 	enteries := []string{"Hello", "Aviral", "Nice", "to", "meet", "you!"}
// 	sorted := sort.StringSlice(enteries)
// 	sorted.Sort()
// 	fmt.Println(enteries)
// }

// IMPLEMENT sort.Interface to make a type sortable!

type Organ struct {
	Name   string
	Weight string
}

type Organs []Organ

func (s Organs) Len() int {
	return len(s)
}

func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByName struct {
	Organs
}

type ByWeight struct {
	Organs
}

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {
	s := []Organ{{"Brain", "600"}, {"Liver", "300"}}
	sort.Sort(ByWeight{s})
	fmt.Println(s)
	sort.Sort(ByName{s})
	fmt.Println(s)
}

// NOTHING IN GO PREVENTS CALLING A METHOD WITH A NIL RECEVIER!!!!

// ADDING SUM OF ELEMENTS IN A LINKED LIST

type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the List elements!!

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
