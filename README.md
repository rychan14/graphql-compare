# graphql-compare
Comparing different GraphQL implementations 

Prerequisites:
1. wrk: https://github.com/wg/wrk
2. Golang(1.10 is used): https://golang.org/dl/
3. nodejs(8.6.0 is used): https://nodejs.org/en/download/

Operation

Run express server
1. `cd express-graphql`
2. `npm install`
3. `node index`
4. server is on localhost:4000
Run go server
1. `cd graphql-go`
2. `go run main.go`
3. server is on localhost:4001
Use wrk to benchmark express:
1. `wrk -t12 -c400 -d30s --timeout 10s "http://localhost:4000/graphql?query={hello}"`
Use wrk to benchmark go:
1. `wrk -t12 -c400 -d30s --timeout 10s "http://localhost:4001/graphql?query={hello}"`
