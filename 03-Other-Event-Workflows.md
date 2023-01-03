# Other Event Workflows
In addition to push events, workflows can be triggered off of a variety of other events as well. In the examples below we will see some in action and understand their caveats. One main caveat of many events is that they must be in the default branch to function.

See [event trigger documentation](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows) for more details information.

## Step 1: Create an issue-triggered workflow
We will start by creating a workflow triggered by opening an issue and see first-hand one of the most important details of non-push event workflows.

1. From the **default** branch of your repository, create a new branch of code called `feature/issue-workflow`
2. Create a new file named `.github/workflows/issue-workflow.yaml`
3. Copy the contents below to the newly created file:

```yaml
name: Issue Events
on:
  issues:
    types: [opened]
jobs:
  doing-more-things:
    name: Doing Things From Issues
    runs-on: ubuntu-latest
    steps:
      - name: Output Issue Information
        env:
          EVENT_NAME: ${{ github.event_name }}
          EVENT_ACTION: ${{ github.event.action }}
        run: echo "The action '${EVENT_ACTION}' was performed against '${EVENT_NAME}'."
```

4. Add & commit your changes, then publish your branch.
5. Go to your repository, and create an issue.

The result will be that **nothing happens** because this event requires the workflow to be defined in the default branch. In our case, we must push this to **default** branch.

## Step 2: Update the workflow to be in the default branch

6. Open a pull request to merge `feature/issue-workflow` into your default branch.
7. Merge the pull request and delete the branch
8. Go to your repository, and create an issue.

The result will be an execution of the Issue Event Workflow that outputs "The action 'opened' was performed against 'issues'".

9. Checkout to the **default** branch on your local repository and pull down the changes.