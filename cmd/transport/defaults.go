package transport

import (
  "fmt"
  net "net/http"
  "os"

  log "github.com/sirupsen/logrus"

  "interview-accountapi/cmd"
)

// A Default of values for SDK client
type Defaults struct {
  Config     cmd.Config
  Middleware Middleware
}

// Preconfigure default values
func Preconfigure() Defaults {
  return Defaults{
    Config:     *config(),
    Middleware: middleware(),
  }
}

func config() *cmd.Config {
  var logger = log.New()
  logger.Formatter = &log.JSONFormatter{
    PrettyPrint: false,
  }
  logger.Out = os.Stdout
  logger.Level = log.InfoLevel
  logger.ReportCaller = false

  return cmd.NewConfig().
    WithEndpoint(fmt.Sprintf("%s://%s:%s/%s", cmd.Scheme, cmd.Domain, cmd.Port, cmd.Version)).
    WithHTTPClient(net.DefaultClient).
    WithLogger(logger)
}

func middleware() Middleware {
  var mdw Middleware

  mdw.Send.Append(DoRequestMiddleware)
  mdw.Body.Append(CopyResponseBodyMiddleware)
  mdw.Build.Append(BaseHeadersMiddleware)
  mdw.Marshal.Append(MarshalMiddleware)
  mdw.Unmarshal.Append(UnmarshalMiddleware)
  mdw.Unmarshal.Append(UnmarshalDiscardBodyMiddleware)
  mdw.Paginator.Append(PaginatorMiddleware)

  return mdw
}
