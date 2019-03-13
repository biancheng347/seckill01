package config

import (
	"seckill01/models"
	"seckill01/structModel"
)

var (
	secKillConf = models.NewSecKillConf()
)


func initRedisBlackConfig() (err error) {
	if err = secKillConf.RedisBlackConf.InitRedisConf("redis_black_addr",
		"redis_black_idle",
		"redis_black_active",
		"redis_black_idle_timeout"); err != nil {
		return
	}
	return
}

func initRedisLayerToProxyConfig() (err error) {
	if err = secKillConf.RedisLayerToProxyConf.InitRedisConf(
		"redis_layerToProxy_addr",
		"redis_layerToProxy_idle",
		"redis_layerToProxy_active",
		"redis_layerToProxy_idle_timeout"); err != nil {
		return
	}
	if err = structModel.AppConfigIntValue(&secKillConf.WriteLayerToProxyGoroutineNum, "write_layerToProxy_goroutine_num"); err != nil {
		return
	}

	if err = structModel.AppConfigIntValue(&secKillConf.ReadLayerToProxyGoroutineNum, "read_layerToProxy_goroutine_num"); err != nil {
		return
	}
	return
}

func initRedisProxyToLayerConfig() (err error) {
	if err = secKillConf.RedisProxyToLayerConf.InitRedisConf(
		"redis_proxyToLayer_addr",
		"redis_proxyToLayer_idle",
		"redis_proxyToLayer_active",
		"redis_proxyToLayer_idle_timeout"); err != nil {
		return
	}
	if err = structModel.AppConfigIntValue(&secKillConf.WriteProxyToLayerGoroutineNum, "write_proxyToLayer_goroutine_num"); err != nil {
		return
	}

	if err = structModel.AppConfigIntValue(&secKillConf.ReadProxyToLayerGoroutineNum, "read_proxyToLayer_goroutine_num"); err != nil {
		return
	}
	return
}

func initLogConfig() (err error) {
	if err = secKillConf.Logs.InitLogConfig(); err != nil {
		return
	}
	if err = structModel.AppConfigStringValue(&secKillConf.CookieSecretKey, "cookie_secretkey"); err != nil {
		return
	}
	return
}

func initLimitConfig() (err error) {
	if err = secKillConf.AccessLimitConf.InitAccessLimitConf(); err != nil {
		return
	}
	return
}

func initEtcdConfig() (err error) {
	if err = secKillConf.EtcdConf.InitEtcdConf(
		"etcd_addr",
		"etcd_timeout",
		"etcd_sec_key_prefix",
		"etcd_product_key"); err != nil {
		return
	}
	return
}

func InitConfig() (err error) {
	//配置黑名单Redis
	if err = initRedisBlackConfig(); err != nil {
		return
	}

	//配置接入层->业务逻辑层
	if err = initRedisProxyToLayerConfig(); err != nil {
		return
	}

	//配置业务逻辑层->接入层
	if err = initRedisLayerToProxyConfig(); err != nil {
		return
	}

	//配置etcd 参数
	if err = initEtcdConfig(); err != nil {
		return
	}

	//配置日志文件相关
	if err = initLogConfig(); err != nil {
		return
	}

	//频率限制
	if err = initLimitConfig(); err != nil {
		return
	}
	return
}
