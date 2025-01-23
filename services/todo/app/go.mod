module github.com/josephhilby/todoApi/services/todo/app

go 1.22.7

replace(
    github.com/josephhilby/todoApi/api => ../../../api
)

require(
    github.com/josephhilby/todoApi/api v0.0.0
)