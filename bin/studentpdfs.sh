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


echo "We're about to convert a Qualtrics CSV file to a bunch of PDFs..."
echo ""

echo "Choose the CSV file you want to convert to PDFs"
CSV=$(gum file .)

echo "Choose the markdown template for creating student PDFs"
TEMPLATE=$(gum file .)

echo "Choose the folder where you want to save the markdown files"
MDOUTPUT=$(gum file --directory .)

echo "Choose the folder where you want to save the PDFs"
PDFOUTPUT=$(gum file --directory .)

CMD_REVIEWS="reviews convert --template=$TEMPLATE --input=$CSV --outputdir=$MDOUTPUT"
if $VERBOSE; then
    echo "About to run: ($CMD_REVIEWS)"
fi

echo "Ready to create markdown files in $MDOUTPUT?"
CHOICE=$(gum choose "YES" "NO")
if [ "$CHOICE" != "YES" ]; then
    echo "Okay, bye!"
    exit 1
fi

gum spin --title "Converting to markdown" -- $CMD_REVIEWS

echo "Read to create PDF files in $PDFOUTPUT? (this is the slow part)"
CHOICE=$(gum choose "YES" "NO")
if [ "$CHOICE" != "YES" ]; then
    echo "Okay, bye!"
    exit 1
fi

for MDFILE in `ls $MDOUTPUT/*.md`;
do
    # make a pdf name based on the markdown filename
    BASEFILENAME=`basename $MDFILE .md`
    PDFNAME="$PDFOUTPUT/$BASEFILENAME.pdf"

    # make short filenames for the spinner display
    SHORTMDNAME=`basename $MDFILE`
    SHORTPDFNAME=`basename $PDFNAME`
    CMD_PANDOC="pandoc $MDFILE -o $PDFNAME -V geometry:landscape --from markdown --template eisvogel --columns=1000"
    if $VERBOSE; then
        echo "Running: ($CMD_PANDOC)"
        $CMD_PANDOC
    else
        gum spin --title "$(printf "Converting %s to %s" $SHORTMDNAME $SHORTPDFNAME)" -- $CMD_PANDOC
    fi
done

echo "Done!"
