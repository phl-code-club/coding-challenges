# coding-challenges

## Seeding Data

The project includes a seeder that can operate in two modes:

### Database Seeding (Default)

Seeds the database with dynamically generated questions, inputs, and answers:

```bash
go run cmd/seed/seed.go
```

## Testing

Once you see the database you can run:

```bash
DB_PATH="../../../test.db" go test ./... -v
```

This will run the tests. As long as they all pass you're good to go!
