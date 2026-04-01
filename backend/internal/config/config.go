package config

import "os"

type Config struct {
	Port                string
	GitHubToken         string
	GitHubWebhookSecret string
	ClaudeAPIKey        string
	SupabaseURL         string
}

func Load() (*Config, error) {

	cfg := &Config{
		Port:                getEnv("PORT", "8080"),
		GitHubToken:         os.Getenv("GITHUB_TOKEN"),
		GitHubWebhookSecret: os.Getenv("GITHUB_WEBHOOK_SECRET"),
		ClaudeAPIKey:        os.Getenv("CLAUDE_API_KEY"),
		SupabaseURL:         os.Getenv("SUPABASE_URL"),
	}
	return cfg, nil
}

func getEnv(key, defaultVal string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return defaultVal
}
