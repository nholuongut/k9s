package ui

import (
	"testing"

	"github.com/nholuongut/k9s/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewCmdUpdate(t *testing.T) {
	defaults, _ := config.NewStyles("")
	v := NewCmdView(defaults, 'T')
	v.update("blee")

	assert.Equal(t, "T> blee\n", v.GetText(false))
}

func TestCmdInCmdMode(t *testing.T) {
	defaults, _ := config.NewStyles("")
	v := NewCmdView(defaults, 'T')
	v.update("blee")
	v.append('!')

	assert.Equal(t, "T> blee!\n", v.GetText(false))
	assert.False(t, v.InCmdMode())
	v.active(true)
	assert.True(t, v.InCmdMode())
}
