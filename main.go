package main

import (
	"Go-Tray/icon"
	"fmt"
	"github.com/getlantern/systray"
	"net"
	"net/http"
	"time"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.GetIconData("blackAndWhite.ico"))
	systray.SetTitle("D-Clo Local Server")
	systray.SetTooltip("D-Clo Local Server")
	mStart := systray.AddMenuItem("Run", "Run D-clo Api Local Server")
	mStop := systray.AddMenuItem("Stop", "Stop D-clo Api Local Server")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	serverHost := "192.168.0.193"
	serverPort := "5000"

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			if checkServerStatus(serverHost, serverPort) {
				systray.SetIcon(icon.GetIconData("default.ico"))
				systray.SetTooltip("D-Clo Local Server - Running")
			} else {
				systray.SetIcon(icon.GetIconData("blackAndWhite.ico"))
				systray.SetTooltip("D-Clo Local Server - Stopped")
			}
		}
	}()

	go func() {
		for {
			select {
			case <-mStart.ClickedCh:
				resp, err := http.Get("http://192.168.0.193:8080/start")
				if err != nil {
					fmt.Println("Error:", err)
					systray.Quit()
					return
				}
				resp.Body.Close()
				break
			case <-mStop.ClickedCh:
				resp, err := http.Get("http://192.168.0.193:8080/stop")
				if err != nil {
					fmt.Println("Error:", err)
					systray.Quit()
					return
				}
				resp.Body.Close()
				break
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
