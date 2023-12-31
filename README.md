# 5feet11

A scalable URL shortener.

## Prepare
- ScyllaDB
- Prometheus*

*: Optional

## Run
Update the config at `etc/fivefeeteleven-api.yaml`.
```sh
# Docker
docker run -it $(docker build -q .)

# Compile
make build && ./5feet11
```


## Usage
```bash
# Shorten a URL
curl -X POST -H "Content-Type: application/json" \
  -d '{"longUrl":"https://news.ycombinator.com"}' \
  http://localhost:5111/api/v1/links

# Shorten a URL that expires after N seconds
curl -X POST -H "Content-Type: application/json" \
  -d '{"longUrl":"https://news.ycombinator.com", "expiresAfter": 30}' \
  http://localhost:5111/api/v1/links

# Expand the URL
curl -iL http://localhost:5111/{id}

# Get information of a specific URL
curl -iL http://localhost:5111/api/v1/links/{id}
```
