# 版权 @2019 凹语言 作者。保留所有权利。

# wasmer
# https://github.com/wasmerio/wasmer/releases

hello:
	go run main.go run _examples/hello

hi:
	go run main.go run _examples/hi

prime:
	go run main.go run _examples/prime

wasm:
	go run main.go wasm hello.wa
	wasmer a.out.wat

clean:
