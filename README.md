# Web3 Template

This is an open source Web3 Go solution from someone who has been working
in the Web2 space for many years. This project uses microservices, containers,
message queue related architecture, and can very comfortably cope with the logic
related to the server side of Web3 business nowadays.

**Framework technology is stable and powerful：**

The use of microservices framework, this framework supports the stable operation of 
[BILIBILI](https://www.bilibili.com/) more than thousands of  service framework, the workplace 100 million DAU.

**Extensive Web3 server experience：**

The project initiator has multiple startups in the Web3 field and has used 
this technology to support multiple project launches and runs.

**Sufficient cases to refer：**

The case will provide a very large number of Web3 functional examples, startup projects 
can save a lot of time using this.


Here is the architecture diagram and examples:

<a href="https://web3-studio.leek.dev/d/demo/web3-studio" target="_blank">
  <img src="https://s.gin.sh/develop/web3/web3_architecture.png" alt="Logo">
</a>


<a href="https://web3-studio.leek.dev/d/demo/web3-studio" target="_blank">
  <img src="https://s.gin.sh/develop/web3/web3-template-demo.png" alt="Logo">
</a>


## Quick Start

- clone repo
```shell
git clone git@github.com:6boris/web3-template.git
cd web3-template
```

- Install [Docker](https://code.visualstudio.com)

- Local run 
```shell
make clean && make start
```
- Open Grafana http://localhost:60005
  - Login with admin/admin

## Development Manual

- Learn about [Go](https://go-kratos.dev) Basic Grammar
- Learn about [Kratos](https://go-kratos.dev) technical documentation
- Prepare a code editor [VS Code](https://code.visualstudio.com) OR [GoLand](https://www.jetbrains.com/go/download)
- Installing forced dependencies
    - [Docker](https://code.visualstudio.com)
    - [Go](https://go.dev)
    - [MySQL](https://www.mysql.com)
    - [Redis](https://redis.io)
    - [RabbitMQ](https://www.rabbitmq.com/)

## Development Information

Some basic application services and basic service ports

|       Service       | HTTP PORT | GRPC PORT |
|:-------------------:|:---------:|:---------:|
| app.interface.web3  |   61011   |   61012   |
|  app.service.demo   |   62011   |   62012   |
|  app.service.data   |   62021   |   62022   |
|  app.job.datawatch  |   63011   |   63012   |
|   app.job.databus   |   63021   |   63022   |
|  app.job.analysis   |   63031   |   63032   |

|         App          |          PORT          |
|:--------------------:|:----------------------:|
|   base.web3.mysql    |         60001          |
|   base.web3.redis    |         60002          |
|  base.web3.rabbitmq  | http://127.0.0.1:60003 |
| base.web3.prometheus | http://127.0.0.1:60004 | 
|  base.web3.grafana   | http://127.0.0.1:60005 | 

## External Dependency

- [Kratos](https://go-kratos.dev) A Go Framework for microservices
- [Web3 Go](https://github.com/6boris/web3-go) A Web3 Go Client

## External Dependency

For other technical questions, you can contact email ....