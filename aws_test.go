package aws

import (
	"testing"

	"github.com/togo-framework/deploy"
)

func TestRegion(t *testing.T) {
	if region(deploy.Spec{}, "x") != "x" {
		t.Fatal("default region")
	}
	if region(deploy.Spec{Region: "y"}, "x") != "y" {
		t.Fatal("override region")
	}
}
