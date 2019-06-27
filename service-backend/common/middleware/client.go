package middleware

import (
	"context"
	"github.com/google/uuid"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	CorrelationId = "CorrelationId"
)

type ClientWrapper struct {
	client.Client
}

func NewClientWrapper(c client.Client) ClientWrapper {
	if c == nil {
		c = client.NewClient()
	}
	return ClientWrapper{
		Client: c,
	}
}

func (c ClientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	id := ctx.Value(CorrelationId)
	logger := GetLogger(ctx)

	if id, ok := id.(string); ok {
		ctx = metadata.NewContext(ctx, map[string]string{
			CorrelationId: id,
		})
	} else {
		ctx = metadata.NewContext(ctx, map[string]string{
			CorrelationId: uuid.New().String(),
		})
	}
	var err error
	defer func(startTime time.Time) {
		if err != nil {
			logger.WithFields(logrus.Fields{
				"TimeTaken": time.Since(startTime),
				"Error":     err.Error(),
			}).Warningln(req.Method() + " failed.")
		} else {
			logger.WithField("TimeTaken", time.Since(startTime).String()).Infoln(req.Method() + " succeeded.")
		}

	}(time.Now())

	err = c.Client.Call(ctx, req, rsp, opts...)
	return err
}
