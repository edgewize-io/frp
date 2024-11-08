package legacy

func ChangeClientConfFromIni(common ClientCommonConf) ClientCommonConf {
	common.TLSEnable = true
	common.DisableCustomTLSFirstByte = false

	common.AdminAddr = "127.0.0.1"
	common.AdminPort = 2333

	return common
}
