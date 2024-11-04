# Contribution Guidelines
* we welcome all contributions to this project
* please use the information below to help configure your local development environment

## Pre-Commit Hooks
* We use pre-commit hooks to perform the following actions:
  * format all go files with `go fmt`
  * fix trailing whitespace
  * standardise EOF newlines for all files
  * enforce the Conventional Commit message standard

### Conventional Commit Messages
* This tool enables a lightweight way to document each commit message
* The format for a commit message is typically `<type>(<scope>): <description>`
  * where type is `feat`, `fix`, `docs`, `style`, `refactor`, `test`, or `chore`
* read more about the message format in the [Conventional Commits Quickstart Guide](https://www.conventionalcommits.org/en/v1.0.0/#summary)

### Installing Pre-Commit
* First you must install [pre-commit](https://pre-commit.com/)
* Then apply the pre-commit and pre-commit-msg hooks for this repo
  * `pre-commit install`
  * `pre-commit install --hook-type commit-msg`

# Working with Protobuf and gRPC
* All RPC services and messages are defined in `proto/lukevenediger/checkers/`
* Regenerate the gRPC code after making changes by running
  * `make proto-gen`
* Please ensure that docker is installed and running
    * The protoc compiler runs in a docker container to ensure consistent results
