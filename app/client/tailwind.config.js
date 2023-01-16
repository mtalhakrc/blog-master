const colors = require("tailwindcss/colors");

module.exports = {
  content: ["./src/components/**/*.{js,vue,ts}", "./src/views/**/*.vue"],
  theme: {
    colors: colors,
    container: false,

    fontFamily: {},
    extend: {
      colors: {
        colortext: "#21243D",
        lightred: "#FF6464",
        lightgray: "#8695A4",
        lightblue: "#EDF7FA",
        darkblue: "#00A8CC",
        dark: "#142850",
      },
    },
  },

  plugins: [],
};
