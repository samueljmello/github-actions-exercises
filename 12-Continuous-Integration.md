# Continuous Integration
A common pattern for continuous integration involves using pull requests to trigger testing & other automation tasks.

In the exercise below, we will walk through adding linting & testing for pull requests.

## Step 1: Create a workflow that triggers on pull request creation
1. Checkout the __default__ branch of your repository
2. Create a new file named `.github/workflows/ci.yaml`
3. Copy the contents below to the newly created file:

```yaml
name: Manual & Shared Workflow
on:
  pull_request:
jobs:
  build:
    name: Build the software
    runs-on: ubuntu-latest
    steps:
      - name: Do A Thing
        run: echo "I've done a thing manually with '${{ inputs.choice-example }}' and '${{ inputs.string-example }}'!"
  lint:
  test:
```