package config

type appCfg struct {
	LTopic  LogTopic `yaml:"l_topic"`
	Kafka   KafkaCfg `yaml:"kafka"`
	Servers Servers  `yaml:"servers"`
	Redis   RedisCfg `yaml:"redis"`
	Mysql   MysqlCfg `yaml:"mysql"`
	Grpc    GrpcCfg  `yaml:"grpc"`
	Etcd    EtcdCfg  `yaml:"etcd"`
}

type LogTopic struct {
	Order   string `yaml:"order"`
	Finance string `yaml:"finance"`
	Gateway string `yaml:"gateway"`
}

type KafkaCfg struct {
	Brokers []string `yaml:"brokers"`
	Timeout int      `yaml:"timeout"`
}

//服务
type Servers struct {
	Order   string `yaml:"order"`
	Finance string `yaml:"finance"`
	Gateway string `yaml:"gateway"`
}

//redis 
type RedisCfg struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	DB          int    `yaml:"db"`
	Password    string `yaml:"password"`
	MaxRetry    int    `yaml:"max_retry"`
	DialTimeout int    `yaml:"dial_timeout"`
	MaxConnAge  int    `yaml:"max_conn_age"`
}

//mysql
type MysqlCfg struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	IsDebug  string `yaml:"is_debug"`
}

//Etcd
type EtcdCfg struct {
	Hosts             []string `yaml:"host"`
	Port              int      `yaml:"port"`
	User              string   `yaml:"user"`
	Password          string   `yaml:"password"`
	DialTimeOut       int32    `yaml:"dial_time_out"`
	DialKeepAliveTime int32    `yaml:"dial_keep_alive_time"`
}

//grpc
type GrpcCfg struct {
	CallTimeOut int32 `yaml:"call_time_out"`
}
