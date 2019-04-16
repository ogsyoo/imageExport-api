package conf

import (
	"ogsyoo/imageExport-api/src/sse"
)

var (
	SseClient   *sse.Server
	DatabaseURL string
	DocPath     string
	RedisURL    string
	RedisTag    string = "traffic"
	PackeDoc    string
	UiDoc       string
)

var Package = struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Repository struct {
		Type string `json:"type"`
		Url  string `json:"url"`
	} `json:"repository"`
	Scripts interface{} `json:"scripts"`
}{}

var BaseInfo = struct {
	Prefix string
}{}
