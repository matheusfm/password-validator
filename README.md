# password-validator

## Build and Run with Docker

#### Build
```shell script
docker -t mfmoraes/password-validator .
```
#### Run
```shell script
docker run --name password-validator -p 8000:8000 mfmoraes/password-validator
```

## Test
```shell script
curl 127.0.0.1:8000/validation -d '{ "password": "foobar" }'
```