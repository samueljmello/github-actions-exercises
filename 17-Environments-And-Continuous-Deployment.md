# Environments & Continuous Deployment
Environments go hand-in-hand with continuous deployment and allow you to add gates and secret overrides to control how & when your deployment occurs.

In a previous exercise, you walked through setting up a continuous integration pipeline that created an app as an artifact (see [13-Continuous-Integration](./13-Continuous-Integration.md)). Later, you modified that process to include continuous delivery which created a Docker image to be deployed (see [16-Packages-And-Continuous-Delivery](./16-Packages-And-Continuous-Delivery.md)).

In this exercise, you will add a `Continuous Deployment` pipeline that deploys once `Continuous Integration & Delivery` completes for the **default** branch (in the examples, `main`). Additionally, you will configure an environment to see the gating functionality.

## Step 1: Create the continuous deployment pipeline
1. From the **default** branch of your repository, create a new branch of code called `feature/deployment`
2. Create a new file named `.github/workflows/cd.yaml`
3. Copy the contents below to the newly created file:

```yaml
name: Continuous Deployment
on:
  workflow_dispatch:
  workflow_run:
    workflows: [Continuous Integration & Delivery]
    branches: [main]
    types:
      - completed
defaults:
  run:
    shell: bash
env:
  IMAGE_ID: ghcr.io/${{ github.repository_owner }}/${{ github.event.repository. name }}
jobs:
  deploy:
    name: Deploy Docker Image
    runs-on: ubuntu-latest
    steps:
      # A real example would have deployment steps for a container, like kubectl commands (for Kubernetes)
      - name: Log In To Package Registry
        run: echo "${{secrets.GITHUB_TOKEN}}" | docker login ghcr.io -u $ --password-stdin
      - name: Pull Down The Image
        run: docker pull $(echo $IMAGE_ID | tr '[A-Z]' '[a-z]'):latest
      - name: Run The Container
        run: docker run $(echo $IMAGE_ID | tr '[A-Z]' '[a-z]'):latest
```

4. Add & commit your changes, then publish your branch.
5. Go to your repository and open a pull request to merge `feature/deployment` to your **default** branch.
6. Click the green `Merge pull request` button on the pull request from step 1.5. This will put your code into the default branch.
7. Go to the `Actions` tab to see the workflow executions.

The result will be the `Continuous Integration & Delivery` workflow executing. Once that completes, the `Continuous Deployment` workflow will execute automatically. The workflow doesn't do a real deployment and simply runs the container itself, but you would replace those steps with real container orchestration commands.

## Step 2: Add an environment and configure
1. Using the [official documentation](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment#creating-an-environment): 
   1. configure an environment named `production` (case sensitive)
   2. Add yourself as a required reiewer (make sure to save the rules).
   3. Limit the deployment branches to `selected branches`, and add your **default** branch (in the examples above, `main`) as the only branch allowed.
2. From the **default** branch of your repository, create a new branch of code called `feature/environment`
3. Replace the contents of `.github/workflows/cd.yaml` with:

```yaml
name: Continuous Deployment
on:
  workflow_dispatch:
  workflow_run:
    workflows: [Continuous Integration & Delivery]
    branches: [main]
    types:
      - completed
defaults:
  run:
    shell: bash
env:
  IMAGE_ID: ghcr.io/${{ github.repository_owner }}/${{ github.event.repository. name }}
jobs:
  deploy:
    name: Deploy Docker Image
    environment: production
    runs-on: ubuntu-latest
    steps:
      # A real example would have deployment steps for a container, like kubectl commands (for Kubernetes)
      - name: Log In To Package Registry
        run: echo "${{secrets.GITHUB_TOKEN}}" | docker login ghcr.io -u $ --password-stdin
      - name: Pull Down The Image
        run: docker pull $(echo $IMAGE_ID | tr '[A-Z]' '[a-z]'):latest
      - name: Run The Container
        run: docker run $(echo $IMAGE_ID | tr '[A-Z]' '[a-z]'):latest
```

The only real difference here is adding the `environment: production` line to the job in the workflow.

4. Add & commit your changes, then publish your branch.
5. Go to your repository and open a pull request to merge `feature/environment` to your **default** branch.
6. Click the green `Merge pull request` button on the pull request from step 1.5. This will put your code into the main branch.
7. Go to the `Actions` tab to see the workflow executions.

The result will be the same as before (`Continuous Integration & Delivery` executes), except the this time `Continuous Delivery` will have a `waiting` status which will require an approval before it will run.

## Step 3: Approve the deployment
1. From the `Actions` tab in your repo, view the `Continuous Delivery` workflow execution that is waiting.
2. You will see a yellow banner with a `Review deployments` button. Click that, check the box for the environment (`production`) and then click `Approve and deploy`.

The result will be the `Continuous Delivery` workflow fully executing, the same as before.

## Step 4: View the environments GUI
1. Navigate to the `Code` tab on your repository
2. On the right side, the `Environments` section will now have `production` listed, with it's current status.
3. Click the `production` environment to see the history.

## Step 5: Clean Up
2. Delete the published branch created in [Step 1](#step-1-create-the-continuous-deployment-pipeline)
2. Switch back to the default branch locally.