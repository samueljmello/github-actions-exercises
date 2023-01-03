# Concurrency Groups

By default, all workflows operate with some level of concurrency defined. Concurrency groups allow you to further define concurrency across workflows or jobs. This will enable you to make entire workflows wait for other workflows to finish, or to cancel previous queued executions.

## Step 1: Create two workflows listening for the same events with concurrency defined

1. From the **default** branch of your repository, create a new branch of code called `feature/concurrency`
2. Create a new file named `.github/workflows/concurrency.yaml`
3. Copy the contents below to the newly created file:

```yaml
name: Concurrent Workflows
on:
  push:
    branches: feature/concurrency
jobs:
  do-things:
    name: Do Things With Concurrency
    runs-on: ubuntu-latest
    concurrency: 
      group: static-group
    steps:
      - name: The Thing I've Done
        run: echo "I've done a thing!"; sleep 15;
```

4. Additionally, create another file named `.github/workflows/concurrency-2.yaml`
5. Copy the contents below to the second newly created file:

```yaml
name: Concurrent Workflows No. 2
on:
  push:
    branches: feature/concurrency
jobs:
  do-things:
    name: Do Other Things With Concurrency
    runs-on: ubuntu-latest
    concurrency: 
      group: static-group
    steps:
      - name: The Other Thing I've Done
        run: echo "I've done another thing!"; sleep 15;
```

6. Add & commit your changes, then publish your branch.
7. Go to your repository, and view the Actions tab to see the execution against your published branch.

The result will be an execution of the workflow whenever any changes are pushed to the `feature/concurrency` branch, but the second workflow (whichever one is picked up second) will wait for the first to complete before allowing execution.

This isn't a real-world use scenario, but it demonstrates to apply concurrency across workflows. Additional executions of a workflow due to additional events being sent already act this way (read: multiple commits pushed up), but if you need to also apply it across other workflows, this is how you accomplish that.

## Step 2: Add cancellation to workflows

1. Modify the workflow file `.github/workflows/concurrency.yaml` with the following contents:

```yaml
name: Concurrency Group
on:
  push:
    branches: feature/concurrency
jobs:
  do-things:
    name: Do Things With Concurrency
    runs-on: ubuntu-latest
    concurrency: 
      group: static-group
      cancel-in-progress: true
    steps:
      - name: The Thing I've Done
        run: echo "I've done a thing!"; sleep 15;
```

2. Modify the second workflow file `.github/workflows/concurrency-2.yaml` with the following contents:

```yaml
name: Concurrency Group No. 2
on:
  push:
    branches: feature/concurrency
jobs:
  do-things:
    name: Do Other Things With Concurrency
    runs-on: ubuntu-latest
    concurrency: 
      group: static-group
      cancel-in-progress: true
    steps:
      - name: The Other Thing I've Done
        run: echo "I've done another thing!"; sleep 15;
```

3. Add & commit your changes, then push to GitHub.

The result will be that one workflow execution (whichever is triggered first) will cancel when the second is queued because they are in the same concurrency group.

## Step 3: Clean Up
1. Delete the published branch created in [Step 1](#step-1-create-two-workflows-listening-for-the-same-events-with-concurrency-defined)
2. Switch back to the default branch locally.