[![codecov](https://codecov.io/gh/schaermu/gopress/branch/main/graph/badge.svg?token=P83Z3KCN0T)](https://codecov.io/gh/schaermu/gopress)
[![Go Report Card](https://goreportcard.com/badge/github.com/schaermu/gopress)](https://goreportcard.com/report/github.com/schaermu/gopress)

# gopress - a presentation builder CLI for developers

Gopress will enable you to build exciting and modern presentations using [impress.js](https://github.com/impress/impress.js) by doing the thing that feels the most natural to us developers: **coding**.

Creating the content is as natural as writing markdown files, building a presentable version of your presentation is done by using a CLI.
Since everything is stored in code, you can even use version control to manage your presentations!

## First steps
1. Start by creating a directory for your presentations: `mkdir -p my-awesome-slides && cd my-awsome-slides`.
2. (optional) Enable source control: `git init`.
3. Create your first presentation: `gopress init gopress-101`.
4. Start writing your content using GitHub-Flavored markdown inside `gopress-101/slides.md`.
5. Build your presentation and open it in your default browser: `gopress present gopress-101`.

## Reference
wip
