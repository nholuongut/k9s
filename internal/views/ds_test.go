package views

import (
	"testing"

	"github.com/nholuongut/k9s/internal/config"
	"github.com/nholuongut/k9s/internal/resource"
	"github.com/stretchr/testify/assert"
)

func TestDaemonSetView(t *testing.T) {
	l := resource.NewDaemonSetList(nil, "fred")
	v := newDaemonSetView("blee", NewApp(config.NewConfig(ks{})), l).(*daemonSetView)

	assert.Equal(t, 3, len(v.hints()))
}
