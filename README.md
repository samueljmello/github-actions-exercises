# GitHub Actions Exercises
This repository exists to provide users of GitHub Actions with very simple and easy-to-follow examples, breaking down more about how Actions works, along with nuances of working with the different features included.

---

## Exercise Prerequisites

### Software Installed
- [Git](https://git-scm.com/downloads)
- A source code editor ([VSCode](https://code.visualstudio.com/download), etc.)

### Environment
1. A GitHub repository with admin permissions
2. Authentication to the repository (OAUTH credentials or PAT token)
3. The repository cloned locally, then the [remote updated](https://docs.github.com/en/get-started/getting-started-with-git/managing-remote-repositories#changing-a-remote-repositorys-url) to your repository from step #1.

### Knowledge
- [How to clone a GitHub repository](https://docs.github.com/en/repositories/creating-and-managing-repositories/cloning-a-repository)
- [How to add, commit, and push changes to a GitHub repository](https://github.com/git-guides/git-commit)
- [How to create & publish a branch](https://github.com/git-guides/git-push)

---

## Outline
1. [Basic workflow creation](./01-Basic-Workflows.md)
2. [Context information & Environment Variables](./02-Context-Information.md)
3. [Other event workflow creation](./03-Other-Event-Workflows.md)
4. [Global environment variables across jobs](./04-Global-Environment-Variables.md)
5. [Workflow defaults](./05-Workflow-Defaults.md)
6. [Dependent jobs creation](./06-Dependent-Jobs.md)
7. [Concurrency groups](./07-Concurrency-Groups.md)
8. [Matrix strategy creation](08-Matrix-Strategy.md)
9. [Manual workflows](09-Manual-Workflow.md)
10. [Reusable workflows](./10-Reusable-Workflow.md)
11. [Using public actions](./11-Using-Actions.md)
12. [Continuous Integration](./12-Continuous-Integration.md)
13. [Dependabot config](./13-Dependabot-Config.md)
14. [Artifacts](./14-Artifacts.md)
15. [Caching](./15-Caching.md)
16. [Packages & Continuous Delivery](./16-Packages-And-Continuous-Delivery.md)
17. [Environments & Continuous Deployment](./17-Environments-And-Continuous-Deployment.md)
18. Releases & Tags

---

## Vocabulary
Below is a quick reference to some of the vocabulary of GitHub Actions

|Term|Description|
|---|---|
|GitHub Actions|The entire product suite that surrounds workflows, actions, etc.|
|Workflow|A set of jobs to be executed on specific events.|
|Job|A group of steps in a job.|
|Steps|The smallest part of a workflow (often referred to as an action), contains command-line code, public action references, or shared workflow references.|
|Action|Can be synonimous with step, but typically refers to reusable actions developed internally or externally, that your steps can reference and use.|
|Runner|The virtual machine, physical machine, or container that workflows execute on.|

---

## Files In This Repository
- `./golang_app` contains all files pertaining to the Golang application we will build in the Continuous Integration exercise.
- `./images` contains any images placed in the exercise Markdown (`./md`) files.
- `./.gitignore` simply ignores specific files that should not exist in the repository.
- `./<##>-<Exercise-Name>.md` are the exercises written in Markdown and can be viewed through the GitHub repository, or opened in a code editor.

---

## Commands Quick-Reference
- Push: `git push`
- Add & commit: `git add <path-to-file>; git commit -m "<your message>"` (replace with file and message)
- Publish a branch: `git push --set-upstream origin <branch>` (replace with branch name)
- Clone a repository: `git clone <repository-url>` (replace with repository URL)
- Checkout to a branch: `git checkout <branch-name>` (replace with branch name)