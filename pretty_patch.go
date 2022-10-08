package pretty

func PrettyString(jsonString string) string {
	ret := Pretty([]byte(jsonString))
	return string(ret)
}

func PrettyStringOptions(jsonString string, opts *Options) string {
	ret := PrettyOptions([]byte(jsonString), opts)
	return string(ret)
}

func PrettyBytesToString(json []byte) string {
	ret := Pretty(json)
	return string(ret)
}

func PrettyBytesToStringWithOptions(json []byte, opts *Options) string {
	ret := PrettyOptions(json, opts)
	return string(ret)
}
