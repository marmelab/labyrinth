# labyrinth

To run this program, you will need `go` installed. You can download it from [https://go.dev/dl/](https://go.dev/dl/).

## CLI

You can run the CLI program using either

```sh
# The following options are available :
# - SAVE_ID=<save-id>       To set the save ID name.
# - BOARD_SIZE=<size>       To set the board size.
# - PLAYER_COUNT=<count>    To set the numbe rof players.
make [options] cli-run
```

## Web application

To run the web application, you will need to create the environment files, and then set the corresponding environment variables:

```sh
cp webapp/.env.dist webapp/.env
cp docker-compose.env.dist docker-compose.env
```

Once the env file has been set,you can either run the app in development mode using

```sh
make develop
```

Or in production mode using
```sh
make run
```

This will start the following applications:
```
webapp              127.0.0.1:9000 (dev) / 0.0.0.0:80 (production)
api-domain          127.0.0.1:9001
```

## Test

You can run the unit tests for both domain and webapp using the following command:

```sh
make test
```

## Clean saves

You can clean existing saves using the following command:

```sh
make cli-clean
```
