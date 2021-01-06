# Test_LBC FizzBuzz

This app expose a FizzBuzz algorithm, either standard or custom using query paramaters. It also expose a frequency handler returning the most called request paramaters.

## Usage

Build the docker image

```bash
make
```

Run it

```bash
make run
```

## API Endpoint

The root path without any parameter returns a default FizzBuzz implementation.
You can also give the API few parameters to customize the output as this :

```bash
http://localhost:8000/?int1=3&int2=5&limit=15&str1=toto&str2=tata
```

ps: as long as you specify one paramater, all five are required

## Tests

```bash
make test
```
