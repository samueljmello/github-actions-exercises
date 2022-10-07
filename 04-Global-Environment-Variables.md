# Global Environment Variables

In exercise 2, we saw how to use environment variables at the step level, but workflows also allow us to define environment varialbles at the global level as well. These variables are then accessable to all jobs.

This exercise will walk you through the usage of that.

## Step 1: See global environment variables in action

1. From the **default** branch of your repository, create a new branch of code called `feature/global-env`
2. Create a new file named `.github/workflows/global-env.yaml`
3. Copy the contents below to the newly created file:

```yaml
name: Environment Variables
on:
  push:
    branches: feature/global-env
env:
  GLOBAL_VAR: test
jobs:
  first-job:
    name: A Job
    runs-on: ubuntu-latest
    steps:
      - name: Output Global Variable
        run: echo "${GLOBAL_VAR}"
  second-job:
    name: Another Job
    runs-on: ubuntu-latest
    steps:
      - name: Output Global Variable
        run: echo "${GLOBAL_VAR}"
```

4. Add & commit your changes, then publish your branch.
5. Go to your repository, and view the Actions tab to see the execution against your published branch.

The result will be a variable available across all jobs.