package minima

/**
@info The request headers structure
@property {string} [key] Key for the header
@property {string} [value] Value of the header
*/
type ReqHeader struct {
	key   string
	value string
}

/**
@info The Incoming header structure
@property {[]*ReqHeader} [headers] Array of request headers
*/
type IncomingHeader struct {
	headers []*ReqHeader
}

/**
@info Gets request header from given key
@property {string} [key] Key for the header
return {string}
*/
func (h IncomingHeader) Get(key string) string {
	var value string
	for _, v := range h.headers {
		if v.key == key {
			value = v.value
		}
	}
	return value
}

/**
@info Declares request header from given key
@property {string} [key] Key for the header
@property {string} [value] Value of the header
*/
func (h *IncomingHeader) Set(key string, v string) {
	h.headers = append(h.headers, &ReqHeader{key: key, value: v})
}
