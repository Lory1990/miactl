# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## Changed

- remove dependency from `golang.org/x/exp`
- better memory allocation for slices
- update uuid to v1.6.0
- update oauth2 to v0.17.0
- marketplace apply: fix output column headers ID to correct OBJECT ID

## [0.12.1] - 2024-02-29

### Changed

- update go version to 1.22.0

### Fixed

- fixed open browser feature on Linux
- fixed typo in flag parameter for triggering a new release

## [0.12.0] - 2024-02-05

### Added

- `project iam list` command
- `project iam edit` command
- `project iam remove-role` command
- `--public` flag to `marketplace list` command

### Changed

- update go version to 1.21.6
- update exp to v0.0.0-20240112132812-db7319d0e0e3
- update oauth2 to v0.16.0
- update sync to v0.6.0

### Fixed

- fixed typos

## [0.11.0] - 2024-01-15

### BREAKING

- move serviceaccount commads under `company iam add`

### Added

- `company iam add user` command
- `company iam edit user` command
- `company iam add group` command
- `company iam add group-member` command
- `company iam edit serviceaccount` command
- `company iam edit group` command
- `company iam remove user` command
- `company iam remove group` command
- `company iam remove serviceaccount` command
- `company iam remove group-member` command

### Changed

- the company iam list commands now return the id of the entities as well
- update go version to 1.21.5
- update logr to v1.4.1
- update uuid to 1.5.0
- update exp to v0.0.0-20231219180239-dc181d75b848

### Fixed

- remove conflicting shortand flag `-v` from `miactl marketplace delete` command
- creation of basic auth service account

## [0.10.0] - 2023-12-20

### BREAKING

- `miactl marketplace delete` does not accept anymore the id as argument, it should be provided to the flag `--object-id`
- `miactl marketplace get` does not accept anymore the id as argument, it should be provided to the flag `--object-id`

### Added

- `company iam list` command
- `company iam list users` command
- `company iam list groups` command
- `company iam list serviceaccounts` command
- marketplace apply: add version to multipart request metadata if present in the item
- add command `miactl marketplace list-versions`
- mark as alpha the features:
  - `marketplace list-versions`
  - `marketplace get` with `--itemId` and `--version` flags
  - `marketplace delete` with `--itemId` and `--version` flags
- verbose log for all the remote HTTP requests

### Changed

- help text of version command
- marketplace apply: consider both itemId and version for duplicate check
- marketplace apply: silently ignore files with invalid extension rather than throwing error
- update go version to 1.21.4
- update oauth2 to v0.15.0
- update exp to v0.0.0-20231127185646-65229373498e
- modified `miactl marketplace delete` command to accept either the `objectId` or a `companyId`-`itemId`-`version`
	tuple that identifies the item to be deleted.
- modified `miactl marketplace get` command to accept either the `objectId` or a `companyId`-`itemId`-`version` tuple
	that identifies the item to be retrieved.

### Fixed

- `runtime logs` now is working correctly for pods with more than one container

## [0.9.0] - 2023-11-15

### Added

- runtime log command

### Changed

- marketplace apply command: item id in output data
- removed beta warning for marketplace commands
- the environment filter is now passed as flag and not as parameter for runtime commands
- update go version to 1.21.3
- update uuid to v1.4.0
- update oauth2 to v0.14.0
- update exp to v0.0.0-20231108232855-2478ac86f678
- update cobra to v1.8.0
- update sync to v0.5.0
- update text to v0.14.0

## [0.8.0] - 2023-10-10

### Changed

- update go version to 1.21.1
- update mergo to v1.0.0
- update exp to v0.0.0-20230905200255-921286631fa9
- update oauth2 to v0.12.0
- update uuid to v1.3.1

### Added

- marketplace list command
- marketplace get command
- environment list command
- runtime list resources command
- runtime create job command
- events command
- version command

## [0.7.0] - 2023-06-26

### Added

- create service account with jwt authentication

## [0.6.1] - 2023-06-12

### Added

- support for authenticate API calls with service account
- `context auth` command for saving the service account in the config file

### Fixed

- error during saving context configuration in particular cases

## [0.6.0] - 2023-06-08

### Changed

- reworked the http connection to the remote servers to allow support for custom installations
- **breaking** new configuration file structure

### Added

- creation of service account for your company

### Fixed

- if you use a url that ends with `/` as endpoint, the remote calls are not broken anymore

## [0.5.0] - 2023-05-10

### Changed

- remove default value for revision flag in the deploy command

### Added

- new command for listing the user companies

## [0.4.0] - 2023-04-07

### Changed

- complete rework of cli
- new login that uses the OIDC flow via the user browser
- new deploy command, now it will wait for the pipeline to finish
- new project list command
- the user now can create contexts for multiple scenarios and consoles

## [0.3.1] - 2022-01-21

### Fixed

- fix warning installing with brew (issue 55)

## [0.3.0] - 2021-06-22

### Added

- login command with username and password
- new command to trigger a new deploy per environment
- get status of a pipeline release
- insecure flag to skip check on certificate authority
- flag to use custom certificate authority
- add support to go v1.16

### Update

- upgrade dependencies

### Changed

- drop support to go v1.13 and v.14

## [0.2.0] - 2020-04-27

- get history of deployments for a specific project environment
- add get projects command

## [0.1.0] - 2020-04-14

- create cli sdk
- create cli renderer

[unreleased]: https://github.com/mia-platform/miactl/compare/v0.12.1...HEAD
[0.12.1]: https://github.com/mia-platform/miactl/compare/v0.12.0...v0.12.1
[0.12.0]: https://github.com/mia-platform/miactl/compare/v0.11.0...v0.12.0
[0.11.0]: https://github.com/mia-platform/miactl/compare/v0.10.0...v0.11.0
[0.10.0]: https://github.com/mia-platform/miactl/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/mia-platform/miactl/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/mia-platform/miactl/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/mia-platform/miactl/compare/v0.6.1...v0.7.0
[0.6.1]: https://github.com/mia-platform/miactl/compare/v0.6.0...v0.6.1
[0.6.0]: https://github.com/mia-platform/miactl/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/mia-platform/miactl/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/mia-platform/miactl/compare/v0.3.1...v0.4.0
[0.3.1]: https://github.com/mia-platform/miactl/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/mia-platform/miactl/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/mia-platform/miactl/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/mia-platform/miactl/releases/tag/v0.1.0
