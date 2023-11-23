# temporal-demo

## Start Temporal

```bash
docker-compose up -d temporal-server temporal-web
```

- server port: `7233`
- web port: `10000`

## Start the API

### Build the image

```bash
pack build myself --buildpack paketo-buildpacks/go --builder paketobuildpacks/builder-jammy-tiny
```

### Run the docker image

```bash
docker-compose up -d myself
```

- port: `8080`
