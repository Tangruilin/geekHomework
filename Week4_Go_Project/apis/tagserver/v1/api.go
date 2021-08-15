package v1

import (
	"context"
)

const (
	APP_KEY    = "eddycjy"
	APP_SECRET = "...."
)

type API struct {
	URL string
}

func (a *API) getAccessToken(ctx context.Context) (string, error) {
	return "", nil
}
