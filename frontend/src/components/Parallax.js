import React, { useRef } from "react";
import { useNavigate } from 'react-router-dom';
import Button from '@mui/material/Button';
import { motion, useScroll, useTransform } from "framer-motion";
import '../css/Parallax.css';
// import fullImage from '../static/image-full.png';
import bottomImage from '../static/image-bottom.png';

import SC7 from '../static/2.png';
import SC8 from '../static/4.png';
import SC10 from '../static/high.png';
import SC13 from '../static/low4.png';

import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../css/MuiThemeProvider.js';


function Parallax() {
  const navigate = useNavigate();

  const handleGetStartedClick = () => {
    navigate('/user-registration')
  };

  const ref = useRef(null);
  const { scrollYProgress } = useScroll({
    target: ref,
    offset: ["start start", "end start"],
  });
  
  const backgroundY = useTransform(scrollYProgress, [0, 1], ["0%", "100%"]); // given to full image
  const textY = useTransform(scrollYProgress, [0, 1], ["0%", "300%"]); // given to text
  // const foregroundY = useTransform(scrollYProgress, [0, 1], ["0%", "100%"]);

  return (
    <ThemeProvider theme={muiTheme}>

    <div
      ref={ref}
      className="p-container"
    >
      <motion.div
        className="parallax-gradient-overlay"
        style={{
          background: `linear-gradient(
            to right,
            rgba(10, 31, 31, 0.78) calc((40vw - 170px) - 100px),
            rgba(10, 31, 31, 0.74) 50%,
            rgba(10, 31, 31, 0.78) 100%
          )`,
          position: "absolute",
          top: 0,
          left: 0,
          width: "100%",
          height: "100%",
        }}
      />

      <motion.div
        className="full-image "
        style={{
          backgroundImage: `url(${SC8})`,
          backgroundPosition: "bottom",
          backgroundSize: "cover",
          y: backgroundY,
        }}
      >

      </motion.div>
      <motion.div
          style={{ y: textY }}
          className="banner-details"
        >
          <h1>Curate. Watch. Repeat.</h1>
          <p>
            Create endless watchlists of your favorite movies, add notes, and track when you finish watching.
          </p>
          <Button
            variant="contained"
            onClick={handleGetStartedClick}
            back="primary"
            sx={{
              boxShadow: `0px 0px 2px rgba(255, 255, 255, 1),
                          0px 0px 6px rgba(255, 255, 255, 0.1),
                          0px 0px 10px rgba(255, 255, 255, 0.05)`,
              border: `1px solid rgba(255, 255, 255, 0.25)`,
            }}
          >
            Get Started
          </Button>
        </motion.div>

      <motion.div
        className="bottom-image"
        style={{
          backgroundImage: `url(${SC13})`,
          backgroundPosition: "bottom",
          backgroundSize: "cover",
        }}
      />
      
    </div>
    </ThemeProvider>
  );
};
export default Parallax;