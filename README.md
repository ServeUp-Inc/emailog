# Emailog
Emailog is a simple email logging server.

## Usage
The server exposes a single `PUT` interface at `/v1` that accepts two string parameters with the key `email` and `msg`. An example request made using the `curl` command is as follows.
```bash
curl -X PUT -d 'email=test1@test.com&msg=asdf' localhost:4000/v1
```
Furthermore, the server application expects the following environment variables to be set as well; note that the variables prefixed with `DB` is used to connect to a MySQL database.
```bash
export DB_NAME="dbname"
export DB_USER="dbuser"
export DB_PASS="supersecretpass!"
export DB_HOST="127.0.0.1"
export DB_PORT=3306

export SERVER_HOST="127.0.0.1"
export SERVER_PORT=4000
```
## Development
Development is facilicated by [Nix](https://nixos.org/). Entering the development environment can be done using the following command. All the required tools for development will be available within the Nix shell. Note that, this repository makes use of [Nix Flakes](https://nixos.wiki/wiki/Flakes) and as such will require that Flakes are enabled.

```bash
nix develop
```
Inside the nix shell, a locked version of the following applications can be found.
- [Curl](https://curl.se/docs/manpage.html)
- [Git](https://git-scm.com/)
- [Grip](https://github.com/joeyespo/grip)
- [Podman](https://podman.io/)
- [Go](https://go.dev/)
- [GnuPG](https://www.gnupg.org/)

After entering the Nix shell, the environment is ready to run and test the go application. Note that a MySQL container is also started with default environment variables set as well.

Within the shell, the developer is able to interact with the container using `podman` commands; for example the following command can be used to enter the container's shell to interact with the database directly.
```bash
podman exec -it sqldb mysql -utestuser -ptestpass
```
## Testing
The following command may be used to recursively run tests.
```bash
go test ./...
```

## Production
The go application can be build using a number of ways. For example, it is possible to build the application using Nix using the following command.
```bash
nix build
```
However it might be required to build the application directly using the native go tools; the following commands may be used to accomplish this. Note that the second command may be used to cross-compile the application.
```bash
go build main.go
env GOOS=linux GOARCH=386 go build main.go
```
In the production environment, since the application will not be run within the Nix development shell, all the required environment variables will not be available. As such, a file named `prodvars.asc` is included in the repository with the required envrionment variables.
The following commands are used to encrypt and decrypt the file respectively. Note that output file can be sourced to easily load the environment variables.
```
gpg --armor --symmetric --cipher-algo AES256 -o prodvars.asc prodvars
gpg -o prodvars -d prodvars.asc
```

## Documentation
The Python application [Grip](https://github.com/joeyespo/grip) is used to render `.md` files; the markdown files can be found at [http://localhost:6419](http://localhost:6419/). 
The following command is used to run the application in the background and supress both `stdout` and `stderr` outputs. Note that when in the Nix shell, the command is already running and does not need to be run again.
```bash
grip > /dev/null 2>&1 &
```
### Development Process
Each task to be completed is tracked through the repository's *GitHub Issues* and through *Github Project*. Furthermore, the specs and designs are detailed in the docs directory.

## FAQ
**Why do I receive the error `...is explicitly required in go.mod, but not marked as explicit in vendor/modules.txt`?**

As outlined in this [discussion](https://discourse.nixos.org/t/buildgomodule-with-local-src-inconsistent-vendoring/8641), this error is a result of a bad vendorSha256; changing the value to `lib.fakeSha256` temporarily will resolve the issue until the real hash value can be used.
