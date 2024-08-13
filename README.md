# nit toy project

## tech stack

- [chi](https://github.com/go-chi/chi)
- [templ](https://templ.guide/)
- [tailwindcss](https://tailwindcss.com/)
- [htmx](https://htmx.org/)

## running the application

before running the application, make sure that [air](https://github.com/air-verse/air), [templ](https://github.com/a-h/templ) and tailwindcss cli tools are installed on your machine

```bash
# run the development server
air
```

## managing migrations

database migrations are handled with golang's [goose](https://github.com/pressly/goose) migration tool. goose commands are defined in the makefile to leverage environment variables and use them with goose since it needs to know which database driver to use as well as any connection string and migration files directory.

```bash
# validate migration files
make db-val

# check migrations (current & previous)
make db-status

# create new migration file
make db-create name=SOMENAME

# apply all migrations
make db-up

# revert last migration
make db-down

# revert all migrations
make db-reset
```
