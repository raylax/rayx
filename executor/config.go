package executor

type Config struct {
	Url     string
	Timeout int
	Headers map[string]string
	Cookies map[string]string
}
