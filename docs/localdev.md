# Local Development

## Prerequisites

* [Docker Desktop](https://docker.com/desktop)
* An HTTP client (Recommendations:)
  * [RapidAPI](https://paw.cloud) (formerly _Paw_)
  * [Insomnia](https://insomnia.rest)

## Start backend services

The local versions of backend services run as containers. From the root of the repository:

```bash
make build-lambda
cd localdev
docker compose up
```

The very first time you run `docker compose up`, the Docker images will need to build. Subsequent runs will leverage the cached completed image.

When you are done, terminate the containers.

```bash
docker compose down
```

Operating Docker Desktop and Docker Compose is outside the scope of these instructions, but you can read the documentation for yourself.

* <https://docs.docker.com/desktop/>
* <https://docs.docker.com/compose/>

## Lambda server (:9000)

In production, we use [Amazon API Gateway v2](https://docs.aws.amazon.com/apigateway/latest/developerguide/welcome.html) sitting in front of [AWS Lambda](https://aws.amazon.com/lambda/).

AWS has [open-sourced their Lambda runtimes](https://github.com/aws/aws-lambda-base-images/tree/provided.al2023) — namely for [Amazon Linux 2023](https://github.com/northwood-labs/lambda-provided-al2023) — so we use that image along with the [AWS Lambda Runtime Interface Emulator](https://github.com/aws/aws-lambda-runtime-interface-emulator) to create a local AWS Lambda environment that is _accurate_.

However, passing payloads directly to AWS Lambda is different from going through API Gateway first.

1. By running `make build-lambda`, then starting/restarting the Docker Compose environment, the application will run as a Lambda function inside the Dockerized Lambda environment.

1. Next, run `make build` to compile and install `devsec-tools`, then run `devsec-tools serve` to run a rudimentary local web server which will accept your API calls, then convert them into API Gateway format before passing them along to the AWS Lambda environment for execution. Multiple hops may be a little slower, but it provides a more accurate development experience.

This server is stateless and ephemeral, so you can start/stop it and nothing is persisted across boots.

## Valkey server (:6379)

[Valkey] is an open-source fork of [Redis](https://redis.io/docs/latest/get-started/), which [ceased to be open-source](https://redis.io/legal/licenses/) in March 2024. AWS provides [ElastiCache Serverless](https://aws.amazon.com/elasticache/what-is-valkey/) with Valkey support, which [devsec.tools](https://devsec.tools) uses for caching results.

When the `devsec-tools` binary is running as a Lambda function it will connect to `localhost:6379` by default for a caching server. When running in production, you can use the `DST_CACHE_HOSTS` environment variable to configure production Valkey hosts.

### Persistent data

This server uses a persistent volume (`localdev_vkdata`), so you can stop it and data will be restored on next restart. If you want to delete the volume data:

1. Run `docker compose down` to terminate the local servers.
1. Run `docker volume rm localdev_vkdata` to delete the persisted data.

The next time you run `docker compose up`, the volume will be recreated. Valkey is only used for caching and cache expiration, so it can be deleted without worrying about loss of important data.

### Environment variable

When **running as a Lambda function**, the value of `DST_CACHE_HOSTS` should be one or more endpoints (delimited by `;`) to use for [Valkey] caching.

```bash
# Example: local dev
DST_CACHE_HOSTS="localhost:6379"

# Example: production
DST_CACHE_HOSTS="server1.host.com:6379;server2.host.com:6379;server3.host.com:6379"
```

## Local CLI server (:8080)

For local testing, the CLI exposes a very simple HTTP/1.1 server at <http://localhost:8080>.

> [!CAUTION]
> **DO NOT** use this web server in **any** production environment.

```bash
devsec-tools serve
```

Re-read _Lambda server_ (above) to get a better understanding of how this works, but think of it like a reverse-proxy to your Lambda environment.

## Endpoints

When launching the local web server, it will tell you which HTTP methods and endpoints are available. It exposes both `GET` and `POST`, as appropriate.

### GET

For `GET` endpoints, any parameters are passed as URL-encoded query string parameters. Using the `/http` endpoint as an example:

```http
GET /http?url=https%3A%2F%2Fapple.com HTTP/1.1
Host: localhost:8080
```

### POST

For `POST` endpoints, parameters are passed as a JSON-encoded request body. Using the `/http` endpoint as an example:

```http
POST /http HTTP/1.1
Host: localhost:8080
Content-Type: application/json; charset=utf-8

{"url":"https://apple.com"}
```

[Valkey]: https://valkey.io