# gorani

echo framework 기반의 go 스켈레톤 프로젝트입니다. go 언어를 이용하여 backend api 서버를 빠르게 만들 수 있도록 하는 것이 목적입니다.

## 개발환경을 위한 설치프로그램 및 설정 (Mac OS X 기준)

1. [Docker Desktop](https://www.docker.com/products/docker-desktop)

2. [Go](https://golang.org/)

3. [Homebrew](https://brew.sh/index_ko)

4. [go-migrate](https://github.com/golang-migrate/migrate)

```shell
$ brew install golang-migrate
```

5. [sqlc](https://sqlc.dev/)

```shell
$ brew install sqlc
```

## Makefile 구성

- Postgresql Database 설치 (docker)

```shell
$ make postgresql
```

- 데이터베이스 생성 / 삭제

```shell
# Makefile 에서 database 이름 수정필요
$ make createdb  # 생성
$ make dropdb    # 삭제
```

- 마이그레이션 파일 생성

```shell
$ make migration
# 입력 프롬프트에서 마이그레이션 파일명 (ex, init_schema) 을 넣으면 자동으로 seq 부여
```

- Migrate Up

```shell
$ make migrateup    # 전체 migration 파일을 일괄 적용할 때 사용
$ make migrateup1   # 최근 1개의 migration 파일만 적용할 때 사용
```

- Migrate Down

```shell
$ make migratedown    # 전체 migration 파일을 일괄 적용할 때 사용
$ make migratedown1   # 최근 1개의 migration 파일만 적용할 때 사용
```

- [sqlc](https://sqlc.dev/) 를 이용한 코드 생성

```shell
$ make sqlc
```

## 프로그램 실행

- 서버 구동

```shell
$ make server
```
