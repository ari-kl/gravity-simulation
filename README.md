# gravity-simulation

This is a simple 2D simulation of gravitational forces between masses. It is written in Go and uses [Ebitengine](https://ebitengine.org) for rendering.

![Screenshot](img/readme.png)

# Usage
A hosted instance of this simulation can be found [here](https://gravity.hexagon.monster).

The controls are as follows:
- `1-9`: Add a mass of `n` at the cursor's current position.
- `0`: Remove all masses.
- `F`: Toggle display of field lines.
- `+/-`: Increase/decrease the strength of the gravitational force (modifies the constant `G`).

# Building
Before building, make sure you have the following dependencies installed:
- Go 1.23 or later
- A working C compiler (for Ebitengine)

## Local
```bash
go mod tidy
go build
```

## WebAssembly
```bash
go mod tidy
GOOS=js GOARCH=wasm go build -o gravity-simulation.wasm
```
