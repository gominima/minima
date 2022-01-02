package minima

type ReqHeader struct {
	key   string
	value string
}
type IncomingHeader struct {
	headers []*ReqHeader
}

func (h IncomingHeader) Get(key string) string {
	var value string
	for _, v := range h.headers {
		if v.key == key {
			value = v.value
		}
	}
	return value
}

func (h *IncomingHeader) Set(key string, v string) {
	h.headers = append(h.headers, &ReqHeader{key: key, value: v})
}
