package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

var pluginCmd *exec.Cmd

func startPlugin(plugin, pluginOpts, ssAddr string, isServer bool) (newAddr string, err error) {
	logf("starting plugin (%s) with option (%s)....", plugin, pluginOpts)
	freePort, err := getFreePort()
	if err != nil {
		return "", fmt.Errorf("failed to fetch an unused port for plugin (%v)", err)
	}
	localHost := "127.0.0.1"
	ssHost, ssPort, err := net.SplitHostPort(ssAddr)
	if err != nil {
		return "", err
	}
	newAddr = localHost + ":" + freePort
	if isServer {
		if ssHost == "" {
			ssHost = "0.0.0.0"
		}
		logf("plugin (%s) will listen on %s:%s", plugin, ssHost, ssPort)
	} else {
		logf("plugin (%s) will listen on %s:%s", plugin, localHost, freePort)
	}
	err = execPlugin(plugin, pluginOpts, ssHost, ssPort, localHost, freePort)
	return
}

func killPlugin() {
	if pluginCmd != nil {
		pluginCmd.Process.Signal(syscall.SIGTERM)
		waitCh := make(chan struct{})
		go func() {
			pluginCmd.Wait()
			close(waitCh)
		}()
		timeout := time.After(3 * time.Second)
		select {
		case <-waitCh:
		case <-timeout:
			pluginCmd.Process.Kill()
		}
	}
}

func execPlugin(plugin, pluginOpts, remoteHost, remotePort, localHost, localPort string) (err error) {
	pluginFile := plugin
	if fileExists(plugin) {
		if !filepath.IsAbs(plugin) {
			pluginFile = "./" + plugin
		}
	} else {
		pluginFile, err = exec.LookPath(plugin)
		if err != nil {
			return err
		}
	}
	logH := newLogHelper("[" + plugin + "]: ")
	env := append(os.Environ(),
		"SS_REMOTE_HOST="+remoteHost,
		"SS_REMOTE_PORT="+remotePort,
		"SS_LOCAL_HOST="+localHost,
		"SS_LOCAL_PORT="+localPort,
		"SS_PLUGIN_OPTIONS="+pluginOpts,
	)
	cmd := &exec.Cmd{
		Path:   pluginFile,
		Env:    env,
		Stdout: logH,
		Stderr: logH,
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	pluginCmd = cmd
	go func() {
		if err := cmd.Wait(); err != nil {
			logf("plugin exited (%v)\n", err)
			os.Exit(2)
		}
		logf("plugin exited\n")
		os.Exit(0)
	}()
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getFreePort() (string, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return "", err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return "", err
	}
	port := fmt.Sprintf("%d", l.Addr().(*net.TCPAddr).Port)
	l.Close()
	return port, nil
}
