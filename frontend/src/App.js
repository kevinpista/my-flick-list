import {BrowserRouter, Route, Routes} from 'react-router-dom'; 
import Home from './components/Home'
import Movie from './components/Movie'
import Test from './components/Test'



function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact ={true} path="/" element={<Home />} />        
        <Route exact ={true} path="/movie" element={<Movie />} />       
        <Route exact ={true} path="/test" element={<Test />} />           
      </Routes>
    </BrowserRouter>
  );
}

export default App;

/*
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <h1>
          Hello world, I'm back!
        </h1>
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;

*/