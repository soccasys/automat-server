# Build Automat

A simple automated software build tool.

## Functionality

This tool is designed to make it easy to automate the build of projects
for which the source code is stored in GIT.

### Project Description

Each project must be described in a JSON formatted file
which includes the following information:

* Name of the Project.
* Components, with for each GIT repository:
  - Name of the sub-directory in which to clone/checkout the repository.
  - URL of the GIT repository to use during the GIT clone operation.
  - Revision to be checked-out, typically master, but any valid GIT branch,
    tag or commit ID can be used.
* Build Steps:
  - Description of the step.
  - Directory in which to run this step, relative to the project build root.
  - Command to be run with all its parameters.
  - Environment, with all the step-level environment variables values.
* Environment, with all the project-level environment variables values.

Here is an example of such project file for this tool itself:

        {
            "name": "build-automat",
            "components": {
                "src/github.com/soccasys/build-automat": {
                    "name": "src/github.com/soccasys/build-automat",
                    "url": "https://github.com/soccasys/build-automat.git",
                    "revision": "master"
                },
                "src/github.com/soccasys/builder": {
                    "name": "src/github.com/soccasys/builder",
                    "url": "https://github.com/soccasys/builder.git",
                    "revision": "master"
                }
            },
            "steps": [
                {
                    "description": "Build All",
                    "directory": ".",
                    "command": [
                        "go",
                        "install",
                        "github.com/soccasys/build-automat"
                    ],
                    "env": {}
                }
            ],
            "env": {}
        }
