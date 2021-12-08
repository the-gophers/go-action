# GitHub Action Using Go
This is a starting point for a GitHub Action based in Go. This repo provides all the structure needed to build
a robust GitHub action using Go and following action development best practices.

## Getting Started
This is a GitHub template repo, so when you click "Use this template", it will create a new copy of this
template in your org or personal repo of choice. Once you have created a repo from this template, you
should be able to clone and navigate to the root of the repository.

### First Build
From the root of your repo, you should be able to run the following to build and test the Go action.
```shell script
go build ./...
go test ./...
```

### What's in Here?

```
.
├── Dockerfile
├── LICENSE
├── README.md
├── action.yaml
├── go.mod
├── go.sum
├── main.go
└── pkg
    └── tweeter
        ├── tweeter.go
        └── tweeter_test.go
```

#### [action.yaml](./action.yaml)
The `action.yml` file contains the metadata which describes our action. This includes, but is not limited
to the following.
- name, description and author
- inputs
- outputs
- branding
- runs

You will see an example structure already exists. The example executes the Dockerfile and provides to it
the arguments described in the `runs` section. We map the sample input to the arguments of the Dockerfile.

By setting `runs.using: docker` we are telling the Actions runtime to execute the Dockerfile when the
Action is used in a workflow.

By setting `runs.image: Dockerfile` we are telling the Actions runtime to build and then execute the
Dockerfile at the entrypoint defined in the Dockerfile. The build for the Dockerfile will happen
each time the Action is executed, which can take a considerable amount of time depending on how long
it takes to build your Go code. Later, we'll change this to a pre-built image for optimization. 

#### [Dockerfile](./Dockerfile)
This Dockerfile should look relatively familiar to folks who use containers to build Go code. We create a
builder intermediate image based on Go 1.15.2, pull in the source, and build the application. After the
application has been built, the statically linked binary is copied into a thin image, which results in 
an image of roughly 8 MB.

#### [go.mod](./go.mod) / [go.sum](./go.sum)
Go module definition which will need to be updated with the name of your module.

#### [main.go](./main.go)
This is the Go entrypoint for your GitHub Action. It is a simple command line application which can be
executed outside the context of the Action by running the following. This is where you will add your Go code for your 
Action.

#### [.github/workflows/tweeter-automation.yaml](.github/workflows/tweeter-automation.yaml)
This is the continuous integration / CLI release automation. The release workflow defined in this automation will
is triggered by tags shaped like `v1.2.3`, and will create a GitHub release for a pushed tag. The release will 
use the [./github/release.yml](.github/release.yml) to automatically generate structured release notes based on
pull request tags.

#### [.github/workflows/action-version.yaml](./.github/workflows/action-version.yaml)
This action triggers when a new release is published. Upon publication of the new release, this will tag the repository
with the shortened major semantic version pointing at the highest semantic version within that major version. For 
example, `v1.2.2` and `v1` point to the same ref `xyz`, a new tag is introduced `v1.2.3` which points to ref `abc`, this 
action will move the `v1` tag to point to `abc` rather than `xyz`. This enables consumers of an action to take a 
"floating" major semantic version dependency, like `my-action@v1`. 

#### [.github/workflows/image-release.yaml](./.github/workflows/image-release.yaml)
This action runs on tags shaped like `image-v1.2.3`, and will build and push a container image to the
GitHub Container Registry.

This action is super useful for optimizing the execution time of your action. By pre-building the
image used in the Action, each invocation of your action can reference the image and not have to
rebuild it for each invocation.

Once you push your first image you will also need to update the Container Registry to allow public access.

## Lab Video
TODO: record and post the first lab walking through creation, execution and optimization

## Lab Instructions
1) Click on "Use this template" on https://github.com/the-gophers/go-action, and create a repo of your own. I'm going 
   to call mine "templated-action", make it public, and click "Create repository from template".
2) Clone your newly crated repo
3) Run `go run . -h`.
5) Run
    ```shell script
    $ go run . --dryRun --message hello
    ::set-output name=sentMessage::hello
    ```
6) Checkout a new branch and let's make this our own GitHub Action. Run `git checkout -b my-action`
7) Update `./action.yml` name, description, and author to something reasonable. The `name` field needs to be
   unique to others in the store.
8) When you are done with your changes, commit them, push your branch to GitHub, and open a pull request. In the PR,
   you should see the CI action run and complete successfully. LGTM! Let's merge these changes. Click the
   "Merge pull request" button, then delete the branch.
9) Check out `main` and pull down the latest changes from GitHub (`git pull`).
11) In `test-repo` click on Actions and run `test-action` with the inputs you desire. Navigate the UI to the running
    action and see that it built the action, built the Dockerfile and executed the entrypoint Go application. Also note
    how long it took to run the action. **Using a Dockerfile will cause it to rebuild that image EACH time the action
    runs!**. We can do better than that. More ahead.
12) Let's tag our first release (`git tag v1.0.0`) and push the tag
    (`git push origin v1.0.0`). This should create our first release in GitHub via the `release action` workflow.
13) Navigate to the `v1.0.0` release and click edit. Within the release edit page, you should see "Publish this Action to the GitHub Marketplace".
    If you check that box, your action will now be publicly advertised to all of GitHub!
14) **PSA:** The rest of this is optional. If you don't care about your action going fast, stop right here.
15) Now we are going to make this **FAST** by pre-baking our container image. Go back to `templated-action` and edit
    `./github/workflows/release-image.yml`. Change `docker.pkg.github.com/owner/` to use your repo owner for `owner`.
    Commit and push the changes.
17) Now tag the repo with `git tag image-v1.0.0` and then push the tag `git push origin image-v1.0.0`. This will
    kick off the image release build.
18) Replace `image: Dockerfile` with `image: docker://ghcr.io/your-repo/your-image:1.0.0` replacing the repo and image name.
    Commit the changes and tag a new release of the Action as done in #12.
19) Rerun the continuous integration and see how much faster the action runs now that it doesn't have to rebuild
    the container image each time.
    


## Contributions
Always welcome! Please open a PR or an issue, and remember to follow the [Gopher Code of Conduct](https://www.gophercon.com/page/1475132/code-of-conduct).
