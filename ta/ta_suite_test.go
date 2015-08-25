package ta_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTa(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ta Suite")
}
