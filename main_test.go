package main

import (
	"fmt"
	"strconv"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewExample(t *testing.T) {
	logger := zap.NewExample()
	defer func() {
		_ = logger.Sync()
	}()

	logger.Debug("dda", zap.String("abc", "123"))
}

func TestAppendInt(t *testing.T) {
	s := []byte{'@'}
	s = strconv.AppendInt(s, 10, 10)
	fmt.Printf("%s\n", s)
}

func TestNewProduction(t *testing.T) {
	zap.NewProductionConfig()
	zap.NewProductionEncoderConfig()
	logger, err := zap.NewProduction()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer func() {
		_ = logger.Sync()
	}()

	logger.Error("ddb", zap.String("def", "456"))
}

func TestNewDep(t *testing.T) {
	zap.NewDevelopmentConfig()
	zap.NewDevelopmentEncoderConfig()
	zap.NewDevelopment()
}

func TestCore(t *testing.T) {
	_ = zapcore.NewCore(nil, nil, nil)
}
