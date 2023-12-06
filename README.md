![go workflow](https://github.com/garbage-collector/temporal-demo/actions/workflows/go.yml/badge.svg)

# temporal-demo

## Start Temporal

Install Temporal CLI:

```bash
curl -sSf https://temporal.download/cli.sh | sh
```

Start Temporal dev server:

```bash
temporal server start-dev
```

- server port: `7233`
- UI port: `8233`

## Start the API

### Build the image

```bash
pack build myself --buildpack paketo-buildpacks/go --builder paketobuildpacks/builder-jammy-tiny
```

### Run the docker image

```bash
docker-compose up -d
```

- port: `8080`
