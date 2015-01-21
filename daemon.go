package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"syscall"
	"time"

	"github.com/bpostlethwaite/go-daemon"
)

var (
	signal = flag.String("s", "", `send signal to the daemon
		quit — graceful shutdown
		stop — fast shutdown
		reload — reloading the configuration file`)
)

const confFile = "danger.json"

func main() {
	flag.Parse()
	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "reload"), syscall.SIGHUP, reloadHandler)

	dconf, err := readConfigs(confFile, []string{"./", "/etc", "/var/lib/danger/"})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	dconf = applyConfigDefaults(dconf)

	if err := os.MkdirAll(dconf.WorkDir, 0755); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cntxt := &daemon.Context{
		PidFileName: path.Base(dconf.PidFile),
		PidFilePerm: 0644,
		LogFileName: path.Base(dconf.LogFile),
		LogFilePerm: 0640,
		WorkDir:     path.Base(dconf.WorkDir),
		Umask:       027,
	}

	if len(daemon.ActiveFlags()) > 0 {
		d, err := cntxt.Search()
		if err != nil {
			log.Fatalln("Unable send signal to the daemon:", err)
		}
		daemon.SendCommands(d)
		return
	}

	// IF the daemon is already running and user specified command args
	// Spin up a tcp server and connect to Daemon and serve commands
	args := flag.Args()
	if len(args) > 0 {
		tcp := NewTCPClient(dconf)
		tcp.Run(args)
		return
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatalln(err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Println("- - - - - - - - - - - - - - -")
	log.Println("daemon started")

	go worker(dconf)

	err = daemon.ServeSignals()
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println("daemon terminated")
}

var (
	stop = make(chan struct{})
	done = make(chan struct{})
)

func worker(conf DangerConfig) {

	dng := NewDanger(conf)

	go dng.Run()

	for {
		time.Sleep(time.Second)
		if _, ok := <-stop; ok {
			break
		}
	}
	done <- struct{}{}
}

func termHandler(sig os.Signal) error {
	log.Println("terminating...")
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
	}
	return daemon.ErrStop
}

func reloadHandler(sig os.Signal) error {
	log.Println("configuration reloaded")
	return nil
}
