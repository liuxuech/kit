package request

import "net/url"

func WithBaseUrl(baseUrl string) Option {
	return func(r *Request) {
		urlParse, _ := url.Parse(baseUrl)
		r.baseUrl = urlParse
	}
}
