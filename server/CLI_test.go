package poker

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

var (
	dummyBlindAlerter = &spyBlindAlerter{}
	dummyPlayerStore  = &StubPlayerStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
)

func TestGameStart(t *testing.T) {
	t.Run("schedules printing of blind alerts for 5 players", func(t *testing.T) {
		blindAlerter := &spyBlindAlerter{}

		game := NewTexasHoldem(blindAlerter, dummyPlayerStore)

		cases := []ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		game.Start(5, ioutil.Discard)

		if len(blindAlerter.alerts) == 0 {
			t.Fatal("no blind alerts were scheduled.")
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("schedules printing of blind alerts for 7 players", func(t *testing.T) {
		blindAlerter := &spyBlindAlerter{}
		game := NewTexasHoldem(blindAlerter, dummyPlayerStore)
		game.Start(7, ioutil.Discard)

		cases := []ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		if len(blindAlerter.alerts) == 0 {
			t.Fatal("no blind alerts were scheduled.")
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})
}

func TestGameStarts(t *testing.T) {
	t.Run("prints out and reads first player prompt", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &SpyGame{}

		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}

	})

}

func TestGameFinish(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		playerStore := &StubPlayerStore{}
		game := NewTexasHoldem(dummyBlindAlerter, playerStore)

		game.Finish("Chris")

		AssertPlayerWin(t, playerStore, "Chris")
	})
}

func TestGameDoesNotStartWithInvalidInput(t *testing.T) {
	t.Run("it prints error when non numeric value entered and does not start game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &SpyGame{}

		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}

		AssertMessagesSentToUser(t, stdout, PlayerPrompt, BadInputErrMsg)

	})
}

type spyBlindAlerter struct {
	alerts []ScheduledAlert
}

func (s *spyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, out io.Writer) {
	s.alerts = append(s.alerts, ScheduledAlert{at, amount})
}

type SpyGame struct {
	StartedWith    int
	StartCalled    bool
	BlindAlert     []byte
	FinishedWith   string
	FinishedCalled bool
}

func (g *SpyGame) Start(numberOfPlayers int, out io.Writer) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
	out.Write(g.BlindAlert)
}

func (g *SpyGame) Finish(winner string) {
	g.FinishedWith = winner
}

func assertScheduledAlert(t testing.TB, got, want ScheduledAlert) {
	t.Helper()
	amountGot := got.Amount
	if amountGot != want.Amount {
		t.Errorf("got amount %d want %d", amountGot, want.Amount)
	}

	gotScheduledTime := got.At
	if gotScheduledTime != want.At {
		t.Errorf("got time %d want %d", gotScheduledTime, want.At)
	}
}

func checkSchedulingCases(t *testing.T, cases []ScheduledAlert, blindAlerter *spyBlindAlerter) {
	t.Helper()
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(blindAlerter.alerts) <= i {
				t.Fatalf("blind alerts %d was not scheduled %v", i, blindAlerter.alerts)
			}

			alert := blindAlerter.alerts[i]

			assertScheduledAlert(t, alert, want)
		})
	}
}

func AssertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func AssertGameStartedWith(t testing.TB, game *SpyGame, n int) {
	t.Helper()

	passed := retryUntil(50*time.Millisecond, func() bool {
		return game.StartedWith == n
	})

	if !passed {
		t.Errorf("start called with %d but got %d", n, game.StartedWith)
	}
}

func AssertFinishCalledWith(t testing.TB, game *SpyGame, name string) {
	t.Helper()
	passed := retryUntil(50*time.Millisecond, func() bool {
		return game.FinishedWith == name
	})

	if !passed {
		t.Errorf("finish called with %q but got %q", name, game.FinishedWith)
	}
}
