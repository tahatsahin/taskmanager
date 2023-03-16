package routers

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestGetTasksNotAuth(t *testing.T) {
	assertResponseBody := func(t *testing.T, s *httptest.Server, expectedBody string) {
		resp, err := s.Client().Get(s.URL + "/tasks")
		if err != nil {
			t.Fatalf("got: %v", err)
		}
		if resp.StatusCode != 401 {
			t.Fatalf("got status: %v", resp.StatusCode)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("error reading body: %v", err)
		}
		if !bytes.Equal(body, []byte(expectedBody)) {
			t.Fatalf("response should be ok, was: %q", string(body))
		}
	}
	router := InitRoutes()
	s := httptest.NewServer(router)
	defer s.Close()
	assertResponseBody(t, s, "{\"error\": \"JWT not valid, no token present in request\"}\n")
}
