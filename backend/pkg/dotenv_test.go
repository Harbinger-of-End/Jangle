package pkg_test

import (
	"jangle/backend/pkg"
	"os"
	"path"
	"testing"
)

func TestLoadDotenv(t *testing.T) {
	cwd, _ := os.Getwd()
	envpath := path.Join(
		path.Dir(cwd),
		"..",
		"..",
		".env",
	)

	if err := pkg.LoadDotenv(envpath); err != nil {
		t.Fatal(err)
	}
}
