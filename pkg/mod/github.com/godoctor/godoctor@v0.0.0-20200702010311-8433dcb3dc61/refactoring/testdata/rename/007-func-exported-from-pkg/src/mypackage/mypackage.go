package mypackage
//Test for renaming an exported function name 
func MyFunction(n int) int {             // <<<<< rename,3,6,3,6,Xyz,pass
	if n == 0 {
		return 1
	} else {
		return n + MyFunction(n-1)
	}
}
