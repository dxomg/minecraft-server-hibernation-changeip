package main

import (
	"fmt"
	"net"

	"msh/lib/config"
	"msh/lib/conn"
	"msh/lib/errco"
	"msh/lib/input"
	"msh/lib/progmgr"
	"msh/lib/servctrl"
	"msh/lib/utility"
)

// contains intro to script and program
var intro []string = []string{
	" _ __ ___  ___| |__  ",
	"| '_ ` _ \\/ __| '_ \\ ",
	"| | | | | \\__ \\ | | | " + progmgr.MshVersion,
	"|_| |_| |_|___/_| |_| " + progmgr.MshCommit,
	"Copyright (C) 2019-2022 gekigek99",
	"github: https://github.com/gekigek99",
	"remember to give a star to this repository!",
}

func main() {
	// print program intro
	// not using errco.Logln since log time is not needed
	fmt.Println(utility.Boxify(intro))

	// load configuration from msh config file
	errMsh := config.LoadConfig()
	if errMsh != nil {
		errco.LogMshErr(errMsh.AddTrace("main"))
		progmgr.AutoTerminate()
	}

	// launch msh manager
	go progmgr.MshMgr()
	// wait for the initial update check
	<-progmgr.ReqSent

	// launch GetInput()
	go input.GetInput()

	// if ms suspension is allowed, pre-warm the server
	if config.ConfigRuntime.Msh.AllowSuspend {
		errco.Logln(errco.LVL_1, "minecraft server will now pre-warm (AllowSuspend is enabled)...")
		errMsh = servctrl.WarmMS()
		if errMsh != nil {
			errco.LogMshErr(errMsh.AddTrace("main"))
		}
	}

	// open a listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.ListenHost, config.ListenPort))
	if err != nil {
		errco.LogMshErr(errco.NewErr(errco.ERROR_CLIENT_LISTEN, errco.LVL_3, "main", err.Error()))
		progmgr.AutoTerminate()
	}

	errco.Logln(errco.LVL_1, "listening for new clients to connect on %s:%d ...", config.ListenHost, config.ListenPort)

	// infinite cycle to accept clients. when a clients connects it is passed to handleClientSocket()
	for {
		clientSocket, err := listener.Accept()
		if err != nil {
			errco.LogMshErr(errco.NewErr(errco.ERROR_CLIENT_ACCEPT, errco.LVL_3, "main", err.Error()))
			continue
		}

		go conn.HandleClientSocket(clientSocket)
	}
}
