# GoFck

Simple [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) interpreter written
in Go. If you want to understand how Brainfuck works I recommend reading
[Basics of Brainfuck](https://gist.github.com/roachhd/dce54bec8ba55fb17d3a).

## Example

```sh
# Build GoFck
go build -o gof gofck.go

# Interpret and execute Brainfuck code
./gof <brainfuck_source.bf>

# Build an executable
./gof --out bf-app <brainfuck_source.bf>
./bf-app
```
