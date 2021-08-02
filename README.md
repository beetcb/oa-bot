# Usage

Download release, Run it with .env file:

```bash
go get github.com/joho/godotenv/cmd/godotenv

oa-bot path_to_pdf_dir pass_for_renv
```

# Build

```bash
godotenv -f .env goreleaser release
```
