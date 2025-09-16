package log

import "context"

type Logger interface {
	Debug(ctx context.Context, arg ...any)
	Debugf(ctx context.Context, templete string, arg ...any)
	Info(ctx context.Context, arg ...any)
	Infof(ctx context.Context, templete string, arg ...any)
	Warn(ctx context.Context, arg ...any)
	Warnf(ctx context.Context, templete string, arg ...any)
	Error(ctx context.Context, arg ...any)
	Errorf(ctx context.Context, templete string, arg ...any)
	Fatal(ctx context.Context, arg ...any)
	Fatalf(ctx context.Context, templete string, arg ...any)
}
