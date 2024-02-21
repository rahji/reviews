# Reviews

This project is only useful for the FSU Department of Art Graduate Director. It is part of a process that takes place after a day of MFA Formal Reviews. Reviews happen three times per semester - once for each cohort of students. The post-review process looks like this:

1. Each faculty member submits their evaluations through a Qualtrics survey. They were previously sent invitation URLs: one for each student that they will review.
2. The survey needs to be closed one week after the review has taken place.
3. The results of all surveys are exported as a CSV file from the "Data & Analysis" tab in Qualtrics.
4. The Grad Director runs the `reviews-docker` Docker container to produce PDFs for distribution to the students and a separate batch of PDFs containing comments directed to the Grad Director and Grad Advisor & Coordinator.

## Why

As in, "Why did you spend so many hours of your life doing this?"

Exporting PDFs in batch is a feature that Qualtrics does not have. It only lets you export one PDF at a time and it doesn't let you choose what's in it. We need this capability so we can return the results of the evaluations to the students after each of their semesterly reviews. 

## Installation

You'll need these things before you get started:

* A unix-like operating system. 
  * On Windows, this means you'll need to install WSL2. You might as well install [Terminal](https://apps.microsoft.com/store/detail/windows-terminal/9N0DX20HK701) while you're at it.
  * MacOS and Linux are already unix-like, so you're all set.
* [Install Docker Desktop](https://www.docker.com/products/docker-desktop/), while signing up for a free account. (If you're on Windows, you'll install Docker Desktop in Windows but [configure it to use WSL2](https://docs.docker.com/desktop/wsl/))

## Usage

### Windows

1. Make sure that Docker is [set up to use WSL2](https://docs.docker.com/desktop/wsl/) by trying the `docker` command while at a linux shell in the Terminal application.
2. Make a new directory in linux and copy the CSV file from Qualtrics to that directory.
3. Change to that directory, then run the following Docker command:

```bash
docker run -u $(id -u):$(id -g) -v .:/app/data reviews-docker /app/data/file_from_qualtrics.csv
```

(Note: the `-u` bit is required or your PDFs will all be owned by the `root` user, which is a pain when if you want to remove them later.)

### MacOS or Linux

Follow steps 2 and 3 from the Windows instructions above.

### Output

No matter which OS you're using, the Docker command will create a series of directories containing all the PDF files you need. After emailing the student PDFs to the individual grads, copy all of the folders into the appropriate folder on Sharepoint for our records.

## Why this is on GitHub

As Graduate Director, I'm trying to make all of the procedures within the program clear and transparent. Everything is in the MFA Handbook, including the criteria that faculty will use to evaluate students in their formal reviews. Although nobody may ever look at this repo, I feel the logical conclusion of making all of our procedures transparent is to make the software tools we use equally open and accessible.

