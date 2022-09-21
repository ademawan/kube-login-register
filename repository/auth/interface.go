package auth

import "kube-login-register/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
	LoginGoogle(email string) (entities.User, error)
}
