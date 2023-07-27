#!/bin/sh

echo "We're about to extract director-only (private) comments from a Qualtrics csv file..."
echo ""

echo "Choose the csv file you want to convert to PDFs"
CSV=$(gum file .)

echo "Choose the markdown template for this director-only (private) comments"
TEMPLATE=$(gum file .)

echo "Choose the folder where you want to save the markdown files"
MDOUTPUT=$(gum file --directory .)

echo "Choose the folder where you want to save the PDFs"
PDFOUTPUT=$(gum file --directory .)

echo "Enter a simple prefix for the markdown and PDF filenames"
PREFIX=$(gum input --placeholder "eg: private, secret, director, etc.")

echo "Ready to create markdown files in $MDOUTPUT?"
CHOICE=$(gum choose "YES" "NO")
if [ "$CHOICE" != "YES" ]; then
    echo "Okay, bye!"
    exit 1
fi

gum spin --title "Converting to markdown" -- reviews director --template=$TEMPLATE --input=$CSV --outputdir=$MDOUTPUT --outputprefix=$PREFIX

echo "Ready to create PDF files in $PDFOUTPUT? (this is the slow part)"
CHOICE=$(gum choose "YES" "NO")
if [ "$CHOICE" != "YES" ]; then
    echo "Okay, bye!"
    exit 1
fi

for MDFILE in `ls $MDOUTPUT/$PREFIX*.md`;
do
    # make a pdf name based on the markdown filename
    BASEFILENAME=`basename $MDFILE .md`
    PDFNAME="$PDFOUTPUT/$BASEFILENAME.pdf"

    # make short filenames for the spinner display
    SHORTMDNAME=`basename $MDFILE`
    SHORTPDFNAME=`basename $PDFNAME`
    MESSAGE=$(printf "Converting %s to %s" $SHORTMDNAME $SHORTPDFNAME)
    gum spin --title "$MESSAGE" -- pandoc $MDFILE -o $PDFNAME -V geometry:landscape --from markdown --template eisvogel
done

echo "Done!"
