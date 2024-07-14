# save-paste
Its a proyect using go, templ and tailwind to practice. 

# install tailwind
## Example for macOS arm64
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
chmod +x tailwindcss-macos-arm64
mv tailwindcss-macos-arm64 tailwindcss
(https://github.com/tailwindlabs/tailwindcss/releases/tag/v3.4.4)
## Create a tailwind.config.js file
./tailwindcss init

## Start a watcher
./tailwindcss -i input.css -o output.css --watch

## Compile and minify your CSS for production
./tailwindcss -i input.css -o output.css --minify