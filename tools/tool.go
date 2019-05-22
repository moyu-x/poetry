package tools

func CheckErr(e error) {
	if nil != e {
		panic(e)
	}
}
