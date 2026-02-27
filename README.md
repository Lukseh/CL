# CL

cl is utility for coloring text output made with go.

## Install

Run command ```go install github.com/Lukseh/CL```

## Usage

You can either

- Use it with filename like ```cl -word X -word Y -word Z -color FFFF00 file.log```

Or

- With pipe like ```cat file.log | cl -word X -word Y -word Z -color FFFF00```

If color is not provided it will default to red.