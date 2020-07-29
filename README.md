# urlenc

Simple URL encoder and decoder for the commandline

## Usage
```sh
# encode
> urlenc "https://github.com/floj/urlenc"
https%3A%2F%2Fgithub.com%2Ffloj%2Furlenc

# decode
> urlenc -d "https%3A%2F%2Fgithub.com%2Ffloj%2Furlenc"
https://github.com/floj/urlenc

# encode from stdin
> echo -n "https://github.com/floj/urlenc" | urlenc
https%3A%2F%2Fgithub.com%2Ffloj%2Furlenc

# decode from stdin
> echo -n "https%3A%2F%2Fgithub.com%2Ffloj%2Furlenc" | urlenc -d
https%3A%2F%2Fgithub.com%2Ffloj%2Furlenc
```

## Build
```sh
CGO_ENABLED=0 go build -o urlenc .
```

## License
MIT