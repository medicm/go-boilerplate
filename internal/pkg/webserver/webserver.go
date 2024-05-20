package webserver

import (
	"boilerplate/internal/app/handlers"
	"boilerplate/internal/pkg/config"
	"boilerplate/internal/pkg/logger"
	"boilerplate/internal/pkg/middlewares/appcontext"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"time"
)

type Webserver struct {
	GinEngine *gin.Engine
}

func NewWebServer(cfg *config.Config, aCtx *appcontext.Context) *Webserver {
	if *cfg.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.New()
	g.Use(logger.GinLoggerMiddleware(aCtx.Log))
	g.Use(appcontext.InjectContext())

	g.Use(func(c *gin.Context) {
		c.Set("AppContext", aCtx)
		c.Next()
	})

	_ = g.SetTrustedProxies(cfg.TrustedProxies)

	var webserver = new(Webserver)
	webserver.GinEngine = g

	return webserver
}

func RegisterHooks(hp handlers.HandlerParams, lc fx.Lifecycle, cfg *config.Config, ws *Webserver) {
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", *cfg.Port),
		Handler:           ws.GinEngine,
		ReadHeaderTimeout: 1 * time.Second,
		MaxHeaderBytes:    5 << 20,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			createHandlers, err := hp.Hr.CreateHandlers()

			if err != nil {
				return err
			}

			for _, h := range createHandlers {
				h.RegisterRoutes(ws.GinEngine)
			}

			go func() {
				hp.Context.Log.Infof("Starting server on %s", server.Addr)
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					hp.Context.Log.Fatalf("Error starting the server: %s", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			hp.Context.Log.Info("Stopping server...")
			shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			return server.Shutdown(shutdownCtx)
		},
	})
}
