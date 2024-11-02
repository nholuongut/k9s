package dialog

import (
	"testing"

	"github.com/nholuongut/tview"
	"github.com/stretchr/testify/assert"
)

func TestPortForwardDialog(t *testing.T) {
	p := tview.NewPages()

	okFunc := func(lport, cport string) {
	}
	ShowPortForward(p, "8080", okFunc)

	d := p.GetPrimitive(portForwardKey).(*tview.ModalForm)
	assert.NotNil(t, d)

	DismissPortForward(p)
	assert.Nil(t, p.GetPrimitive(portForwardKey))
}

func TestStripPort(t *testing.T) {
	uu := map[string]struct {
		port, e string
	}{
		"full": {
			"fred:8000", "8000",
		},
		"port": {
			"8000", "8000",
		},
		"protocol": {
			"dns:53╱UDP", "53",
		},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, stripPort(u.port))
		})
	}
}
