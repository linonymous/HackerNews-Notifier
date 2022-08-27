package hackernews

import (
	"fmt"
	"log"
	"net/http"
)

type HttpServiceProvider interface {
	MakeRequest(method, url string, response interface{}) error
}

type HackerNewsService struct {
	httpServiceProvider HttpServiceProvider
	config              Config
}

func NewHackerNewsService(provider HttpServiceProvider, config Config) *HackerNewsService {
	if provider == nil || !isValidConfig(config) {
		return nil
	}
	return &HackerNewsService{httpServiceProvider: provider, config: config}
}

func isValidConfig(config Config) bool {
	return config.TopStoriesURLPath != "" && config.StoryInfoURLPath != ""
}

func (h *HackerNewsService) GetTopStories() *TopStories {
	topStoriesURL := h.config.TopStoriesURLPath
	log.Println(topStoriesURL)
	var topStories TopStories
	err := h.httpServiceProvider.MakeRequest(http.MethodGet, topStoriesURL, &topStories)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil
	}
	topStories = topStories[:h.config.MaxLimit]
	return &topStories
}

func (h *HackerNewsService) GetStoryInfo(id int64) *StoryInfo {
	storyInfoURL := fmt.Sprintf(h.config.StoryInfoURLPath, id)
	var storyInfo StoryInfo
	err := h.httpServiceProvider.MakeRequest(http.MethodGet, storyInfoURL, &storyInfo)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil
	}
	return &storyInfo
}
