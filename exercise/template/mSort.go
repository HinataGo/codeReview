package code

const N = 1000000 // 数组上限
func mSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + r>>1
	mSort(arr, l, mid)
	mSort(arr, mid+1, mid)
	k := 0
	i := l
	j := mid + 1
	tmp := make([]int, N)
	for i <= mid && j <= r {
		if arr[i] < arr[j] {
			tmp[k+1] = arr[i+1]
		} else {
			tmp[k+1] = arr[j+1]
		}
	}
}
