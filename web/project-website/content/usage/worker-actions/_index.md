---
title: Worker Actions
---

The Worker Actions menu can be found in the *Workers* tab of Flamenco Manager's
web interface. By default it shows *Choose an action...*, and the *Apply* button
will be disabled until a specific action is chosen.

![Screenshot of the Worker Actions menu in the Flamenco Manager web interface](worker_actions.webp)

The available actions are:

<style>
  sup {
    color: red;
  }
</style>

| Menu Item                                    | Effect                                                                                                                                   |
|----------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|
| Shut Down (after task is finished)           | Wait until the current task is done, then stop the worker. It will exit with status code `0`, indicating a clean shutdown.               |
| Shut Down (immediately)                      | Abort the current task, and return it to the Manager for requeueing. Then stop the worker process like described above.                  |
| Restart (after task is finished)<sup>*</sup> | Wait until the current task is done, then stop the worker. It will exit with the configured status code, indicating a desire to restart. |
| Restart (immediately)<sup>*</sup>            | Abort the current task, and return it to the Manager for requeueing. Then restart the worker process like described above.               |
| Send to Sleep (after task is finished)       | Let the worker sleep after finishing its current task.                                                                                   |
| Send to Sleep (immediately)                  | Let the worker sleep immediately. Its current task is aborted and requeued by the Manager.                                               |
| Wake Up                                      | Wake the worker up. A sleeping worker can take a minute to respond.                                                                      |

<sup>*</sup> The 'Restart' options are only available when the selected worker is marked as 'restartable'. See below.

## Shut Down & Restart actions

Both the 'Shut Down' and 'Restart' actions stop the Worker process.

Shutting down the worker will make it exit succesfully, with status code `0`.

Restarting the worker is only possible if it was started or configured with a
'restart exit code'. This can be done by using the `-restart-exit-status 47`
commandline option, or by settings this code in the [worker config file][wconfig].
Requesting a worker restart will make it exit with the configured status code.

It is up to the process management system (for example [systemd][systemd]) to
respond to these exit status code correctly. Here is an example systemd service
unit file that shows how to set this up on Linux:

```systemd
[Unit]
Description=Flamenco Worker connecting to Manager on localhost
Documentation=https://localhost:8080/
After=network.target

[Service]
Type=simple
CPUSchedulingPolicy=idle
Nice=19

WorkingDirectory=/home/flamenco
# Tell the Worker that it should exit with status code 47 in order to restart.
ExecStart=/home/flamenco/flamenco-worker -manager http://localhost:8080/ -restart-exit-code 47

User=flamenco
Group=flamenco

# Make systemd restart the service on exit code 47, as well as
# 'failure' codes (such as hard crashes).
RestartForceExitStatus=47
Restart=on-failure

EnvironmentFile=-/etc/default/locale

[Install]
WantedBy=multi-user.target
```

[wconfig]: {{< ref "/usage/worker-configuration" >}}
[systemd]: https://systemd.io/

{{< hint type=note >}}
The 'Shut Down' and 'Restart' actions only relate to the Flamenco Worker process. They do **not** shut down or restart the computer itself.
{{< /hint >}}
