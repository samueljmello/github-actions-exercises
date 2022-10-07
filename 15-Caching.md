# Caching
When working with workflows, you might find yourself performing the same set of tasks multiple times across workflows or even jobs. This is commonly seen in build scenarios where dependencies are downloaded. Because Actions runners are ephemeral, the files from a previous execution are deleted.

This inevitable can cause builds to run longer, especially when the number of dependencies used are numerous.

In the previous exercise on `Continuous Integration`, you built a workflow that is performing automation when Pull Requests are opened to merge to the **default** branch. In this exercise, we will reuse that workflow to demonstrate a bit of caching.

## Prerequisite Exercises
- [13-Continuous-Integration](./13-Continuous-Integration.md)

## Step 1: Add the cache action
1. From the **default** branch of your repository, create a new branch of code called `feature/cache`
2. Open the file named `.github/workflows/ci.yaml`
3. Replace the contents of the file with:

```yaml
name: Continuous Integration
on:
  pull_request:
  workflow_dispatch:
defaults:
  run:
    shell: bash
    working-directory: golang_app
jobs:
  build-and-test:
    name: Cache, Build, Lint, & Test
    runs-on: ubuntu-latest
    steps:
      - name: Path Setup
        id: go-paths
        working-directory: ${{github.workspace}}
        run: |
          echo "::set-output name=gomodcache::$(go env GOMODCACHE)"
          echo "::set-output name=gocache::$(go env GOCACHE)"
      - name: Clone
        uses: actions/checkout@v3.1.0
      - name: Cache
        uses: actions/cache@v3.0.10
        with:
          path: |
            ${{ steps.go-paths.outputs.gomodcache }}
            ${{ steps.go-paths.outputs.gocache }}
          key: ${{ runner.os }}-gomodcache-${{ hashFiles('**/go.sum') }}
      - name: Get Dependencies
        run: go get app
      - name: Build
        run: go build
      - name: Run Linting
        uses: golangci/golangci-lint-action@v3
        with:
          working-directory: golang_app
      - name: Run Tests
        run: go test
      - name: Store Artifact
        uses: actions/upload-artifact@v3.1.0
        with:
          name: golang_app
          path: golang_app/app
```

4. Replace the contents of the file [golang_app/main.go](./golang_app/main.go) with the contents in [./golang_replacements/15-main.go](./golang_replacements/15-main.go)
5. Replace the contents of the file [golang_app/go.mod](./golang_app/go.mod) with the contents in [./golang_replacements/15-go.mod](./golang_replacements/15-go.mod)
6. Replace the contents of the file [golang_app/go.sum](./golang_app/go.sum) with the contents in [./golang_replacements/15-go.sum](./golang_replacements/15-go.sum)
7. Add & commit your changes, then publish your branch.
8. Go to your repository, and view the Pull Requests tab.
9. Create a pull request to merge `feature/cache` into your **default** branch.

In the steps above, you added several steps to the workflow file. These steps help to define the Go paths for caching, and then pass those paths to the `actions/cache` action that handles the downloading & updating of caches.

No additional step is required to save your cache, because the action has post-execution hooks that handle it for you.

The result won't be a lot of time saved because the number of dependencies are very few, but in a real-world scenario this could save a lot, reducing dependency downloads from the 10's of minutes to just a few seconds.

## Additional Documentation
- [GitHub Actions Caching](https://github.com/actions/cache)