/* @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{html,js}","./public/**/*.{html,js}","!./public/js/app.js"],
  corePlugins: {
    backdropOpacity: false,
    backgroundOpacity: false,
    borderOpacity: false,
    divideOpacity: false,
    ringOpacity: false,
    textOpacity: false,
  },
 
  theme: {
    extend: {},
  },
   plugins: [
    require('@tailwindcss/container-queries'),
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ]
}
