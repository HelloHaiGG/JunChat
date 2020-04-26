package config

type appCfg struct {
	Servers Servers      `yaml:"servers"`
	Redis   RedisCfg     `yaml:"redis"`
	Mysql   MysqlCfg     `yaml:"mysql"`
	Grpc    GrpcCfg      `yaml:"grpc"`
	Etcd    EtcdCfg      `yaml:"etcd"`
	CN      ConnectNodes `yaml:"connect_port"`
	CC      ConnectNodes `yaml:"core_port"`
	JC      JunChat      `yaml:"jun_chat"`
}

//服务
type Servers struct {
	Gateway string `yaml:"gateway"`
	Core    string `yaml:"core"`
	Connect string `json:"connect"`
	Queue   string `json:"queue"`
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

//connect 端口
type ConnectNodes struct {
	Nodes []string `yaml:"nodes"`
}

//JunChat 端口
type JunChat struct {
	Port string `yaml:"port"`
}
