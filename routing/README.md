routing
=======

[![Build Status](http://img.shields.io/travis/brsuite/lnd.svg)](https://travis-ci.org/brsuite/lnd) 
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/brsuite/lnd/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/brsuite/lnd/routing)

The routing package implements authentication+validation of channel
announcements, pruning of the channel graph, path finding within the network,
sending outgoing payments into the network and synchronizing new peers to our
channel graph state.

## Installation and Updating

```shell
⛰  go get -u github.com/brsuite/lnd/routing
```
