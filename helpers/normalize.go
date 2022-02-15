package helpers

import "strings"

// Normalize passed fields. This is used in models.User.Create.
func NormalizeUserCreate(n string, e string, p string) (string, string, string) {
	n = strings.TrimSpace(n)
	e = strings.ToLower(e)
	e = strings.TrimSpace(e)
	p = strings.TrimSpace(p)
	return n, e, p
}

// Normalize passed fields. This is used in models.User.Authenticate.
func NormalizeUserAuth(e string, p string) (string, string) {
	e = strings.ToLower(e)
	e = strings.TrimSpace(e)
	p = strings.TrimSpace(p)
	return e, p
}
