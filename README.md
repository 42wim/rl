# rl

Rate limit messages from stdin. (keep or drop those exceeding the limit)

# Use case
* When sending data to mattermost (or slack), it's convenient to limit the amount of data that can be sent.

E.g. when using [mattertee](https://github.com/42wim/matterstuff/tree/master/mattertee) to send output to mattermost, you can now just pipe it through rl so that your channels aren't being flooded.

* Tailing -f a very quickly increasing logfile when in tmux remotely ;)

# Installing
## Binaries
Binaries can be found [here] (https://github.com/42wim/rl/releases/)

## Building
Go 1.6+ is required. Make sure you have [Go](https://golang.org/doc/install) properly installed, including setting up your [GOPATH] (https://golang.org/doc/code.html#GOPATH
)

```
cd $GOPATH
go get github.com/42wim/rl
```

You should now have rl binary in the bin directory:

```
$ ls bin/
rl
```

# Usage
```
Usage of ./rl:
  -k    keep the messages instead of dropping them
  -r int
        limit to r messages per second (drops those exceeding the limit) (default 5)


(the number of dropped messages will be sent to stderr, when not using the -k switch)
```


# Example
```
journalctl -f | rl -k -r 5
```
