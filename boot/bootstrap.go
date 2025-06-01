package boot

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"new-blog/app/admin"
	"new-blog/core/config"
	"new-blog/core/event"
	"new-blog/core/http"
	logging "new-blog/core/logger"
	"new-blog/core/orm"
)

func bootstrap(
	lifecycle fx.Lifecycle,
	server *http.Service,
	db *gorm.DB,
	event *event.Service,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				zap.S().Infoln("启动Web服务器...", server.Server.Addr)
				err := server.Server.ListenAndServe()
				if err != nil {
					_ = closeDb(db)
					return
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.S().Error("停止Web服务器")
			_ = closeDb(db)
			event.Shutdown()
			return server.Server.Shutdown(ctx)
		},
	})
}

func closeDb(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Errorln("数据库实例获取失败！", err)
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		zap.S().Errorln("无法关闭数据库", err)
		return err
	}
	return nil
}

var Module = fx.Options(
	// 配置模块
	config.Module,
	// 日志模块
	logging.Module,
	// 事件模块
	event.Module,
	orm.Module,
	// http模块
	http.Module,

	admin.Module,

	fx.Invoke(bootstrap),
)
