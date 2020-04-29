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

## Custom Validations
You can implement your password validation:
```go
func maxLength(min int) validator.Validation {
	return func(pwd string) error {
		if len(pwd) < min {
			return fmt.Errorf("must be no more than %v characters", min)
		}
		return nil
	}
}
```
and pass it as argument to validator: `validator.New(maxLength(128))`