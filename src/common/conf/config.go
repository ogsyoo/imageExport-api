package conf

var (
	DatabaseURL string
	DocPath     string
	RedisURL    string
	RedisTag    string = "traffic"
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
