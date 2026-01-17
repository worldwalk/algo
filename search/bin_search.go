package search

import (
	"fmt"
	"gopkg/sort"
)

func BinarySearchTest() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 6}
	sort.QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr) // [1 2 3 4 5 6 7 8]

	fmt.Println(binarySearch(arr, 4)) // 输出：3
	fmt.Println(binarySearch(arr, 9)) // 输出：-1
	fmt.Println(binarySearch(arr, 1)) // 输出：0
	fmt.Println(binarySearch(arr, 8)) // 输出：7
}

// binarySearch 标准二分查找（查找目标值，如果存在返回索引，否则返回-1）
// 写法1：左闭右闭区间 [l, r]，使用 l <= r
func binarySearch(arr []int, val int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// binarySearch2 标准二分查找的另一种写法
// 写法2：左闭右开区间 [l, r)，使用 l < r
func binarySearch2(arr []int, val int) int {
	l, r := 0, len(arr) // 注意：r 初始化为 len(arr)，不包含
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			r = mid // 右开区间，所以 r = mid（不包含 mid）
		} else {
			l = mid + 1
		}
	}
	return -1
}

// binarySearchFirst 查找第一个等于目标值的位置（处理重复元素）
// 如果有多个重复的 target，返回索引值最小的一个
func binarySearchFirst(arr []int, val int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		if arr[l] == val {
			return l
		}
		mid := l + (r-l)/2
		if arr[mid] < val {
			l = mid + 1
		} else if arr[mid] == val {
			r = mid // 重点：相等时继续向左查找，不返回
		} else {
			r = mid - 1
		}
	}
	return -1
}

// binarySearchLast 查找最后一个等于目标值的位置（处理重复元素）
// 如果有多个重复的 target，返回索引值最大的一个
func binarySearchLast(arr []int, val int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		if arr[r] == val {
			return r
		}
		mid := l + (r-l)/2
		if arr[mid] < val {
			l = mid + 1
		} else if arr[mid] == val {
			l = mid // 重点：相等时继续向右查找，不返回
		} else {
			r = mid - 1
		}
	}
	return -1
}

// binarySearchLowerBound 查找第一个大于等于目标值的位置（下界）
// 如果目标值存在，返回第一个等于目标值的位置
// 如果目标值不存在，返回第一个大于目标值的位置
func binarySearchLowerBound(arr []int, val int) int {
	l, r := 0, len(arr) // 右开区间
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] < val {
			l = mid + 1
		} else {
			r = mid // arr[mid] >= val，继续向左查找
		}
	}
	return l // 返回 l，可能是目标值的位置，也可能是插入位置
}

// binarySearchUpperBound 查找第一个大于目标值的位置（上界）
// 返回第一个大于目标值的位置，如果所有元素都小于等于目标值，返回 len(arr)
func binarySearchUpperBound(arr []int, val int) int {
	l, r := 0, len(arr) // 右开区间
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] <= val {
			l = mid + 1
		} else {
			r = mid // arr[mid] > val，继续向左查找
		}
	}
	return l
}

/** // 有多个重复的target的时候，返回索引值最小的一个
    #寻找开始位置，即找到第一个
    def find_first_index(self,nums,size,target):
        left=0
        right=size-1
        while left<=right:
            if nums[left]==target:
                return left

            mid=left+(right-left)//2
            if nums[mid]<target:
                left=mid+1
            elif nums[mid]==target:
                right=mid        // 重点!   如果是返回索引值最大的，则此处改为： left=mid
            else:
                right=mid-1

        return -1

————————————————
原文链接：https://blog.csdn.net/weixin_45666566/article/details/114213485
*/
