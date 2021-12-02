package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

// stack represents a stack of program counters.
type stack []uintptr

type withStack struct {
	error
	*stack
}

type withMessage struct {
	cause error
	msg   string
}

type TaskMsg struct {
	ActionType int   `json:"actionType"`
	ActionId   int64 `json:"actionId"`
	Mid        int64 `json:"mid"`
	Ts         int64 `json:"date"`
}

func main() {
	var taskMsg TaskMsg
	msg := TaskMsg{
		1,
		2,
		122,
		time.Now().Unix(),
	}
	msgStr, _ := json.Marshal(msg)
	err := json.Unmarshal(msgStr, &taskMsg)
	fmt.Println(err)
	fmt.Println(taskMsg)
}

func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }
func (w *withMessage) Cause() error  { return w.cause }

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	err = &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
	return &withStack{
		err,
		callers(),
	}
}

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   message,
	}
}
