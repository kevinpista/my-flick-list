import React, { useRef } from "react";
import { useNavigate } from 'react-router-dom';
import Button from '@mui/material/Button';
import { motion, useScroll, useTransform } from "framer-motion";
import '../css/Parallax.css';

import { easeIn } from "framer-motion"

import MovieCollage from '../static/MovieCollage.png';
// import SC10 from '../static/high.png';
// import SC14 from '../static/wat.png';

// import Guy1 from '../static/Guy1.png';
//import Guy2 from '../static/Guy2.png';
// import Crowd1 from '../static/Crowd1.png'; // close to content
import Crowd from '../static/Crowd5.png';


import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../css/MuiThemeProvider.js';


function Parallax() {
  const navigate = useNavigate();

  const handleGetStartedClick = () => {
    navigate('/user-login')
  };

  const ref = useRef(null);
  const { scrollYProgress } = useScroll({
    target: ref,
    offset: ["start start", "end start"],
  });
  
  const backgroundY = useTransform(scrollYProgress, [0, 1], ["0%", "100%"], { ease: easeIn });
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
          backgroundImage: `url(${MovieCollage})`,
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

        <motion.img
        className="bottom-image"
        src={Crowd}

      />
      
    </div>
    </ThemeProvider>
  );
};
export default Parallax;