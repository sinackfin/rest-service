package config

type PGStoreConfg struct {
	DB_HOST				string 	`mapstructure:"DB_HOST"`
	DB_PORT				string	`mapstructure:"DB_PORT"`
	DB_USER 			string	`mapstructure:"DB_USER"`
	DB_PASS				string 	`mapstructure:"DB_PASS"`
	DB_DBNAME 			string	`mapstructure:"DB_DBNAME"`
}