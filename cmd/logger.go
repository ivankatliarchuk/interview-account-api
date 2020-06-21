package cmd

type Logger interface {
  Info(...interface{})
  Debug(...interface{})
  Error(...interface{})
}
