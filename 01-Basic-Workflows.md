# Basic Workflow Executions

This exercise will walk you through setting up workflow on your current repository that listens for repository push events. Workflows consume actions and are at the heart of using Actions in your repository.

**GitHub Actions** run off of workflow files that are managed and maintained in your repository. You can see workflow executions and their job history from the Actions tab on your repository. Additionally, you can see the status of those executions (successful, failing, running).

Removing a workflow from your repository does not remove it's history either. Logs are retained based off of the settings in your enterprise & organization.

## Step 1: Create a basic workflow

1. From the **default** branch of your repository, create a new branch of code called `feature/basic-workflow`
2. Create a new file named `.github/workflows/basic-workflow.yaml`
3. Copy the contents below to the newly created file:

```yaml
name: Basic Workflow
on:
  push:
    branches: feature/basic-workflow
jobs:
  do-things:
    name: Do Things
    runs-on: ubuntu-latest
    steps:
      - name: The Thing I've Done
        run: echo "I've done a thing!"
```

4. Add & commit your changes, then publish your branch.
5. Go to your repository, and view the Actions tab to see the execution against your published branch.

The result will be an execution of the workflow whenever any changes are pushed to the `feature/basic-workflow` branch. The results will be a successful execution of the `do-things` job with one step executed.

In this example, the `run` property is used to execute a simple CLI command. In our case, we are using Linux (ubuntu) and therefore are writing bash commands.

## Step 2: Update the workflow to be failing
We will now intentionally update a workflow to make it fail so that we can simulate what you would expect to see if somethin went wrong.

1. Update the workflow file you created with the contents below:

```yaml
name: Basic Workflow
on:
  push:
    branches: feature/basic-workflow
jobs:
  do-things:
    name: Do Things
    runs-on: ubuntu-latest
    steps:
      - name: The Thing I've Done
        run: echo "I've done a thing!"
      - name: An Error I made
        run: |
          echo "I've made an error!"
          exit 1
```

2. Add & commit your changes, then push the updates.
3. Go back to the Actions tab and see the failed execution.
4. Check out to your **default** branch, but leave the `feature/basic-workflow` branch.

The results will be a "failed" execution.

## Step 3: Open a PR to merge your code
Additionally, we will see the failed results in a corresponding pull request. This is, again, to simulate a failing scenario and how it might affect your team.

1. Go to the pull requests tab in your GitHub repository.
2. Open a pull request to merge `feature/basic-workflow` to your **default** branch.
3. Review the status checks area right above the merge option.

The result will be that the pull request shows the details of status checks above the merge option.

Additionally, you've seen how to run multiple commands in one `run` value using the `|` character (pipe).

## Step 4: Clean Up
1. Delete the PR created in [Step 3](#step-3-open-a-pr-to-merge-your-code)
2. Delete the published branch created in [Step 1](#step-1-create-a-basic-workflow)
2. Switch back to the default branch locally.