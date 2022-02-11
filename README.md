# Auto Slack Check-in

<p align="center">
  <a href="https://github.com/mhughdo/auto-slack-checkin/actions/workflows/release.yml/badge.svg">
    <img src="https://github.com/mhughdo/auto-slack-checkin/actions/workflows/release.yml/badge.svg" alt="Build Status" />
  </a>
  <a href="https://godoc.org/github.com/mhughdo/auto-slack-checkin">
    <img src="https://godoc.org/github.com/mhughdo/auto-slack-checkin?status.svg" alt="GoDoc" />
  </a>
  <a href="https://goreportcard.com/report/github.com/mhughdo/auto-slack-checkin">
    <img src="https://goreportcard.com/badge/github.com/mhughdo/auto-slack-checkin" alt="Go Report Card" />
  </a>
</p>

```
auto-slack-checkin is a command line interface (CLI) that sends check-in or check-out message for you while you are sleeping 🛌.

Usage:
  auto-slack-checkin [command]

Available Commands:
  config          Show all configs
  cron            Run a cron job that automatically sends a message at a specified time.
  config set      Set configs.

Flags:
  --help              Get helo about any command (work with all commands)

Use "auto-slack-checkin [command] -h" for more information about a command.
```

- [Auto Slack Check-in](#auto-slack-check-in)
  - [Installing `auto-slack-checkin`](#installing-auto-slack-checkin)
    - [Downloading a Release from GitHub](#downloading-a-release-from-github)
  - [Configuring Values](#configuring--values)
    - [How to generate cron expression: https://crontab.guru/](#how-to-generate-cron-expression-httpscrontabguru)
    - [How to get token, cookie, channel-id](#how-to-get-token-cookie-channel-id)
      - [Token & Cookie](#token--cookie)
      - [Channel id](#channel-id)
  - [Examples](#examples)

## Installing `auto-slack-checkin`

### Downloading a Release from GitHub

Visit the [Releases page](https://github.com/mhughdo/auto-slack-checkin/releases) for the
[`auto-slack-checkin` GitHub project](https://github.com/mhughdo/auto-slack-checkin), and find the appropriate archive
for your operating system and architecture. Download the archive from from your browser or copy its URL and retrieve it
to your home directory with `wget` or `curl`.

For example, with `wget`:

```
cd ~
wget https://github.com/mhughdo/auto-slack-checkin/releases/download/v<version>/auto-slack-checkin_<version>_linux_amd64.tar.gz
```

Or with `curl`:

```
cd ~
curl -OL https://github.com/mhughdo/auto-slack-checkin/releases/download/v<version>/auto-slack-checkin_<version>_linux_amd64.tar.gz
```

Extract the binary:

```
tar xf ~/auto-slack-checkin_<version>_linux_amd64.tar.gz
```

where `<version>` is the full semantic version, e.g., `1.1.2`.

Window version is not supported yet.

## Configuring Values

The `auto-slack-checkin` configuration file is stored at \$HOME/.auto-slack-checkin.json

> Note: You need to run any command, e.g. ./auto-slack-checkin config to init config file.

You can directly change the config file but make sure file format stays the same.

Save and close the file. The next time you use `auto-slack-checkin`, the new values you set will be in effect.

You need to set `token`, `cookie`, `message`, `channel-id` for CLI to work. `cron-expr` is set to "0 8 \* \* \*" by
default which is everyday at 8AM. Full command:

```
./auto-slack-checkin config set --cron-expr '0 8 * * *' --token xoxc-3058611573620-3056323014531-3086306920976-dacaa15d71a468bae5380db258aae5dbb08d2996f119d665a7c73bd7fb19550c --cookie d=WjrniNOtYeO7vSHo%2BDvBzV2iKRhz%2FbnvEDLB7Nf%2BQk15FHNpcZjVsa0ZGqoFnQcLrLklTB9EvxtkhPRs5vJbWso9%2F6cHyVUqKD%2Biuytj9g7W2eDD7d4vUzspyX9ABroxqin4Qbab7qAi4%2BEF3YtYhWfvU6%2FaNEmgK00EadzfGxz7EB92DV2ZEkPtKw%3D%3D --message "Helloooo" --channel-id C032C1GC1Q8

```

### How to generate cron expression: https://crontab.guru/

### How to get token, cookie, channel-id

1. Open workspace in the browser.
2. Open DevTools by pressing Control+Shift+J or Command+Option+J (Mac). The Console panel opens.

#### Token & Cookie

1. Open a channel that you can freely send anything.
2. Send a random message.
3. Click the Network tab in the DevTools.
4. Find the request that contains `chat.postMessage`
5. Click on the request
6. Click the Payload tab, scroll down a bit and you will see token:
   ![image](https://user-images.githubusercontent.com/15611134/153596766-ca46d933-2604-4549-be8c-70288aedc172.png)
7. Click the Headers tab, scroll down a bit. At the Request Headers section, cookie being sent by request should be
   there. Copying only the cookie whose key is `d` is enough.
   ![image](https://user-images.githubusercontent.com/15611134/153597140-fb29c9aa-b31c-4164-a63d-abf89137ebc2.png)

#### Channel id

You can easily find the channel id by clicking on the channel you want to take id and looking at the URL. The
`channel-id` should be the last path parameter.

![image](https://user-images.githubusercontent.com/15611134/153597499-2fb50ec4-1bba-4051-ace6-78626bcf7000.png)

## Examples

Below are a few common usage examples.

- Run the CLI

```
./auto-slack-checkin cron
```

- Run the CLI in the background

```
nohup ./auto-slack-checkin cron &
```

- Show all configs

```
./auto-slack-checkin config
```

- Set config

```
./auto-slack-checkin config set --token <value> --cookie <value>
```
