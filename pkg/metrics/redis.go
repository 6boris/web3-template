package metrics

import (
	"context"
	"github.com/redis/go-redis/v9"
	"net"
	"time"
)

type RedisMetricsHook struct {
	Name            string
	ServiceID       string
	ServiceInstance string
}

func (h RedisMetricsHook) DialHook(hook redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		conn, err := hook(ctx, network, addr)
		return conn, err
	}
}

func (h RedisMetricsHook) ProcessHook(hook redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		startTime := time.Now()

		err := hook(ctx, cmd)
		status := "success"
		if cmd.Err() != nil {
			status = "fail"
		}
		RedisMetricRequests.WithLabelValues(h.Name, h.ServiceID, h.ServiceInstance, cmd.Name(), status).Inc()
		RedisMetricSeconds.WithLabelValues(h.Name, h.ServiceID, h.ServiceInstance, cmd.Name(), status).Observe(time.Since(startTime).Seconds())
		return err
	}
}

func (h RedisMetricsHook) ProcessPipelineHook(hook redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		err := hook(ctx, cmds)
		for _, cmd := range cmds {
			status := "success"
			if cmd.Err() != nil {
				status = "fail"
			}
			RedisMetricRequests.WithLabelValues(cmd.Name(), h.ServiceID, h.ServiceInstance, status).Inc()
		}
		return err
	}
}
