// Home.js 
import { Container, Row, Col } from 'react-bootstrap';
import Button from '@mui/material/Button';
import { useNavigate } from 'react-router-dom';
import '../css/HomeTest.css';
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
import SC7 from '../static/first-image.png';


function HomeTest() {
  const navigate = useNavigate();

  const handleGetStartedClick = () => {
    navigate('/user-registration')
  };

  return (
    <div className="home">
      <NavBar />
      <Parallax />
        <div className="features-container">
            <div className="features">

                <div className="first-feature">
                    <div className="feature-info">
                    <div class="feature-subtitle">Get Started</div>
                        <div class="feature-headline">
                            Search for Any Movie
                        </div>
                        <div class="feature-body">
                            Determining what level of hiker you are can be an important
                            tool when planning future hikes. This hiking level guide will
                            help you plan hikes according to different hike ratings set by
                            various websites like All Trails and Modern Hiker. What type
                            of hiker are you – novice, moderate, advanced moderate,
                            expert, or expert backpacker?
                        </div>

                        <div class="feature-info-footer">
                            <a href="#">
                                <button>
                                    Try Search Tool
                                </button></a>
                        </div>   

                    </div>
                    
                    <div className="feature-image">
                        <img src={SC7} />
                    </div>
                </div>


            </div>
            
        </div>
    </div>

  );
}

export default HomeTest;