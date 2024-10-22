package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"time"

	"github.com/dreamsxin/go-seo/events"
	"github.com/dreamsxin/go-seo/models"
	"github.com/dreamsxin/go-sitemap"
	"github.com/dreamsxin/go-sitemap/crawl"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx    context.Context
	config models.Config
}

// NewApp creates a new App application struct
func NewApp() *App {

	a := &App{
		config: models.Config{
			Port:        9000,
			AutoProxy:   true,
			SaveLogFile: false,
			Sitemap: models.SitemapConfig{
				Dsturl:         "https://",
				Savepath:       ".",
				Filename:       "sitemap.xml",
				Concurrency:    10,
				Crawltimeout:   10,
				Requesttimeout: 10,
			},
		},
	}

	return a
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

func (a *App) GetSitemapConfig() models.SitemapConfig {
	return a.config.Sitemap
}

func (a *App) GenerateSitemap(config models.SitemapConfig) {
	a.config.Sitemap = config

	log.Println("GenerateSitemap", config)
	go func() {
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
				MaxIdleConns:        config.Concurrency,
				MaxIdleConnsPerHost: config.Concurrency,
				MaxConnsPerHost:     config.Concurrency,
				IdleConnTimeout:     crawl.DefaultKeepAlive,
			},
			Timeout: time.Duration(config.Requesttimeout) * time.Second,
		}

		oldurls := make(map[string]*sitemap.URL)

		siteMap, siteMapErr := crawl.CrawlDomain(
			config.Dsturl,
			crawl.SetMaxConcurrency(config.Concurrency),
			crawl.SetCrawlTimeout(time.Duration(config.Crawltimeout)*time.Second),
			crawl.SetKeepAlive(crawl.DefaultKeepAlive),
			crawl.SetTimeout(time.Duration(config.Requesttimeout)*time.Second),
			crawl.SetClient(client),
			crawl.SetSitemapURLS(oldurls),
			crawl.SetCrawlValidator(func(v *sitemap.URL) bool {
				return true
			}),
			crawl.SetEventCallbackReadLink(func(hrefResolved *url.URL, linkReader *crawl.LinkReader) {

			}),
		)

		if siteMapErr != nil {
			a.FireErrorEvent(events.EVENT_CODE_SITEMAP, fmt.Sprintf("生成失败: %s", siteMapErr.Error()))
			return
		}
		//siteMap.WriteMap(os.Stdout)

		file, err := os.OpenFile(filepath.Join(config.Savepath, config.Filename), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
		if err != nil {
			a.FireErrorEvent(events.EVENT_CODE_SITEMAP, fmt.Sprintf("生成失败: %s", err.Error()))
			return
		}
		defer file.Close()

		// 初始化
		sm := sitemap.New()

		urls := siteMap.GetURLS()
		for _, v := range urls {
			sm.Add(v)
		}

		//排序
		sort.Slice(sm.URLs, func(i, j int) bool {
			return sm.URLs[i].Priority >= sm.URLs[j].Priority
		})

		_, err = sm.WriteTo(file)
		if err != nil {
			a.FireErrorEvent(events.EVENT_CODE_SITEMAP, fmt.Sprintf("生成失败: %s", err.Error()))
		} else {
			a.FireEvent(events.EVENT_CODE_SITEMAP, "生成成功")
			exec.Command(`explorer`, `/select,`, filepath.Join(config.Savepath, config.Filename)).Run()
		}
	}()
}

func (a *App) FireEvent(code int, msg string) {

	runtime.EventsEmit(a.ctx, events.EVENT_TYPE_RESPONSE, &events.Event{Type: events.GENERAL, Code: code, Message: msg})
}

func (a *App) FireErrorEvent(code int, msg string) {
	log.Println("FireErrorEvent", code, msg)
	runtime.EventsEmit(a.ctx, events.EVENT_TYPE_ERROR, &events.Event{Type: events.ERROR, Code: code, Message: msg})
}
