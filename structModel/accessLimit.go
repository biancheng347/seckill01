package structModel

type AccessLimitConf struct {
	IPSecAccessLimit   int
	UserSecAccessLimit int
	IPMinAccessLimit   int
	UserMinAccessLimit int
}


func (p *AccessLimitConf) InitAccessLimitConf() (err error) {
	if err = appConfigIntValue(&p.IPSecAccessLimit, "ip_sec_access_limit"); err != nil {
		return
	}
	if err = appConfigIntValue(&p.UserSecAccessLimit, "user_sec_access_limit"); err != nil {
		return
	}
	if err = appConfigIntValue(&p.IPMinAccessLimit, "ip_min_access_limit"); err != nil {
		return
	}
	if err = appConfigIntValue(&p.UserMinAccessLimit, "user_min_access_limit"); err != nil {
		return
	}
	return
}