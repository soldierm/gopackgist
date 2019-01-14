package monolog

import (
	"io"
	golog "log"
	"os"
)

const (
	Debug     = 100 //Detailed debug information
	Info      = 200 //Interesting events
	Notice    = 250 //Uncommon events
	Warning   = 300 //Exceptional occurrences that are not errors
	Error     = 400 //Runtime errors
	Critical  = 500 //Critical conditions
	Alert     = 550 //Action must be taken immediately
	Emergency = 600 //Urgent alert.
)

var translate = map[int]string{
	Debug:     "【DEBUG】",
	Info:      "【INFO】",
	Notice:    "【DEBUG】",
	Warning:   "【WARNING】",
	Error:     "【ERROR】",
	Critical:  "【CRITICAL】",
	Alert:     "【ALERT】",
	Emergency: "【EMERGENCY】",
}

type Monolog struct {
	log   *golog.Logger
	level int
}

func (this *Monolog) setLevel(level int) {
	this.level = level
}

func (this *Monolog) setPrefix() {
	this.log.SetPrefix(translate[this.level])
}

func New(file string) *Monolog {
	var errFile io.Writer
	var err error
	if file != "" {
		errFile, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			golog.Fatalln("打开日志文件失败：", err)
		}
	} else {
		errFile = os.Stdout
	}
	return &Monolog{
		log:   golog.New(errFile, "", golog.LstdFlags|golog.Lshortfile),
		level: Info,
	}
}

func (this *Monolog) Debug(msg string) {
	this.Log(Debug, msg)
}

func (this *Monolog) Info(msg string) {
	this.Log(Info, msg)
}

func (this *Monolog) Notice(msg string) {
	this.Log(Notice, msg)
}

func (this *Monolog) Warning(msg string) {
	this.Log(Warning, msg)
}

func (this *Monolog) Error(msg string) {
	this.Log(Error, msg)
}

func (this *Monolog) Critical(msg string) {
	this.Log(Critical, msg)
}

func (this *Monolog) Alert(msg string) {
	this.Log(Alert, msg)
}

func (this *Monolog) Emergency(msg string) {
	this.Log(Emergency, msg)
}

func (this *Monolog) Log(level int, msg string) {
	this.setLevel(level)
	this.setPrefix()
	this.log.Println(msg)
}
