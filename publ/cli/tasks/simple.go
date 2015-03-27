package tasks

func TaskSimple(args ...string) []string {
	ret := []string{}
	for _, i := range args {
		ret = append(ret, i)
	}
	return ret
}
