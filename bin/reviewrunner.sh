#!/bin/sh

#REVIEWSDIR=/app/reviews
#DATADIR=/app/data
REVIEWSDIR=~/git/reviews
DATADIR=~/git/reviews/data

# Accept input file from local folder
INPUT_FILE=$1

# Check if input file is provided
if [ -z "$INPUT_FILE" ]; then
  echo "Input file not provided."
  exit 1
fi

# Check if the input file exists
if [ ! -f "$INPUT_FILE" ]; then
  echo "Input file $INPUT_FILE not found."
  exit 1
fi

# Check if the input file is a CSV file
if [ "${INPUT_FILE##*.}" != "csv" ]; then
  echo "Error: Input file is not a CSV file."
  exit 1
fi

# Create student markdown files
$GOBIN/reviews convert --template=$REVIEWSDIR/templates/student_template.md \
  --input="$INPUT_FILE" --outputdir=$DATADIR/student/markdown
echo "Created student markdown files from input file"

# Create student PDF files
for MDFILE in `ls $DATADIR/student/markdown/*.md`;
do
    # make a pdf name based on the markdown filename
    BASEFILENAME=`basename $MDFILE .md`
    PDFNAME="$DATADIR/student/pdf/$BASEFILENAME.pdf"
    pandoc $MDFILE -o $PDFNAME -V geometry:landscape --from markdown --template eisvogel
    echo "Created student PDF: $PDFNAME"
done

# Create director-only markdown files
$GOBIN/reviews director --template=$REVIEWSDIR/templates/directorcomments_template.md \
  --input="$INPUT_FILE" --outputdir=$DATADIR/director/markdown --outputprefix=private
echo "Created director-only markdown files from input file"

# Create director-only PDF files
for MDFILE in `ls $DATADIR/director/markdown/*.md`;
do
    # make a pdf name based on the markdown filename
    BASEFILENAME=`basename $MDFILE .md`
    PDFNAME="$DATADIR/director/pdf/$BASEFILENAME.pdf"
    pandoc $MDFILE -o $PDFNAME -V geometry:landscape --from markdown --template eisvogel
    echo "Created director-only PDF: $PDFNAME"
done


