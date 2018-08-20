---
title: "IPInfo"
date: 2018-06-01T23:18:48-07:00
draft: false
weight: 20
---

<img class="screenshot" src="/imgs/modules/ipinfo.png" width="320" height="199" alt="ipinfo screenshot" />

Displays your current IP address information, from ipinfo.io.

**Note:** IPInfo.io has a free-plan rate limit of 1000 requests per day.


## Source Code

```bash
wtf/ipinfo/
```

## Configuration

```yaml
ipinfo:
  colors:
    name: red
    value: white
  enabled: true
  position:
    top: 1
    left: 2
    height: 1
    width: 1
  refreshInterval: 150
```

### Attributes

`colors.name` <br />
The default colour for the row names. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11 color</a> name.

`colors.value` <br />
The default colour for the row values. <br />
Values: Any <a href="https://en.wikipedia.org/wiki/X11_color_names">X11 color</a> name.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed. <br />

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.
