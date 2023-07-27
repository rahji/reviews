# Reviews

This project is only useful for the FSU Department of Art Graduate Director. It is part of a process that takes place after a day of MFA Formal Reviews. Reviews happen three times per semester - once for each cohort of students. The post-review process looks like this:

1. Each faculty member submits their evaluations through a Qualtrics survey. They were previously sent invitation URLs: one for each student that they will review.
2. The results of all surveys are exported as a CSV file from Qualtrics.
3. The Grad Director runs the interactive [studentpdfs.sh](bin/studentpdfs.sh) shell script, which uses the `reviews convert` command and [pandoc](https://pandoc.org/) to create a collection of PDFs from the data: one for each student/faculty pair.
4. The Grad Director also runs the interactive [directorpdfs.sh](bin/directorpdfs.sh) shell script, which uses the `reviews director` command and pandoc to create a collection of PDFs. Each PDF includes all comments (including comments meant only for the Grad Director) for a single student.

# Why

Exporting PDFs in batch is a feature that Qualtrics does not have - it lets you export one PDF at a time and doesn't let you choose what's in it. We need this capability so we can return the results of the evaluations to the students after each of their semesterly reviews. 

## Installation

Since this project is not meant for wide consumption, I have not done work to make it easy to install. You'll need these things before you get started:

* A unix-like operating system. 
  * On Windows, this means you'll need to install WSL2. You might as well install [Terminal](https://apps.microsoft.com/store/detail/windows-terminal/9N0DX20HK701) while you're at it.
  * On a Mac, you will be able to do most of this from the Terminal app. In WSL (Ubuntu) you'll be able to install most things using `apt`. On a Mac, you'll have to figure out which package manager to use.
* [git](https://git-scm.com/downloads), which you may already have. You might as well [configure your name and email address](https://git-scm.com/book/en/v2/Getting-Started-First-Time-Git-Setup) if you've never used it before.
* [go](https://go.dev/dl/)
* texlive-full (`apt-get install texlive-full` on WSL/Ubuntu)
* [pandoc](https://pandoc.org)
* [Eisvogel pandoc template](https://github.com/Wandmalfarbe/pandoc-latex-template/releases). Move the template file to your pandoc templates folder (see the Eisvogel instructions for more info).
* [Glow](https://github.com/charmbracelet/glow) is optional. It lets you read markdown files from the command-line.

Follow these steps to install the `reviews` project:

1. Clone this repo
2. Run `go install github.com/rahji/reviews`

## Usage

```bash
reviews convert
```

> PS: [Glow](https://github.com/charmbracelet/glow) is a fun way to read Markdown documents in the terminal.

## Running pandoc

To convert a single markdown file to PDF, use the command:

```bash
pandoc file.md -o file.pdf -V geometry:landscape --from markdown --template eisvogel --columns=1000
```

Converting a folder of markdown files to PDFs will look something like this:

```bash
for f in `ls *md`; do \
pandoc $f -o `basename --suffix=.pdf $f`.pdf \
        --from markdown --template eisvogel --columns=1000; \
done
```
