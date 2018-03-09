## Back-end Developer Test

### Devcenter Backend Developer Test I

The purpose of this test is not only to quickly gauge an applicant's abilities with writing codes, but also their approach to development.

Applicants may use whatever language they want to achieve the outcome.

## Task

Build a bot that extracts the following from people’s Twitter bio (on public/open accounts), into a Google spreadsheet:

* Twitter profile name 
* Number of followers

Target accounts using either of these criteria:
* Based on hashtags used
* Based on number of followers; Between 1,000 - 50,000

## How to complete the task

1. Fork this repository into your own public repo.

2. Complete the project and commit your work. Send the URL of your own repository to @seun on the Slack here bit.ly/dcs-slack.

## Show your working

If you choose to use build tools to compile your CSS and Javascript (such as SASS of Coffescript) please include the original files as well. You may update this README file outlining the details of what tools you have used.

## Clean code

This fictitious project is part of a larger plan to reuse templates for multiple properties. When authoring your CSS ensure that it is easy for another developer to find and change things such as fonts and colours.


## Good luck!

We look forward to seeing what you can do. Remember, although it is a test, there are no specific right or wrong answers that we are looking for - just do the job as best you can. Any questions - create an issue in the panel on the right (requires a Github account).


## Demo
![screen shot](https://user-images.githubusercontent.com/8668661/33088863-330b4250-ceef-11e7-9e9c-b4fd9ca299d8.gif)


## Implementation

* run `cp config.example.json config.json`.
* Populate `config.json` with values from your twitter app account.
* run `cp client_secret.example.json client_secret.json`.
* populate `client_secret.json` with the values from your service account.
* Change The spreadsheet ID to the desired
* run `go run main.go`

## Update

* Improve on the Spreadsheet Creation to make it accessible to the project owner
* Make sure duplicate user information is not happening
* Track the go-routines to prevent them from overwritting values in cells.
* Find an easier and more scalable ways to retrieve the last row of the specified sheet