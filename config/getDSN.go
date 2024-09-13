package config

func GetDSN() (string, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return "", err
	}
	dsn := "host=" + cfg.DB_Host + "user=" + cfg.DB_User + "password" + cfg.DB_Password + "dbname=" + cfg.DB_Name + "port=" + cfg.DB_Port + "sslmode=" + cfg.SSLMode
	return dsn, nil
}
