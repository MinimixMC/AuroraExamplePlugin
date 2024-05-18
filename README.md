⚠️ Highly experimental

<h1 align="center">Aurora Example Plugin 🌟</h1>

> An example plugin for Aurora<br />

## Usage

Aurora plugins can be built as Go plugins, or can be built directly into Aurora.

See platform compatability here: [AuroraAPI](https://github.com/MinimixMC/AuroraAPI/blob/main/README.md)

This example focuses on external plugins, but still applies to internal ones. If you want to know more check the [wiki](https://github.com/MinimixMC/AuroraAPI/wiki)

## Building

```sh
go build -buildmode=plugin -o build/example.so .
```