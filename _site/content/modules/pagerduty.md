---
title: "PagerDuty"
date: 2018-08-20T20:53:40+02:00
draft: false
weight: 170
---

<img class="screenshot" src="/imgs/modules/pagerduty.png" width="320" height="389" alt="pagerduty screenshot" />

Connects to the PagerDuty API and displays who's currently on call.

## Source Code

```bash
wtf/pagerduty/
```

## Configuration

```yaml
pagerduty:
  apiKey: "gNzSe5ABCDwe1234jeej"
  enabled: true
  escalationPolicyIDs: ["JU5U5CO", "PZAJZ1Z"]
  position:
    top: 2
    left: 0
    height: 2
    width: 2
  refreshInterval: 300
  userIDs: ["PAI7I3U"]
```

### Attributes

`apiKey` < br />
Value: Your <a href="https://v2.developer.pagerduty.com/docs/authentication">PagerDuty API</a> token.

`enabled` <br />
Whether or not this module is executed and if its data is displayed onscreen. <br />
Values: `true`, `false`.

`escalationPolicyIDs` <br />
A filter of the escalation policies that will be displayed. <br />
Value: A list of Escalation Policy IDs.

`position` <br />
Where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`userIDs` <br />
A filter of the users that will be displayed. <br />
Value: A list of User IDs.
