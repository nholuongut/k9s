package dialog

import (
	"testing"

	"github.com/nholuongut/tview"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDialog(t *testing.T) {
	p := tview.NewPages()

	okFunc := func(c, f bool) {
		assert.True(t, c)
		assert.True(t, f)
	}
	caFunc := func() {
		assert.True(t, true)
	}
	ShowDelete(p, "Yo", okFunc, caFunc)

	d := p.GetPrimitive(deleteKey).(*tview.ModalForm)
	assert.NotNil(t, d)

	dismissDelete(p)
	assert.Nil(t, p.GetPrimitive(deleteKey))
}
