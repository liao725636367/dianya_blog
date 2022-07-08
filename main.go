package main

import (
	_ "blog/boot"
	"blog/internal/cmd"
	_ "blog/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
