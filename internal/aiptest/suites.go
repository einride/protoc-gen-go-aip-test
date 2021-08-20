package aiptest

import (
	"github.com/einride/protoc-gen-go-aip-test/internal/aiptest/batchget"
	"github.com/einride/protoc-gen-go-aip-test/internal/aiptest/create"
	"github.com/einride/protoc-gen-go-aip-test/internal/aiptest/get"
	"github.com/einride/protoc-gen-go-aip-test/internal/aiptest/list"
	"github.com/einride/protoc-gen-go-aip-test/internal/aiptest/update"
	"github.com/einride/protoc-gen-go-aip-test/internal/suite"
)

var Suites = []suite.Suite{
	create.Suite,
	get.Suite,
	batchget.Suite,
	update.Suite,
	list.Suite,
}
