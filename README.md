# Zarldev Module Template

This is a template for creating new modules, it uses the [gonew](https://go.dev/blog/gonew) [tool](https://pkg.go.dev/golang.org/x/tools/cmd/gonew) to generate the module.

## Usage

```bash
$ gonew github.com/zarldev/template-module ./<module-name>
```

## Features

- Automatic installation of development dependencies
- Use of tools.go to manage development dependencies
- Use of Makefile to manage the module including setting up the git hooks and repository

## License

[GNU GPL 3.0](https://www.gnu.org/licenses/gpl-3.0.en.html)