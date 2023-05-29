package handlers

import (
	"net/http"
	"testing"
)

func TestRedirect(t *testing.T) {
	store := StubURLsStore{
		map[string]string{
			"x4Adb": "https://google.com",
			"p0zm2": "https://youtube.com",
		},
	}
	server := RedirectHandler{Store: &store}
	method := http.MethodGet

	t.Run("returns a full Url for x4Adb", func(t *testing.T) {
		request, response := getRequestResponse("/x4Adb", method)

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusFound)
		assertResponseBody(t, response.Header().Get("Location"), "https://google.com")
	})

	t.Run("returns a full Url for p0zm2", func(t *testing.T) {
		request, response := getRequestResponse("/p0zm2", method)

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusFound)
		assertResponseBody(t, response.Header().Get("Location"), "https://youtube.com")
	})

	t.Run("returns 404 on non-existing url", func(t *testing.T) {
		request, response := getRequestResponse("/aftsra", method)

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})

}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

type StubURLsStore struct {
	urls map[string]string
}

func (s *StubURLsStore) GetFullUrl(short string) string {
	full := s.urls[short]
	return full
}
