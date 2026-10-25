- Go agent that polls Github for newer releases
- Also support detecting content changes from a CMS
- Agent needs to be a ready-to-use binary that can be downlaoded into a server, and a YAML file that contains configuration for all applications
- Github does the building, testing, installing, and finally creating an artifact for releases
- Use systemd with a timer to run the agent. It has more benefits.
- Automate systemd installation (Built-in installation commands)¨
- sudo deploy-agent install 

## Releases

Release API https://docs.github.com/en/rest/releases/releases

For example, to retrieve the latest published release:

```
GET /repos/{owner}/{repo}/releases/latest`
```

For a private repository, you must authenticate.
Example using curl

```
curl \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer YOUR_GITHUB_TOKEN" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  https://api.github.com/repos/mycompany/backend/releases/latest
  ```
