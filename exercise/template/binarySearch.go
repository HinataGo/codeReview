package code

type binarySearch struct {
	Data []int
}

// 前提数组是要有序的
// 查找v在有序数组array中的位置
func (bs *binarySearch) IndexOf(v int) (index int) {
	if bs.Data == nil || len(bs.Data) == 0 {
		return -1
	}
	begin := 0
	end := len(bs.Data)
	for begin < end {
		mid := (begin + end) >> 1
		if bs.Data[mid] > v {
			end = mid
		} else if bs.Data[mid] < v {
			begin = mid + 1
		} else {
			index = mid
			return
		}

	}
	return -1
}

// 查找v在有序数组array中待插入位置
func (bs *binarySearch) Find(v int) (index int) {
	if bs.Data == nil || len(bs.Data) == 0 {
		return -1
	}
	begin := 0
	end := len(bs.Data)
	for begin < end {
		mid := (begin + end) >> 1
		if bs.Data[mid] > v {
			end = mid
		} else if bs.Data[mid] < v {
			begin = mid + 1
		}

	}
	return begin
}
