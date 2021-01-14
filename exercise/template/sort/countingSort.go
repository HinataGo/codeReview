package sort

// 计数排序
// 作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。
// 当输入的元素是 n 个 0 到 k 之间的整数时，它的运行时间是 Θ(n + k)。计数排序不是比较排序，排序的速度快于任何比较排序算法。
// （1）找出待排序的数组中最大和最小的元素
// （2）统计数组中每个值为i的元素出现的次数，存入数组C的第i项
// （3）对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）
// （4）反向填充目标数组：将每个元素i放在新数组的第C(i)项，每放一个元素就将C(i)减去1

func countingSort(arr []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen) // 初始为0的数组

	sortedIndex := 0
	length := len(arr)

	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return arr
}
