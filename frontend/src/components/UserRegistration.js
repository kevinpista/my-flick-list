import * as React from 'react';
import { useState } from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { registerUser } from '../api/userRegistrationAPI';
import * as errorConstants from '../api/errorConstants';



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

const defaultTheme = createTheme();


export default function UserRegistration() {

    const [formData, setFormData] = useState({
        name: '',
        email: '',
        password: '',
        });

    const handleSubmit = async (event) => {
    event.preventDefault();

    try {
        const response = await registerUser(formData); // i believe response var here = the backend json body already
        // can just access the body like response.name or response.id etc. No need for response.data.name as 
        // that is handled in the axios code file when returned as response.data

        // If there's a response with no error code, successful
        if (response) {
            console.log('Registration successful!!')
            // TODO - add action such as logging user in or redirecting
        } 
        // If error was thrown by API request
        // TODO - nice message pop up action, not a console.log()
        } catch(error) {
            if (error.message === errorConstants.ERROR_EMAIL_EXISTS) {
                console.log('Email address is already in use!')
            } else if (error.message === errorConstants.ERROR_INVALID_EMAIL) {
                console.log('Email format is not valid!')
            } else if (error.message === errorConstants.ERROR_PASSWORD_WHITESPACE) {
                console.log('Password cannot have any whitespace!')
            } else if (error.message === errorConstants.ERROR_PASSWORD_EMPTY) {
                console.log('Password cannot be empty!')
            } else if (error.message === errorConstants.ERROR_INVALID_NAME) {
                console.log('Name cannot be empty!')
            } else if (error.message === errorConstants.ERROR_BAD_REQUEST) {
                console.log('Bad request')
            } else if (error.message === errorConstants.ERROR_SERVER) {
                console.log('Server connection error.')
            } else {
                console.log("else statement executed")
            }       
        } 
    };

    const handleInputChange = (e) => {
        const fieldName = e.target.name;
        const fieldValue = e.target.value;
        
            setFormData({
            ...formData,
            [fieldName]: fieldValue,
            });
        
        };


    // RENDER COMPONENT
    return (
        <ThemeProvider theme={defaultTheme}>
            <Container component="main" maxWidth="xs">
            <CssBaseline />
            <Box
                sx={{
                marginTop: 8,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                }}
            >
                <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
                <LockOutlinedIcon />
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
                    />
                    </Grid>

                </Grid>
                <Button
                    type="submit"
                    fullWidth
                    variant="contained"
                    sx={{ mt: 3, mb: 2 }}
                >
                    Sign Up
                </Button>
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
