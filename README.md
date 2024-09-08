# LIBRARY BACKEND

## Instructions

============== /home/anuragm/go/bin/ ====================

## Prerequits

**DOCKER**
**Makefile**
**GO**

### 1. Install golang migrate

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 2. Create migration files (already created in repo)

```bash
migrate create -ext sql -dir ./db/migrations -seq init_library
```

### 3. Install postgres from docker

```bash
make postgres
```

### 4. Install pgAdmin from docker

```bash
make pgadmin
```

### 5. Start postgres and pgAdmin from docker

```bash
make start
```

### 6. Run create DB

```bash
make createdb
```

### 7. Run migrate up to create tables in db

```bash
make migrateup
```

### 3. Install sqlc golang

```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

### 7. Generate code from sqlc

```bash
make sqlc
```
