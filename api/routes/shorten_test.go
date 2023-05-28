package routes

import (
	"fwdme/api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShorten(t *testing.T) {
	store := StubShortenStore{map[string]string{}}
	server := &handlers.ShortenHandler{Store: &store}
	method := http.MethodPost
	path := "/shorten"

	t.Run("returns a string", func(t *testing.T) {
		request, response := getRequestResponse(path, method)

		server.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("expected HTTP status %d, got %d", http.StatusOK, response.Code)
		}

		got := response.Body.String()

		if len(got) == 0 {
			t.Errorf("expected a non-empty string")
		}

	})

	t.Run("Shorten the url", func(t *testing.T) {
		request, response := getRequestResponse(path, method)

		testUrl := "https://www.example.com/very/long/url/that/needs/shortening"

		q := request.URL.Query()
		q.Add("url", testUrl)
		request.URL.RawQuery = q.Encode()

		server.ServeHTTP(response, request)

		got := response.Body.String()

		if len(testUrl) < len(got) {
			t.Errorf("String is not shorter")
		}

	})

	t.Run("returns unique urls", func(t *testing.T) {
		testUrl := "https://example.com/some/url/very/long"
		testUrl2 := "https://example.com/some/different/very/long/url"

		request, response := getRequestResponse(path, method)
		q := request.URL.Query()
		q.Add("url", testUrl)
		request.URL.RawQuery = q.Encode()

		server.ServeHTTP(response, request)
		url1 := response.Body.String()

		request, response = getRequestResponse(path, method)
		q = request.URL.Query()
		q.Add("url", testUrl2)
		request.URL.RawQuery = q.Encode()

		server.ServeHTTP(response, request)
		url2 := response.Body.String()

		if url1 == url2 {
			t.Errorf("Returned URLs are not unique")
		}

	})
}

func getRequestResponse(path, method string) (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(method, path, nil)
	response := httptest.NewRecorder()

	return request, response
}

type StubShortenStore struct {
	urls map[string]string
}

func (s *StubShortenStore) GetShortenedUrl(full string) string {
	shortened := s.urls[full]
	return shortened
}
