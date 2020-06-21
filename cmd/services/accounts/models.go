package accounts

import (
  "github.com/google/uuid"
)

type (
  Account struct {
    Id             uuid.UUID  `json:"id" validate:"required"`
    OrganisationId uuid.UUID  `json:"organisation_id" validate:"required"`
    Attributes     Attributes `json:"attributes" validate:"required"`
    Type           string     `json:"type" validate:"required"`
    CreatedOn      string     `json:"created_on"`
    Version        int        `json:"version"`
  }
  Attributes struct {
    Country       string `json:"country" validate:"required"`
    BIC           string `json:"bic,omitempty"`
    IBAN          string `json:"iban,omitempty"`
    AccountNumber string `json:"account_number,omitempty"`
    BankId        string `json:"bank_id,omitempty"`
    BankIdCode    string `json:"bank_id_code,omitempty"`
  }
)
