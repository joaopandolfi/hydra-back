package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/flosch/pongo2"
	"github.com/gorilla/sessions"
	"github.com/joaopandolfi/blackwhale/configurations"
	"github.com/unrolled/secure"
)

type Configs struct {
	Token           string
	Debug           bool
	TLSCert         string
	TLSKey          string
	DefaultPassword string
}

var Config Configs

func Load(args []string) configurations.Configurations {
	var confFile map[string]string
	confFile = configurations.LoadJsonFile("./config.json")

	Config = Configs{
		Token:           "$238#!%s@233**#sd*",
		Debug:           confFile["DEBUG"] == "true",
		TLSCert:         confFile["TLS_CERT"],
		TLSKey:          confFile["TLS_KEY"],
		DefaultPassword: confFile["HYDRA_DEFAULT_PASSWORD"],
	}

	if len(args) == 3 {
		if args[0] == "ssl" {
			Config.Debug = false
			Config.TLSCert = args[1]
			Config.TLSKey = args[2]
		}
	}

	mongoPool, _ := strconv.Atoi(confFile["MONGO_POOL"])
	bcryptCost, _ := strconv.Atoi(confFile["BCRYPT_COST"])
	tokenValidity, _ := strconv.Atoi(confFile["TOKEN_VALIDITY_MINUTES"])

	return configurations.Configurations{
		Name: "Hydra Back - GO",

		MysqlUrl: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			confFile["MYSQL_USER"],
			confFile["MYSQL_PASSWORD"],
			confFile["MYSQL_HOST"],
			confFile["MYSQL_PORT"],
			confFile["MYSQL_DB"]),

		MongoUrl:  confFile["MONGO_URL"],
		MongoDb:   confFile["MONGO_DB"],
		MongoPool: mongoPool,

		Port: confFile["PORT"],
		CORS: confFile["CORS"],

		Timeout: configurations.Timeout{
			Write: 60 * time.Second,
			Read:  60 * time.Second,
		},

		MaxSizeMbUpload: 10 << 55, // min << max

		BCryptSecret: confFile["BCRYPT_SECRET"],
		ResetHash:    confFile["RESET_HASH"],

		// Session
		Session: configurations.SessionConfiguration{
			Name:  confFile["SESSION_NAME"],
			Store: sessions.NewCookieStore([]byte(confFile["SESSION_STORE"])),
			Options: &sessions.Options{
				Path:     "/",
				MaxAge:   3600 * 2, //86400 * 7,
				HttpOnly: true,
			},
		},

		Security: configurations.Opsec{
			Options: secure.Options{
				BrowserXssFilter:   true,
				ContentTypeNosniff: false, // Da pau nos js
				SSLHost:            "locahost:443",
				SSLRedirect:        false,
			},
			BCryptCost:    bcryptCost,
			TLSCert:       "",
			TLSKey:        "",
			JWTSecret:     confFile["JWT_SECRET"],
			TokenValidity: tokenValidity,
		},

		Templates: make(map[string]*pongo2.Template),

		// Slack
		SlackToken:   "",
		SlackWebHook: []string{"", ""},
		SlackChannel: "",
	}
}
