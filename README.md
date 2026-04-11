# pokedexcli

`pokedexcli` is a Go command-line tool for looking up Pokemon data from [PokeAPI](https://pokeapi.co/).

## Requirements

- Go `1.25.6` or newer
- Internet access

## Run Without Installing

From the repository root:

```bash
go run . help
```

You can replace `help` with any supported command.

## Build The CLI

```bash
go build -o pokedexcli .
```

Then run it with:

```bash
./pokedexcli help
```

## Usage

```bash
pokedexcli <command> [arguments]
```

If no command is provided, the CLI prints the usage text and available commands.

## Commands

### `help`

Show the available commands.

```bash
./pokedexcli help
```

### `explore <location-area>`

List Pokemon that can be encountered in a specific location area.

```bash
./pokedexcli explore viridian-forest-area
```

### `inspect <pokemon> [--json]`

Show a Pokemon's height, weight, stats, and types.

Plain text output:

```bash
./pokedexcli inspect pikachu
```

JSON output:

```bash
./pokedexcli inspect pikachu --json
```

### `weak <pokemon>`

List the damage types the Pokemon is weak against.

```bash
./pokedexcli weak bulbasaur
```

### `super <pokemon>`

List the damage types the Pokemon is super effective against.

```bash
./pokedexcli super charmander
```

### `find <pokemon> [game ...]`

List location areas where a Pokemon can be encountered.

Without game filters:

```bash
./pokedexcli find pikachu
```

With one or more game filters:

```bash
./pokedexcli find pikachu red blue yellow
```

Only locations that match one of the provided game names are returned.

## Notes

- Pokemon and area names should use the API naming format, such as `pikachu` or `viridian-forest-area`.
- Data is fetched from PokeAPI and cached by the application to reduce repeated requests.
