package write

// 合并两个有序数组
// A末尾有足够缓冲空间 容纳B
// 从末尾扫描最大的直接添加,正向扫描太过复杂

func merge(A []int, m int, B []int, n int) {
	i, j := m-1, n-1
	k := len(A) - 1

	for i >= 0 && j >= 0 {
		if A[i] > B[j] {
			A[k] = A[i]
			i--
		} else {
			A[k] = B[j]
			j--
		}
		k--
	}
	// 等于0,  特点情况
	// [0]
	// 0
	// [1]
	// 1
	if j >= 0 {
		copy(A[:j+1], B[:j+1])
	}
}
