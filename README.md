clippy.io Server: Golang Edition [![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)
=============

# Setup

Requires:

- Redis

| ENV             | description                                                                                                                                                                                         |
|-----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `PORT`          | Port the web server runs at. Defaults to `9001`                                                                                                                                                     |
| `REDISTOGO_URL` | Heroku: The connection information for the Redis Service provisioned by Redis To Go is stored as a URL in the `REDISTOGO_URL` config var. If not provided it will default to the default to `:6379` |

# Building and Running

```bash
godep go install
export PORT=9001
clippy-api-go
```
