package integration_test

import (
  "testing"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "interview-accountapi/cmd/services/accounts"
  "interview-accountapi/cmd/transport"
)

func TestBooks(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Integration Suite")
}

var _ = AfterSuite(func() {
  tr := transport.DefaultClient()
  service := accounts.New(tr)
  result, _ := service.List()
  for _, account := range result {
    _ = service.Delete(account.Id.String(), account.Version)
  }
})
