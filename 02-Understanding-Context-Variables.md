# Understanding Context Variables
Every workflow execution has access to context information during runtime. This exercise will show you how to access the different contexts available, and what they mean.

Additionally, you will learn how to use environment variables within a step.

## Note
In all the examples below, replace `main` with your **default** branch name:

## Step 1: Create a new workflow to see different contexts
First, let's take a look at all the different context objects available in the runner. We will do this by converting the objects to JSON and storing them in an environment variable on each step. Lastly, we'll output that environment variable with a simple `echo` statement.

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

The result will be an execution of the workflow whenever any changes are pushed EXCEPT on the `main` branch. Context information is viewable in the output so that you can understand how to utilize those values in your workflows.

## Step 2: Update the workflow to with step output
Previously, we saw all the data available in the various context objects, but the `steps` context variable was empty. Here, we will learn how to populate the `steps` context variable, so we can pass data to subsequent steps. We will also purposely ommit a setting required to complete this step as it can be a common mistake.

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
        run: echo "TRUE_STATEMENT=Possums are terrible." >> $GITHUB_OUTPUT
      - name: Dump Step Information
        env:
          CONTEXT_ITEM: ${{ toJson(steps) }}
        run: echo "${CONTEXT_ITEM}"
```

4. Add & commit your changes, then push your branch.
5. Go to your repository, and view the Actions tab to see the execution.

The result will be that `steps` is still empty.

## Step 3: Update the workflow with an `id` for the step
Finally, to complete what we learned in [Step 2](#step-2-update-the-workflow-to-with-step-output), we will apply the necessary property `id` so we can pass data between steps.


1. Follow the same process as [Step 2](#step-2-update-the-workflow-to-with-step-output), but use this content:

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
        run: echo "TRUE_STATEMENT=Possums are terrible." >> $GITHUB_OUTPUT
      - name: Dump Step Information
        env:
          CONTEXT_ITEM: ${{ toJson(steps) }}
        run: echo "${CONTEXT_ITEM}"
```

The result will be that the `steps` context now has the data from the previous step.

## Step 4: Add a second output variable
It's possible to pass multiple values in a single step. This will teach you how to accomplish that and then access both values.

1. Follow the same process as [Step 3](#step-3-update-the-workflow-with-an-id-for-the-step), but use this content:

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
          echo "TRUE_STATEMENT=Possums are terrible." >> $GITHUB_OUTPUT
          echo "FALSE_STATEMENT=Possums are great." >> $GITHUB_OUTPUT
      - name: Dump Step Information
        env:
          CONTEXT_ITEM: ${{ toJson(steps) }}
        run: echo "${CONTEXT_ITEM}"
```

The result will be that the `steps` context now has even more data than the previous step.

## Step 5: Clean Up

1. Delete the published branch created in [Step 1](#step-1-create-a-new-workflow-to-see-different-contexts)
2. Switch back to the default branch locally.
