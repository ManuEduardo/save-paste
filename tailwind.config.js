/** @type {import('tailwindcss').Config} */
export const content = ["./src/views/*.html", "./src/views/*.templ", 
  "./src/views/layouts/*.templ", "./src/views/layouts/*.html", 
  "./src/views/components/*.templ", "./src/views/components/*.html"];
export const theme = {
  extend: {},
};
export const plugins = [];

export const darkMode = ["class"]