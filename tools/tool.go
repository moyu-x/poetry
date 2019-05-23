package tools

// CheckErr 使用通用的错误检查工具
func CheckErr(e error) {
	if nil != e {
		panic(e)
	}
}
