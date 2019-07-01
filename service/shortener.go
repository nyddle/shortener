package service

// Shortener implements business logic for the shortener service
type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type ShortenerService struct {
	Shortener
}

// NewShortenerService creates a new shortener service
func NewShortenerService(shortener Shortener) *ShortenerService {
	return &ShortenerService{shortener}
}
