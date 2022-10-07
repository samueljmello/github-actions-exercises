# GitHub Actions Exercises
This repository exists to provide users of GitHub Actions with very simple and easy-to-follow examples, breaking down more about how Actions works, along with nuances of working with the different features included.

## Exercise Prerequisites

### Software Installed
- Git
- A source code editor (VSCode, etc.)

### Environment
- A GitHub repository with admin permissions
- Authentication to the repository (OAUTH credentials or PAT token)
- The repository cloned locally

### Knowledge
- How to clone a GitHub repository
- How to add, commit, and push changes to a GitHub repository
- How to create & publish a branch

### Commands Quick-Reference
- Push: `git push`
- Add & commit: `git add <path-to-file>; git commit -m "<your message>"` (replace with file and message)
- Publish a branch: `git push --set-upstream origin <branch>` (replace with branch name)
- Clone a repository: `git clone <repository-url>` (replace with repository URL)
- Checkout to a branch: `git checkout <branch-name>` (replace with branch name)

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
12. [Continuous integration](./12-Continuous-Integration.md)
13. [Dependabot config](./13-Dependabot-Config.md)
14. Artifacts (under construction)
15. Caching (under construction)
16. Packages (under construction)
17. Using releases to trigger workflows (under construction)
18. Environments (under construction)