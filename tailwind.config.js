plugin = require('tailwindcss/plugin')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templ/*.templ", "./handler/*.go"],
  theme: {
    extend: {},
  },
  plugins: [ 
    require('@tailwindcss/typography'),
    plugin(function({ addVariant }) {
      addVariant("htmx-settling", ["&.htmx-settling", ".htmx-settling &"]);
      addVariant("htmx-request", ["&.htmx-request", ".htmx-request &"]);
      addVariant("htmx-swapping", ["&.htmx-swapping", ".htmx-swapping &"]);
      addVariant("htmx-added", ["&.htmx-added", ".htmx-added &"]);
      addVariant("htmx-indicator", ["&.htmx-indicator", ".htmx-indicator &"]);
    }),
  ],
};
