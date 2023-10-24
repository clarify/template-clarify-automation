# template-clarify-automation

This is a template repo for running [Clarify.io](https://clarify.io) automation routines using the [clarify-go](https://github.com/clarify/clarify-go) SDK. It currently include examples for the following routines:

- `publish/example`: Automatically publish items from signals by filtering which signals to publish and applying transforms.
- `evaluate/high-temperature`: Detect the high flank when the inside temperature goes above 26 degrees and send a message on Slack.
- `debug/send-to-slack`: Test sending messages to Slack.

## Getting started

This is a _template repository_. To use this code and run it on GitHub, you can start by creating a new (private) repository by clicking "Use this template" button. Next up, you can clone the repository to your local computer and make your own changes.

If your organization does not use GitHub for hosting code, don't worry. Simply create a clone of this repo on your organizations code management server, or just clone it down locally. Just make sure that you DO NOT push reviews regarding your organization _back_ to the [template repository](https://github.com/clarify/template-clarify-automation/), or otherwise make your company code public by accident. Pulling down changes as the template repository evolve, is still fine and recommended. Be aware of breaking changes.

## Run routines locally

To run this code locally, you need a pair of credentials with access to the _admin_ namespace. For security reasons, do not grant wider permissions than you need. See [our docs](https://docs.clarify.io/developers/quickstart/create-integration) for how to set-up an integration in Clarify and generate new credentials. We recommend that you create a separate integration for using with this repo. For running this locally, we are going to use a _credentials file_. Once you have set-up the credentials, download it and either remember where you placed it, or move it to the root of this repository. You are now ready to start playing with automation routines.

For information on all available commands and options, you can use the `-help` flag:

```sh
- go run . -help  # Help on available options for running routines
```

Quick examples:

```sh
- go run . -credentials=credentials.json # Run all routines
- go run . -credentials=credentials.json publish # Run all publish routines
- go run . -credentials=credentials.json publish/example # Run specific routine

```

You can also use `*` to match all routines routine names at a given level.

## Modify routines

All routines in this repo are configured in the `routines.go` file. When forking the repository, you should modify this file to create your own routine tree structure. File names in Go does not contain any semantic meaning, so if you need to organize your routines across multiple files, you are free to do so as well. You can also write completely custom routines by implementing the [Routine](https://pkg.go.dev/github.com/clarify/clarify-go@v0.3.0-pre.1/automation#Routine) interface from the clarify-go automation package. Be sure to update your GitHub actions strict to run the right routines.

In the template repository, then the `publish_transforms.go` contain transforms for item publishing. Transforms allow cleaning up meta-data before publishing your items. You may want to write more.

## Run routines in GitHub actions

This repository comes with templates that allows it to run directly in [GitHub Actions][ga]. GitHub Actions is just a very convenient way to run your code without setting up a separate infrastructure to run it in. If you prefer to run your code somewhere else, you can always disable it.

The template is set-up to run the automation routines on:

- New pull-request (using dry-run mode)
- Merge to main (not using dry-run)
- On a fixed schedule (every 24 hours by default)
- On manual trigger

In order to do this, you must:

- Enable GitHub Actions for your cloned repository.
- Generate username/password credentials for an "automation" integration in Clarify.
- Copy the values and add [secrets][ga-secrets] `CLARIFY_USERNAME` and `CLARIFY_PASSWORD` to the repository.

[ga-secrets]: https://docs.github.com/en/rest/actions/secrets
[ga]: https://github.com/features/actions

If you also want to enable posing messages to slack for the `actionSendToSlack`

## Run routines elsewhere

If you want to run your routines elsewhere, you can build a static binary and deploy it to your preferred destination.
