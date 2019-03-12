package config

import (
	"fmt"
	"github.com/astaxie/beego"
	"seckill01/models"
	"strings"
)

var (
	secKillConf = models.NewSecKillConf()
	appconfg    = beego.AppConfig
)

func appConfigString(key string) (str string, err error) {
	str = appconfg.String(key)
	if len(str) == 0 {
		err = fmt.Errorf("app config string failed,key: %v", key)
		return
	}
	return
}

func appConfigInt(key string) (i int, err error) {
	i, err = appconfg.Int(key)
	if err != nil {
		err = fmt.Errorf("app config int failed,key: %v", key)
		return
	}
	return
}

func appConfigIntValue(num *int, key string) (err error) {
	i, err := appConfigInt(key)
	if err != nil {
		return
	}
	*num = i
	return
}

func appConfigStringValue(str *string, key string) (err error) {
	s, err := appConfigString(key)
	if err != nil {
		return
	}
	*str = s
	return
}

func initRedisConfig(redisConf *models.RedisConf, keys ...string) (err error) {
	for _, v := range keys {
		if strings.HasSuffix(v, "addr") {
			if err = appConfigStringValue(&redisConf.RedisAddr, v); err != nil {
				break
			}
		} else if strings.HasSuffix(v, "idle") {
			if err = appConfigIntValue(&redisConf.RedisMaxIdle, v); err != nil {
				break
			}
		} else if strings.HasSuffix(v, "active") {
			if err = appConfigIntValue(&redisConf.RedisMaxActive, v); err != nil {
				break
			}
		} else if strings.HasSuffix(v, "timeout") {
			if err = appConfigIntValue(&redisConf.RedisIdleTimeout, v); err != nil {
				break
			}
		}
	}
	return
}

func initRedisBlackConfig() (err error) {
	if err = initRedisConfig(&secKillConf.RedisBlackConf,
		"redis_black_addr",
		"redis_black_idle",
		"redis_black_active",
		"redis_black_idle_timeout"); err != nil {
		return
	}
	return
}

func initRedisLayerToProxyConfig() (err error) {
	if err = initRedisConfig(&secKillConf.RedisLayerToProxyConf,
		"redis_layerToProxy_addr",
		"redis_layerToProxy_idle",
		"redis_layerToProxy_active",
		"redis_layerToProxy_idle_timeout"); err != nil {
		return
	}

	if err = appConfigIntValue(&secKillConf.WriteLayerToProxyGoroutineNum, "write_layerToProxy_goroutine_num"); err != nil {
		return
	}

	if err = appConfigIntValue(&secKillConf.ReadLayerToProxyGoroutineNum, "read_layerToProxy_goroutine_num"); err != nil {
		return
	}
	return
}

func initRedisProxyToLayerConfig() (err error) {
	if err = initRedisConfig(&secKillConf.RedisProxyToLayerConf,
		"redis_proxyToLayer_addr",
		"redis_proxyToLayer_idle",
		"redis_proxyToLayer_active",
		"redis_proxyToLayer_idle_timeout"); err != nil {
		return
	}

	if err = appConfigIntValue(&secKillConf.WriteProxyToLayerGoroutineNum, "write_proxyToLayer_goroutine_num"); err != nil {
		return
	}

	if err = appConfigIntValue(&secKillConf.ReadProxyToLayerGoroutineNum, "read_proxyToLayer_goroutine_num"); err != nil {
		return
	}
	return
}

func initLogConfig() (err error) {
	if err = appConfigStringValue(&secKillConf.LogPath, "log_path"); err != nil {
		return
	}

	if err = appConfigStringValue(&secKillConf.LogLevel, "log_level"); err != nil {
		return
	}

	if err = appConfigStringValue(&secKillConf.CookieSecretKey, "cookie_secretkey"); err != nil {
		return
	}
	return
}

func initLimitConfig() (err error) {

	if err = appConfigIntValue(&secKillConf.AccessLimitConf.IPSecAccessLimit, "ip_sec_access_limit"); err != nil {
		return
	}

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
