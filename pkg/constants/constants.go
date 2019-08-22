package constants

import "time"

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

const FPath = "conf/conf.toml"

const (
	JwtExpireDuration = 3 * time.Hour
	JwtIssuer         = "market"
)
