package integration

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "interview-accountapi/cmd/services/accounts"
  "interview-accountapi/cmd/transport"
)

const (
  firstAccount = 0
)

var _ = Describe("Account", func() {
  var (
    service *accounts.Service
    tr *transport.Client
  )

  BeforeEach(func() {
    tr = transport.DefaultClient()
    service = accounts.New(tr)
  })

  Describe("Account Create", func() {
    Context("Single Account Create Will Begun", func() {
      When("register an existing bank account with organisation.", func() {
        var account *accounts.Account
        numberOfAccounts := 1

        It("should create a bank account", func() {
          var err error
          accInput := PrepareAccounts(numberOfAccounts)[firstAccount]
          account, err = service.Create(accInput)
          Expect(err).To(BeNil())
        })

        It("should have version '0'", func() {
          Expect(account.Version).Should(BeZero())
        })

        It("should be created today", func() {
          today := TimeToday()
          Expect(account.CreatedOn).Should(ContainSubstring(today))
        })

        It("should delete an account", func() {
          err := service.Delete(account.Id.String(), account.Version)
          Expect(err).To(BeNil())
        })
      })
    })
  })

  Describe("Account Fetch", func() {
    Context("Single Account Fetch Will Begun", func() {
      var actualId string
      numberOfAccounts := 1

      BeforeEach(func() {
        accInput := PrepareAccounts(numberOfAccounts)[firstAccount]
        actualAcc, _ := service.Create(accInput)
        actualId = actualAcc.Id.String()
      })

      AfterEach(func() {
        _ = service.Delete(actualId, 0)
      })

      It("should 'Get' a single account using the account ID.", func() {
        acc, err := service.Fetch(actualId)
        Expect(err).To(BeNil())
        Expect(acc.Id.String()).Should(Equal(actualId))
      })
    })
  })

  Describe("Account Paginator", func() {
    Context("Single Account Fetch Will Begun", func() {
      var accounts []*accounts.Account
      numberOfAccounts := 83

      BeforeEach(func() {
        acc := PrepareAccounts(numberOfAccounts)
        for _, el := range acc {
          actualAcc, _ := service.Create(el)
          accounts = append(acc, actualAcc)
        }
      })

      AfterEach(func() {
        for _, acc := range accounts {
          _ = service.Delete(acc.Id.String(), acc.Version)
        }
      })

      It("Should be able to fetch page: '4' and size: '5'", func() {
        actualSize := 5
        result, err := service.Paginate("4", actualSize)
        Expect(err).To(BeNil())
        Expect(len(result)).To(Equal(actualSize))
      })

      It("Should be able to fetch page: 'first' and size: '7'", func() {
        actualPage := "first"
        actualSize := 7
        result, err := service.Paginate(actualPage, actualSize)
        Expect(err).To(BeNil())
        Expect(len(result)).To(Equal(actualSize))
      })

      It("Should be able to fetch page: 'last' and size: '7'", func() {
        var lastPageElements int
        actualPage := "last"
        actualSize := 7
        if numberOfAccounts % actualSize == 0 {
          lastPageElements = actualSize
        } else {
          lastPageElements = numberOfAccounts % actualSize
        }
        result, err := service.Paginate(actualPage, actualSize)
        Expect(err).To(BeNil())
        Expect(len(result)).To(Equal(lastPageElements))
      })
    })
  })
})
