// package afero_test tests my arrangement of the afero
// code, not afero itself
package afero_test

import (
	"testing"

	_ "github.com/spf13/afero"
	myAfero "github.com/tbhartman/afero-lite"
	fullAfero "github.com/tbhartman/afero-lite/full"
	"github.com/tbhartman/afero-lite/mem"
)

func TestMyPackage(t *testing.T) {
	var _ myAfero.Fs = mem.NewMemMapFs()
	var _ myAfero.Fs = fullAfero.NewOsFs()
	// doesn't yet work with real afero
	// var _ myAfero.Fs = realAfero.NewOsFs()
}
