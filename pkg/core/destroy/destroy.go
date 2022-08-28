package destroy

import (
	"go.uber.org/zap"
	"leinao-os-adapter/pkg/core/event_manage"
	"leinao-os-adapter/pkg/global/consts"
	"leinao-os-adapter/pkg/global/variable"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	//Listening of system exit signals
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		received := <-c
		variable.ZapLog.Warn(consts.ProcessKilled, zap.String("信号值", received.String()))
		(event_manage.CreateEventManageFactory()).FuzzyCall(variable.EventDestroyPrefix)
		close(c)
		os.Exit(1)
	}()

}
