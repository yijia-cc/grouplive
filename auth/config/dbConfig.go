package config

var UserDBConfig = struct {
	Driver   string
	Hostname string
	Port     string
	DBname   string
	Username string
	Password string
}{
	Driver:   "mysql",
	Hostname: "allgame.fun",
	Port:     "18300",
	DBname:   "auth-staging",
	Username: "auth-staging",
	Password: "6CbmK37rDtX62va2kHyrzAzGhG4eSH",
}
