# urlenc

Simple URL encoder and decoder for the commandline

## Usage
```sh
> urlenc "https://github.com/floj/urlenc"
https%3A%2F%2Fgithub.com%2Ffloj%2Furlenc

> urlenc -d "https%3A%2F%2Fgithub.com%2Ffloj%2Furlenc"
https://github.com/floj/urlenc
```

## Build
```sh
CGO_ENABLED=0 go build -o urlenc .
```

## License
MIT