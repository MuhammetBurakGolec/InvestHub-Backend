// api/models/user.go

package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // Gerçek uygulamada şifre hash olarak saklanmalı
}
