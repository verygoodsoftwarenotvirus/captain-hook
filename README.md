# Captain-Hook
I wanted to build a simple server in Go that would respond to Github webhooks and run some shell scripts, but I ran into the issue of how to authenticate the requests so evil (Russian/Chinese/Idahoan) hackers didn't constantly trigger builds via the route (if they were to discover it).

When I finally figured out how to do it, I decided I should make this public so nobody else has to go through the pain of figuring it out.

So here it is, already-figured-out Github webhook authentication in Golang.

Huge thanks to [SimonWaldherr](https://github.com/SimonWaldherr/) for [this file](https://github.com/SimonWaldherr/golang-examples/blob/master/beginner/hashing.go)
