# hot

## Installation

```
$ go get github.com/naoty/hot
```

## Usage

`hot` command lists files in order of the commited count.

```
$ hot
15: todo.go
15: README.md
10: main.go
5: undone.go
5: done.go
5: list.go
4: delete.go
3: add.go
2: Makefile
2: rename.go
2: add_test.go
2: clear.go
2: delete_test.go
2: formatter.go
2: move.go
1: .travis.yml
1: LICENSE
1: clear_test.go
1: done_test.go
1: move_test.go
1: rename_test.go
1: todo_file.go
1: undone_test.go
```

```
$ hot 5
15: todo.go
15: README.md
10: main.go
5: undone.go
5: done.go
```

## Author

[naoty](https://github.com/naoty)

