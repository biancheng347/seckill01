package config

import (
	"fmt"
	"github.com/astaxie/beego"
	"seckill01/models"
	"strings"
)

var (
	secKillConf = models.NewSecKillConf()
	appconfg = beego.AppConfig
)

func appConfigString(key string)(str string, err error) {
	str = appconfg.String(key)
	if len(str) == 0 {
		err = fmt.Errorf("app config string failed,key: %v",key)
		return
	}
	return
}

func appConfigInt(key string)(i int,err error) {
	i,err = appconfg.Int(key)
	if err != nil {
		err = fmt.Errorf("app config int failed,key: %v",key)
		return
	}
	return
}



func initRedisConfig(redisConf *models.RedisConf,keys ...string) (err error) {
	for _,v := range keys {
		if strings.HasSuffix(v,"addr") {
			str,err := appConfigString(v)
			if err != nil {
				break
			}
			redisConf.RedisAddr= str
		}else if strings.HasSuffix(v,"idle") {
			i,err  := appConfigInt(v)
			if err != nil {
				break
			}
			redisConf.RedisMaxIdle = i
		}else if strings.HasSuffix(v,"active") {
			i,err  := appConfigInt(v)
			if err != nil {
				break
			}
			redisConf.RedisMaxActive = i
		}else if strings.HasSuffix(v,"timeout") {
			i,err  := appConfigInt(v)
			if err != nil {
				break
			}
			redisConf.RedisIdleTimeout = i
		}
	}
	return
}

func initRedisBlackConfig() (err error) {
	err = initRedisConfig(&secKillConf.RedisBlackConf,
		"redis_black_addr",
		"redis_black_idle",
		"redis_black_active",
		"redis_black_idle_timeout")
	if err != nil {
		return
	}
	return
}










func initRedisLayerToProxyConfig() (err error) {
	err = initRedisConfig(&secKillConf.RedisProxyToLayerConf,
		"redis_layerToProxy_addr",
		"redis_layerToProxy_idle",
		"redis_layerToProxy_active",
		"redis_layerToProxy_idle_timeout")
	if err != nil {
		return
	}

	i,err  := appConfigInt("write_layerToProxy_goroutine_num")
	if err != nil {
		return
	}
	secKillConf.WriteLayerToProxyGoroutineNum = i

	i,err  = appConfigInt("read_layerToProxy_goroutine_num")
	if err != nil {
		return
	}
	secKillConf.ReadLayerToProxyGoroutineNum = i

	return
}



func initRedisProxyToLayerConfig() (err error) {
	redisProxyToLayerAddr := beego.AppConfig.String("redis_proxyToLayer_addr")
	if len(redisProxyToLayerAddr) == 0 {
		err = fmt.Errorf("initRedisProxyToLayerConfig redisProxyToLayerAddr failed,err")
		return
	}
	secKillConf.RedisProxyToLayerConf.RedisAddr = redisProxyToLayerAddr

	redisProxyToLayerIdle, err := beego.AppConfig.Int("redis_proxyToLayer_idle")
	if err != nil {
		err = fmt.Errorf("initRedisProxyToLayerConfig redisProxyToLayerIdle failed,err:%v", err)
		return
	}
	secKillConf.RedisProxyToLayerConf.RedisMaxIdle = redisProxyToLayerIdle

	redisProxyToLayerActvie, err := beego.AppConfig.Int("redis_proxyToLayer_active")
	if err != nil {
		err = fmt.Errorf("initRedisProxyToLayerConfig redisProxyToLayerActvie failed,err:%v", err)
		return
	}
	secKillConf.RedisProxyToLayerConf.RedisMaxActive = redisProxyToLayerActvie

	redisProxyToLayerIdleTimeout, err := beego.AppConfig.Int("redis_proxyToLayer_idle_timeout")
	if err != nil {
		err = fmt.Errorf("initRedisProxyToLayerConfig redisProxyToLayerIdleTimeout failed,err:%v", err)
		return
	}
	secKillConf.RedisProxyToLayerConf.RedisIdleTimeout = redisProxyToLayerIdleTimeout

	writeProxyToLayerGoroutineNum, err := beego.AppConfig.Int("write_proxyToLayer_goroutine_num")
	if err != nil {
		err = fmt.Errorf("initRedisProxyToLayerConfig writeProxyToLayerGoroutineNum failed,err:%v", err)
		return
	}
	secKillConf.WriteProxyToLayerGoroutineNum = writeProxyToLayerGoroutineNum

	readProxyToLayerGoroutineNum, err := beego.AppConfig.Int("read_proxyToLayer_goroutine_num")
	if err != nil {
		err = fmt.Errorf("initRedisProxyToLayerConfig readProxyToLayerGoroutineNum failed,err:%v", err)
		return
	}
	secKillConf.ReadProxyToLayerGoroutineNum = readProxyToLayerGoroutineNum

	return
}

func initLogConfig() (err error) {
	logPath := beego.AppConfig.String("log_path")
	if len(logPath) == 0 {
		err = fmt.Errorf("initLogConfig logPath failed,err")
		return
	}
	secKillConf.LogPath = logPath

	logLevel := beego.AppConfig.String("log_level")
	if len(logLevel) == 0 {
		err = fmt.Errorf("initLogConfig logLevel failed,err")
		return
	}
	secKillConf.LogLevel = logLevel

	cookieSecretKey := beego.AppConfig.String("cookie_secretkey")
	if len(cookieSecretKey) == 0 {
		err = fmt.Errorf("initLogConfig cookieSecretKey failed,err")
		return
	}
	secKillConf.CookieSecretKey = cookieSecretKey

	return
}



func initLimitConfig() (err error) {
	ipSecAccessLimit, err := beego.AppConfig.Int("ip_sec_access_limit")
	if err != nil {
		err = fmt.Errorf("initLimitConfig ipSecAccessLimit failed,err:%v", err)
		return
	}
	secKillConf.AccessLimitConf.IPSecAccessLimit = ipSecAccessLimit

	userSecAccessLimit, err := beego.AppConfig.Int("user_sec_access_limit")
	if err != nil {
		err = fmt.Errorf("initLimitConfig userSecAccessLimit failed,err:%v", err)
		return
	}
	secKillConf.AccessLimitConf.UserSecAccessLimit = userSecAccessLimit

	ipMinAccessLimit, err := beego.AppConfig.Int("ip_min_access_limit")
	if err != nil {
		err = fmt.Errorf("initLimitConfig ipMinAccessLimit failed,err:%v", err)
		return
	}
	secKillConf.AccessLimitConf.IPMinAccessLimit = ipMinAccessLimit

	userMinAccessLimit, err := beego.AppConfig.Int("user_min_access_limit")
	if err != nil {
		err = fmt.Errorf("initLimitConfig userMinAccessLimit failed,err:%v", err)
		return
	}
	secKillConf.AccessLimitConf.UserMinAccessLimit = userMinAccessLimit

	return
}

func InitConfig() (err error) {
	//配置黑名单Redis
	err = initRedisBlackConfig()
	if err != nil {
		return
	}

	//配置接入层->业务逻辑层
	err = initRedisProxyToLayerConfig()
	if err != nil {
		return
	}

	//配置业务逻辑层->接入层
	err = initRedisLayerToProxyConfig()
	if err != nil {
		return
	}

	////配置etcd 参数
	//err = initEtcdConfig()
	//if err != nil {
	//	return
	//}


	//配置日志文件相关
	err = initLogConfig()
	if err != nil {
		return
	}

	//频率限制
	err = initLimitConfig()
	if err != nil {
		return
	}

	return
}
