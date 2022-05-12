package wzqb

import "math/rand"

// 冒泡排序 (bubble sort)
// 原理：遍历数组，比较并将大的元素与下一个元素交换位置， 在一轮的循环之后，可以让未排序i的最大元素排列到数组右侧。在一轮循环中，如果没有发生元素位置交换，那么说明数组已经是有序的，此时退出排序。
func BubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}

// 插入排序 (insertion sort)
// 原理：数组先看成两部分，排序序列和未排序序列。排序序列从第一个元素开始，该元素可以认为已经被排序。遍历数组， 每次将扫描到的元素与之前的元素相比较，插入到有序序列的适当位置。
func IntertionSort(arr []int) []int {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		temporary := arr[currentIndex]
		iterator := currentIndex
		for ; iterator > 0 && arr[iterator-1] >= temporary; iterator-- {
			arr[iterator] = arr[iterator-1]
		}
		arr[iterator] = temporary
	}
	return arr
}

// 三向切分快速排序 (quick sort)
// 原理：快速排序也是分治法的一个应用，先随机拿到一个基准 pivot，通过一趟排序将数组分成两个独立的数组，左子数组小于或等于 pivot，右子数组大于等于 pivot。 然后可在对这两个子数组递归继续以上排序，最后使整个数组有序。
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < pivot:
			lowPart = append(lowPart, item)
		case item == pivot:
			middlePart = append(middlePart, item)
		case item > pivot:
			highPart = append(highPart, item)
		}
	}
	lowPart = QuickSort(lowPart)
	highPart = QuickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}
