package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}

	server, err := NewPlayerServer(&store, &SpyGame{})
	if err != nil {
		t.Errorf("unable to initialize server, %v", err)
	}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server, err := NewPlayerServer(&store, &SpyGame{})
	if err != nil {
		t.Errorf("unable to initialize server, %v", err)
	}

	t.Run("returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response, http.StatusAccepted)

		assertPlayerWin(t, &store, player)

	})
}

func TestLeague(t *testing.T) {
	wantedLeague := League{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	store := StubPlayerStore{nil, nil, wantedLeague}
	server, err := NewPlayerServer(&store, &SpyGame{})
	if err != nil {
		t.Errorf("unable to initialize server, %v", err)
	}

	t.Run("returns 200 on /league", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertStatus(t, response, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, JsonContentType)
	})

}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server, err := NewPlayerServer(&StubPlayerStore{}, &SpyGame{})
		if err != nil {
			t.Errorf("unable to initialize server, %v", err)
		}
		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
	})

	t.Run("start game with 3 players and finish game with 'Ruth' as winner", func(t *testing.T) {
		tenMS := 10 * time.Millisecond
		wantedBlindAlert := "Blind is 100"
		game := &SpyGame{BlindAlert: []byte(wantedBlindAlert)}
		winner := "Ruth"
		playerServer := MustMakePlayerServer(t, &StubPlayerStore{}, game)

		server := httptest.NewServer(playerServer)
		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws := MustDialWS(t, wsURL)

		defer ws.Close()
		defer server.Close()

		WriteWSMessage(t, ws, "3")
		WriteWSMessage(t, ws, winner)

		// time.Sleep(tenMS)
		AssertGameStartedWith(t, game, 3)
		AssertFinishCalledWith(t, game, winner)

		within(t, tenMS, func() { assertWebSocketGotMsg(t, ws, wantedBlindAlert) })
	})
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/game", nil)
	return request
}

func getLeagueFromResponse(t testing.TB, body io.Reader) League {
	t.Helper()

	league, err := NewLeague(body)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return league
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of application/json got %v", response.Result().Header)
	}
}

func assertStatus(t testing.TB, got *httptest.ResponseRecorder, want int) {
	if got.Code != want {
		t.Errorf("status code is wrong: got status %d want %d", got.Code, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong; got %q, want %q", got, want)

	}
}

func assertLeague(t testing.TB, got, want League) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}

func MustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}
	return ws
}

func MustMakePlayerServer(t *testing.T, store PlayerStore, game Game) *PlayerServer {
	playerServer, err := NewPlayerServer(store, game)

	if err != nil {
		t.Errorf("unable to initialize server, %v", err)
	}

	return playerServer
}

func WriteWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection , %v", err)
	}
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)
	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}

func assertWebSocketGotMsg(t *testing.T, ws *websocket.Conn, want string) {
	_, wsMsg, _ := ws.ReadMessage()
	got := string(wsMsg)
	if got != want {
		t.Errorf("got blind alert %q, want %q", got, want)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
