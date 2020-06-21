package integration

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "interview-accountapi/cmd/services/health"
  "interview-accountapi/cmd/transport"
)

var _ = Describe("Health", func() {
  var service *health.Service

  BeforeEach(func() {
    service = health.New(transport.DefaultClient())
  })

  Context("account API is running", func() {
    When("the API returns 'UP'", func() {
      It("service returns `true`", func() {
        result := service.IsUP()
        Expect(result).To(BeTrue())
      })
    })
  })
})
