package views

import (
	"testing"

	"github.com/nholuongut/k9s/internal/config"
	"github.com/nholuongut/k9s/internal/resource"
	"github.com/stretchr/testify/assert"
)

func TestDeployView(t *testing.T) {
	l := resource.NewDeploymentList(nil, "fred")
	v := newDeployView("blee", NewApp(config.NewConfig(ks{})), l).(*deployView)

	assert.Equal(t, 3, len(v.hints()))
}
