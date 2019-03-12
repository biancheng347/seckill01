package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"seckill01/models"
)

var (
	secKillConf = models.NewSecKillConf()
)

func initRedisBlackConfig() (err error) {
	redisBlackAddr := beego.AppConfig.String("redis_black_addr")
	if len(redisBlackAddr) == 0 {
		err = fmt.Errorf("initRedisBlackConfig redisBlackAddr failed,err")
		return
	}
	secKillConf.RedisBlackConf.RedisAddr = redisBlackAddr

	redisBlackIdle, err := beego.AppConfig.Int("redis_black_idle")
	if err != nil {
		err = fmt.Errorf("initRedisBlackConfig redisBlackIdle failed,err:%v", err)
		return
	}
	secKillConf.RedisBlackConf.RedisMaxIdle = redisBlackIdle

	redisBlackActvie, err := beego.AppConfig.Int("redis_black_active")
	if err != nil {
		err = fmt.Errorf("initRedisBlackConfig redisBlackActvie failed,err:%v", err)
		return
	}
	secKillConf.RedisBlackConf.RedisMaxActive = redisBlackActvie

	redisBlackIdleTimeout, err := beego.AppConfig.Int("redis_black_idle_timeout")
	if err != nil {
		err = fmt.Errorf("initRedisBlackConfig redisBlackIdleTimeout failed,err:%v", err)
		return
	}
	secKillConf.RedisBlackConf.RedisIdleTimeout = redisBlackIdleTimeout

	return
}

func initRedisLayerToProxyConfig() (err error) {
	redisLayerToProxyAddr := beego.AppConfig.String("redis_layerToProxy_addr")
	if len(redisLayerToProxyAddr) == 0 {
		err = fmt.Errorf("initRedisLayerToProxyConfig redisLayerToProxyAddr failed,err")
		return
	}

	redisLayerToProxyIdle, err := beego.AppConfig.Int("redis_layerToProxy_idle")
	if err != nil {
		err = fmt.Errorf("initRedisLayerToProxyConfig redisLayerToProxyIdle failed,err:%v", err)
		return
	}
	secKillConf.RedisLayerToProxyConf.RedisIdleTimeout = redisLayerToProxyIdle

	redisLayerToProxyActvie, err := beego.AppConfig.Int("redis_layerToProxy_active")
	if err != nil {
		err = fmt.Errorf("initRedisLayerToProxyConfig redisLayerToProxyActvie failed,err:%v", err)
		return
	}
	secKillConf.RedisLayerToProxyConf.RedisMaxActive = redisLayerToProxyActvie

	redisLayerToProxyIdleTimeout, err := beego.AppConfig.Int("redis_layerToProxy_idle_timeout")
	if err != nil {
		err = fmt.Errorf("initRedisLayerToProxyConfig redisLayerToProxyIdleTimeout failed,err:%v", err)
		return
	}
	secKillConf.RedisLayerToProxyConf.RedisIdleTimeout = redisLayerToProxyIdleTimeout

	writeLayerToProxyGoroutineNum, err := beego.AppConfig.Int("write_layerToProxy_goroutine_num")
	if err != nil {
		err = fmt.Errorf("initRedisLayerToProxyConfig writeLayerToProxyGoroutineNum failed,err:%v", err)
		return
	}
	secKillConf.WriteLayerToProxyGoroutineNum = writeLayerToProxyGoroutineNum

	readLayerToProxyGoroutineNum, err := beego.AppConfig.Int("read_layerToProxy_goroutine_num")
	if err != nil {
		err = fmt.Errorf("initRedisLayerToProxyConfig readLayerToProxyGoroutineNum failed,err:%v", err)
		return
	}
	secKillConf.ReadLayerToProxyGoroutineNum = readLayerToProxyGoroutineNum
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

func initConfig() (err error) {
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
