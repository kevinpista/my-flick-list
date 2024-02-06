import {BrowserRouter, Route, Routes} from 'react-router-dom'; 
import Home from './components/Home'
import ParaTest from './components/ParaTest'
import HomeTest from './components/HomeTest'
import Movie from './components/Movie'
import Test from './components/Test'
import TestTwo from './components/TestTwo'
import UserRegistration from './components/UserRegistration'
import UserLogin from './components/UserLogin'
import JwtTest from './components/JwtTest';
import Watchlist from './components/watchlist/Watchlist';
import ListOfWatchlists from './components/watchlist/ListOfWatchlists';
import MovieSearch from './components/MovieSearch';


function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact ={true} path="/" element={<Home />} />      
        <Route exact ={true} path="/para-test" element={<ParaTest />} />     
        <Route exact ={true} path="/home-test" element={<HomeTest />} />               
        <Route exact ={true} path="/movie/:movieID" element={<Movie />} />       
        <Route path="/test" element={<Test />} />  
        <Route exact ={true} path="/test-two" element={<TestTwo />} />  
        <Route exact ={true} path="/user-registration" element={<UserRegistration />} />                 
        <Route exact ={true} path="/user-login" element={<UserLogin />} />           
        <Route exact ={true} path="/jwt-test" element={<JwtTest />} />
        <Route exact ={true} path="/watchlist" element={<ListOfWatchlists />} />
        <Route exact={true} path="/watchlist/:watchlistID" element={<Watchlist />} />        
        <Route path="/movie-search" element={<MovieSearch />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;

//        <Route exact ={true} path="/watchlist" element={<ListOfWatchlists />} />     
