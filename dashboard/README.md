# Calendar Service

## Prerequisite

## Getting Started

Build Docker image

```bash
docker build -t dashboard .
```

Create and update `.env` file

```
cp .env.dist .env
```

Start local server

```bash
docker run --env-file .env -p 8080:8080 --network="host" dashboard
```

## Testing

## Deployment

### Staging

Merge a PR into `master`  branch will automatically deploy the service to `staging` environment.

You can try out the service at [http://dashboard.api.staging.allgame.fun](http://dashboard.api.staging.allgame.fun) after deployment.

###  Production

Merge a PR from `master` branch to `production` branch will automatically deploy the service to `production` environment.

You can try out the service at [http://dashboard.api.allgame.fun](http://dashboard.api.allgame.fun) after deployment.

## Authors

## License
