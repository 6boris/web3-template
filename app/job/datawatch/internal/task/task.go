package task

import (
	"context"
	"google.golang.org/appengine/log"
)

type ConsumerData struct {
	ConsumerName string `json:"consumer_name"`
	RawData      string `json:"raw_data"`
}

type ConsumerHook func(ctx context.Context, data ConsumerData) error
type WatchHook func(ctx context.Context, ch chan *ConsumerData) error

type Consumer interface {
	Name() string
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Reload(ctx context.Context) error
}

var GlobalConsume = []Consumer{
	NewWatchEthereumNewBlockHeader(),
	NewWatchPolygonNewBlockHeader(),
	NewWatchBScNewBlockHeader(),
	NewWatchAvalancheNewBlockHeader(),
	NewWatchFantomNewBlockHeader(),
	NewWatchAptosNewBlockHeader(),
}

func (s *Server) _startWatchTask(ctx context.Context) error {
	for _, v := range GlobalConsume {
		tmpV := v
		go func() {
			loopErr := tmpV.Start(ctx)
			if loopErr != nil {
				log.Errorf(ctx, "[Task] _startWatchTask Name:%+v err:%+v", tmpV.Name(), loopErr)
			}
		}()
	}
	return nil
}
func (s *Server) _stopWatchTask(ctx context.Context) error {
	for _, v := range GlobalConsume {
		loopErr := v.Stop(ctx)
		if loopErr != nil {
			log.Errorf(ctx, "[Task] _stopWatchTask Name:%+v err:%+v", v.Name(), loopErr)
		}
	}
	return nil
}
