# Opportunities to create Rust-based replacements for classic Linux utilities

A list of popular Linux commands that have already been rewritten in Rust.
The goal of this list is to help you transition to Rust: you can either choose a Linux command and rewrite it yourself, estimating the workload by checking the number of lines of code in its C repository, or find ongoing projects and contribute to them.
It also shows which popular TUIs are commonly used in Rust when rewriting commands from C to Rust.

| Command (C) | Repository (C) | Stars (C) | LoC (C) | Command (Rust) | Repository (Rust) | Stars (Rust) | LoC (Rust) | Alternatives | Search more |
|-------------|----------------|-----------|---------|----------------|-----------------|------------|---------|-------------|-------------|
| cat | [cat](https://github.com/coreutils/coreutils/blob/master/src/cat.c) | 4700 | 695 | bat | [bat](https://github.com/sharkdp/bat) | 54048 | 0 |  | [GitHub](https://github.com/search?l=Rust&q=%22cat%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22cat%22+AND+%22Rust%22) |
| chmod | [chmod](https://github.com/coreutils/coreutils/blob/master/src/chmod.c) | 4700 | 566 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22chmod%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22chmod%22+AND+%22Rust%22) |
| cp | [cp](https://github.com/coreutils/coreutils/blob/master/src/cp.c) | 4700 | 1116 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22cp%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22cp%22+AND+%22Rust%22) |
| echo | [echo](https://github.com/coreutils/coreutils/blob/master/src/echo.c) | 4700 | 247 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22echo%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22echo%22+AND+%22Rust%22) |
| du | [du](https://github.com/coreutils/coreutils/blob/master/src/du.c) | 4700 | 963 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22du%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22du%22+AND+%22Rust%22) |
| head | [head](https://github.com/coreutils/coreutils/blob/master/src/head.c) | 4700 | 930 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22head%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22head%22+AND+%22Rust%22) |
| hostname | [hostname](https://github.com/coreutils/coreutils/blob/master/src/hostname.c) | 4700 | 94 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22hostname%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22hostname%22+AND+%22Rust%22) |
| kill | [kill](https://github.com/coreutils/coreutils/blob/master/src/kill.c) | 4700 | 269 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22kill%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22kill%22+AND+%22Rust%22) |
| ls | [ls](https://github.com/coreutils/coreutils/blob/master/src/ls.c) | 4700 | 4821 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22ls%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22ls%22+AND+%22Rust%22) |
| mkdir | [mkdir](https://github.com/coreutils/coreutils/blob/master/src/mkdir.c) | 4700 | 267 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22mkdir%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22mkdir%22+AND+%22Rust%22) |
| mv | [mv](https://github.com/coreutils/coreutils/blob/master/src/mv.c) | 4700 | 494 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22mv%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22mv%22+AND+%22Rust%22) |
| pwd | [pwd](https://github.com/coreutils/coreutils/blob/master/src/pwd.c) | 4700 | 333 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22pwd%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22pwd%22+AND+%22Rust%22) |
| rm | [rm](https://github.com/coreutils/coreutils/blob/master/src/rm.c) | 4700 | 328 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22rm%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22rm%22+AND+%22Rust%22) |
| sleep | [sleep](https://github.com/coreutils/coreutils/blob/master/src/sleep.c) | 4700 | 122 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22sleep%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22sleep%22+AND+%22Rust%22) |
| tee | [tee](https://github.com/coreutils/coreutils/blob/master/src/tee.c) | 4700 | 294 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22tee%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22tee%22+AND+%22Rust%22) |
| tail | [tail](https://github.com/coreutils/coreutils/blob/master/src/tail.c) | 4700 | 2166 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22tail%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22tail%22+AND+%22Rust%22) |
| touch | [touch](https://github.com/coreutils/coreutils/blob/master/src/touch.c) | 4700 | 375 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22touch%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22touch%22+AND+%22Rust%22) |
| true | [true](https://github.com/coreutils/coreutils/blob/master/src/true.c) | 4700 | 67 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22true%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22true%22+AND+%22Rust%22) |
| yes | [yes](https://github.com/coreutils/coreutils/blob/master/src/yes.c) | 4700 | 113 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22yes%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22yes%22+AND+%22Rust%22) |
| wc | [wc](https://github.com/coreutils/coreutils/blob/master/src/wc.c) | 4700 | 855 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22wc%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22wc%22+AND+%22Rust%22) |
| ffmpeg | [ffmpeg](https://github.com/FFmpeg/FFmpeg) | 52539 | 1538205 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22ffmpeg%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22ffmpeg%22+AND+%22Rust%22) |
| f3 | [f3](https://github.com/AltraMayor/f3) | 2937 | 6368 |  |  |  |  |  | [GitHub](https://github.com/search?l=Rust&q=%22f3%22+language%3ARust&type=repositories) / [Google](https://www.google.com/search?q=site%3Agithub.com+%22f3%22+AND+%22Rust%22) |

Status: ⚠️ Work in progress

The article will be:
- 🇺🇦 First published in Ukrainian on [DOU](https://dou.ua/lenta/tags/Rust/)
- 🌍 Later translated and published in English on [Reddit's r/rust](https://www.reddit.com/r/rust/)

Links will be added here after publication.

## How to Contribute

We welcome contributions! If you would like to add a new command to the list, please follow these rules:

1. The command must be a **well-known command written in C**.
2. There must be an **existing Rust repository** implementing this command with **100+ GitHub stars**.
3. All changes should be made **only in** [`internal/data/commands.go`](https://github.com/doutivity/which-linux-commands-to-rewrite-in-rust/blob/main/internal/data/commands.go).
