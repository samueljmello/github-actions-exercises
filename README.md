# GitHub Actions Exercises
Have you ever wanted to use GitHub Actions, but aren't really sure where to start? Then this repo is for you! In the [Outline](#outline) below, you will find a list of exercises to perform that will increase your knowledge of Actions with practical application and real-world examples. Most exercises can be performed alone, but a few build on each-other and say so in the prerequisites for the exercise. You will also learn about GitHub features that tie-in to Actions and see how to use them first-hand.

Don't see an exercise you were hoping to find? Open an `issue` on this repo to suggest it.

This work is entirely open source and comes with no support or guarantee, especially as GitHub features change over time.

---

## Exercise Prerequisites

### Software Installed
- [Git](https://git-scm.com/downloads)
- A source code editor ([VSCode](https://code.visualstudio.com/download), etc.)

### Environment
1. An empty (unitialized) GitHub repository that you have admin access to.
2. Authentication to GitHub (OAUTH credentials or PAT token)
3. The repository cloned locally, then the [remote updated](https://docs.github.com/en/get-started/getting-started-with-git/managing-remote-repositories#changing-a-remote-repositorys-url) to your repository from step 1.

### Knowledge
- [How to create a repository on GitHub](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-new-repository)
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
12. [Dependabot config](./13-Dependabot-Config.md)
13. [Continuous integration](./12-Continuous-Integration.md)
14. [Artifacts](./14-Artifacts.md)
15. [Caching](./15-Caching.md)
16. [Packages & continuous delivery](./16-Packages-And-Continuous-Delivery.md)
17. [Environments & continuous deployment](./17-Environments-And-Continuous-Deployment.md)
18. [Releases & Tags](./18-Releases-And-Tags.md)
19. Creating Your Own Actions (under construction)

---

## Files In This Repository
- `./docker_image` contains files pertaining to the Docker image creation in the Continuous Delivery exercise.
- `./golang_app` contains all files pertaining to the Golang application we will build in the Continuous Integration exercise.
- `./golang_replacments` contains files pertaining to the Golang application modification in several exercises.
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