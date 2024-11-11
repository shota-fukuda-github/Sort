package sort


import (
    "slices"
)

// バブルソート
func BubbleSort(array []int, defineSortFunc func(int, int) bool) []int {
    for i := range array {
        for j := i+1; j < len(array); j++ {
            if defineSortFunc(array[i], array[j]) {
                tmp := array[j]
                array[j] = array[i]
                array[i] = tmp
            }
        }
    }
    return array
}

// 選択ソート
func SelectSort(array []int, defineSortFunc func(int, int) bool) []int {
    for i := range array {
        tmp := array[i]
        idx := i
        for j:= i+1; j < len(array); j++ {
            if defineSortFunc(tmp, array[j]) {
                tmp = array[j]
                idx = j
            }
        }
        array[idx] = array[i]
        array[i]   = tmp
    }
    return array
}

// クイックソート
func QuickSort(array []int, defineSortFunc func(int, int) bool) []int {
    count := len(array)
    if count <= 1 {
        return array
    }
    // 境界は先頭の要素を使用
    pivot := array[0]
    left  := []int{}
    right := []int{}
    for i:= 1; i < count; i++ {
        if defineSortFunc(pivot, array[i]) {
            left = append(left, array[i])
        } else {
            right = append(right, array[i]) 
        }
    }
    SortedLeft  := QuickSort(left, defineSortFunc)
    SortedRight := QuickSort(right, defineSortFunc)
    return slices.Concat(append(SortedLeft, pivot), SortedRight)
}


// マージソート
func MergeSort(array []int, defineSortFunc func(int, int) bool) []int {
    count := len(array)
    if count <= 1 {
        return array
    }
    divNum := count/2
    left   := array[:divNum]
    right  := array[divNum:]

    sortedLeft  := MergeSort(left, defineSortFunc)
    sortedRight := MergeSort(right, defineSortFunc)
    sorted := make([]int, 0, count)

    // 右の要素番号保存
    j := 0
    for i := range sortedLeft {
        for ;j < len(sortedRight); j++ {
            if defineSortFunc(sortedLeft[i], sortedRight[j]) {
                sorted = append(sorted, sortedRight[j])
                // 右の要素が最後だった場合、残る左をすべて入れてreturn
                if j+1 == len(sortedRight) {
                    return slices.Concat(sorted, sortedLeft[i:]);
                }
            } else {
                sorted = append(sorted, sortedLeft[i])
                // 左の要素が最後だった場合、残る右をすべて入れてreturn
                if i+1 == len(sortedLeft) {
                    return slices.Concat(sorted, sortedRight[j:]);
                }
                // 次の右の要素も見るのをやめて、左の要素を進める
                break
            }
        }
    }
    return sorted
}

// 昇順用の関数
func SortDesc(num1 int, num2 int) bool {
    return num1 < num2
}

// 降順用の関数
func SortAsc(num1 int, num2 int) bool {
    return num1 > num2
}