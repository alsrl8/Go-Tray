package main

import (
	"Go-Tray/icon"
	"fmt"
	"github.com/getlantern/systray"
	"net"
	"net/http"
	"time"
)

type serverType int

const (
	dclo serverType = iota
	admin
)

type serverInfo struct {
	serverType
	host      string
	port      string
	isRunning bool
}

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.GetIconData("blackAndWhite.ico"))
	systray.SetTitle("D-Clo Local Server")
	systray.SetTooltip("D-Clo Local Server")
	dcloStart := systray.AddMenuItem("Start Dclo", "Run D-clo Api Local Server")
	dcloStop := systray.AddMenuItem("Stop Dclo", "Stop D-clo Api Local Server")
	adminStart := systray.AddMenuItem("Start Admin", "Run Admin Api Local Server")
	adminStop := systray.AddMenuItem("Stop Admin", "Stop Admin Api Local Server")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	dcloServerInfo := serverInfo{
		host:       "192.168.0.193",
		port:       "5000",
		serverType: dclo,
		isRunning:  false,
	}
	adminServerInfo := serverInfo{
		host:       "192.168.0.193",
		port:       "5001",
		serverType: admin,
		isRunning:  false,
	}

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			dcloServerInfo.isRunning = checkServerStatus(dcloServerInfo)
			adminServerInfo.isRunning = checkServerStatus(adminServerInfo)

			if !dcloServerInfo.isRunning && !adminServerInfo.isRunning {
				systray.SetIcon(icon.GetIconData("blackAndWhite.ico"))
				systray.SetTooltip("Every server is stopped")
				dcloStart.Show()
				dcloStop.Hide()
				adminStart.Show()
				adminStop.Hide()
				continue
			} else {
				systray.SetIcon(icon.GetIconData("default.ico"))
			}

			if dcloServerInfo.isRunning {
				dcloStart.Hide()
				dcloStop.Show()
				if adminServerInfo.isRunning {
					adminStart.Hide()
					adminStop.Show()
					systray.SetTooltip("Servers(Dclo, Admin) are running")
				} else {
					adminStart.Show()
					adminStop.Hide()
					systray.SetTooltip("Dclo server is running")
				}
			} else {
				dcloStart.Show()
				dcloStop.Hide()
				adminStart.Hide()
				adminStop.Show()
				systray.SetTooltip("Admin server is running")
			}
		}
	}()

	go func() {
		for {
			select {
			case <-dcloStart.ClickedCh:
				resp, err := http.Get("http://192.168.0.193:8080/start")
				if err != nil {
					fmt.Println("Error:", err)
					systray.Quit()
					return
				}
				resp.Body.Close()
				break
			case <-dcloStop.ClickedCh:
				resp, err := http.Get("http://192.168.0.193:8080/stop")
				if err != nil {
					fmt.Println("Error:", err)
					systray.Quit()
					return
				}
				resp.Body.Close()
				break
			case <-adminStart.ClickedCh:
				resp, err := http.Get("http://192.168.0.193:8080/start-admin")
				if err != nil {
					fmt.Println("Error:", err)
					systray.Quit()
					return
				}
				resp.Body.Close()
				break
			case <-adminStop.ClickedCh:
				resp, err := http.Get("http://192.168.0.193:8080/stop-admin")
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

func checkServerStatus(info serverInfo) bool {
	timeout := 1 * time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(info.host, info.port), timeout)
	if err != nil {
		return false
	}
	if conn != nil {
		defer conn.Close()
		return true
	}
	return false
}
