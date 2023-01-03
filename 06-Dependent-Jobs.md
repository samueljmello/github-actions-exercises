# Dependent Jobs

Jobs by default run in parallel. This allows for the execution of multiple tasks as the same time. Additionally, jobs can be chained in series using the `needs` property.

This exercise will walk you through creating a chained set of jobs the depend on each other.

## Step 1: Create separate parrellel jobs
First we will create two jobs independent of each other (running parallel).

1. From the **default** branch of your repository, create a new branch of code called `feature/dependent`
2. Create a new file named `.github/workflows/dependent.yaml`
3. Copy the contents below to the newly created file:

```yaml
name: Dependent Jobs
on:
  push:
    branches: feature/dependent
jobs:
  first-job:
    name: First Job
    runs-on: ubuntu-latest
    steps:
      - name: Dump Job Information
        env:
          CONTEXT_ITEM: ${{ toJson(job) }}
        run: echo "${CONTEXT_ITEM}"
  second-job:
    name: Second Job
    runs-on: ubuntu-latest
    steps:
      - name: Dump Job Information
        env:
          CONTEXT_ITEM: ${{ toJson(job) }}
        run: echo "${CONTEXT_ITEM}"
```

4. Add & commit your changes, then publish your branch.
5. Go to your repository, and view the Actions tab to see the execution against your published branch.

The result will be an execution that has two jobs which executed at the same time. The execution graph shows that jobs separated.

## Step 2: Modify the workflow to make specific jobs using "needs"
Second we will adjust our workflow to have two jobs in series, and a third in parallel.

1. Replace the contents of the workflow file from the previous step:

```yaml
name: Dependent Jobs
on:
  push:
    branches: feature/dependent
jobs:
  first-job:
    name: First Job (parallel)
    runs-on: ubuntu-latest
    steps:
      - name: Dump Job Information
        env:
          CONTEXT_ITEM: ${{ toJson(job) }}
        run: echo "${CONTEXT_ITEM}"
  second-job:
    name: Second Job (series)
    runs-on: ubuntu-latest
    needs: first-job
    steps:
      - name: Dump Job Information
        env:
          CONTEXT_ITEM: ${{ toJson(job) }}
        run: echo "${CONTEXT_ITEM}"
  third-job:
    name: Third Job (parallel)
    runs-on: ubuntu-latest
    steps:
      - name: Dump Job Information
        env:
          CONTEXT_ITEM: ${{ toJson(job) }}
        run: echo "${CONTEXT_ITEM}"

```

4. Add & commit your changes, then push your branch.
5. Go to your repository, and view the Actions tab to see the execution.

The result will be `second-job` waiting for `first-job` to complete, while `third-job` runs in parallel.

## Step 3: Clean Up
1. Delete the published branch created in [Step 1](#step-1-create-separate-parrellel-jobs)
2. Switch back to the default branch locally.