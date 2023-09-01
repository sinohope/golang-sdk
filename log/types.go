package log

type App struct {
	Name string `toml:"name"`
}

type Log struct {
	Stdout struct {
		Enable bool `toml:"enable"`
		Level  int  `toml:"level"`
	} `toml:"stdout"`
	File struct {
		Enable bool   `toml:"enable"`
		Level  int    `toml:"level"`
		Path   string `toml:"path"`
		MaxAge int    `toml:"max-age"`
	} `toml:"file"`
}
