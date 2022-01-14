# ✨ Contributing to Telego

Hello stranger 👋

> Thanks for taking the time to contribute to this project and improve it.
> I really appreciate this. 🙂

The following is a set of guidelines for contributing to Telego. These are mostly guidelines, not rules, but still, it
will be great if you followed them. If you feel that something should be changed in this document (or in any other place
of the project) fill free to propose those changes in issue or pull request.

## 🤨 What should I know before I get started?

### Main idea of Telego

This library was initially created to have a one-to-one implementation of
[Telegram Bot API](https://core.telegram.org/bots/api) with a user-friendly interface. Since using bare methods and
types not so easy, helper and utility methods were added, also processing updates just in for loop not so great too, so
handlers were added.

### Packages & folders

Project divided into packages & folders with their own purpose and new contributions should follow this structure:

- `telego` - main functionality with methods, types, and core helper functions
- `telegoapi` - core methods for communication with Telegram
- `telegoutil` - utility methods for quality of life improvements (other packages should not depend on this one)
- `telegohandler` - methods for handling updates in a way similar to `net/http`


- `internal/generator-v2` - a mechanism that is responsible for the generation of code from Telegram docs
- `internal/generator` - old (and crappy) version of generator
- `internal/test` - tests of different parts of the library (should not be considered as proper code, used only for
  basic proof work)


- `examples` - list of usage examples
- `docs` - documentation and other useful assets

### Always releasable

The goal is to have `main` branch always "releasable", that means that while new functional is added or old is updated
no changes should break existing code. If feature was working before update, after update it should work the same or
with expected changed behaviour.

This is main reason why this project doesn't have `dev` branch.

## 🧐 How can I contribute?

### Reporting bugs

You can help the project by reporting bugs or issues that you found, but before creating a bug report, please check the
following:

- Read Telegram [docs](https://core.telegram.org/bots/api) and check if your bug can be related to Telegram and not to
  Telego
- Read Telego docs to ensure that this is not expected behavior

Also, when submitting bugs, please provide as much information as possible.

### Suggesting new features

If you see that something is missing, or you think some functionality can be extended, feel free to propose that. Also,
any part of docs or comments can be improved, and you can create issue or pull request for that.

In case you want some new functionality, adding usage (or even implementation) examples will be great. A full detailed
description will also help a lot.

### Code contribution

Code contribution is also welcomed. You can pick unresolved issue or just create new feature or add documentation,
basically any help will be great. Still note that your code should meet some basic quality and style guidelines
(see below).

How to contribute step by step:

1. Fork repo
2. Clone `git clone https://github.com/<username>/telego.git`
3. Create new branch `git checkout -b my-new-feature`
4. Make your changes, then add them `git add .`
5. Commit `git commit -m "New feature added"`
6. Push `git push origin my-new-feature`
7. Create pull request in Telego repo

How to run tests & linter locally:

- Run tests: `make test`
- Run linter: `make lint`
    - Install linter: `make lint-install`


- Run both tests and linter: `make pre-commit`

To see full usage of [Makefile](../Makefile) use: `make help` or just `make`.

## 🎨 Style guidelines

### Commit messages

No specific requirements, but all commit messages should start with a capital letter, and verbs should be in the past
tense. Message should contain a brief description of what you've done, no need for full text, but just `Fix` won't be
enough.

Bad:

- `fix`
- `added tests`
- `move function`

Good:

- `Fixed function X`
- `Added unit tests for X`
- `Moved X from Y to Z`

### Code style

Your code should meet general Go standards, like camel case naming, capitalized abbreviations, Go style comments, etc.
All described can be read in [Effective Go](https://go.dev/doc/effective_go) and strongly recommended following.

Also, using panic is generally not allowed, but may be used in some cases like initialization of handler with nil handle
function or in unit tests. All errors should be handled (returned to the user or logged).

Documentation of new functionality is essential, and should be added before merging pull request.

### CI testing & linters

Your contribution should pass code quality gates, which means that new code should be covered with unit tests at list by
80%, no new code smells should be added and pass linters.

Unit tests should be independent of each other and if needed use mocks.
