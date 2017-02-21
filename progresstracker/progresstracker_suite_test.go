package progresstracker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProgresstracker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Progresstracker Suite")
}
