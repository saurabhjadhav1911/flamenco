---
title: SocketIO
weight: 50
---

[SocketIO v2](https://socket.io/docs/v2/) is used for sending updates from
Flamenco Manager to the web frontend. Version 2 of the protocol was chosen,
because that has a mature Go server implementation readily available.

SocketIO messages have an *event name* and *room name*.

- **Web interface clients** send messages to the server with just an *event
  name*. These are received in handlers set up by
  `internal/manager/webupdates/webupdates.go`, function
  `registerSIOEventHandlers()`.
- **Manager** typically sends to all clients in a specific *room*. Which client
  has joined which room is determined by the Manager as well. By default every
  client joins the "job updates" and "chat" rooms. This is done in the
  `OnConnection` handler defined in `registerSIOEventHandlers()`.  Clients can
  send messages to the Manager to change which rooms they are in.
- Received messages (regardless of by whom) are handled based only on their
  *event name*. The *room name* only determines *which client* receives those
  messages.

## Technical Details

The following files & directories are relevant to the SocketIO broadcasting
system on the Manager/backend side:

`internal/manager/webupdates`
: package for the SocketIO broadcasting system

`internal/manager/webupdates/sio_rooms.go`
: contains the list of predefined SocketIO *rooms* and *event types*. Note that
 there are more rooms than listed in that file; there are  dynamic room name
 like `job-fa48930a-105c-4125-a7f7-0aa1651dcd57` that cannot be listed there as
 constants.

`internal/manager/webupdates/job_updates.go`
: sending job-related updates.

`internal/manager/webupdates/worker_updates.go`
: sending worker-related updates.

For a relatively simple example of a job update broadcast, see
`func (f *Flamenco) SetJobPriority(...)` in `internal/manager/api_impl/jobs.go`.
