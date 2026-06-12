package gocui

import (
	"testing"
	"time"

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

func TestDuplicateTcellKeyEventsAreIgnored(t *testing.T) {
	g := &Gui{}
	when := time.Now()
	first := tcell.NewEventKeyEx(tcell.KeyRune, "a", tcell.ModNone, true, tcell.Key('a'), 1)
	first.SetEventTime(when)
	second := tcell.NewEventKeyEx(tcell.KeyRune, "a", tcell.ModNone, true, tcell.Key('a'), 1)
	second.SetEventTime(when)

	assert.False(t, g.isDuplicateTcellKey(first))
	assert.True(t, g.isDuplicateTcellKey(second))
}

func TestSameTcellKeyWithDifferentTimestampIsHandled(t *testing.T) {
	g := &Gui{}
	first := tcell.NewEventKeyEx(tcell.KeyRune, "a", tcell.ModNone, true, tcell.Key('a'), 1)
	first.SetEventTime(time.Now())
	second := tcell.NewEventKeyEx(tcell.KeyRune, "a", tcell.ModNone, true, tcell.Key('a'), 1)
	second.SetEventTime(first.When().Add(time.Nanosecond))

	assert.False(t, g.isDuplicateTcellKey(first))
	assert.False(t, g.isDuplicateTcellKey(second))
}
