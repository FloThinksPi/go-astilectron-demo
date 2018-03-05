# Mouse and Keyboard Sharing based on Golang and Electron

!!! This is WIP and not Functional !!!

I dont know a good name yet, maybe go-harmony but anyway, this will be a try to rewrite
a Keyboard and Mouse sharing service between mutliple PCs (mac/linux/win).

This was initiated because of the messy development of https://github.com/symless/synergy-core with
the decision to move to closed source for their GUI https://www.reddit.com/r/linux/comments/72gmua/synergy_20_goes_closed_source/.

### Features

Still Thinking about them aswell as a name

## Development Setup

Check your `$GOPATH` and don't forget to add `$GOPATH/bin` to your `$PATH`.

#### Install Dependencies
```
go get -u github.com/asticode/go-astilectron/...
go get -u github.com/asticode/go-astilectron-bootstrap/...
go get -u github.com/asticode/go-astilectron-bundler/...
go get -u github.com/asticode/go-astilog/...
```
#### Bundle/Compile the Application

1. Go to the root of the git repo
2. run `astilectron-bundler -v`


#### Execute the App

The result is in the `output/<your os>-<your arch>` folder and is waiting for you to test it!

#### Deploy for more Environments

To bundle the app for more environments, add an `environments` key to the bundler configuration (`bundler.json`):

```json
"environments": [
  {"arch": "amd64", "os": "linux"},
  {"arch": "386", "os": "windows"}
]
```

and repeat **Bundle/Compile the Application**.
