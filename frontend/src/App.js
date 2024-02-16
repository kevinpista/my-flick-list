import {BrowserRouter, Route, Routes} from 'react-router-dom'; 
import Home from './components/Home'
import Movie from './components/Movie'
import UserRegistration from './components/UserRegistration'
import UserLogin from './components/UserLogin'
import Watchlist from './components/watchlist/Watchlist';
import ListOfWatchlists from './components/watchlist/ListOfWatchlists';
import MovieSearch from './components/MovieSearch';
import About from './components/About';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact ={true} path="/" element={<Home />} />      
        <Route exact ={true} path="/movie/:movieID" element={<Movie />} />       
        <Route exact ={true} path="/user-registration" element={<UserRegistration />} />                 
        <Route exact ={true} path="/user-login" element={<UserLogin />} />           
        <Route exact ={true} path="/watchlist" element={<ListOfWatchlists />} />
        <Route exact={true} path="/watchlist/:watchlistID" element={<Watchlist />} />        
        <Route path="/movie-search" element={<MovieSearch />} />
        <Route path="/about" element={<About />} />
      </Routes>
    </BrowserRouter>
  );
}
export default App;