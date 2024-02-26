/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,jsx,ts,tsx,vue}'],
  theme: {
    extend: {
      animation: {
        'background-shine': 'background-shine 2.5s linear infinite',
        'text-gradient': 'text-gradient 2.5s linear infinite'
      },
      keyframes: {
        'background-shine': {
          from: {
            backgroundPosition: '0 0'
          },
          to: {
            backgroundPosition: '-200% 0'
          }
        },
        'text-gradient': {
          to: {
            backgroundPosition: '200% center'
          }
        }
      }
    }
  },
  plugins: []
}
