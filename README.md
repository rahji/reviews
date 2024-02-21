# Reviews

This project is only useful for the FSU Department of Art Graduate Director. It is part of a process that takes place after a day of MFA Formal Reviews. Reviews happen three times per semester - once for each cohort of students. The post-review process looks like this:

1. Each faculty member submits their evaluations through a Qualtrics survey. They were previously sent invitation URLs: one for each student that they will review.
2. The survey needs to be closed one week after the review has taken place.
3. The results of all surveys are exported as a CSV file from the "Data & Analysis" tab in Qualtrics.
4. The Grad Director runs the Docker command described below to produce PDFs for distribution to the students and a separate batch of PDFs containing comments directed to the Grad Director and Grad Advisor & Coordinator.

## Why

As in, "Why did you spend so many hours of your life doing this?"

Exporting PDFs in batch is a feature that Qualtrics does not have. It only lets you export one PDF at a time and it doesn't let you choose what's in it. We need this capability so we can return the results of the evaluations to the students after each of their semesterly reviews. 

## Installation

You'll need a unix-like operating system before you get started.
On Windows, this means you'll need to install WSL2. You might as well install [Terminal](https://apps.microsoft.com/store/detail/windows-terminal/9N0DX20HK701) while you're at it.
MacOS and Linux are already unix-like, so you're all set.

Follow these steps to get read to run the Docker container. Even though we'll use Docker from the terminal, it makes the most sense to install Docker Desktop:

1. [Install Docker Desktop](https://www.docker.com/products/docker-desktop/), while signing up for a free account.  If you're on Windows, you'll install Docker Desktop in Windows but [configure it to use WSL2](https://docs.docker.com/desktop/wsl/). To make sure it's working, try typing `docker version` at the terminal. You shouldn't get any errors.
2. From the terminal, run `docker pull robduarte/reviews:v1`

## Usage

1. Start the terminal application (make sure it's WSL if you're on Windows).
2. Make a new directory and copy the CSV file from Qualtrics to that directory.
2. Change to the new directory, then run the following Docker command:

```bash
docker run -u $(id -u):$(id -g) -v .:/app/data robduarte/reviews:v1 /app/data/YOURFILE.CSV
```
**Enter the command exactly as above, except for YOURFILE.CSV, which should be replaced with your CSV filename**

*(Note: the `-u` bit is required or your PDFs will all be owned by the `root` user, which is a pain when if you want to remove them later.)*

The above Docker command will create a series of directories containing all the PDF files you need. It also creates a corresponding markdown file for each PDF. 

After emailing the student PDFs to the individual grads, copy all of the folders into the appropriate folder on Sharepoint for our records.

## Why this is on GitHub

As Graduate Director, I'm trying to make all of the procedures within the program clear and transparent. Everything is in the MFA Handbook, including the criteria that faculty will use to evaluate students in their formal reviews. Although nobody may ever look at this repo, I feel the logical conclusion of making all of our procedures transparent is to make the software tools we use equally open and accessible.

