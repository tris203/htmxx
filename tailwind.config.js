/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templ/*.templ", "./handler/*.go"],
  theme: {
    extend: {},
  },
  plugins: [ 
    require('@tailwindcss/typography'),
  ],
};
