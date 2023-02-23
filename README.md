# labyrinth

To run this program, you will need `go` installed. You can download it from [https://go.dev/dl/](https://go.dev/dl/).

## Usage

You can run the program using either

```sh
make SAVE_ID=<save-id> run
```

or

```sh
go run labyrinth.go -s <save-id>
```

Note: To start a new game, just pick a new save id

Note: You can use the `Noto Color Emoji` font to have an almost square display.

## Test

You can run the tests using either

```sh
make test
```

or

```sh
go test -race ./...
```


## Clean saves

You can clean existing saves using the following command:

```sh
make clean
```
