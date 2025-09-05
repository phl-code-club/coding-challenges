# coding-challenges

## Seeding Data

The project includes a seeder that can operate in two modes:

### Database Seeding (Default)
Seeds the database with dynamically generated questions, inputs, and answers:
```bash
go run cmd/seed/seed.go
```

### File Generation
Generates inputs and answers to files in a `generated/` directory:
```bash
go run cmd/seed/seed.go -generate
```

This creates:
- `questionX_input.txt` - The input data for each question
- `questionX_answers.txt` - The answers for both parts
- `questionX_details.txt` - Question name, intro, and part descriptions