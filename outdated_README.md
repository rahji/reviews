# Reviews

*This is the old version of the README that suggests running the interactive shell scripts. The new method uses Docker instead and should be a lot simpler.*

This project is only useful for the FSU Department of Art Graduate Director. It is part of a process that takes place after a day of MFA Formal Reviews. Reviews happen three times per semester - once for each cohort of students. The post-review process looks like this:

1. Each faculty member submits their evaluations through a Qualtrics survey. They were previously sent invitation URLs: one for each student that they will review.
2. The results of all surveys are exported as a CSV file from Qualtrics.
3. The Grad Director runs the interactive [studentpdfs.sh](bin/studentpdfs.sh) shell script, which uses the `reviews convert` command and [pandoc](https://pandoc.org/) to create a collection of PDFs from the data: one for each student/faculty pair.
4. The Grad Director also runs the interactive [directorpdfs.sh](bin/directorpdfs.sh) shell script, which uses the `reviews director` command and pandoc to create a collection of PDFs. Each PDF includes all comments for a single student, including comments directed to the Grad Director and Grad Advisor & Coordinator.

## Why

As in, "Why did you spend a week of your life doing this?"

Exporting PDFs in batch is a feature that Qualtrics does not have - it lets you export one PDF at a time and doesn't let you choose what's in it. We need this capability so we can return the results of the evaluations to the students after each of their semesterly reviews. 

## Installation

You'll need these things before you get started:

* A unix-like operating system. 
  * On Windows, this means you'll need to install WSL2. You might as well install [Terminal](https://apps.microsoft.com/store/detail/windows-terminal/9N0DX20HK701) while you're at it.
  * I haven't tried any of this on a Mac, but there's no reason it shouldn't work.
* [git](https://git-scm.com/downloads), which you may already have. You might as well [configure your name and email address](https://git-scm.com/book/en/v2/Getting-Started-First-Time-Git-Setup) if you've never used it before.
* [go](https://go.dev/dl/)
* texlive-full (`apt-get install texlive-full` on WSL/Ubuntu; you'd have to use a different package manager on MacOS)
* [pandoc](https://pandoc.org)
* [Eisvogel pandoc template](https://github.com/Wandmalfarbe/pandoc-latex-template/releases). (The template is also included in this repo, if you don't want to download it.) Move the template file to your pandoc templates folder (`~/.pandoc/templates/` on Linux).
* [gum](https://github.com/charmbracelet/gum)

To install the `reviews` project, run the following command from your terminal application:

```bash
go install github.com/rahji/reviews
```

That installs the binary, but you'll also want to clone this repo so you have the templates, scripts, etc. (see below)

## Usage

Once all of the surveys have been submitted and the CSV file has been exported from Qualtrics, copy it to the repository's root folder. Change to that root folder and run the command `bin/studentpdfs.sh`. You'll be asked for the following, by the script:

* The CSV file you want to convert to PDFs
* The markdown template for creating student PDFs
* The folder where you want to save the markdown files
* The folder where you want to save the PDFs

To get PDFs of all of the comments left by faculty (one PDF per student), run the command `bin/directorpdfs.sh` from the repository's root folder. It will ask for this info:

* The CSV file you want to convert to PDFs
* The markdown template for creating grad director PDFs
* The folder where you want to save the markdown files
* The folder where you want to save the PDFs
* A simple prefix for the markdown and PDF filenames (eg: "comments", "director", etc)

Use the arrow keys to navigate while answering questions in these scripts. To abandon the process at any time, press `CTRL+C` a couple times.

BTW, both scripts call `review` (to create Markdown files from the CSV) and `pandoc` (to create PDFs from the Markdown files). If you're interested in seeing how `review` works (for use outside of the scripts above), type the command below in your terminal application. You can also pass a `-v` flag to either of the above scripts, to see the underlying commands being run.

```bash
review --help
```

## Why this is on GitHub

As Graduate Director, I'm trying to make all of the procedures within the program clear and transparent. Everything is in the MFA Handbook, including the criteria that faculty will use to evaluate students in their formal reviews. Although nobody may ever look at this repo, I feel the logical conclusion of making all of our procedures transparent is to make the software tools we use equally open and accessible.
