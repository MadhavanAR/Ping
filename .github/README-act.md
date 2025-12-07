# Running GitHub Actions Locally with Act

This guide helps you test your GitHub Actions workflows locally using [act](https://github.com/nektos/act).

## Quick Start

1. **Run the helper script:**
   ```bash
   chmod +x .github/act-local.sh
   ./github/act-local.sh
   ```

2. **Or run act directly:**
   ```bash
   # List all workflows
   act -l
   
   # Run the ping-ci workflow (skips Docker push)
   act push -W .github/workflows/ping-ci.yml --skip-push
   
   # Run specific job
   act -j build-and-deploy --skip-push
   ```

## Common Commands

### List workflows
```bash
act -l
```

### Run workflow on push event
```bash
act push -W .github/workflows/ping-ci.yml
```

### Run without Docker push (recommended for local testing)
```bash
act push -W .github/workflows/ping-ci.yml --skip-push
```

### Run specific job
```bash
act -j build-and-deploy --skip-push
```

### Run with secrets (optional)
Create a `.secrets` file in the project root:
```
DOCKERHUB_USERNAME=your_username
DOCKERHUB_TOKEN=your_token
```

Then run:
```bash
act --secret-file .secrets
```

### Run with environment variables
```bash
act --env NODE_ENV=production
```

### Use a specific Docker image
```bash
act -P ubuntu-latest=catthehacker/ubuntu:act-latest
```

### Apple Silicon (M1/M2/M3) users
If you're on Apple Silicon, use the linux/amd64 architecture:
```bash
act push -W .github/workflows/ping-ci.yml --skip-push --container-architecture linux/amd64
```

## Notes

- **Docker Push**: The workflow includes a Docker push step. Use `--skip-push` to skip it during local testing.
- **Secrets**: Docker Hub credentials are not needed for local testing unless you want to test the push step.
- **Performance**: Local runs may be slower than GitHub Actions runners.
- **Differences**: Some GitHub-specific features may not work exactly the same locally.

## Troubleshooting

### MODULE_NOT_FOUND errors
If you encounter module errors, ensure all dependencies are installed:
```bash
cd webapp && npm install
cd channels && npm install
```

### Docker issues
Make sure Docker is running:
```bash
docker ps
```

### Large images
Act uses Docker images. First run may take time to download images. Subsequent runs will be faster.

## Resources

- [Act Documentation](https://github.com/nektos/act)
- [Act User Guide](https://github.com/nektos/act#usage)

