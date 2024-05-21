## A simple Key-Value API In-Memory Store written in Go for learning purposes

To test the api, use `./test.http` with [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) in Visual Studio Code.

### Related blog post

https://developerschallenges.com/2020/10/25/run-go-application-in-azure-container-instance/

### Docker

``` sh
docker build -t go-key-value-api .
docker run --rm -p 9000:9000 go-key-value-api
docker rmi go-key-value-api
```