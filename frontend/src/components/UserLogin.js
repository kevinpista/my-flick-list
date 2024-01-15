import * as React from 'react';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import NavBar from './NavBar';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { ThemeProvider } from '@mui/material/styles';
import { loginUser } from '../api/userLoginAPI';
import * as errorConstants from '../api/errorConstants';
import * as themeStyles from '../styling/ThemeStyles';
import Alert from '@mui/material/Alert';

function Copyright(props) {
    return (
      <Typography variant="body2" color="text.secondary" align="center" {...props}>
        {'Copyright © '}
        <Link color="inherit" href="https://mui.com/">
          My Flick List
        </Link>{' '}
        {new Date().getFullYear()}
        {'.'}
      </Typography>
    );
  }

export default function UserLogin() {

  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  const [emailError, setEmailError] = useState(false);
  const [passwordError, setPasswordError] = useState(false);

  const [showSuccessAlert, setShowSuccessAlert] = useState(false);
  const [errorAlertMessage, setErrorAlertMessage] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (!formData.email && !formData.password) {
      setEmailError(true);
      setPasswordError(true);
      return;
    }

    if (!formData.email) {
      setEmailError(true);
      return;
    }

    if (!formData.password) {
      setPasswordError(true);
      return;
    }

    try {
        const response = await loginUser(formData);

        if (response) {
          setErrorAlertMessage('');
          setShowSuccessAlert(true);

          setTimeout(() => {
            navigate('/watchlist');
          }, 2000); // Redirect to user's watchlist page after 2 seconds delay

        }
    } catch (error) {
        if (error.message === errorConstants.ERROR_INVALID_EMAIL) {
          setErrorAlertMessage('Email format not valid.');
          setEmailError(true);
      } else if (error.message === errorConstants.ERROR_INVALID_LOGIN) {
          setErrorAlertMessage('Invalid login credentials');
          setEmailError(true); // Set error status for both form fields
          setPasswordError(true);
      } else {
          setErrorAlertMessage('An unexpected error occurred');
      }
    };
  };
  
  const handleInputChange = (e) => {
    const fieldName = e.target.name;
    const fieldValue = e.target.value;

    // Reset error state when the user starts typing in a field
    if (fieldName === 'email') {
      setEmailError(false);
    } else if (fieldName === 'password') {
      setPasswordError(false);
    }

    setFormData({
      ...formData,
      [fieldName]: fieldValue,
    });

  };

  return (
    <ThemeProvider theme={themeStyles.formTheme}>
      <NavBar/>
      <Container component="main" maxWidth="xs">
        {/* Display alert based  on login success or error */}
        {showSuccessAlert && (
          <Alert severity="success">
            <strong>Successful Credentials</strong> - Logging you in...
          </Alert>
        )}

        {errorAlertMessage && (
          <Alert severity="error">
            <strong>Error</strong> - {errorAlertMessage}
          </Alert>
        )}

        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: 'primary.main' }}>
            <AccountCircleIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Account Login
          </Typography>
          <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
            <TextField
              margin="normal"
              required
              fullWidth
              id="email"
              label="Email Address"
              name="email"
              autoComplete="email"
              autoFocus
              onChange={handleInputChange}
              error={emailError} // Applying error style conditionally with useState
            />
            <TextField
              margin="normal"
              required
              fullWidth
              id="password"
              label="Password"
              name="password"
              type="password"
              autoComplete="current-password"
              onChange={handleInputChange}
              error={passwordError}
            />
            <FormControlLabel
              control={<Checkbox value="remember" color="primary" />}
              label="Remember me"
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Sign In
            </Button>
            <Grid container>
              <Grid item xs>
                <Link href="#" variant="body2">
                  Forgot password?
                </Link>
              </Grid>
              <Grid item>
                <Link href="http://localhost:3000/user-registration" variant="body2">
                  {"Don't have an account? Sign Up"}
                </Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
        <Copyright sx={{ mt: 5}} />
      </Container>
    </ThemeProvider>
  );
}