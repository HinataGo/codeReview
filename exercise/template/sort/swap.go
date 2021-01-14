package sort

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func SwapGoAdvanced(list []int, i, j int) {
	list[i] = list[i] + list[j]
	list[j] = list[i] - list[j]
	list[i] = list[i] - list[j]
}
