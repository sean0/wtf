---
title: "Jira"
date: 2018-05-10T10:44:35-07:00
draft: false
weight: 140
---

<img class="screenshot" src="/imgs/modules/jira.png" width="640" height="188" alt="jira screenshot" />

Displays all Jira issues assigned to you for the specified project.

## Source Code

```bash
wtf/jira/
```

## Keyboard Commands

<span class="caption">Key:</span> `[return]` <br />
<span class="caption">Action:</span> Open the selected issue in the browser.

<span class="caption">Key:</span> `j` <br />
<span class="caption">Action:</span> Select the next item in the list.

<span class="caption">Key:</span> `k` <br />
<span class="caption">Action:</span> Select the previous item in the list.

<span class="caption">Key:</span> `↓` <br />
<span class="caption">Action:</span> Select the next item in the list.

<span class="caption">Key:</span> `↑` <br />
<span class="caption">Action:</span> Select the previous item in the list.

## Configuration

### Single Jira Project

```yaml
jira:
  apiKey: "3276d7155dd9ee27b8b14f8743a408a9"
  colors:
    rows:
      even: "lightblue"
      odd: "white"
  domain: "https://umbrellacorp.atlassian.net"
  email: "chriscummer@me.com"
  enabled: true
  jql: "issueType = Story"
  position:
    top: 4
    left: 1
    height: 1
    width: 2
  project: "ProjectA"
  refreshInterval: 900
  username: "chris.cummer"
  verifyServerCertificate: true
```

### Multiple Jira Projects

If you want to monitor multiple Jira projects, use the following
configuration (note the difference in `project`):

```yaml
jira:
  apiKey: "3276d7155dd9ee27b8b14f8743a408a9"
  colors:
    rows:
      even: "lightblue"
      odd: "white"
  domain: "https://umbrellacorp.atlassian.net"
  email: "chriscummer@me.com"
  enabled: true
  jql: "issueType = Story"
  position:
    top: 4
    left: 1
    height: 1
    width: 2
  project: ["ProjectA", "ProjectB"]
  refreshInterval: 900
  username: "chris.cummer"
  verifyServerCertificate: true
```

### Attributes

`apiKey` <br />
Value: Your <a href="https://confluence.atlassian.com/cloud/api-tokens-938839638.html">Jira API</a> key.

`colors.rows.even` <br />
Define the foreground color for even-numbered rows. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`colors.rows.odd` <br />
Define the foreground color for odd-numbered rows. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11
color name</a>.

`domain` <br />
Your Jira corporate domain. <br />
Values: A valid URI.

`email` <br />
The email address associated with your Jira account. <br />
Values: A valid email address string.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`jql` <br />
_Optional_ <br />
Custom JQL to be appended to the search query. <br />
Values: See <a href="https://confluence.atlassian.com/jiracore/blog/2015/07/search-jira-like-a-boss-with-jql">Search Jira like a boss with JQL</a> for details.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`project` <br />
The Jira project to fetch information for. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`username` <br />
Your Jira username. <br />

`verifyServerCertificate` <br />
_Optional_ <br />
Determines whether or not the server's certificate chain and host name are verified. <br />
Values: `true`, `false`.
