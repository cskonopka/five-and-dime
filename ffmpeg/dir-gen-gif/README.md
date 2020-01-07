
# dir-gen-gif

![gif](https://i.ibb.co/6sxD6Ys/dir-gen-gif.gif) 

Generate GIFs from a directory of video content using FFmpeg. When prompted, provide a target directory and specify the output format.

# Background
As a [video artist](www.cskonopka.com) I deal with many physical formats (analog composite to 4K) and digital formats (.mp4, .mov, Motion JPEG, .gif, etc.). Going between formats can be a nightmare, and I like to build tools that can ease any potential stress. Over the summer, I wanted to convert my [analog video synthesis](vimeo.com/cskonopka) into .gif files to upload to [Giphy](https://giphy.com/cskonopka/). So I made this bash script that ingests a directory of video content and asks for a target file format.

# Breakdown

Create a new bash script named *dir-gen-gif.sh*.

```bash
# dir-gen-gif.sh
#!/bin/bash
```

Ask for a directory of video content. 
```bash
echo "Provide a directory please :"
read dir
```

Ask for a video file format.
```bash
echo "Filetype?"
read filetype
```

Create a loop that iterates over the directory and finds the target video file format.
```bash
for f in $dir/*.$filetype
do
    echo "Processing $f"
    # ... ffmpeg code
done
```

Add the FFmpeg scripts. First create a .png analysis file and then generate a .gif file.
```bash
    # Generate .png analysis file
    ffmpeg -y -ss 0 -t 11 -i $f -filter_complex "[0:v] palettegen" "${f%.*}.png"
    # Create .gif file
    ffmpeg -ss 0 -t 11 -i $f -filter_complex "[0:v] fps=24,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1" "${f%.*}.gif"
```

And the process will look something like this ...

![gif](https://i.ibb.co/8NNNX1t/dir-gen-gif3.gif)

# Reference
```bash
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
```
