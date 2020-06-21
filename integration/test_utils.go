package integration

import (
  "math/rand"
  "time"

  "interview-accountapi/cmd/services/accounts"
)

func PrepareAccounts(count int) []*accounts.Account {
  var result []*accounts.Account
  for i := 0; i < count; i++ {
    attr := RandomCountry()
    acc := accounts.NewDefaultAccount(attr)
    result = append(result, acc)
  }
  return result
}

func RandomCountry() string {
  defaults := accounts.DefaultAttributes
  countries := make([]string, 0)
  for country := range defaults {
    countries = append(countries, country)
  }
  rnd := rand.Intn(len(countries))
  return countries[rnd]
}

func TimeToday() string {
  layout := "2006-01-02"
  return time.Now().Format(layout)
}
