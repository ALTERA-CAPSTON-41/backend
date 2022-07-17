# ✨ Contribution Guides
Thanks to all contributors and their mentor that helps to finish up this project. There was a contribution guide to explain contribution rules.

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#commit-convention">Commit Convention</a></li>
    <li><a href="#contributing-guide">Contributing Guide</a></li>
    <li><a href="#appendices">Appendices</a></li>
  </ol>
</details>

## Getting Started
Our source code is developed trough hexagonal architecture that developed by Uncle Bob. We've implemented component as a subset from a feature. Development scopes will be starts on feature development. 

### Directory Structure 
```
- 📂 libs
  - 📂 api-spec
    - 📝 api-spec.yml     <- OpenAPI v3 specification file
  - 📂 dbml
    - 📝 clinic.dbml      <- DBML template file

- 📂 public
  - 📂 api-spec
    - 📝 api-spec.html    <- api-spec.yml viewer

- 📂 src
  - 📂 adapters
    - 📝 ca.go            <- CA adapter mapper
  - 📂 app                <- Features - aka. scopes path
  - 📂 database           <- Database configuration
  - 📂 middleware         <- Middleware path
  - 📂 routes             <- Route mapper
  - 📂 types              <- Public scope used types, like interfaces, enums or types itself 
  - 📂 utils              <- The helpers functions

- 📝 docker-compose.yml   <- Docker 
- 📝 Dockerfile           <- Docker
- 📝 go.mod               <- List of dependencies
- 📝 go.sum               <- List of dependencies
- 📝 main.go              <- Entry file
```

## Commit Convention
This commit conventions was followed angular commit conventions. So that, it can be detected with auto-releaser automatically. 

### Formats
```
<type>(<scope>): <short summary>
  │       │             │
  │       │             └─⫸ Summary in present tense. Not capitalized. No period at the end.
  │       │
  │       └─⫸ Commit Scope: api-spec | dbml | (and all of path name in
                             src/app folder)
  │
  └─⫸ Commit Type: build|ci|docs|feat|fix|perf|refactor|test
```

The `<type>` and `<summary>` fields are mandatory, the `(<scope>)` field is optional.

#### Type
Must be one of the following:

- **build**: Changes that affect the build system or external dependencies (example scopes: go, docker)  
- **ci**: Changes to our CI configuration files and scripts (examples: CircleCi, SauceLabs)  
- **docs**: Documentation only changes  
- **feat**: A new feature  
- **fix**: A bug fix  
- **perf**: A code change that improves performance  
- **refactor**: A code change that neither fixes a bug nor adds a feature  
- **test**: Adding missing tests or correcting existing tests

#### Scopes 
The scopes should be a libs' child folders or src/app child folder:
- from libs
  - `api-spec`
  - `dbml`
- from src/app (aka. features)
  - `account`
  - `admin`
  - `apispec`
  - `dashboard`
  - `doctor`
  - `icd10`
  - `medical_record`
  - `nurse`
  - `patient`
  - `polyclinic`
  - `prescription`
  - `queue`

#### Summary 
Use the summary field to provide a succinct description of the change:
- use the imperative, present tense: "change" not "changed" nor "changes"
- don't capitalize the first letter
- no dot (.) at the end

## Contributing Guide
### Make Changes
#### Changes Locally 
1. Install Git and Golang
2. Clone repo and install dependencies
3. Create a working branch and start develop!

### Changes in Codespace and in the UI
Currently, changes in the codespace or in the UI is not allowed, except on PR change request, you can apply them with following the commit conventions.

### Commit a Changes 
We're recommend you (as a contributor) to commit by context with your work in your working branch. You shouldn't create a huge changes in a commit. Create a commit message descriptively and follow the commit conventions. When you've finished to change, you can publish your working branch and open a Pull Request.

### Pull Request 
When you finished a change, create a pull request (aka. PR):
- Creating a PR Overview is recommended;
- Solve the conflict first by rebasing from the master branch and do a force push;
- Please include a testing report if you working trough feature, bugfix or doing a performance increase;
- You have to request to the other member(s) to review your change, except unneeded reviews to PR, eg. CI/CD, re-style the code, huge refactor or minor documentation changes;
- We may ask your changes before it can be approved, either using suggested changes or pull request coments. You can apply the suggestions trough the UI;
- Mark as resolved if you solve a request changes.

### Merge Branch
- You can merge your PR by yourself since this approved;
- You must use merge PR and change the commit title with what you're doing and followed by PR number, eg. `docs(dbml): change dbml entities to draw out trough the app (#7)`.

## Appendices
### EOL's
This project use `CRLF` type endline or `\r\n` in regex. Because most of our developers are use windows to contribute.

### Development Environments
This project is working well in both linux and windows. We can't test this on mac environment since we're don't have the apple devices.
