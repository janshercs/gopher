package poker_test

import (
	"bytes"
	"io"
	"poker"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &poker.SpyGame{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()
		poker.AssertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		poker.AssertGameStartedWith(t, game, 3)
		poker.AssertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and finish game with 'Cleo' as winner", func(t *testing.T) {
		game := &poker.SpyGame{}
		stdout := &bytes.Buffer{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()
		poker.AssertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		poker.AssertGameStartedWith(t, game, 8)
		poker.AssertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("prints error when non numeric value is first answered", func(t *testing.T) {
		game := &poker.SpyGame{}
		stdout := &bytes.Buffer{}

		in := userSends("pies", "Cleo wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()
		if game.StartCalled {
			t.Errorf("game should not have started")
		}

		poker.AssertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadInputErrMsg)
	})
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}
