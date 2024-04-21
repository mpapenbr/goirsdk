package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/mpapenbr/goirsdk/irsdk"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	for {
		if irsdk.CheckIfSimIsRunning() {
			break
		}
		time.Sleep(time.Second * 2)
	}

	api := irsdk.NewIrsdk()
	api.WaitForValidData()
	headers := api.GetVarHeaders()
	fmt.Println("|Name|Desc|Unit|Type|")
	fmt.Println("|-|-|-|-|")
	for _, h := range headers {
		fmt.Printf("|%s|%s|%s|%s\n",
			extractCString(h.Name[:]),
			extractCString(h.Desc[:]), extractCString(h.Unit[:]),
			resolveType(h.Type, int(h.Count)))
	}

	return nil
}

func resolveType(t irsdk.VarType, count int) string {
	pureType := func(t irsdk.VarType) string {
		switch t {
		case irsdk.IrsdkTypeChar:
			return "char"
		case irsdk.IrsdkTypeBool:
			return "bool"
		case irsdk.IrsdkTypeInt:
			return "int"
		case irsdk.IrsdkTypeBitField:
			return "bitfield"
		case irsdk.IrsdkTypeFloat:
			return "float"
		case irsdk.IrsdkTypeDouble:
			return "double"
		case irsdk.IrsdkTypeETCount:
			return "etcount"
		default:
			return "unknown"
		}
	}
	if count > 1 {
		return fmt.Sprintf("%s[%d]", pureType(t), count)
	} else {
		return pureType(t)
	}
}

func extractCString(buf []byte) string {
	for i, v := range buf {
		if v == 0 {
			return string((buf)[:i])
		}
	}
	return ""
}
