## Development

Development is facilicated by [Nix](https://nixos.org/). Entering the development environment can be done using the following command. All the required tools for development will be available within the Nix shell. Note that, this repository makes use of [Nix Flakes](https://nixos.wiki/wiki/Flakes) and as such will require that Flakes are enabled.
```bash
nix develop
```
Inside the nix shell, a locked version of the following applications can be found.
- [Git](https://git-scm.com/)
- [Grip](https://github.com/joeyespo/grip)
- [Podman](https://podman.io/)
- [Go](https://go.dev/)

### Development Build
```bash
podman machine init
podman machine start
```
```bash
pnpm dev

```
### Production Build
```bash
pnpm run build
```

### Documentation
The Python application [Grip](https://github.com/joeyespo/grip) is used to render `.md` files; the markdown files can be found at [http://localhost:6419](http://localhost:6419/).
```bash
grip
```
Use the following command to run the application in the background and supress both `stdout` and `stderr` outputs.
```bash
grip > /dev/null 2>&1 &
```
### Testing
TODO

### Development Process
Each task to be completed is tracked through the repository's *GitHub Issues* and through *Github Project*. Furthermore, the specs and designs are detailed in the docs directory.

### Sections
#### [Emails](docs/emailing.md)

## Usage
