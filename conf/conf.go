package conf

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Log struct {
	Level     string
	Formatter string
	FileName  string
	Caller    bool
}

var LogConf = &Log{}
