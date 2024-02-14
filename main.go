package main

import (
	"Go-Tray/icon"
	"github.com/getlantern/systray"
	"net"
	"time"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.GetIconData(icon.BlackAndWhiteIconPath()))
	systray.SetTitle("D-Clo Local Server")
	systray.SetTooltip("D-Clo Local Server")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	serverHost := "192.168.0.193"
	serverPort := "5000"

	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				if checkServerStatus(serverHost, serverPort) {
					systray.SetIcon(icon.GetIconData(icon.DefaultIconPath()))
					systray.SetTooltip("D-Clo Local Server - Running")
				} else {
					systray.SetIcon(icon.GetIconData(icon.BlackAndWhiteIconPath()))
					systray.SetTooltip("D-Clo Local Server - Stopped")
				}
			case <-mQuit.ClickedCh:
				ticker.Stop()
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
}

func checkServerStatus(host string, port string) bool {
	timeout := 1 * time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return false
	}
	if conn != nil {
		defer conn.Close()
		return true
	}
	return false
}
