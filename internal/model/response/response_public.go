package response

import "github.com/yqchilde/gint/internal/model"

type SignIn struct {
	User      *model.User `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expires_at"`
}
