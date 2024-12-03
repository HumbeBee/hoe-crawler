package interfaces

type CloudflareBypasser interface {
	GetCookies(url string) ([]byte, error)
}
