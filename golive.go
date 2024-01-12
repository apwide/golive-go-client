package main

import (
	"context"
	"fmt"
	"net/http"
)

type securityProviderBearerToken struct {
	token string
}

func newSecurityProviderBearerToken(token string) (*securityProviderBearerToken, error) {
	return &securityProviderBearerToken{
		token: token,
	}, nil
}

func (s securityProviderBearerToken) Intercept(ctx context.Context, req *http.Request) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	return nil
}

func Golive(server string, token string) (*ClientWithResponses, error) {
	apiKeyProvider, _ := newSecurityProviderBearerToken(token)
	return NewClientWithResponses(server, WithRequestEditorFn(apiKeyProvider.Intercept))
}
