# Goecho

This is intended as an example for Computer Networks 2 class at Poznań University of Technology. Mostly usable by my students. Goecho lousy implements Echo protocol over TCP described in [RFC862](https://tools.ietf.org/html/rfc862).

For a daytime, check [here](https://github.com/tdi/godaytime/)

## Install

```
go get -u github.com/tdi/goecho
```

## Usage 
```
goecho [-h] [-H HOSTNAME] [-p PORT]
```

By default godaytime listens on port localhost:3333. 

## Example run 

```
➜  nc localhost 2222
hello there
hello there
```

## Author and licence

Dariusz Dwornikowski, MIT 