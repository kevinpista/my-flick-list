import {BrowserRouter, Route, Routes} from 'react-router-dom'; 
import Home from './components/Home'
import Movie from './components/Movie'
import Test from './components/Test'
import UserRegistration from './components/UserRegistration'
import UserLogin from './components/UserLogin'


function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact ={true} path="/" element={<Home />} />        
        <Route exact ={true} path="/movie" element={<Movie />} />       
        <Route exact ={true} path="/test" element={<Test />} />  
        <Route exact ={true} path="/user-registration" element={<UserRegistration />} />                 
        <Route exact ={true} path="/user-login" element={<UserLogin />} />                    
   
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