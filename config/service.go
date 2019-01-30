package config

type Service struct {
	Port string `envconfig:"serviceport,default=8080" json:"Port"`

	DbType string `envconfig:"dbtype,default=memory" json:"DBType"` // [memory | mssql]
}
