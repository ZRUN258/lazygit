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
