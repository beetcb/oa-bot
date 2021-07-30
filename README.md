# Usage

Download release, Run it with .env file:

```bash
go get github.com/joho/godotenv/cmd/godotenv

godotenv -f .env oa-bot path_to_pdf_dir
```

# Build

```bash
godotenv -f .env goreleaser release --snapshot
```
