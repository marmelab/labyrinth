# labyrinth

To run this program, you will need `go` installed. You can download it from [https://go.dev/dl/](https://go.dev/dl/).

## Usage

You can run the program using either

```sh
# The following options are available :
# - SAVE_ID=<save-id>       To set the save ID name.
# - BOARD_SIZE=<size>       To set the board size.
# - PLAYER_COUNT=<count>    To set the numbe rof players.
make [options] run
```

or

```sh
# The following options are available :
# - --save=<save-id>        To set the save ID name.
# - --size=<size>           To set the board size.
# - --players=<count>       To set the numbe rof players.
go run ./cli [options]
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
