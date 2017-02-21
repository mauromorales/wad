package filemanager_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWad(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filemanager Suite")
}
