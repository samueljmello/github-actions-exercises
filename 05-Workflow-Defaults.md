# Job Defaults

Default options exist for jobs and steps within a workflow. Currently you can define the `shell` and `working-directory`.

In this exercise, you will see how to control at both the top-level (for all jobs) workflow, as well as for jobs individually.

## Current shell compatiblity

|OS|Shell|Default|
|---|---|---|
|Windows, Linux|pwsh|Windows|
|Windows, Linux|python|
|Windows, Linux|bash|Linux|
|Linux, macOS|sh|macOS|
|Windows|cmd|
|Windows|PowerShell|


## Step 1: Manually assign shell & working directory to different steps
We will start by creating a workflow that sets a bunch of `shell` and `working-directory` values on each step, and then later adjust to reduce those values.

_Note: The first step has to create a source directory so that later steps can use it. By default, we aren't cloing the repository, so I manually chose to create a directory (using: `mkdir src`) to demonstrate._

1. From the **default** branch of your repository, create a new branch of code called `feature/defaults`
2. Create a new file named `.github/workflows/defaults.yaml`
3. Copy the contents below to the newly created file:

```yaml
name: Using Defaults
on:
  push:
    branches: feature/defaults
jobs:
  first-job:
    name: First Job
    runs-on: ubuntu-latest
    steps:
      - name: Create Source Directory
        run: mkdir src
        shell: bash
      - name: Use Python
        run: import os; print("I'm running python! Hissssss! " + os.getcwd());
        shell: python
        working-directory: src
      - name: Use Bash
        run: echo "I'm running hum-drum bash in $(pwd). Booo."
        shell: bash
        working-directory: src
      - name: Use Bash Also
        run: echo "I'm running bash also, but elsewhere in $(pwd). Booo."
        shell: bash
  second-job:
    name: Second Job
    runs-on: ubuntu-latest
    steps:
      - name: Create Source Directory
        run: mkdir src
        shell: bash
      - name: Use Bash
        run: echo "I'm running bash in $(pwd). So sad."
        shell: bash
        working-directory: src
```

4. Add & commit your changes, then publish your branch.
5. Go to your repository, and view the Actions tab to see the execution against your published branch.

The result will be an execution which utilizes bash for most of the steps within the jobs.

## Step 2 - Utilizing defaults at the job level
Now we will see how `jobs` can have `defaults`.

1. Replace the contents of the workflow file from the previous step:

```yaml
name: Using Defaults
on:
  push:
    branches: feature/defaults
jobs:
  first-job:
    name: First Job
    runs-on: ubuntu-latest
    defaults:
      run:
        shell: bash
        working-directory: src
    steps:
      - name: Create Source Directory
        run: mkdir src
        working-directory: .
      - name: Use Python
        run: import os; print("I'm running python! Hissssss! " + os.getcwd())
        shell: python
      - name: Use Bash
        run: echo "I'm running hum-drum bash in $(pwd). Booo."
      - name: Use Bash Also
        run: echo "I'm running bash also, but elsewhere in $(pwd). Booo."
        working-directory: ..
  second-job:
    name: Second Job
    runs-on: ubuntu-latest
    steps:
      - name: Create Source Directory
        run: mkdir src
        shell: bash
        working-directory: .
      - name: Use Bash
        run: echo "I'm running bash in $(pwd). So sad."
        shell: bash
        working-directory: src
```

4. Add & commit your changes, then push your branch.
5. Go to your repository, and view the Actions tab to see the execution.

The result will be less code in your workflow, while the execution still performed the same. It used `bash` as the default `shell` and provided a default working directory of `src`

## Step 3 - Utilizing defaults at the workflow level
Finally, we will see how the entire workflow has `defaults`.

1. Follow the same process as step 2, but use this content:

```yaml
name: Job Default
on:
  push:
    branches: feature/defaults
defaults:
  run:
    shell: bash
    working-directory: src
jobs:
  first-job:
    name: First Job
    runs-on: ubuntu-latest
    steps:
      - name: Create Source Directory
        run: mkdir src
        working-directory: .
      - name: Use Python
        run: import os; print("I'm running python! Hissssss! " + os.getcwd())
        shell: python
      - name: Use Bash
        run: echo "I'm running hum-drum bash in $(pwd). Booo."
      - name: Use Bash Also
        run: echo "I'm running bash also, but elsewhere in $(pwd). Booo."
        working-directory: ..
  second-job:
    name: Second Job
    runs-on: ubuntu-latest
    steps:
      - name: Create Source Directory
        run: mkdir src
        working-directory: .
      - name: Use Bash
        run: echo "I'm running bash in $(pwd). So sad."
```

The result will be even less code, making `bash` the default `shell` and `src` the `working-directory` for all job steps.

## Step 4 - Clean Up

1. Delete the published branch created in [Step 1](#step-1-manually-assign-shell--working-directory-to-different-steps)
2. Switch back to the default branch locally.
