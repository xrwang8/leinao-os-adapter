package config

import (
	_ "leinao-os-adapter/pkg/core/destroy" // 监听程序退出信号，用于资源的释放
	"leinao-os-adapter/pkg/global/my_errors"
	"leinao-os-adapter/pkg/global/variable"
	"leinao-os-adapter/pkg/http/validator/common/register_validator"
	"leinao-os-adapter/pkg/service/sys_log_hook"
	"leinao-os-adapter/pkg/utils/validator_translation"
	"leinao-os-adapter/pkg/utils/yml_config"
	"leinao-os-adapter/pkg/utils/zap_factory"
	"log"
)

func init() {
	//初始化表单参数验证器，注册在容器（Web、Api共用容器）
	register_validator.WebRegisterValidator()
	register_validator.ApiRegisterValidator()
	// 启动针对配置文件(confgi.yml)变化的监听， 配置文件操作指针，初始化为全局变量
	variable.ConfigYml = yml_config.CreateYamlFactory()
	variable.ConfigYml.ConfigFileChangeListen()
	// 初始化全局日志句柄，并载入日志钩子处理函数
	variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)
	//全局注册 validator 错误翻译器,zh 代表中文，en 代表英语
	if err := validator_translation.InitTrans("zh"); err != nil {
		log.Fatal(my_errors.ErrorsValidatorTransInitFail + err.Error())
	}
}
