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

func TestHeapSort(t *testing.T) {
    t.Parallel()
	tests := getCommonSortTestStruct()

	for _, tt := range tests {
        tt := tt
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got := HeapSort(tt.array, tt.sort)
            if !isEqualIntSlice(got, tt.want) {
                t.Errorf("HeapSort(%v) == %v, want %v", tt.array, got, tt.want)
            }
        })
	}
}

func getCommonSortTestStruct() []TestStruct{
	return []TestStruct {
		{"配列の要素0: ASC", []int{}, Asc, []int{}},
		{"配列の要素1: ASC", []int{1}, Asc, []int{1}},
		{"配列の要素2: ASC", []int{1,2}, Asc, []int{1,2}},
		{"配列の要素3: ASC", []int{3,2,1}, Asc, []int{1,2,3}},
		{"同一値あり : ASC", []int{3,2,3}, Asc, []int{2,3,3}},
		{"マイナス値 : ASC", []int{-3,0,-5,1}, Asc, []int{-5,-3,0,1}},
		{"配列の要素0: DESC", []int{}, Desc, []int{}},
		{"配列の要素1: DESC", []int{2}, Desc, []int{2}},
		{"配列の要素2: DESC", []int{1,2}, Desc, []int{2,1}},
		{"配列の要素3: DESC", []int{3,2,1}, Desc, []int{3,2,1}},
		{"同一値あり : DESC", []int{3,2,3}, Desc, []int{3,3,2}},
		{"マイナス値 : DESC", []int{-3,0,-5,1}, Desc, []int{1,0,-3,-5}},
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
		array := makeRandomSlice(10000)
		BubbleSort(array, Asc)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		array := makeRandomSlice(10000)
		SelectSort(array, Asc)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		array := makeRandomSlice(10000)
		QuickSort(array, Asc)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		array := makeRandomSlice(10000)
		MergeSort(array, Asc)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		array := makeRandomSlice(10000)
		HeapSort(array, Asc)
	}
}

func makeRandomSlice(n int) []int{
	array := make([]int, n)
	for range n {
		array = append(array, rand.Intn(n * 3))
	}
	return array
}