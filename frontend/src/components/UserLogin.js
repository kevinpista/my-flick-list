import * as React from 'react';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Avatar from '@mui/material/Avatar';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import NavBar from './NavBar';
import { Container, Typography, Button, Tooltip } from '@mui/material';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import { ThemeProvider } from '@mui/material/styles';
import { loginUser } from '../api/userLoginAPI';
import * as errorConstants from '../api/errorConstants';
import * as themeStyles from '../styling/ThemeStyles';
import Alert from '@mui/material/Alert';
import LoadingButton from '@mui/lab/LoadingButton';
import { CopyToClipboard } from 'react-copy-to-clipboard';
import '../css/UserLogin.css';

function Copyright(props) {
    return (
      <Typography variant="body2" color="text.secondary" align="center" {...props}>
        {'Copyright © '}
        <Link color="inherit" href="/">
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
  const [loading, setLoading] = useState(false);
  const [emailError, setEmailError] = useState(false);
  const [passwordError, setPasswordError] = useState(false);

  const [showSuccessAlert, setShowSuccessAlert] = useState(false);
  const [errorAlertMessage, setErrorAlertMessage] = useState('');

  const demoEmail = 'demo@test.com';
  const demoPassword = 'demo123!';

  const [tooltipMessage, setTooltipMessage] = useState('Click to copy');

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
        setLoading(true)

        if (response) {
          setErrorAlertMessage('');
          setShowSuccessAlert(true);

          setTimeout(() => {
            navigate('/watchlist');
          }, 1600); // Redirect to user's watchlist page
          setLoading(false)

        }
    } catch (error) {
        if (error.message === errorConstants.ERROR_INVALID_EMAIL) {
          setErrorAlertMessage('Email format not valid.');
          setLoading(false)
          setEmailError(true);
      } else if (error.message === errorConstants.ERROR_INVALID_LOGIN) {
          setErrorAlertMessage('Invalid login credentials');
          setLoading(false)
          setEmailError(true); // Set error status for both form fields
          setPasswordError(true);
      } else {
          setErrorAlertMessage('An unexpected error occurred');
          setLoading(false)
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

  const handleCopy = (text) => {
    copyToClipboard(text);
    setTooltipMessage('Copied!');
    setTimeout(() => setTooltipMessage('Click to copy'), 1000);
  };

  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text).then(() => {
      // Clipboard writing successful
    }, (err) => {
      console.error('Failed to copy text:', err);
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
            marginTop: 4,
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

            <LoadingButton 
              type="submit"
              fullWidth
              variant="contained"
              loading={loading}
              style={{marginTop: "5px", marginBottom: "10px"}}
              
              >
              SIGN IN
            </LoadingButton>
            <Grid container>
              <Grid item xs>
                <Link href="#" variant="body2">
                  Forgot password?
                </Link>
              </Grid>
              <Grid item>
                <Link href="https://myflicklist-fa78f7f017a1.herokuapp.com/user-registration" variant="body2">
                  {"Don't have an account? Sign Up"}
                </Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
        <Copyright sx={{ mt: 2.5}} />

        <Box
          sx={{
            marginTop: 3,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            border: '2px solid #0b1d26',
          }}
        >
          <div className="demo-container">
            <h3>
              Demo Account
            </h3>
            <h6>
              Use these credentials to log in.
            </h6>

            <div className="demo-credentials">
              <p style={{marginBottom: '12px'}}>
                <b>Email:</b> {demoEmail}
                <span style={{marginLeft: '10px'}}>
                <CopyToClipboard text={demoEmail} onCopy={() => handleCopy(demoEmail)}>
                <Tooltip title={tooltipMessage}>
                  <Button 
                    style={{maxWidth: '20px', maxHeight: '20px', minWidth: '20px', minHeight: '20px', padding: '10px', marginLeft: '0.30rem'}}
                  >
                    <ContentCopyIcon />
                  </Button>
                </Tooltip>
                </CopyToClipboard>
                </span>
              </p>

              <p>
                <b style={{marginTop: '20px'}}>Password:</b> {demoPassword} 
                <span style={{marginLeft: '10px'}}>
                  <CopyToClipboard text={demoPassword} onCopy={() => handleCopy(demoPassword)}>
                  <Tooltip title={tooltipMessage}>
                    <Button 
                      style={{maxWidth: '20px', maxHeight: '20px', minWidth: '20px', minHeight: '20px', padding: '10px', marginLeft: '1.2rem'}}
                    >
                      <ContentCopyIcon />
                    </Button>
                  </Tooltip>
                  </CopyToClipboard>
                </span>
              </p>

            </div>
          </div>
        </Box>
      </Container>
    </ThemeProvider>
  );
}