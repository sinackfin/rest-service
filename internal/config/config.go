package config

type Config struct {
	AppPort				uint	`mapstructure:"APP_PORT"`
	NatoinalizeAPI_URL 	string	`mapstructure:"NATIONALITY_API_URL"`
	GenderizeAPI_URL 	string	`mapstructure:"GENDERIZE_API_URL"`
	AgifyAPI_URL		string	`mapstructure:"AGIFY_API_URL"`
	PgConf				*PGStoreConfg
}

