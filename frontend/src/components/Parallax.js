import React, { useRef } from "react";
import { useNavigate } from 'react-router-dom';
import NavBar from './NavBar.js';
import { motion, useScroll, useTransform } from "framer-motion";
import '../css/Parallax.css';
import fullImage from '../static/image-full.png';
import bottomImage from '../static/image-bottom.png';

import SC6 from '../static/Showcase_13_small.jpg';


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
      <motion.h1 
        style={{ y: textY }}
        className="center-header"
      >
        PARALLAX
      </motion.h1>

      <motion.div
        className="full-image "
        style={{
          backgroundImage: `url(${fullImage})`,
          backgroundPosition: "bottom",
          backgroundSize: "cover",
          y: backgroundY,
        }}
      />

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