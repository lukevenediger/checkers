# Contribution Guidelines

We welcome all contributions to this project. Please use the guidelines below to set up your development environment.

## Prerequisites

The following tools are required to contribute to this project:

* [Go](https://golang.org/dl/)
* [Docker](https://docs.docker.com/get-docker/)
* [pre-commit](https://pre-commit.com/)
* [make](https://www.gnu.org/software/make/)

## Pre-Commit Hooks

Our pre-commit hooks help maintain code quality and consistency by performing the following actions:
* Formatting all Go files with `go fmt`.
* Removing trailing whitespace.
* Standardizing end-of-file (EOF) newlines.
* Enforcing the Conventional Commit message standard.

### Conventional Commit Messages

We follow a lightweight structure for commit messages using the Conventional Commits standard. The format is `<type>(<scope>): <description>`, where `type` can be one of `feat`, `fix`, `docs`, `style`, `refactor`, `test`, or `chore`.

For more details, see the [Conventional Commits Quickstart Guide](https://www.conventionalcommits.org/en/v1.0.0/#summary).

### Installing Pre-Commit

1. Install [pre-commit](https://pre-commit.com/).
2. Apply the pre-commit and commit message hooks to this repository:
   ```sh
   pre-commit install
   pre-commit install --hook-type commit-msg
   ```

# Working with Protobuf and gRPC
All RPC services and messages are defined in `proto/lukevenediger/checkers/` and used as part of the gRPC implementation.
Please ensure that docker is installed and running. The protoc compiler runs in a docker container to ensure consistent results

Regenerate the gRPC code after making changes by running:

```sh
make proto-gen
```
