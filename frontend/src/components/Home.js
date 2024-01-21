// Home.js 
import { Button, Container, Row, Col } from 'react-bootstrap';
import '../css/Home.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import NavBar from './NavBar';
import MovieSearchBar from './MovieSearchBar.js';

function Home() {

  return (
    <div className="home">
      <NavBar />

      {/* Custom Jumbotron */}
      <div className="jumbotron">
        <h1>My Flick List </h1>
        <p>
          Welcome to my flick list. A movie watch list app that allows you to seamlessly
          add movies to a watchlist, track your notes, and track off when you finish watching it.
        </p>
        <Button variant="primary" size="lg">
          Get Started 
        </Button>
      </div>

      {/* React Bootstrap */}
      <Container>
        <MovieSearchBar />
        <Row>
          <Col md={4}>
            <h2>Search</h2>
            <p>Search for new movies to add to your watchlist.</p>
          </Col>

          <Col md={4}>
            <h2>Watchlist</h2>
            <p>View and manage your personal watchlist.</p>
          </Col>

          <Col md={4}>
            <h2>History</h2>
            <p>See what movies you've already watched.</p>
          </Col>
        </Row>
      </Container>

    </div>
  );
}

export default Home;