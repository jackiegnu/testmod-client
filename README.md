# testmod-client
testmode-client is used to test testmod which define a module

# using local cached package 
## export GOPROXY=http://localhost:9908
 * python3 -m http.server 9908
 * go run proxy.go -http :9908 cache
## export GOPROXY=file:///path/to/cache
