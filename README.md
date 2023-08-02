# UNDER CONSTRUCTION 
- need access Teams :-( before I can continue, and it seem there is no free Teams 
- does it support emoji?
- how to auth ? (token? / user+pass)
- setup team to allow the app to post message ?
-- add disable config ? use ENV values to auth ?
Info:
````
- https://devblogs.microsoft.com/microsoft365dev/building-go-applications-with-the-microsoft-graph-go-sdk/ 
- https://github.com/infracloudio/msbotbuilder-go/blob/develop/samples/echobot/main.go 
- https://pkg.go.dev/github.com/atc0005/go-teams-notify/v2 
- https://stackoverflow.com/questions/74543384/msgraph-sdk-go-just-responds-with-odata-context-not-the-data-itself 
```

# end-to-teams-go
send message to a configured Microsoft Teams channel

## usage

```
usage: send-to-teams [-h|--help] [-c|--configFile "<value>"]
                     [-m|--message "<value>" [-m|--message "<value>" ...]]
                     [-e|--emoji "<value>"] [-q|--quiet]
                     [-v|--version]

                     Simple script send a message to a Microsoft
                     Teams channel

Arguments:

  -h  --help        Print help information
  -c  --configFile  Configuration file to be use. Default:
                    /usr/local/etc/send-to-teams/config.ini
  -m  --message     Message to be sent between double quotes or single quotes,
                    required
  -e  --emoji       Emoji to use.
  -q  --quiet       quiet mode. Default: false
  -v  --version     Show version
```

# how to build

```
go mod init
go mod tidy
go build -o send-to-teams send-to-teams.go
```
