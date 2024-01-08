import { createTheme } from '@mui/material/styles';

// Styling for buttons 
export const muiTheme = createTheme({
    palette: {
      primary: {
        main: '#032541',
        // light: will be calculated from palette.primary.main,
        // dark: will be calculated from palette.primary.main,
        // contrastText: will be calculated to contrast with palette.primary.main
      },
      // Not currently using
      secondary: {
        main: '#E0C2FF',
        light: '#F5EBFF',
        // dark: will be calculated from palette.secondary.main,
        contrastText: '#47008F',
      },
      typography: {
        fontFamily: "'Source Sans Pro', Arial, sans-serif",
      },
    },
  });
