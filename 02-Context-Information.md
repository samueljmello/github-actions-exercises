# Context Information

Every workflow execution has access to context information during runtime. This exercise will show you how to access the different contexts available, and what they mean.

Additionally, you will learn how to use environment variables within a step.

## Note
In all the examples below, replace `main` with your **default** branch name:

## Step 1: Create a new workflow to see different contexts

1. From the **default** branch of your repository, create a new branch of code called `feature/context`
2. Create a new file named `.github/workflows/context.yaml`
3. Copy the contents below to the newly created file:

```yaml
name: Context Information
on:
  push:
    branches-ignore: main
jobs:
  show-context:
    name: Show Context
    runs-on: ubuntu-latest
    steps:
      - name: Dump Event Information
        env:
          CONTEXT_ITEM: ${{ toJson(github) }}
        run: echo "${CONTEXT_ITEM}"
      - name: Dump Job Information
        env:
          CONTEXT_ITEM: ${{ toJson(job) }}
        run: echo "${CONTEXT_ITEM}"
      - name: Dump Runner Information
        env:
          CONTEXT_ITEM: ${{ toJson(runner) }}
        run: echo "${CONTEXT_ITEM}"
      - name: Dump Step Information
        env:
          CONTEXT_ITEM: ${{ toJson(steps) }}
        run: echo "${CONTEXT_ITEM}"
```

4. Add & commit your changes, then publish your branch.
5. Go to your repository, and view the Actions tab to see the execution against your published branch.

The result will be an execution of the workflow whenever any changes are pushed EXCEPT on the `main` branch. Context information for `github` and `job` are viewable in the output.

## Step 2: Update the workflow to with step output

1. Replace the contents of the workflow file from the previous step:

```yaml
name: Context Information
on:
  push:
    branches-ignore: main
jobs:
  show-context:
    name: Show Context
    runs-on: ubuntu-latest
    steps:
      - name: Provide Some Step Outputs
        run: echo "::set-output name=TRUE_STATEMENT::Possums are terrible."
      - name: Dump Step Information
        env:
          CONTEXT_ITEM: ${{ toJson(steps) }}
        run: echo "${CONTEXT_ITEM}"
```

4. Add & commit your changes, then push your branch.
5. Go to your repository, and view the Actions tab to see the execution.

The result will be that `steps` is still empty.

## Step 3: Update the workflow with an `id` for the step

1. Follow the same process as step 2, but use this content:

```yaml
name: Context Information
on:
  push:
    branches-ignore: main
jobs:
  show-context:
    name: Show Context
    runs-on: ubuntu-latest
    steps:
      - name: Provide Some Step Outputs
        id: step-outputs
        run: echo "::set-output name=TRUE_STATEMENT::Possums are terrible."
      - name: Dump Step Information
        env:
          CONTEXT_ITEM: ${{ toJson(steps) }}
        run: echo "${CONTEXT_ITEM}"
```

The result will be that the `steps` context now has the data from the previous step.

## Step 4: Add a second output variable

1. Follow the same process as step 3, but use this content:

```yaml
name: Context Information
on:
  push:
    branches-ignore: main
jobs:
  show-context:
    name: Show Context
    runs-on: ubuntu-latest
    steps:
      - name: Provide Some Step Outputs
        id: step-outputs
        run: |
          echo "::set-output name=TRUE_STATEMENT::Possums are terrible."
          echo "::set-output name=FALSE_STATEMENT::Possums are great."
      - name: Dump Step Information
        env:
          CONTEXT_ITEM: ${{ toJson(steps) }}
        run: echo "${CONTEXT_ITEM}"
```

The result will be that the `steps` context now has even more data than the previous step.
