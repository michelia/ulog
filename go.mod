module github.com/michelia/ulog

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190506204251-e1dfcc566284
	golang.org/x/net => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190507053917-2953c62de483
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190506145303-2d16b83fe98c
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/rs/zerolog v1.14.3
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.2.2 // indirect
)
