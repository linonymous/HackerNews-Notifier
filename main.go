package main

import (
	"HackerNews/hackernews"
	http2 "HackerNews/http"
	"HackerNews/notifier"
	"log"
	"net/http"
)

const (
	HackerNews = "HackerNews"
	ConfigPath = "./"
)

func main() {
	notifier := notifier.NewNotifier()
	config, err := hackernews.LoadConfig(ConfigPath)
	if err != nil {
		log.Fatal("could not read config", err)
	}
	httpProvider := http2.NewHttpService(http.Client{})
	hackerNewsService := hackernews.NewHackerNewsService(httpProvider, *config)
	topStories := hackerNewsService.GetTopStories()
	if topStories == nil {
		log.Fatal("could not get top stories")
	}
	for _, story := range *topStories {
		storyInfo := hackerNewsService.GetStoryInfo(story)
		if storyInfo == nil {
			continue
		}
		err := notifier.Notify(HackerNews, storyInfo.Title, storyInfo.URL)
		log.Println(err)
	}

}
