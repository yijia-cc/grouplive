package obs

import "context"

type LogLevel int

const (
	Fatal LogLevel = iota
	Error
	Warn
	Info
	Debug
)

type Logger struct {
	level LogLevel
	layer int
}

// Fatal records errors that lead to crash of the service or the whole system.
func (l Logger) Fatal(ctx context.Context, err error) {

}

// Error records errors which is fatal to the specific operation.
func (l Logger) Error(ctx context.Context, err error) {

}

// Warn records issues that are automatically recovered
func (l Logger) Warn(ctx context.Context, message string) {

}

// Info records info that is generally helpful to developers and others
func (l Logger) Info(ctx context.Context, message string) {

}

// Debug records info which help people diagnose issues.
func (l Logger) Debug(ctx context.Context, message string) {

}

func (l Logger) NextLayer() Logger {
	return Logger{layer: l.layer + 1}
}

func NewLogger(level LogLevel) Logger {
	return Logger{level: level}
}
