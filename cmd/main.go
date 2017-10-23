package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/rfyiamcool/go_pubsub/config"
	"github.com/rfyiamcool/go_pubsub/server"
)

var configFile *string = flag.String("config", "etc/config.ini", "idgo config file")
var logLevel *string = flag.String("log-level", "", "log level [debug|info|warn|error], default error")

const (
	sysLogName = "sys.log"
	MaxLogSize = 1024 * 1024 * 1024
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	if len(*configFile) == 0 {
		fmt.Println("must use a config file")
		return
	}

	cfg, err := config.ParseConfigFile(*configFile)
	if err != nil {
		fmt.Printf("parse config file error:%v\n", err.Error())
		return
	}

	server.GlobalConf = cfg

	//when the log file size greater than 1GB, kingtask will generate a new file
	if len(cfg.LogPath) != 0 {
	}

	if *logLevel != "" {
		setLogLevel(*logLevel)
	} else {
		setLogLevel(cfg.LogLevel)
	}

	var s *server.Server
	s, err = server.NewServer(cfg)
	if err != nil {
		s.Close()
		return
	}

	err = s.Init()
	if err != nil {
		s.Close()
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		sig := <-sc
		fmt.Println(sig)
		s.Close()
	}()
	s.Serve()
}

func setLogLevel(level string) {
}
