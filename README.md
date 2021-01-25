## Dependencies

### chromapring

Chromaprint needs to be installed and its C bindings available

```shell
$ brew install chromaprint
```

### ffprobe

`ffprobe` needs to be installed (part of ffmpeg) and available on PATH

```shell
$ brew install ffmpeg
```

### fftw3

```shell
$ wget http://fftw.org/fftw-3.3.9.tar.gz
$ tar xzvf fftw-3.3.9.tar.gz
$ cd fftw-3.3.9.tar.gz
$ ./configure
$ make && make install
```

## Runnning

```shell
$ go mod download
$ go run .
```
