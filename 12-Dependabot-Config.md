# Dependabot Config
Within every GitHub repository exists the ability to configure Dependabot. It is a feature that GitHub offers that can remediate out-of-date dependencies by notifying and/or opening pull requests automatically.

To see this in action, we previously used a public action (actions/checkout@v2) which is an older version (see [11-Using-Actions.md](./11-Using-Actions.md) exercise if you haven't completed it) and will enable dependabot for GitHub Actions.

An extensive list of supported dependency ecosystems can be found in the documentation section.

## Prerequisite Exercises
- [11-Using-Actions.md](./11-Using-Actions.md)

## Step 1: Add dependabot config

1. Checkout the **default** branch of your repository
2. Create a new file named `.github/dependabot.yaml`
3. Copy the contents below to the newly created file:

```yaml
version: 2
updates:
  - package-ecosystem: github-actions
    directory: "/"
    schedule:
      interval: daily
    open-pull-requests-limit: 10
```

4. Add, commit, and push your changes to the default branch.
5. Go to your repository, and view the Pull Requests tab to see the pull request Dependabot opens (may take several minutes)

The result will be a pull request with full documentation on what is being changed. Any workflows that trigger on pull request will also execute, validating the changes.

## Documentation
- [Dependabot Config Options](https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file)