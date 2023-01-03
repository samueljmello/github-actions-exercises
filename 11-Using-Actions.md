# Using Actions
Up until this point, all previous exercsies focused on using the `run` property for a step. This allowed us to execute CLI commands in the environment we were using (Ubuntu Linux). As a result, we witnessed simple bash commands (`echo`, etc.).

The following exercise will walk you through using actions (these can be from a public or internal repository) which differs from simple `run` commands because they utilize the Action creation syntax and rules, which we will cover in the next step.

## Step 1: Create a new workflow file

1. From the **default** branch of your repository, create a new branch of code called `feature/using-actions`
2. Create a new file named `.github/workflows/using-actions.yaml`
3. Copy the contents below to the newly created file:

```yaml
name: Using Actions
on:
  push:
    branches: feature/using-actions
jobs:
  do-things:
    name: Doing Things With Actions
    runs-on: ubuntu-latest
    steps:
      - name: List Directory Contents
        run: ls -latr
      - name: Checkout The Code
        uses: actions/checkout@v3
      - name: List Directory Contents Again
        run: ls -latr
```

In previous examples, we were running commands in the runner, but the runner (by default) does not have the repository contents cloned down. In the above configuration, we use the `actions/checkout` action (owned by GitHub in the organization `actions`) to clone the code down to the runner.

4. Add & commit your changes, then publish your branch.
5. Go to your repository, and view the Actions tab to see the execution against your published branch.

The result will be an execution that lists the contents of an empty directory, clones the repository down, and then lists the contents again (this time, with files).

Similarly to reusable workflows, the syntax for an action is `<owner>/<repo>@<release-tag>`

## Step 2: Adding inputs for an action

When using an action, you may be required to give various inputs. This is unique to each action. To see a list of inputs, simply visit `https://github.com/<owner>/<repo>`, and look at the source of the `action.yml` file.

For example: `https://github.com/actions/checkout/blob/main/action.yml`

The `inputs` property defines all possible inputs and whether they are required.

1. Modify the workflow file you just created: `.github/workflows/using-actions.yaml`
2. Replace the contents with:

```yaml
name: Using Actions
on:
  push:
    branches: feature/using-actions
jobs:
  do-things:
    name: Doing Things With Actions
    runs-on: ubuntu-latest
    steps:
      - name: List Directory Contents
        run: ls -latr
      - name: Checkout The Code
        uses: actions/checkout@v3
        with:
          clean: true
      - name: List Directory Contents Again
        run: ls -latr
```

4. Add & commit your changes, then push them up.
5. Go to your repository, and view the Actions tab to see the execution.

The result will be the same as before, but this time we've added the `clean` input to perform some additional clean tasks when cloning the repository (defined in the documentation)

## Step 3: Using `script-actions` to write quick actions for interaction with the GitHub API

These exercises don't cover creating your own actions, but for the sake of understanding some of the underlying process and getting a bit more familiar with inputs, we will attempt to use the popular action `actions/github-script`.

This utilizes Node.js and [OctoKit](https://github.com/octokit) to execute the script you pass to it, much like you would in creating your own Action with Node.js. This simplifies the process by taking care of instantiating OctoKit and any additional set-up work involved.

1. Modify the workflow file you just created: `.github/workflows/using-actions.yaml`
2. Replace the contents with:

```yaml
name: Using Actions
on:
  push:
    branches: feature/using-actions
jobs:
  do-things:
    name: Doing Things With Actions
    runs-on: ubuntu-latest
    steps:
      - name: List Directory Contents
        run: ls -latr
      - name: Checkout The Code
        uses: actions/checkout@v2
        with:
          clean: true
      - name: List Directory Contents Again
        run: ls -latr
      - name: Use GitHub Script Action
        uses: actions/github-script@v6
        with:
          script: |
            const create = await github.rest.issues.create({
              owner: context.repo.owner,
              repo: context.repo.repo,
              title: "Learning To Use Script-Actions!",
              body: "Hey, we used script actions! Here's the commit SHA that triggered this: ${{github.sha}}"
            })
            return create.data.number
```

4. Add & commit your changes, then push them up.
5. Go to your repository and view the Actions tab to see the execution, then visit the Issues tab.

The result will be an issue created in your repository.

## Step 4: Merge the pull request into the **default** branch for later use

In the previous steps, we implemented public actions, but specifically used an out-of-date version. We want this to be put into the **default** branch so we can see it in a later step.

1. Go to your repository and navigate to the "Pull Requests" tab
2. Create a new pull request to merge `feature/using-actions` into your **default** branch
3. Merge the pull request created.
4. Delete the published branch created in [Step 1](#step-1-create-a-new-workflow-file)