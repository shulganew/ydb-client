package config

import "flag"

type Config struct {
	DSN      string
	Series   uint64
	Seasons  uint64
	Episodes uint64
}

func GetConfig() Config {
	cfg := Config{}
	flag.StringVar(&cfg.DSN, "dsn", "", "DB DSN")
	flag.Uint64Var(&cfg.Series, "serials", 100, "total serials")
	flag.Uint64Var(&cfg.Seasons, "seasons", 4, "total seasons in each serail")
	flag.Uint64Var(&cfg.Episodes, "episodes", 10, "total episodes in each season")
	flag.Parse()

	//cfg.EmployeesCount = int64(count)
	return cfg
}
