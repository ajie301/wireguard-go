/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2019 WireGuard LLC. All Rights Reserved.
 */

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/ipc"

	"golang.zx2c4.com/wireguard/tun"

	"strings"
	"bufio"

	"golang.zx2c4.com/wireguard/windows/elevate"
	"golang.zx2c4.com/wireguard/windows/tunnel"
	"golang.zx2c4.com/wireguard/windows/conf"
)

const (
	ExitSetupSuccess = 0
	ExitSetupFailed  = 1
)

const testInput = `
[Interface]
PrivateKey = iC7PijUwVv417WdPC3qvIE8vJ4sE88cqyOya+3P5SU0=
Address = 10.253.0.4/32
DNS = 1.1.1.1, 8.8.8.8

[Peer]
PublicKey = 3Qps8Q72H1pemY3moErXVkQDUmPZ1qLoWpFHLbgJcFE=
AllowedIPs = 0.0.0.0/0
Endpoint = 42.159.91.157:51820
PersistentKeepalive = 10`

func TestRuntWireguard() error {

	interfaceName := "wg_test"

	logger := device.NewLogger(
		device.LogLevelDebug,
		fmt.Sprintf("(%s) ", interfaceName),
	)
	logger.Info.Println("Starting wireguard-go version", device.WireGuardGoVersion)
	logger.Debug.Println("Debug log enabled")

	// create watcher
	watcher := tunnel.NewWatcher()
	if watcher == nil {
		logger.Error.Println("Failed to create watcher")
		os.Exit(ExitSetupFailed)
	}
	logger.Info.Println("watcher created")

	wintun, err := tun.CreateTUN(interfaceName, 0)
	if err == nil {
		realInterfaceName, err2 := wintun.Name()
		if err2 == nil {
			interfaceName = realInterfaceName
		}
	} else {
		logger.Error.Println("Failed to create TUN device:", err)
		os.Exit(ExitSetupFailed)
	}

	device := device.NewDevice(wintun, logger)
	device.Up()
	logger.Info.Println("Device started")

	uapi, err := ipc.UAPIListen(interfaceName)
	if err != nil {
		logger.Error.Println("Failed to listen on uapi socket:", err)
		os.Exit(ExitSetupFailed)
	}

	// set config
	conf, err := conf.FromWgQuick(testInput, "test")
	uapiConf, err := conf.ToUAPI()
	if err != nil {
		logger.Error.Println("Failed to read uapi config :", err)
		os.Exit(ExitSetupFailed)
	}

	logger.Info.Println("set config info")
	device.IpcSetOperation(bufio.NewReader(strings.NewReader(uapiConf)))
	device.Up()

	logger.Info.Println("configure watcher, watcher run")
	nativeTun := wintun.(*tun.NativeTun)
	watcher.Run(device, conf, nativeTun)

	errs := make(chan error)
	term := make(chan os.Signal, 1)

	go func() {
		for {
			conn, err := uapi.Accept()
			if err != nil {
				errs <- err
				return
			}
			go device.IpcHandle(conn)
		}
	}()
	logger.Info.Println("UAPI listener started")

	// wait for program to terminate

	signal.Notify(term, os.Interrupt)
	signal.Notify(term, os.Kill)
	signal.Notify(term, syscall.SIGTERM)

	select {
	case <-term:
	case <-errs:
	case <-device.Wait():
	}

	// clean up

	uapi.Close()
	device.Close()

	logger.Info.Println("Shutting down")
	return nil
}

func main(){
	err := elevate.DoAsSystem(TestRuntWireguard)
	fmt.Println(err)
	return
}
