package sort

import (
	"math/rand"
	"testing"
)

type TestStruct[T Number] struct {
	name  string
	array []T
	sort  func(T, T) bool
	want  []T
}

func TestIntBubbleSort(t *testing.T) {
	t.Parallel()
	tests := getCommonIntSortTestStruct()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BubbleSort(tt.array, tt.sort)
			if !isEqualSlice(got, tt.want) {
				t.Errorf("BubbleSort(%v) == %v, want %v", tt.array, got, tt.want)
			}
		})
	}
}

func TestFolatBubbleSort(t *testing.T) {
	t.Parallel()
	tests := getCommonFloatSortTestStruct()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BubbleSort(tt.array, tt.sort)
			if !isEqualSlice(got, tt.want) {
				t.Errorf("BubbleSort(%v) == %v, want %v", tt.array, got, tt.want)
			}
		})
	}
}

func TestIntSelectSort(t *testing.T) {
	t.Parallel()
	tests := getCommonIntSortTestStruct()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := SelectSort(tt.array, tt.sort)
			if !isEqualSlice(got, tt.want) {
				t.Errorf("SelectSort(%v) == %v, want %v", tt.array, got, tt.want)
			}
		})
	}
}

func TestFloatSelectSort(t *testing.T) {
	t.Parallel()
	tests := getCommonFloatSortTestStruct()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := SelectSort(tt.array, tt.sort)
			if !isEqualSlice(got, tt.want) {
				t.Errorf("SelectSort(%v) == %v, want %v", tt.array, got, tt.want)
			}
		})
	}
}

func TestIntQuickSort(t *testing.T) {
	t.Parallel()
	tests := getCommonIntSortTestStruct()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := QuickSort(tt.array, tt.sort)
			if !isEqualSlice(got, tt.want) {
				t.Errorf("QuickSort(%v) == %v, want %v", tt.array, got, tt.want)
			}
		})
	}
}

func TestFloatQuickSort(t *testing.T) {
	t.Parallel()
	tests := getCommonFloatSortTestStruct()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := QuickSort(tt.array, tt.sort)
			if !isEqualSlice(got, tt.want) {
				t.Errorf("QuickSort(%v) == %v, want %v", tt.array, got, tt.want)
			}
		})
	}
}

func TestIntMergeSort(t *testing.T) {
	t.Parallel()
	tests := getCommonIntSortTestStruct()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MergeSort(tt.array, tt.sort)
			if !isEqualSlice(got, tt.want) {
				t.Errorf("MergeSort(%v) == %v, want %v", tt.array, got, tt.want)
			}
		})
	}
}

func TestFloatMergeSort(t *testing.T) {
	t.Parallel()
	tests := getCommonFloatSortTestStruct()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := MergeSort(tt.array, tt.sort)
			if !isEqualSlice(got, tt.want) {
				t.Errorf("MergeSort(%v) == %v, want %v", tt.array, got, tt.want)
			}
		})
	}
}

func TestIntHeapSort(t *testing.T) {
	t.Parallel()
	tests := getCommonIntSortTestStruct()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := HeapSort(tt.array, tt.sort)
			if !isEqualSlice(got, tt.want) {
				t.Errorf("HeapSort(%v) == %v, want %v", tt.array, got, tt.want)
			}
		})
	}
}

func TestFloatHeapSort(t *testing.T) {
	t.Parallel()
	tests := getCommonFloatSortTestStruct()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := HeapSort(tt.array, tt.sort)
			if !isEqualSlice(got, tt.want) {
				t.Errorf("HeapSort(%v) == %v, want %v", tt.array, got, tt.want)
			}
		})
	}
}

func getCommonIntSortTestStruct() []TestStruct[int] {
	return []TestStruct[int]{
		{"配列の要素0: ASC", []int{}, Asc[int], []int{}},
		{"配列の要素1: ASC", []int{1}, Asc[int], []int{1}},
		{"配列の要素2: ASC", []int{1, 2}, Asc[int], []int{1, 2}},
		{"配列の要素3: ASC", []int{3, 2, 1}, Asc[int], []int{1, 2, 3}},
		{"同一値あり : ASC", []int{3, 2, 3}, Asc[int], []int{2, 3, 3}},
		{"マイナス値 : ASC", []int{-3, 0, -5, 1}, Asc[int], []int{-5, -3, 0, 1}},
		{"配列の要素0: DESC", []int{}, Desc[int], []int{}},
		{"配列の要素1: DESC", []int{2}, Desc[int], []int{2}},
		{"配列の要素2: DESC", []int{1, 2}, Desc[int], []int{2, 1}},
		{"配列の要素3: DESC", []int{3, 2, 1}, Desc[int], []int{3, 2, 1}},
		{"同一値あり : DESC", []int{3, 2, 3}, Desc[int], []int{3, 3, 2}},
		{"マイナス値 : DESC", []int{-3, 0, -5, 1}, Desc[int], []int{1, 0, -3, -5}},
	}
}

func getCommonFloatSortTestStruct() []TestStruct[float64] {
	return []TestStruct[float64]{
		{"配列の要素0: ASC", []float64{}, Asc[float64], []float64{}},
		{"配列の要素1: ASC", []float64{1.0}, Asc[float64], []float64{1.0}},
		{"配列の要素2: ASC", []float64{1.1, 2.1}, Asc[float64], []float64{1.1, 2.1}},
		{"配列の要素3: ASC", []float64{3, 2, 1}, Asc[float64], []float64{1, 2, 3}},
		{"同一値あり : ASC", []float64{3.3, 2.33, 3.3}, Asc[float64], []float64{2.33, 3.3, 3.3}},
		{"マイナス値 : ASC", []float64{-3.22, 0.22, -5.22, 1.22}, Asc[float64], []float64{-5.22, -3.22, 0.22, 1.22}},
		{"配列の要素0: DESC", []float64{}, Desc[float64], []float64{}},
		{"配列の要素1: DESC", []float64{2.44}, Desc[float64], []float64{2.44}},
		{"配列の要素2: DESC", []float64{1.1234, 2.1234}, Desc[float64], []float64{2.1234, 1.1234}},
		{"配列の要素3: DESC", []float64{3.123, 2.123, 1.123}, Desc[float64], []float64{3.123, 2.123, 1.123}},
		{"同一値あり : DESC", []float64{3.123, 2.123, 3.123}, Desc[float64], []float64{3.123, 3.123, 2.123}},
		{"マイナス値 : DESC", []float64{-3.4321, 0.43, -5.1, 1.123}, Desc[float64], []float64{1.123, 0.43, -3.4321, -5.1}},
	}
}

func isEqualSlice[T comparable](arrayA, arrayB []T) bool {
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

func makeRandomSlice(n int) []float64 {
	array := make([]float64, n)
	for range n {
		array = append(array, rand.Float64())
	}
	return array
}
