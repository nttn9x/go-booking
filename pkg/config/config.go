package config

import "github.com/alexedwards/scs/v2"

type AppConfig struct {
	InProduction bool
	Session      *scs.SessionManager
}

var App *AppConfig

func Init(s *scs.SessionManager) {
	App = &AppConfig{
		InProduction: false,
		Session:      s,
	}
}
