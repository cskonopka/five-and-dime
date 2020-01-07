# dir-gen-gif.sh
#!/bin/bash

# Provide a directory containing video files 
echo "Provide a directory please :"
read dir

# Filetype target
echo "Filetype?"
read filetype

# Iterate over the directory, look for target filetype
for f in $dir/*.$filetype
do
    echo "Processing $f"
    # Generate .png analysis file
    ffmpeg -y -ss 0 -t 11 -i $f -filter_complex "[0:v] palettegen" "${f%.*}.png"
    # Create .gif file
    ffmpeg -ss 0 -t 11 -i $f -filter_complex "[0:v] fps=24,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1" "${f%.*}.gif"
done