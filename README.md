# nook-server-app
Nook Server App is a basic Go application that creates a HTTP server and content from the [nook](https://github.com/lindsaygelle/nook) package.

Nook Server App listens on `:8080`.

## Git
Getting the code.

```sh
git clone https://github.com/lindsaygelle/nook-server-app.git
```

## Docker
This code has been built for Docker.

### Building 
Building the Container.

```sh
docker build . -t nook-server-app
```

### Running
Developing and running Go from within the Container.

```sh
docker run -it --rm --name nook-server-app nook-server-app
```

### Compose
Starting the application from within a Docker Container.

```sh
docker-compose up --build -d
```

## Contents
Below are some snippets from the different routes.

### All
Content served for the default route. This has been truncated for brevity. 

```json

  {
    "animal": "Alligator",
    "character": "Alfonso",
    "link": "/alligator/alfonso",
    "special": false
  },
  {
    "animal": "Alligator",
    "character": "Del",
    "link": "/alligator/del",
    "special": false
  },
  {
    "animal": "Alligator",
    "character": "Sly",
    "link": "/alligator/sly",
    "special": false
  },
  {
    "animal": "Alligator",
    "character": "Pironkon",
    "link": "/alligator/pironkon",
    "special": false
  },
  {
    "animal": "Alligator",
    "character": "Alli",
    "link": "/alligator/alli",
    "special": false
  },
  {
    "animal": "Alligator",
    "character": "Boots",
    "link": "/alligator/boots",
    "special": false
  },
  {
    "animal": "Alligator",
    "character": "Drago",
    "link": "/alligator/drago",
    "special": false
  },
  {
    "animal": "Alligator",
    "character": "Gayle",
    "link": "/alligator/gayle",
    "special": false
  },
  {
    "animal": "Alligator",
    "character": "Liz",
    "link": "/alligator/liz",
    "special": false
  }
]

```

### Resident
Content served for a `nook.Resident`.

```json
{
  "animal": "Dog",
  "character": "Resident",
  "birthday": 20,
  "birthday_month": 12,
  "gender": "Isabelle",
  "language": "en-US",
  "name": "Female",
  "ok": true,
  "special": true
}
```

### Villager
Content served for a `nook.Villager`.

```json
{
  "animal": "Bear",
  "character": "Villager",
  "birthday": 31,
  "birthday_month": 3,
  "gender": "Klaus",
  "language": "en-US",
  "name": "Male",
  "ok": true,
  "special": false,
  "phrase": "strudel"
}
```