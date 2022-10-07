# Matrix Strategy Workflows

Defining a matrix strategy in a workflow allow you to scale out job execution based off of variables you define. This is incredibly useful for repeating the same job but with slightly different values, like if you needed to test software against different versions of a language, or build software against multiple operating systems.

The below example walks you through using a matrix strategy to execute jobs on different runners.

## Step 1: Create a matrix workflow to scale across runners

1. From the __default__ branch of your repository, create a new branch of code called `feature/matrix`
2. Create a new file named `.github/workflows/matrix.yaml`
3. Copy the contents below to the newly created file:

```yaml
name: Matrix Strategy
on:
  push:
    branches: feature/matrix
jobs:
  do-things:
    name: Do Things In A Matrix - ${{ matrix.os }}
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [macos-latest, ubuntu-latest, windows-latest]
    steps:
      - name: The Thing I've Done - ${{ matrix.os }}
        run: echo "I've done a thing on ${{ matrix.os }}!"
```

4. Add & commit your changes, then publish your branch.
5. Go to your repository, and view the Actions tab to see the execution against your published branch.

The result will be an execution of not just one job, but three. The strategy defined one variable with three values (matrix.os) that were then used to scale out. For each variable value you add, the number of jobs increases.

## Step 2: Scale out the matrix strategy even further

1. Replace the contents of the workflow file from the previous step:

```yaml
name: Matrix Strategy
on:
  push:
    branches: feature/matrix
jobs:
  do-things:
    name: Do Things In A Matrix - ${{ matrix.os }}, ${{ matrix.node-version }}
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [macos-latest, ubuntu-latest, windows-latest]
        node-version: [10.x, 12.x, 14.x, 15.x]
    steps:
      - name: Use Node.js ${{ matrix.node-version }}
        uses: actions/setup-node@v3
        with:
          node-version: ${{ matrix.node-version }}
```

4. Add & commit your changes, then push your branch.
5. Go to your repository, and view the Actions tab to see the execution.

The result will be twelve concurrent jobs running across the different runners and setting up the corresponding versions of node.