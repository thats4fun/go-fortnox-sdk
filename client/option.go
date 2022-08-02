package client

import "net/http"

type AccessTokenOptions struct {
	BaseURL    string
	HTTPClient *http.Client
}

type Options struct {
	ClientID         string
	AccessToken      string
	RefreshToken     string
	ClientSecret     string
	BaseURL          string
	AutoRefreshToken bool
	HTTPClient       *http.Client
}

type OptionFunc func(co *Options)

func WithClientIDOpt(clientID string) OptionFunc {
	return func(co *Options) {
		co.ClientID = clientID
	}
}

func WithOpt(token, secret string) OptionFunc {
	return func(co *Options) {
		co.AccessToken = token
		co.ClientSecret = secret
	}
}

func WithAuthOpt(token, secret string) OptionFunc {
	return func(co *Options) {
		co.AccessToken = token
		co.ClientSecret = secret
	}
}

func WithRefreshOpt(refreshToken string) OptionFunc {
	return func(co *Options) {
		co.RefreshToken = refreshToken
	}
}

func WithURLOpt(url string) OptionFunc {
	return func(co *Options) {
		co.BaseURL = url
	}
}

func WithAutoRefreshTokenOpt(autoRefresh bool) OptionFunc {
	return func(co *Options) {
		co.AutoRefreshToken = autoRefresh
	}
}
