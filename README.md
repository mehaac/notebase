# notebase

Notebase is my yet another iteration on the "Everything App". This time I am going for my markdown notes.

There are several key points, why I develop this project:

- I need a mobile-friendly web interface for my notes (Obsidian Web is not good enough, and Obsidian App is too heavy)
- I want a well-established query language, like SQL to work with my notes (Dataview and Datacore are still great, but see previous point)
- I want to replace multiple tools I use for everyday life and I always thought they fit nicely into my markdown notes, just need a better interface (like todo lists, finance goals, tv show tracker, habit tracker, etc.)

I document some of my decisions regarding this project in [docs/adrs](./docs/adrs).

## Usage

Notebase is intented to run as a Docker container. You can try it out together with the example data like this:

```bash
docker build -t notebase:latest .
docker run -it -p 8080:8080 notesbase
```

You can further customize the container by setting environment variables:

```bash
docker run -it \
  -p 8080:8080 \
  -e SUPERUSER_EMAIL=admin@example.com \
  -e SUPERUSER_PASSWORD=supersecret \
  -e NOTES_ROOT=/tmp/notes \
  -v $PWD/data:/tmp/example/notes \
  notesbase \
  serve --http=0.0.0.0:8080 --dev
```

- make sure to add `.notebase.yml` to your root. You can check out my current config in [/examples/biozz_notebase_config.yml](./examples/biozz_notebase_config.yml)
- `SUPERUSER_EMAIL` and `SUPERUSER_PASSWORD` are optional, but if you don't set them, you will be prompted to create a superuser account
- `--dev` is also optional, you can set it if want to see query and access logs

After creating it, go to http://localhost:8080.

## Development

You are going to need Go (whichever version is provided in `go.mod`) and Node/pnpm.

There is an `example/` directory with a bunch of notes and a config file.

```
NOTES_ROOT=./example/notes go run . serve
```

If you feel ambitious after thoroughly inspecting the source code, you can do `NOTES_ROOT=/abs/path/to/your/vault go run . serve` and try it on the real data. :)

And if you do, don't forget to add `.notebase.yml` to your root. You can check out my current config in [/examples/biozz_notebase_config.yml](./examples/biozz_notebase_config.yml).

`serve` command also starts a background `sync` job. `serve` will created `pb_data` directory, which should not be checked into the VCS and is meant to be disposable. That means taht if there is any error or migration conflict you delete `pb_data` and run the server again.

Follow PocketBase instructions to create an account.

Then you can run frontend separately (assuming you installed dependencies):

```
pnpm dev
```

## Contributing

I might be interested in some contributions, but because it is an experimental project, I would prefer discussions and maybe code-reviews of the non-optimal or non-conventional parts of the codebase.
