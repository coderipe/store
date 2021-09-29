package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/srikrsna/protoc-gen-store/pkg/lang/dart"
	"github.com/srikrsna/protoc-gen-store/pkg/module"
)

func main() {
	pgs.Init().RegisterModule(module.New()).RegisterPostProcessor(dart.DartFmt()).Render()
}
