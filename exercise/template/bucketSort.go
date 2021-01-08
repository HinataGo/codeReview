package code

// 桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。为了使桶排序更加高效，我们需要做到这两点：
// 在额外空间充足的情况下，尽量增大桶的数量
// 使用的映射函数能够将输入的 N 个数据均匀的分配到 K 个桶中
// 同时，对于桶中元素的排序，选择何种比较排序算法对于性能的影响至关重要
//  什么时候最快 --- 当输入的数据可以均匀的分配到每一个桶中。
//  什么时候最慢 --- 当输入的数据被分配到了同一个桶中

func bucketSort(list []int) []int {
	max := max(list)
	min := min(list)
	base := 0
	if min < 0 {
		base = -min
	} else {
		base = min
	}
	max = (max + base) / 10
	min = (min + base) / 10
	bucket := make([][]int, max-min+1)
	var result []int
	for _, value := range list {
		i := (int)((value + base) / 10)
		bucket[i] = append(bucket[i], value)
	}
	for _, value := range bucket {
		if len(value) > 0 {
			quickSort(value, 0, len(value)-1)
		}
	}
	for _, value := range bucket {
		if len(value) > 0 {
			for _, v := range value {
				result = append(result, v)
			}
		}
	}
	return result
}
func max(list []int) int {
	max := list[0]
	for _, value := range list {
		if value > max {
			max = value
		}
	}
	return max
}
func min(list []int) int {
	min := list[0]
	for _, value := range list {
		if value < min {
			min = value
		}
	}
	return min
}
