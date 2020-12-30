package code

// 递归版本
func qSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	// 确定边界
	p := l - 1
	q := r + 1
	mid := arr[l+r>>1]
	for p < q {
		p++
		for {
			p++
			if !(arr[p] < mid) {
				break
			}
			// ...
		}
		for {
			q--
			if !(arr[q] > mid) {
				break
			}
			// ...
		}
	}
	qSort(arr, l, p)
	qSort(arr, p+1, r)
}

// 非递归
// 1.找枢轴,默认第一个 2.对比数字 3.i,j交换 4.i,j移动 重叠退出 返回结果
func quickSort(arr []int, left, right int) []int {
	if left < right {
		i := qs(arr, left, right)
		quickSort(arr, left, i-1)
		quickSort(arr, i+1, right)
	}
	return arr
}

func qs(arr []int, left, right int) int {
	l := left
	r := l + 1
	for i := r; i <= right; i++ {
		if arr[i] < arr[l] {
			swap(arr, i, r)
			r += 1
		}
	}
	swap(arr, l, r-1)
	return r - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
