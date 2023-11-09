// Theme styling for material UI <ThemeProvider> component

import { createTheme } from '@mui/material/styles';

export const formTheme = createTheme({
    palette: {
        primary: {
          main: '#032541', // primary colo - Dark Navy Blue - button etc. 
        },
        secondary: {
          main: '#2196f3', // secondary color - Light blue 'sign up' color
        },
      },
      typography: {
        fontSize: 16,
        fontFamily: "'Source Sans Pro', Arial, sans-serif",
      },
      spacing: 7, // spacing unit
      shape: {
        borderRadius: 6, // button + field border radius
      },
});

