# Project Startup Instructions

## When opening project next time:

### 1. Navigate to project directory

```bash
cd /mnt/c/Users/asus/Documents/LearnCode/{GoProject#5}/coffee-server
```

### 2. Start Docker container

```bash
make start_container
```

### 3. Verify container is running

```bash
docker ps
```

### 4. Run Go application

```bash
go run cmd/server/main.go
```

---

## Fresh Setup (only first time or after deleting container)

```bash
make create_container
make create_db
make migup
```

---

## Quick Reference

**Every restart:**

- `cd coffee-server`
- `make start_container` - starts containers
- `make run` - builds binary and run the code
