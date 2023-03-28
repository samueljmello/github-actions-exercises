# Packages
GitHub has a built-in feature for package registries called "Packages". It supports various registries such as `npm`, `mvn`, `Docker`, and more (see [documentation](https://docs.github.com/en/packages/learn-github-packages/introduction-to-github-packages) for full list).

In this exercise, you will learn how to use the GitHub package registry for Docker ([https://ghcr.io](https://ghcr.io)) with to publish a Docker image as part of a Continuous Delivery pipeline.

## Prerequisite Exercises
- [13-Continuous-Integration](./13-Continuous-Integration.md)
- [14-Artifacts](./14-Artifacts.md)
- [15-Caching](.15-Caching.md)

## Step 1: Add the continuous delivery job that creates the package
1. From the **default** branch of your repository, create a new branch of code called `feature/packages`
2. Open the file named `.github/workflows/ci-cd.yaml`
3. Replace the contents of the file with:

```yaml
name: Continuous Integration & Delivery
on:
  pull_request:
  workflow_dispatch:
  push:
    branches: main
defaults:
  run:
    shell: bash
jobs:
  ci:
    name: Continuous Integration
    runs-on: ubuntu-latest
    defaults:
      run:
        working-directory: golang_app
    steps:
      - name: Path Setup
        id: go-paths
        working-directory: ${{github.workspace}}
        run: |
          echo "gomodcache=$(go env GOMODCACHE)" >> $GITHUB_OUTPUT
          echo "gocache=$(go env GOCACHE)" >> $GITHUB_OUTPUT
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
  cd:
    name: Continuous Delivery
    needs: ci
    runs-on: ubuntu-latest
    defaults:
      run:
        working-directory: docker_image
    steps:
      - name: Clone
        uses: actions/checkout@v3.1.0
      - name: Get Artifact
        uses: actions/download-artifact@v3.0.0
        with:
          name: golang_app
          path: docker_image
      - name: Log In To Package Registry
        run: echo "${{secrets.GITHUB_TOKEN}}" | docker login ghcr.io -u $ --password-stdin
      - name: Build & Push
        run: |
          IMAGE_ID=ghcr.io/${{ github.repository_owner }}/${{ github.event.repository.name }}
          IMAGE_ID=$(echo $IMAGE_ID | tr '[A-Z]' '[a-z]')
          VERSION=$(echo "${{github.ref}}" | sed -e 's,.*/\(.*\),\1,')
          [ "$VERSION" == "main" ] && VERSION=latest
          docker build . --tag $IMAGE_ID:$VERSION --label "runnumber=${{github.run_id}}" --file Dockerfile
          docker push $IMAGE_ID:$VERSION
```

4. Add & commit your changes, then publish your branch.
5. Go to your repository and open a pull request to merge `feature/packages` to your **default** branch.
6. Go to your `Actions` tab and watch the resulting workflow execution for `Continuous Integration & Delivery`.

The result will be an artifact created (the Golang app) like we experienced before, and an additional Docker image will be uploaded to the package registry.

7. Go to your `Code` tab on your repository and click the `Packages` link on the right-side navigation.

Now you have a fully deployable Docker image, ready to be ran as a container on any orchestration platform, or to be used as a `FROM` reference in another Dockerfile.


## Step 2: Merge the changes to your default branch and update your local repository

1. Click the green `Merge pull request` button on the pull request from step 1.5. This will put your code into the main branch.
2. Delete the published branch created in [Step 1](#step-1).
3. Checkout to your default branch locally and pull down the changes.