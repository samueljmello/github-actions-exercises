# Reusable Workflows

In the previous exercise [09-Manual-Workflow.md](./09-Manual-Workflow.md), you created a workflow that could be triggered manually.

In this exercise, we will modify that same workflow to be called by other workflows. This model can be used with public or internal repository workflows in order to create a centralized set of workflows. This can reduce workflow configuration across your organization, saving time and effort.

## Note
In any examples below, replace `main` with your **default** branch name:

## Step 1: Modify the manual workflow and add in changes

1. Checkout the **default** branch of your repository
2. Open the file named `.github/workflows/manual.yaml` (if you don't have this file, follow the instructions in [09-Manual-Workflow.md](./09-Manual-Workflow.md))
3. Replace the contents of the file with:

```yaml
name: Manual & Shared Workflow
on:
  workflow_dispatch:
    inputs:
      choice-example:
        description: Choice Example
        required: true
        default: warning
        type: choice
        options:
        - info
        - warning
        - debug
      string-example:
        description: String Example
        required: true
        default: input
        type: string
  workflow_call:
    inputs:
      choice-example:
        description: Choice Example
        required: true
        default: warning
        type: string
      string-example:
        description: String Example
        required: true
        default: input
        type: string
jobs:
  do-things:
    name: Do Things Manually
    runs-on: ubuntu-latest
    steps:
      - name: Do A Thing
        run: echo "I've done a thing manually with '${{ inputs.choice-example }}' and '${{ inputs.string-example }}'!"
```

*Notice that the `choice-example` input cannot be of type `choice` for `workflow_call`.*

4. Add, commit, and push your changes to the default branch.
5. Go to your repository, and view the Actions tab to see the workflow you created (`Manual & Shared Workflow`)

The result will be a button a workflow that is still manually executable

## Step 2: Add a new workflow that uses the `Manual & Shared` workflow

1. Checkout the **default** branch of your repository
2. Create a new file named `.github/workflows/shared.yaml`
3. Replace the contents of the file with:

```yaml
name: Using Shared Workflow
on:
  push:
    branches: main
jobs:
  centralized-job:
    uses: ./.github/workflows/manual.yaml
    with:
      choice-example: debug
      string-example: a thing
```
4. Add, commit, and push your changes to the default branch.
5. Go to your repository, and view the Actions tab to see the workflow you created (`Manual & Shared Workflow`)

The result will be an execution of the workflow `Manual & Shared Workflow`, that will output the values passed to it from the `with` option.

## Step 3: Clean Up
No further steps are needed as you committed directly to the default branch.

## Notes
- When using a shared workflow from a **public** or **internal** repository, the syntax will be slightly different:
`<owner>/<repo>/<path-to-workflow>@<release-tag>`
- Workflows shared from an **internal** repository must have the appropriate repository settings (see [instructions here](https://docs.github.com/en/enterprise-cloud@latest/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-github-actions-settings-for-a-repository#allowing-access-to-components-in-an-internal-repository))