package client

import "net/http"

type AccessTokenOptions struct {
	BaseURL    string
	HTTPClient *http.Client
}

type Options struct {
	AccessToken  string
	ClientSecret string
	BaseURL      string
	HTTPClient   *http.Client
}

type OptionFunc func(co *Options)

func WithAuthOpt(token, secret string) OptionFunc {
	return func(co *Options) {
		co.AccessToken = token
		co.ClientSecret = secret
	}
}

func WithURLOpt(url string) OptionFunc {
	return func(co *Options) {
		co.BaseURL = url
	}
}
