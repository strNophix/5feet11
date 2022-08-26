# 5feet11

A URL shortener that generates links even shorter than 5'11.

## Usage
```bash
# Shorten a URL
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"redirectUrl":"https://news.ycombinator.com"}' \
  http://localhost:5111/redirect

# Expand the URL
curl -iL http://localhost:5111/{id}
```
