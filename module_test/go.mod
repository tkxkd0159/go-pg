module module_test

go 1.16

require (
	rsc.io/quote/v3 v3.1.0
	rsc.io/sampler v1.99.99 // indirect

	github.com/tkxkd0159/go-chain/module_test/hello v1.0.0
	github.com/tkxkd0159/go-chain/module_test/hello/nums v1.0.0
)


replace (
	github.com/tkxkd0159/go-chain/module_test/hello => "./hello"
	github.com/tkxkd0159/go-chain/module_test/hello/nums => "./hello/nums"
)