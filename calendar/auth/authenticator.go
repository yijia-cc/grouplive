package auth

type Authenticator interface {
	VerifyIdentity(authToken string) (string, error)
}

