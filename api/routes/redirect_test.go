package routes

import (
	"fwdme/api/handlers"
	"net/http"
	"testing"
)

func TestRedirect(t *testing.T) {
	store := StubRedirectStore{
		map[string]string{
			"x4Adb": "https://google.com",
			"p0zm2": "https://youtube.com",
		},
	}
	server := &handlers.RedirectHandler{Store: &store}
	method := http.MethodGet

	t.Run("returns a redirect response", func(t *testing.T) {
		request, response := getRequestResponse("/", method)

		server.ServeHTTP(response, request)

		if response.Code != http.StatusFound {
			t.Errorf("Expected a redirect")
		}
	})

	t.Run("returns a full Url for x4Adb", func(t *testing.T) {
		request, response := getRequestResponse("/x4Adb", method)

		server.ServeHTTP(response, request)

		got := response.Header().Get("Location")
		want := "https://google.com"

		if got != want {
			t.Errorf("Expected redirect URL %s, got %s", want, got)
		}
	})

	t.Run("returns a full Url for p0zm2", func(t *testing.T) {
		request, response := getRequestResponse("/p0zm2", method)

		server.ServeHTTP(response, request)

		got := response.Header().Get("Location")
		want := "https://youtube.com"

		if got != want {
			t.Errorf("Expected redirect URL %s, got %s", want, got)
		}
	})
}

type StubRedirectStore struct {
	urls map[string]string
}

func (s *StubRedirectStore) GetFullUrl(short string) string {
	full := s.urls[short]
	return full
}
