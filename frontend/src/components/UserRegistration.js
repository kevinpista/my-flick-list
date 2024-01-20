import * as React from 'react';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Avatar from '@mui/material/Avatar';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import NavBar from './NavBar';
import LockOpenIcon from '@mui/icons-material/LockOpen';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { ThemeProvider } from '@mui/material/styles';
import { registerUser } from '../api/userRegistrationAPI';
import * as errorConstants from '../api/errorConstants';
import * as themeStyles from '../styling/ThemeStyles';
import Alert from '@mui/material/Alert';
import LoadingButton from '@mui/lab/LoadingButton';


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

// TODO remove, this demo shouldn't need to reset the theme.

export default function UserRegistration() {

    const [formData, setFormData] = useState({
        name: '',
        email: '',
        password: '',
        });

    // State variables for error <Alert> and status for form fields
    const [nameError, setNameError] = useState(false);
    const [emailError, setEmailError] = useState(false);
    const [passwordError, setPasswordError] = useState(false);
    const [loading, setLoading] = useState(false);

    const [showSuccessAlert, setShowSuccessAlert] = useState(false);
    const [errorAlertMessage, setErrorAlertMessage] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (event) => {
        event.preventDefault();
        setLoading(true)

    try {
        const response = await registerUser(formData); // i believe response var here = the backend json body already
        // can just access the body like response.name or response.id etc. No need for response.data.name as 
        // that is handled in the axios code file when returned as response.data

        // If there's a response with no error code, successful
        if (response) {
            setErrorAlertMessage('');
            setShowSuccessAlert(true);

            setTimeout(() => {
                navigate('/watchlist');
              }, 1600); // Redirect to user's watchlist page
        } 

        // If error was thrown by API request
        } catch(error) {
            setLoading(false)
            if (error.message === errorConstants.ERROR_EMAIL_EXISTS) {
                setErrorAlertMessage('Email address is already in use.');
                setEmailError(true);
            } else if (error.message === errorConstants.ERROR_INVALID_EMAIL) {
                setErrorAlertMessage('Email format is not valid.');
                setEmailError(true);
            } else if (error.message === errorConstants.ERROR_PASSWORD_WHITESPACE) {
                setErrorAlertMessage('Password cannot have any whitespace.');
                setPasswordError(true);
            } else if (error.message === errorConstants.ERROR_PASSWORD_EMPTY) {
                setErrorAlertMessage('Password cannot be empty.');
                setPasswordError(true);
            } else if (error.message === errorConstants.ERROR_INVALID_NAME) {
                setErrorAlertMessage('Name cannot be empty.');
                setNameError(true);
            } else if (error.message === errorConstants.ERROR_BAD_REQUEST) {
                setErrorAlertMessage('Bad request.');
            } else {
                setErrorAlertMessage('An unexpected error occurred.');
            }       
        } 
    };

    const handleInputChange = (e) => {
        const fieldName = e.target.name;
        const fieldValue = e.target.value;

        // Reset error state when the user starts typing in a field
        if (fieldName === 'name') {
            setNameError(false);
        } else if (fieldName === 'email') {
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
            {/* Display alert based on registration status */}
            {showSuccessAlert && (
                <Alert severity="success">
                    <strong>Successful Sign Up</strong> - Logging you in...
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
                <LockOpenIcon />
                </Avatar>
                <Typography component="h1" variant="h5">
                Create Account
                </Typography>
                <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
                <Grid container spacing={2}>
                    <Grid item xs={12}>
                    <TextField
                        required
                        fullWidth
                        id="name"
                        label="Full Name"
                        name="name"
                        autoComplete="name"
                        onChange={handleInputChange}
                        error={nameError} // Applying error style conditionally with useState
                    />
                    </Grid>
                    <Grid item xs={12}>
                    <TextField
                        required
                        fullWidth
                        id="email"
                        label="Email Address"
                        name="email"
                        autoComplete="email"
                        onChange={handleInputChange}
                        error={emailError}
                    />
                    </Grid>
                    <Grid item xs={12}>
                    <TextField
                        required
                        fullWidth
                        id="password"
                        label="Password"
                        name="password"
                        type="password"
                        autoComplete="new-password"
                        onChange={handleInputChange}
                        error={passwordError}
                    />
                    </Grid>
                </Grid>
                <LoadingButton 
                    type="submit"
                    fullWidth
                    variant="contained"
                    loading={loading}
                    style={{marginTop: "20px", marginBottom: "15px"}}
                    
                    >
                    SIGN UP
                </LoadingButton>
                <Grid container justifyContent="center">
                    <Grid item>
                    <Link href="http://localhost:3000/user-login" variant="body2">
                        Already have an account? Sign in
                    </Link>
                    </Grid>
                </Grid>
                </Box>
            </Box>
            <Copyright sx={{ mt: 5 }} />
            </Container>
        </ThemeProvider>
    );
}
