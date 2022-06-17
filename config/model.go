package config


//此文件保存公共的结构

const (
	SmsTplDelAccountAsk = "DelAccountAsk"
)

type Mysql struct {
	Host            string `json:"host"`
	Host_rw         string `json:"host_rw"`         // 读写分离连接
	Host_admin_read string `json:"host_admin_read"` // 读写分离连接
	Db              string `json:"db"`
	MaxCon          int    `json:"max_con"`
}

type Redis struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type Log struct {
	SavePath	string	`json:"save_paht"`
}


type CommonConf struct {
	AppName string `json:"app_name"`
	Env     string `json:"env"` //alpha beta pre prod
	Debug   bool   `json:"debug"` //alpha beta pre prod
	Log     Log    `json:"log"`
	Mysql   Mysql  `json:"mysql"`

	MysqlLog   Mysql         `json:"mysql_log"`
	Redis      Redis         `json:"redis"`
	Oss        Oss           `json:"oss"`
	AiConf     *AiConfig     `json:"ai_conf"`
	U3dConf    *U3dConfig    `json:"u3d_conf"`
	UploadConf *UploadConfig `json:"upload_conf"`

	TarsSuffix                    string                      `json:"tars_suffix"`
}



//
type Oss struct {
	Endpoint     string `json:"endpoint"`      //数据中心域名
	AccessKey    string `json:"access_key"`    //校验key
	AccessSecret string `json:"access_secret"` //校验密钥
	Bucket       string `json:"bucket"`        //存储空间名称
	Env          string `json:"env"`           //环境 alpha beta prod
}


// AI配置
type AiConfig struct {
	ApiHost      string `json:"api_host"`
	ApiHostV2    string `json:"api_host_v2"`
}

// u3d配置
type U3dConfig struct {
	ApiHost      string `json:"api_host"`
}

// 上传配置配置
type UploadConfig struct {
	CdnHost      string `json:"cdn_host"`
	UploadDir    string `json:"upload_dir"`
}

