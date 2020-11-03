// @Author : DAIPENGYUAN
// @File : logger_test
// @Time : 2020/10/21 9:50 
// @Description : zap日志测试

package log

import "testing"

func TestLogger1(t *testing.T) {
	Logger.Info("this is a common INFO msg")
	Logger.Debug("this is a common DEBUG msg")
	GinLogger.Info("this is a gin INFO msg")
	GinLogger.Debug("this is a gin DEBUG msg")
	SetLevel("INFO")
	Logger.Info("this is a common INFO msg")
	Logger.Error("this is a common error msg")
	GinLogger.Info("this is a gin INFO msg")
	GinLogger.Debug("this is a gin DEBUG msg")
}
