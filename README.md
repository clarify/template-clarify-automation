# template-clarify-automation

This is a template repo for running [Clarify.io](https://clarify.io) automation routines using the [clarify-go](https://github.com/clarify/clarify-go) SDK. It currently include examples for the following routines:

- `publish`: Automatically publish items from signals by filtering which signals to publish and applying transforms.

## Getting started

This is a _template repository_. To use this code and run it on GitHub, you can start by creating a new (private) repository by clicking "Use this template" button. Next up, you can clone the repository to your local computer and make your own changes.

If your organization does not use GitHub for hosting code, don't worry. Simply create a clone of this repo on your organizations code management server, or just clone it down locally. Just make sure that you DO NOT push reviews regarding your organization _back_ to the [template repository](https://github.com/clarify/template-clarify-automation/). Pulling down changes as the template repository evolve, is still fine.

## Run routines locally

To run this code locally, you need a pair of credentials with access to the _admin_ namespace. For security reasons, do not grant wider permissions than you need. See [our docs](https://docs.clarify.io/developers/quickstart/create-integration) for how to set-up an integration in Clarify and generate new credentials. We recommend that you create a separate integration for using with this repo. For running this locally, we are going to use a _credentials file_. Once you have set-up the credentials, download it and either remember where you placed it, or move it to the root of this repository. You are now ready to start playing with automation routines.

For information on all available commands and options, you can use the `-help` flag:

```sh
- go run . -help          # Help on global flags and available sub-commands.
- go run . publish -help  # Help regarding the publish command.
```

### Run the publish routine

1. Navigate to Clarify, and copy down the integration ID(s) that you want to publish signals from.
2. Edit the file `publish_rules.go`

You are now ready tur run your automation; to see what it's planning to do, run:

```sh
go run . -credentials clarify-credentials.json -v publish -dry-run
```

To run a sub-set of your publish rules, you can specify the rule names.

```sh
go run . -credentials clarify-credentials.json -v publish -dry-run -rules my-rules
```

When you are happy with the results, you can run it without the `-dry-run` flag. You can also skip the `-v` flag if you want less details.

## Run routines in GitHub actions

This repository comes with templates that allows it to run directly in [GitHub Actions][ga]. GitHub Actions is just a very convenient way to run your code without setting up a separate infrastructure to run it in. If you prefer to run your code somewhere else, you can always disable it.

The template is set-up to run the automation routines on:

- New pull-request (using dry-run mode)
- Merge to main (not using dry-run)
- On a fixed schedule (every 24 hours by default)
- On manual trigger

In order to do this, you must:

- In your cloned repository settings, ensure GitHub Actions are enabled (usually on by default).
- Generate username/password credentials for an "automation" integration in Clarify.
- Copy the values and add [secrets][ga-secrets] `CLARIFY_USERNAME` and `CLARIFY_PASSWORD` to the repository.

[ga-secrets]: https://docs.github.com/en/rest/actions/secrets
[ga]: https://github.com/features/actions

## Run routines elsewhere

If you want to run your routines elsewhere, you can build a static binary and deploy it to your preferred destination.
