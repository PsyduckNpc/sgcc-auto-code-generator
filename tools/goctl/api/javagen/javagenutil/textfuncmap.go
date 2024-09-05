package javagenutil

func CheckSliceLastElem(slice []any, elem any) bool {
	if slice == nil || elem == nil {
		return false
	}
	if slice[len(slice)-1] == elem {
		return true
	}
	return false
}

func CheckSliceFirstElem(slice []any, elem any) bool {
	if slice == nil || elem == nil {
		return false
	}
	if slice[0] == elem {
		return true
	}
	return false
}
