package config

type AppConf struct {
	AppPort       uint `mapstructure:"APP_PORT"`
	DBConnTimeout uint `mapstructure:"DB_CONN_TIMEOUT"`
	//**************external API*****************//
	NatoinalizeAPI_URL string `mapstructure:"NATIONALITY_API_URL"`
	GenderizeAPI_URL   string `mapstructure:"GENDERIZE_API_URL"`
	AgifyAPI_URL       string `mapstructure:"AGIFY_API_URL"`
	//******************************************//
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_DBNAME"`
}
