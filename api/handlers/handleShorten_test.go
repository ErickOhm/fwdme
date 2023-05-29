package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShorten(t *testing.T) {
	path := "/shorten"
	method := http.MethodPost
	store := StubURLsStore{
		map[string]string{},
	}

	server := ShortenHandler{Store: &store}

	t.Run("it saves on POST", func(t *testing.T) {
		request, response := getRequestResponse(path, method)

		testUrl := "https://example.com/long/url/to/shorten"
		queryParams := request.URL.Query()
		queryParams.Add("url", testUrl)
		request.URL.RawQuery = queryParams.Encode()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		short := response.Body.String()

		if _, exists := store.urls[short]; !exists {
			t.Errorf("full url was not saved")
		}
	})
}

func getRequestResponse(path, method string) (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(method, path, nil)
	response := httptest.NewRecorder()

	return request, response
}

func (s *StubURLsStore) GetShortenedUrl(url string) string {
	s.urls["a"] = url
	return "a"
}
