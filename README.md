# Reviews

This program takes a CSV file of survey results exported from Qualtrics and creates a separate Markdown file for each survey. Exporting PDFs in batch is a feature that Qualtrics does not have - it lets you export one PDF at a time and doesn't let you choose what's in it. I needed this capability since we need to return the results of the evaluations to the students after their semesterly reviews. Several faculty members meet with each student, so the students receive multiple forms after each review. The workaround is to export all of the results from Qualtrics as one CSV file, then feed that file to this program.

## Installation

`reviews` will run on linux, MacOS, or Windows. More information about installation is needed here.

## Usage

```bash
reviews convert
```

Todo:

* need a way to specify template file
* also the csv file
* also needs to create markdown *files* instead sending markdown to STDOUT

> PS: [Glow](https://github.com/charmbracelet/glow) is a fun way to read Markdown documents in the terminal.

## Markdown to PDF

This program exports Markdown because it's easy to deal with. Specifically, the template file is easier to edit than an HTML template would be. But in the end, we want to send PDFs to the students, so the next step is to convert the collection of Markdown files to PDFs. For this, we can use [pandoc](https://pandoc.org/). 

### Installing pandoc

Installing pandoc requires a bit of effort. The (very large) `texlive-full` linux package is required and can be easily installed with the following command on Ubuntu/Debian. If you're on Windows, you'll want to install WSL (Windows Subsystem for Linux). You'll have to figure out the equivalent way of installing the `texlive-full` package if you're using MacOS.

```bash
apt-get install texlive-full
```

Unfortunately, pandoc makes *ugly* PDFs by default, so you'll also want to [download the Eisvogel template](https://github.com/Wandmalfarbe/pandoc-latex-template/releases) and follow the instructions for installing it (which just means placing the template file in a folder where pandoc can find it.)

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
