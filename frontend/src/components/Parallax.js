import React, { useRef } from "react";
import { useNavigate } from 'react-router-dom';
import NavBar from './NavBar.js';
import { motion, useScroll, useTransform } from "framer-motion";
import '../css/Parallax.css';
import fullImage from '../static/image-full.png';
import bottomImage from '../static/image-bottom.png';

import SC7 from '../static/2.png';
import SC8 from '../static/4.png';
import SC9 from '../static/4.1.png'; // 4k version 


function Parallax() {
  const navigate = useNavigate();

  const ref = useRef(null);
  const { scrollYProgress } = useScroll({
    target: ref,
    offset: ["start start", "end start"],
  });
  
  const backgroundY = useTransform(scrollYProgress, [0, 1], ["0%", "100%"]); // given to full image
  const textY = useTransform(scrollYProgress, [0, 1], ["0%", "300%"]); // given to text
  // const foregroundY = useTransform(scrollYProgress, [0, 1], ["0%", "100%"]);

  return (
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
      />

      <motion.h1 
        style={{ y: textY }}
        className="center-header"
      >
        PARALLAX
      </motion.h1>

      <motion.div
        className="bottom-image"
        style={{
          backgroundImage: `url(${bottomImage})`,
          backgroundPosition: "bottom",
          backgroundSize: "cover",
        }}
      />
      
    </div>
  );
};
export default Parallax;