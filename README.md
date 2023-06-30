# Dice Roller

A simple dice rolling library, written in go.

## Features

- Roll one or many dice of any denomination using a subset of [Standard Dice Notation](https://en.wikipedia.org/wiki/Dice_notation)
- Advantage & Disadvantage (keep highest roll or keep lowest roll)
- Exploding dice (roll an additional die for each die that rolls it's maximum value)

## Examples

Roll a single die

```go
result, _ := Roll("1d6")
fmt.Printf("%+v\n", result)
```

Output

```go
{result:3 rolls:[3]}

Roll multiple dice

```go
result, _ := Roll("4d10")
fmt.Printf("%+v\n", result)
```

Output

```go
{result:18 rolls:[8 6 1 3]}
```

Roll with Advantage (Roll two dice and keep the highest)

```go
result, _ := Roll("2d20kh")
fmt.Printf("%+v\n", result)
```

Output

```go
{result:12 rolls:[4 12]}
```

Roll with Disadvantage (Roll two dice and keep the lowest)

```go
result, _ := Roll("2d20kl")
fmt.Printf("%+v\n", result)
```

Output

```go
{result:4 rolls:[4 12]}
```

Roll Exploding Dice

```go
result, _ := Roll("2d4!")
fmt.Printf("%+v\n", result)
```

Output

```go
{result:16 rolls:[4 1 4 4 3]}
```