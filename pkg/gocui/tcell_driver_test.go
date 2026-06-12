package gocui

import (
	"testing"

	"github.com/gdamore/tcell/v3"
	"github.com/stretchr/testify/assert"
)

func TestNewGuiUsesLegacyKeyboardProtocol(t *testing.T) {
	g := newTestGui(t)

	assert.Equal(t, tcell.LegacyKeyboard, g.screen.KeyboardProtocol())
}

func TestTcellKeyReleaseEventsAreIgnored(t *testing.T) {
	ev := gocuiEventFromTcellKey(tcell.NewEventKeyEx(tcell.KeyRune, "a", tcell.ModNone, false, tcell.Key('a'), 1))

	assert.Equal(t, eventNone, ev.Type)
}

func TestTcellKeyPressEventsAreHandled(t *testing.T) {
	ev := gocuiEventFromTcellKey(tcell.NewEventKeyEx(tcell.KeyRune, "a", tcell.ModNone, true, tcell.Key('a'), 1))

	assert.Equal(t, eventKey, ev.Type)
	assert.True(t, ev.Key.Equals(NewKeyRune('a')))
}
