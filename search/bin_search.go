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
