#!/bin/sh

# use -v option to show full commands being run

VERBOSE='false'
while getopts ':v' 'OPTKEY'; do
    case ${OPTKEY} in
        'v')
            VERBOSE='true'
            ;;
    esac
done
if $VERBOSE; then
    echo "Running in VERBOSE mode!"
    echo
fi


echo "We're about to extract comments (including those meant for the grad Director) from a Qualtrics CSV file..."
echo ""

echo "Choose the CSV file you want to convert to PDFs"
CSV=$(gum file .)

echo "Choose the markdown template for creating grad director PDFs"
TEMPLATE=$(gum file .)

echo "Choose the folder where you want to save the markdown files"
MDOUTPUT=$(gum file --directory .)

echo "Choose the folder where you want to save the PDFs"
PDFOUTPUT=$(gum file --directory .)

echo "Enter a simple prefix for the markdown and PDF filenames"
PREFIX=$(gum input --placeholder "eg: comments, director, etc.")

echo "Ready to create markdown files in $MDOUTPUT?"
CHOICE=$(gum choose "YES" "NO")
if [ "$CHOICE" != "YES" ]; then
    echo "Okay, bye!"
    exit 1
fi

CMD_REVIEWS="reviews director --template=$TEMPLATE --input=$CSV --outputdir=$MDOUTPUT --outputprefix=$PREFIX"
if $VERBOSE; then
    echo "About to run: ($CMD_REVIEWS)"
fi

gum spin --title "Converting to markdown" -- $CMD_REVIEWS

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
    CMD_PANDOC="pandoc $MDFILE -o $PDFNAME -V geometry:landscape --from markdown --template eisvogel"
    if $VERBOSE; then
        echo "Running: ($CMD_PANDOC)"
        $CMD_PANDOC
    else
        gum spin --title "$(printf "Converting %s to %s" $SHORTMDNAME $SHORTPDFNAME)" -- $CMD_PANDOC
    fi
done

echo "Done!"
