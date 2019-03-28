
## Configuration

Add a `.env` file to the root directory

```bash
HAWK_SERVER_ADDR=:PORT
HAWK_DB_USER=
HAWK_DB_PASSWORD=
HAWK_DB_NAME=
```

Generate keys

```bash
$ make keys
```

## Running

```bash
$ make serve
```

## Commands

```bash
make serve
make keys
make migrate
make drop_all
```

