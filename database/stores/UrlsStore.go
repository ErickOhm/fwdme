package stores

func NewInMemoryUrlStore() *InMemoryUrlStore {
	return &InMemoryUrlStore{
		map[string]string{},
	}
}

type InMemoryUrlStore struct {
	store map[string]string
}

func (i *InMemoryUrlStore) GetShortenedUrl(long string) string {
	return ""
}

func (i *InMemoryUrlStore) GetFullUrl(short string) string {
	return ""
}
