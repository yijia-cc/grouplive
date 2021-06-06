# Calendar Service

## Prerequisite

- Docker v20.10.6

## Getting Started

Build Docker image

```bash
docker build -t web .
```

Create and update `.env` file

```
cp .env.dist .env
```

Start local server

```bash
docker run -p 8080:80 --network="host" web
```

You can visit your service at [http://127.0.0.1:8080](http://127.0.0.1:8080).

## Testing

## Deployment

### Staging

Merge a PR into `master`  branch will automatically deploy the service to `staging` environment.

You can try out the service at [http://payment.api.staging.allgame.fun](http://payment.api.staging.allgame.fun) after deployment.

###  Production

Merge a PR from `master` branch to `production` branch will automatically deploy the service to `production` environment.

You can try out the service at [http://payment.api.allgame.fun](http://payment.api.allgame.fun) after deployment.

## Authors

## License
