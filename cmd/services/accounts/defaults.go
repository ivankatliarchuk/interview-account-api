package accounts

import "github.com/google/uuid"

func NewDefaultAccount(country string) *Account {
  return &Account{
    Id:             uuid.New(),
    OrganisationId: uuid.New(),
    Type:           "accounts",
    Attributes:     DefaultAttributes[country],
  }
}

var DefaultAttributes = map[string]Attributes{
  "GB": {Country: "GB", BIC: "MONZGB2L", IBAN: "GB29NWBK60161331926819", AccountNumber: "41426819", BankId: "040004", BankIdCode: "GBDSC"},
  "AU": {Country: "AU", BIC: "AUAUAU01", IBAN: "", AccountNumber: "8390273849", BankId: "082902", BankIdCode: "AUBSB"},
  "BE": {Country: "BE", BIC: "", IBAN: "", AccountNumber: "19304483", BankId: "BE1", BankIdCode: ""},
  "CA": {Country: "CA", BIC: "CACACA01", IBAN: "", AccountNumber: "", BankId: "0912837485", BankIdCode: "CACPA"},
  "FR": {Country: "FR", BIC: "FRFRFR99", IBAN: "FR1420041010050500013M02606", AccountNumber: "0500013M026", BankId: "3849302918", BankIdCode: "FR"},
  "DE": {Country: "DE", BIC: "", IBAN: "DE23100000001234567890", AccountNumber: "1234567", BankId: "DABA950121", BankIdCode: "DEBLZ"},
  "GR": {Country: "GR", BIC: "HEBACY2N", IBAN: "GR00200128000001200527600", AccountNumber: "12344678901928394", BankId: "HEBACY2N", BankIdCode: "GRBIC"},
}
