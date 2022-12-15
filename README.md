# tcl &nbsp; [![Go Reference](https://pkg.go.dev/badge/github.com/1H0/tcl.svg)](https://pkg.go.dev/github.com/1H0/tcl) [![Latest Release](https://img.shields.io/github/v/release/1H0/tcl?style=flat-square)](https://github.com/1H0/tcl/releases/latest)


> **tcl** stands for *terminal component library* and provides some "compoents" for nicely formatted terminal outputs.

<details>

<summary>Badges</summary>

[![last-commit](https://img.shields.io/github/last-commit/1H0/tcl?style=flat-square)](https://github.com/1H0/tcl/commits)
[![release-date](https://img.shields.io/github/release-date/1H0/tcl?style=flat-square)](https://github.com/1H0/tcl/releases/latest)
[![issues](https://img.shields.io/github/issues/1H0/tcl?style=flat-square)](https://github.com/1H0/tcl/issues)
[![prs](https://img.shields.io/github/issues-pr/1H0/tcl?style=flat-square)](https://github.com/1H0/tcl/pulls)
[![contributors](https://img.shields.io/github/contributors/1H0/tcl?style=flat-square)](https://github.com/1H0/tcl/graphs/contributors)
![size](https://img.shields.io/github/languages/code-size/1H0/tcl?style=flat-square)
![loc](https://img.shields.io/tokei/lines/github/1H0/tcl?style=flat-square)
![activity](https://img.shields.io/github/commit-activity/m/1H0/tcl?style=flat-square)

</details>

## Table of content

- [tcl    ](#tcl---)
  - [Table of content](#table-of-content)
  - [Installation](#installation)
  - [Documentation](#documentation)
  - [Examples](#examples)
    - [Footer](#footer)
    - [Headers](#headers)
    - [Key-Value Pairs](#key-value-pairs)
    - [Lists](#lists)
      - [Single Unordered List Items](#single-unordered-list-items)
      - [Single Ordered List Items](#single-ordered-list-items)
      - [Unordered List by Array](#unordered-list-by-array)
      - [Ordered List by Array](#ordered-list-by-array)
    - [Messages](#messages)
    - [Quotes](#quotes)
    - [Rulers](#rulers)
    - [Text](#text)
    - [Titles](#titles)
    - [Tables](#tables)
  - [ToDo](#todo)
  - [Credits](#credits)
  - [License](#license)

## Installation

```bash
go get github.com/1H0/tcl
```

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/1H0/tcl.svg)](https://pkg.go.dev/github.com/1H0/tcl)

## Examples

### Footer

```go
tcl.Footer()
tcl.FooterC("1H0")
```

![footers](https://user-images.githubusercontent.com/37246258/172071945-517e85fb-aefc-4a90-86a2-2d9c5a59c195.png)

### Headers

```go
tcl.Header("your app @ v 0.1.0")
```

![headers](https://user-images.githubusercontent.com/37246258/172071944-61076df6-ff20-4b23-b66f-9dbdee20e6e1.png)

### Key-Value Pairs

```go
tcl.KeyValue("Lorem", "Ipsum", 0)
tcl.KeyValue("Dolor", "Sit", 0)
tcl.ListItemU("Lorem", 0)
tcl.KeyValue("Amet", "Consectetur", 1)
```

![key-value](https://user-images.githubusercontent.com/37246258/172071943-0caf3661-ff60-442f-8182-27406d0eb375.png)

### Lists

#### Single Unordered List Items

```go
tcl.ListItemU("Lorem", 0)
tcl.ListItemU("Ipsum", 0)
tcl.ListItemU("Dolor", 0)
tcl.ListItemU("Sit", 1)
tcl.ListItemU("Amet", 1)
tcl.ListItemU("Consectetur", 2)
```

![ul-items](https://user-images.githubusercontent.com/37246258/172071941-129af705-e3ce-490b-806a-db60941f09c7.png)

#### Single Ordered List Items

```go
tcl.ListItemO("Lorem", 0, 1)
tcl.ListItemO("Ipsum", 0, 2)
tcl.ListItemO("Dolor", 0, 3)
tcl.ListItemO("Sit", 1, 1)
tcl.ListItemO("Amet", 1, 2)
tcl.ListItemO("Consectetur", 2, 1)
```

![ol-items](https://user-images.githubusercontent.com/37246258/172071939-2a4824e7-5c35-4095-96f0-9547c5445f37.png)

#### Unordered List by Array

```go
tcl.ListU([]string{"Lorem", "Ipsum", "Dolor"}, 0)
tcl.ListU([]string{"Sit", "Amet"}, 1)
tcl.ListU([]string{"Consectetur"}, 2)
```

![ul-list](https://user-images.githubusercontent.com/37246258/172071938-3356d415-59dc-4914-8efc-737c20c3272e.png)

#### Ordered List by Array

```go
tcl.ListO([]string{"Lorem", "Ipsum", "Dolor"}, 0)
tcl.ListO([]string{"Sit", "Amet"}, 1)
tcl.ListO([]string{"Consectetur"}, 2)
```

![ol-list](https://user-images.githubusercontent.com/37246258/172071937-da167e3f-3456-4357-9a39-3ffbea2678c5.png)

### Messages

```go
tcl.Message("Lorem", "This is a 'black' message", "black")
tcl.Message("Lorem", "This is a 'blue' message", "blue")
tcl.Message("Lorem", "This is a 'cyan' message", "cyan")
tcl.Message("Lorem", "This is a 'green' message", "green")
tcl.Message("Lorem", "This is a 'magenta' message", "magenta")
tcl.Message("Lorem", "This is a 'red' message", "red")
tcl.Message("Lorem", "This is a 'white' message", "white")
tcl.Message("Lorem", "This is a 'yellow' message", "yellow")
tcl.Message("Lorem", "This is a 'default' message", "")
```

![messages](https://user-images.githubusercontent.com/37246258/172071936-c3144240-b621-472b-a90f-380b51fa175c.png)

### Quotes

```go
tcl.Quote("Lorem ipsum dolor sit amet.")
```

![quotes](https://user-images.githubusercontent.com/37246258/172071935-69bbda61-4897-47d7-8b5d-4eaea329433d.png)

### Rulers

```go
tcl.Ruler("/")
tcl.Ruler("- ")
tcl.Ruler("=")
tcl.Ruler(" - ")
tcl.Ruler(" = ")
tcl.Ruler("")
```

![rulers](https://user-images.githubusercontent.com/37246258/172071934-1a78774f-dcd3-413f-8f3f-9741f2ba9910.png)

### Text

```go
tcl.Text("Lorem Ipsum Dolor Sit Amet.")
```

![texts](https://user-images.githubusercontent.com/37246258/172071933-215a22a8-8bd6-476a-9e28-68143b6bfa41.png)

### Titles

```go
tcl.Title("The normal title")
tcl.TitleC("The centered title")
tcl.FullTitle("A Title that spans the whole width")
tcl.SubTitle("The subtitle")
```

![titles](https://user-images.githubusercontent.com/37246258/172071930-41e9ed3b-6ca0-4f69-bd7a-27dfb4ff824f.png)

### Tables

```go
tcl.Table([]string{"Col 1", "Col 2", "Col 3"}, [][]string{
  {"Lorem", "Ipsum", "Dolor"},
  {"Lorem", "Ipsum", "Dolor"},
  {"Lorem", "Ipsum", "Dolor"},
})
```

![Tables](https://user-images.githubusercontent.com/37246258/207857387-2e4a4cad-297b-4063-9931-42bf502bce72.png)

## ToDo

- [x] Nested Lists
- [ ] Tests
- [ ] Ability to loop over Struct and print entrys as key-value pairs
- [x] Tables
- [ ] Customizable Colors and other settings like indentation

## Credits

- [@fatih](https://github.com/fatih) and the maintainers of the [`fatih/color`](https://github.com/fatih/color) for providing the main dependency of this package

## License

The *GNU GENERAL PUBLIC LICENSE* - see [`LICENSE`](./LICENSE)
