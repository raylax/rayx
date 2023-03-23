package cmd

type Config struct {
	Url     string
	Headers map[string]string
	Cookies map[string]string
}
