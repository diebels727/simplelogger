package simplelogger
import(
  "log"
  "io"
)

type SimpleLogger struct {
  logger *log.Logger
}

func New(out io.Writer,prefix string,flag int) (*SimpleLogger){
  logger := log.New(out,prefix,flag)
  simple_logger := SimpleLogger{logger}
  return &simple_logger
}

func (s *SimpleLogger) Info(str string) {
  s.logger.Println(str)

  // (*s).Println(str)
}