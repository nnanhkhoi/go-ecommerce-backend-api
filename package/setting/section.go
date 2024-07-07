package setting

type Config struct {
	Mysql MYSQLSetting `mapstructure:"mysql"`
}

type MYSQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mastructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"ConnMaxLifetime"`
}

/*
Config
Redis
Mysql
*/
