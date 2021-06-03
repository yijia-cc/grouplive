# Auth Service

## Prerequisite

- Docker v20.10.6

## Getting Started

Build Docker image

```bash
docker build -t auth .
```

Create and update `.env` file

```
cp .env.dist .env
```

Start local server

```bash
docker run --env-file .env -p 8080:8080 --network="host" auth
```

You can visit your service at [http://127.0.0.1:8080](http://127.0.0.1:8080).

## Testing

## Deployment

### Staging

Merge a PR into `master`  branch will automatically deploy the service to `staging` environment.

You can try out the service at [http://auth.api.staging.allgame.fun](http://auth.api.staging.allgame.fun) after deployment.

###  Production

Merge a PR from `master` branch to `production` branch will automatically deploy the service to `production` environment.

You can try out the service at [http://auth.api.allgame.fun](http://auth.api.allgame.fun) after deployment.

## Authors

## License
