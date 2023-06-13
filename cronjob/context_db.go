package cronjob

import (
	"context"

	"github.com/devarchi33/goutils/ctxdb"
	"github.com/devarchi33/goutils/echomiddleware"
	"github.com/devarchi33/goutils/kafka"

	"xorm.io/xorm"
)

func ContextDB(service string, xormEngine *xorm.Engine, kafkaConfig kafka.Config) Middleware {
	ctxdb := ctxdb.New(xormEngine, service, kafkaConfig)

	return func(job, action string, next HandlerFunc) HandlerFunc {
		return func(ctx context.Context) error {
			session := ctxdb.NewSession(ctx)
			defer session.Close()

			ctx = context.WithValue(ctx, echomiddleware.ContextDBName, session)

			return next(ctx)
		}
	}
}
