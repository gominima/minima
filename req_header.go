package minima

/**
* Minima is a free and open source software under Mit license

Copyright (c) 2021 gominima

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

* Authors @apoorvcodes @megatank58
* Maintainers @Panquesito7 @savioxavier @Shubhaankar-Sharma @apoorvcodes @megatank58
* Thank you for showing interest in minima and for this beautiful community
*/

/**
 * @info The request headers structure
 * @property {string} [key] Key for the header
 * @property {string} [value] Value of the header
*/
type ReqHeader struct {
	key   string
	value string
}

/**
 * @info The Incoming header structure
 * @property {[]*ReqHeader} [headers] Array of request headers
*/
type IncomingHeader struct {
	headers []*ReqHeader
}

/**
 * @info Gets request header from given key
 * @property {string} [key] Key for the header
return {string}
*/
func (h IncomingHeader) Get(key string) string {
	var value string

	for _, v := range h.headers {
		if v.key == key {
			value = v.value
			break
		}
	}

	return value
}

/**
 * @info Declares request header from given key
 * @property {string} [key] Key for the header
 * @property {string} [value] Value of the header
*/
func (h *IncomingHeader) Set(key string, v string) {
	h.headers = append(h.headers, &ReqHeader{key: key, value: v})
}
