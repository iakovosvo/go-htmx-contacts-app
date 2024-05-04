/** @type {import('tailwindcss').Config} */
export default {
  content: ["./views/**/*.templ"],
  theme: {
    theme: {
      extend: {
        colors: {
          "custom-gray": "#1a1a1a",
        },
      },
    },
  },
  plugins: [],
};
