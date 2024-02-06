import { Container, Row, Col } from 'react-bootstrap';
import Button from '@mui/material/Button';
import { useNavigate } from 'react-router-dom';
import '../css/ParaTest.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import NavBar from './NavBar';
import MovieSearchBar from './MovieSearchBar.js';
import Rectangle_Movies from '../static/Rectangle_Movies.jpg';
import Sqaure_Movies from '../static/Square_Movies.jpg';
import Square from '../static/Square.jpg';
import { ThemeProvider } from '@mui/material/styles';
import { muiTheme } from '../css/MuiThemeProvider.js';
import Parallax from './Parallax.js';

import SC1 from '../static/Showcase_1_big.jpg';
import SC2 from '../static/Showcase_2_big.jpg';
import SC3 from '../static/Showcase_3_small.jpg';
import SC4 from '../static/Showcase_11_small.jpg';
import SC5 from '../static/Showcase_12_small.jpg';
import SC6 from '../static/Showcase_13_small.jpg';


function ParaTest() {
  const navigate = useNavigate();

  const handleGetStartedClick = () => {
    navigate('/user-registration')
  };

  return (
    <main>
      <NavBar />
      <Parallax />
      <div className="background-parallax">
        <div className="text-details">
          <p>
            Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ipsam quae
            earum nobis quasi repellat. Amet facere nulla dolorum accusantium
            sit dolores odio excepturi facilis laboriosam officiis dolorem,
            nobis reprehenderit molestiae.
          </p>
          <p>
            Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ipsam quae
            earum nobis quasi repellat. Amet facere nulla dolorum accusantium
            sit dolores odio excepturi facilis laboriosam officiis dolorem,
            nobis reprehenderit molestiae.
          </p>
          <p>
            Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ipsam quae
            earum nobis quasi repellat. Amet facere nulla dolorum accusantium
            sit dolores odio excepturi facilis laboriosam officiis dolorem,
            nobis reprehenderit molestiae.
          </p>
          <p>
            Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ipsam quae
            earum nobis quasi repellat. Amet facere nulla dolorum accusantium
            sit dolores odio excepturi facilis laboriosam officiis dolorem,
            nobis reprehenderit molestiae.
          </p>
          <p>
            Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ipsam quae
            earum nobis quasi repellat. Amet facere nulla dolorum accusantium
            sit dolores odio excepturi facilis laboriosam officiis dolorem,
            nobis reprehenderit molestiae.
          </p>
          <p>
            Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ipsam quae
            earum nobis quasi repellat. Amet facere nulla dolorum accusantium
            sit dolores odio excepturi facilis laboriosam officiis dolorem,
            nobis reprehenderit molestiae.
          </p>
          <p>
            Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ipsam quae
            earum nobis quasi repellat. Amet facere nulla dolorum accusantium
            sit dolores odio excepturi facilis laboriosam officiis dolorem,
            nobis reprehenderit molestiae.
          </p>
        </div>
      </div>
    </main>
  );
};

export default ParaTest;