# GroupLive

An open source Public Community Management System based on Service-oriented architecture (SOA) in Go & Java. 

## Web Preview
https://user-images.githubusercontent.com/76569613/122740912-4bb42c80-d239-11eb-946f-d44be87bcef6.mov

## Getting Started
### Accessing the Source Code
```
git clone https://github.com/yijia-cc/grouplive.git
```

### Prerequisites
- [Docker](https://docs.docker.com/docker-for-mac/install/) v20.10.6

### Database
```
docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password --network="host" -d mysql:8.0.25
```

### Launch Application
Please start services at different ports and update `.env` files with the correct ports so that services can communicate with each other.

- [Web UI](https://github.com/yijia-cc/grouplive/tree/master/web) 
- [Auth Service](https://github.com/yijia-cc/grouplive/tree/master/auth)
- [Calendar Service](https://github.com/yijia-cc/grouplive/tree/master/calendar)
- [Discussion Service](https://github.com/yijia-cc/grouplive/tree/master/discussion)
- Real-time Chat Service (In Progress)
- Dashboard Service (In Progress)
- Payment Service (In Progress)

Assuming `web` is exposed at 8080, visit `http://localhost:8080` to view the website.

## Tech Stack
- gRPC
- GraphQL
- Go
- Java Springboot
- MySQL
- React
- Docker
- Nginx

## Authors
- **yijia-cc**: Infrastructure(Docker, Continuous Delivery, Database), Auth Service(current), Calendar Service
- **wl328Weiminli**: Frontend, Web UI
- **isabellakqq**: Dashboard Service, Auth Service(first draft)
- **jackiewang5566**: Discussion Service
- **MEtoCS**: Chat Service
- **yuranranran**: Chat Service
- **wikiwu24**: Chat Service
- **Zanzan666**: Payment Service

## License
This project is maintained under MIT license.




