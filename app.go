package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/dreamsxin/go-netsniffer/models"
	"github.com/google/martian/v3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	//"net/http/cookiejar"
)

const authorityName string = "GoNetSniffer Proxy Authority"

// App struct
type App struct {
	ctx      context.Context
	config   models.Config
	serve    *martian.Proxy
	lock     sync.Mutex
	dataChan chan *models.Packet
}

// NewApp creates a new App application struct
func NewApp() *App {

	a := &App{
		config: models.Config{
			Port:        9000,
			AutoProxy:   true,
			SaveLogFile: false,
		},
		dataChan: make(chan *models.Packet, 1000),
	}

	go a.RunLoop()
	return a
}

func (a *App) RunLoop() {

	file, err := os.OpenFile(fmt.Sprintf("log%s.txt", time.Now().Format(time.DateOnly)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 循环读取 dataChan
	for packet := range a.dataChan {
		// 处理数据
		if a.config.FilterHost != "" {
			if !strings.Contains(packet.Host, a.config.FilterHost) {
				continue
			}
		}

		runtime.EventsEmit(a.ctx, "Packet", packet)
		if a.config.SaveLogFile {
			b, err := json.Marshal(packet)
			if err != nil {
				log.Println("json.Marshal", err)
				continue
			}
			// 追加内容
			file.Write(b)
			file.WriteString("\n\n")
		}
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	b, err := os.ReadFile("config.json")
	if err != nil {
		log.Println("Read config.json", err)
		return
	}
	err = json.Unmarshal(b, &a.config)
	if err != nil {
		log.Println("Unmarshal config.json", err)
		return
	}
}

func (a *App) shutdown(ctx context.Context) {

	b, err := json.Marshal(a.config)
	if err != nil {
		log.Println("Marshal config.json", err)
		return
	}
	err = os.WriteFile("config.json", b, 0644)
	if err != nil {
		panic(err)
	}
}

func (a *App) GetConfig() models.Config {
	return a.config
}

func (a *App) SetConfig(field string, config models.Config) {
	a.config = config
	log.Println("SetConfig", field, config)
}

func (a *App) OpenDirectoryDialog() (string, error) {
	directoryPath, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: ".",
		Title:            "选择目录",
	})
	if err != nil {
		return "", fmt.Errorf("failed opening dialog - %s", err.Error())
	}
	return directoryPath, nil
}
