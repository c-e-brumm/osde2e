package osde2e_scale

import (
	"testing"

	"github.com/openshift/osde2e/common"

	// import suites to be tested
	_ "github.com/openshift/osde2e/test/scale"
)

func TestScale(t *testing.T) {
	common.RunE2ETests(t)
}
