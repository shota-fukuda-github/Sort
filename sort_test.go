package sort

import (
	"testing"
	"math/rand"
)

type TestStruct struct {
	name string
	array []int
	sort func(int, int) bool
	want []int
}

func TestBubbleSort(t *testing.T) {
    t.Parallel()
	tests := getCommonSortTestStruct()

	for _, tt := range tests {
        tt := tt
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got := BubbleSort(tt.array, tt.sort)
            if !isEqualIntSlice(got, tt.want) {
                t.Errorf("BubbleSort(%v) == %v, want %v", tt.array, got, tt.want)
            }
        })
	}
}

func TestSelectSort(t *testing.T) {
    t.Parallel()
	tests := getCommonSortTestStruct()

	for _, tt := range tests {
        tt := tt
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got := SelectSort(tt.array, tt.sort)
            if !isEqualIntSlice(got, tt.want) {
                t.Errorf("SelectSort(%v) == %v, want %v", tt.array, got, tt.want)
            }
        })
	}
}

func TestMergeSort(t *testing.T) {
    t.Parallel()
	tests := getCommonSortTestStruct()

	for _, tt := range tests {
        tt := tt
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got := MergeSort(tt.array, tt.sort)
            if !isEqualIntSlice(got, tt.want) {
                t.Errorf("MergeSort(%v) == %v, want %v", tt.array, got, tt.want)
            }
        })
	}
}

func TestQuickSort(t *testing.T) {
    t.Parallel()
	tests := getCommonSortTestStruct()

	for _, tt := range tests {
        tt := tt
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got := QuickSort(tt.array, tt.sort)
            if !isEqualIntSlice(got, tt.want) {
                t.Errorf("QuickSort(%v) == %v, want %v", tt.array, got, tt.want)
            }
        })
	}
}

func getCommonSortTestStruct() []TestStruct{
	return []TestStruct {
		{"配列の要素0: ASC", []int{}, SortAsc, []int{}},
		{"配列の要素1: ASC", []int{1}, SortAsc, []int{1}},
		{"配列の要素2: ASC", []int{1,2}, SortAsc, []int{1,2}},
		{"配列の要素3: ASC", []int{3,2,1}, SortAsc, []int{1,2,3}},
		{"同一値あり : ASC", []int{3,2,3}, SortAsc, []int{2,3,3}},
		{"マイナス値 : ASC", []int{-3,0,-5,1}, SortAsc, []int{-5,-3,0,1}},
		{"配列の要素0: DESC", []int{}, SortDesc, []int{}},
		{"配列の要素1: DESC", []int{2}, SortDesc, []int{2}},
		{"配列の要素2: DESC", []int{1,2}, SortDesc, []int{2,1}},
		{"配列の要素3: DESC", []int{3,2,1}, SortDesc, []int{3,2,1}},
		{"同一値あり : DESC", []int{3,2,3}, SortDesc, []int{3,3,2}},
		{"マイナス値 : DESC", []int{-3,0,-5,1}, SortDesc, []int{1,0,-3,-5}},
	}
}

func isEqualIntSlice(arrayA, arrayB []int) bool {
	if len(arrayA) != len(arrayB) {
		return false
	}

	for i, v := range arrayA {
		if v != arrayB[i] {
			return false
		}
	}
	return true
}


func BenchmarkBubbleSort(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		array := makeRandomSlice(1000)
		BubbleSort(array, SortAsc)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		array := makeRandomSlice(1000)
		SelectSort(array, SortAsc)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		array := makeRandomSlice(1000)
		QuickSort(array, SortAsc)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		array := makeRandomSlice(1000)
		MergeSort(array, SortAsc)
	}
}

func makeRandomSlice(n int) []int{
	array := make([]int, n)
	for range n {
		array = append(array, rand.Intn(n * 3))
	}
	return array
}