install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

up:
	goose mysql "root:root@tcp(localhost:3306)/doutok?parseTime=true" up

down:
	goose mysql "root:root@tcp(localhost:3306)/doutok?parseTime=true" down

# 在终端使用命令：make create name=v0_x_0
create:
	goose mysql "root:root@tcp(localhost:3306)/doutok?parseTime=true" create ${name} sql
