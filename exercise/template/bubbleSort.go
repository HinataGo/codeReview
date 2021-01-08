package code

// 旧版
// func bubbleSort(arr []int) []int {
// 	length := len(arr)
// 	for i := 0; i < length; i++ {
// 		for j := 0; j < length-1-i; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// 	return arr
// }

// 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 针对所有的元素重复以上的步骤，除了最后一个。
// 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
// 最快( 正序,此时要他何用? )  最慢 (反序)
// 改进版
func bubbleSort(arr []int) []int {

	for i := len(arr) - 1; i > 0; i-- { // 反向遍历
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
