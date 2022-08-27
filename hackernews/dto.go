package hackernews

type StoryInfo struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

type TopStories []int64

type Config struct {
	TopStoriesURLPath string `mapstructure:"TOP_STORIES_URL_PATH"`
	StoryInfoURLPath  string `mapstructure:"STORY_INFO_URL_PATH"`
	MaxLimit          int64  `mapstructure:"MAX_LIMIT"`
}
