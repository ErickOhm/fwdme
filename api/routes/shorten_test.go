package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShorten(t *testing.T) {
	t.Run("returns a string", func(t *testing.T) {
		request, response := getRequestResponse()

		Shorten(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("expected HTTP status %d, got %d", http.StatusOK, response.Code)
		}

		got := response.Body.String()

		if len(got) == 0 {
			t.Errorf("expected a non-empty string")
		}

	})

	t.Run("Shorten the url", func(t *testing.T) {
		request, response := getRequestResponse()

		testUrl := "https://www.example.com/very/long/url/that/needs/shortening"

		q := request.URL.Query()
		q.Add("url", testUrl)
		request.URL.RawQuery = q.Encode()

		Shorten(response, request)

		got := response.Body.String()

		if len(testUrl) < len(got) {
			t.Errorf("String is not shorter")
		}

	})

	t.Run("returns unique urls", func(t *testing.T) {
		testUrl := "https://example.com/some/url/very/long"

		request, response := getRequestResponse()
		q := request.URL.Query()
		q.Add("url", testUrl)
		request.URL.RawQuery = q.Encode()

		Shorten(response, request)
		url1 := response.Body.String()

		request, response = getRequestResponse()
		q = request.URL.Query()
		q.Add("url", testUrl)
		request.URL.RawQuery = q.Encode()

		Shorten(response, request)
		url2 := response.Body.String()

		if url1 == url2 {
			t.Errorf("Returned URLs are not unique")
		}

	})
}

func getRequestResponse() (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(http.MethodPost, "/shorten", nil)
	response := httptest.NewRecorder()

	return request, response
}
