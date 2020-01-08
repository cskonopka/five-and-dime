# #!/bin/bash
# echo input png?
# read input

# echo output histogram location?
# read output

# convert $input -format %c histogram:info:- > $output


# create-histogram.sh
#!/bin/bash

# Provide a ".png" file
echo "PNG please ~~~~~"
read PNG

# Remove ".png" from PNG
SPLICE=${PNG%????}

# Create the EXPORTFILE as ".txt"
EXPORTFILE="${SPLICE}.txt"
echo $EXPORTFILE

# Export a histogram ".txt" file based on the ".png"
convert $PNG -format %c histogram:info:- > $EXPORTFILE