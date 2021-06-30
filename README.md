# GroupLive

An open source Public Community Management System implemented with Microservices in Go & Java. 

Core functions include 
- [Auth Services](https://github.com/yijia-cc/grouplive/tree/master/auth) for Sign In/Sign Up
- [Dashboard Service](https://github.com/yijia-cc/grouplive/tree/master/dashboard) for news aggregation 
- [Discussion Service](https://github.com/yijia-cc/grouplive/tree/master/discussion) for Content Post and Share 
- [Calendar Service](https://github.com/yijia-cc/grouplive/tree/master/calendar) for Reservation Creation and Modification 


## Web Preview

https://user-images.githubusercontent.com/76569613/122740912-4bb42c80-d239-11eb-946f-d44be87bcef6.mov




## Getting Started
### Accessing the Source Code
`git clone https://github.com/yijia-cc/grouplive.git`

### Prerequisites
#### Development
[Go](https://golang.org/doc/install) v1.15.6

[JDK](https://www.oracle.com/java/technologies/javase-jdk11-downloads.html) v11.0.2

[Maven](https://maven.apache.org/install.html) v3.8.1

[Java](https://www.oracle.com/java/technologies/javase/jdk14-archive-downloads.html) v14

[Node.js](https://nodejs.org/en/download/) v14.16.1


#### Infrastructure
[Docker](https://docs.docker.com/docker-for-mac/install/) v20.10.6

[MySQL](https://dev.mysql.com/downloads/mysql/) v8.0.25

#### API Framework
[gRPC](https://github.com/yijia-cc/grouplive/blob/master/proto/scripts/grpc-install.sh) for Go

[gRPC](https://github.com/yijia-cc/grouplive/tree/master/proto/java) for Java

[graphQL](https://github.com/graphql-go/graphql) for Go

## Service Architecture
GroupLive adopts [Microservices Architecture](https://microservices.io/patterns/microservices.html) to organize dependent services and to enable independent deployment of each service.





