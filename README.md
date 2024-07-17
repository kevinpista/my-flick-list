<a name="readme-top"></a>

<br />
<div align="center">
  <a href="https://myflicklist-fa78f7f017a1.herokuapp.com/">
    <img src="frontend/public/logo_circle.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">My Flick List</h3>
  <p>
     <a href="https://myflicklist-fa78f7f017a1.herokuapp.com/"> > VIEW LIVE DEMO < </a>
  </p>
  <p align="center">
      <b>'My Flick List'</b> is a web application created as a personal project. It allows a registered user to organize movies within
      personalized watchlists, track what they've watched, and add personal notes for each movie.
      The web app allows a user to search for any movie thanks to the data provided by <a href="https://www.themoviedb.org/"><b>'TMDb API'</b></a>.
      Every movie will contain information such as the title, overview, release date, trailer video,
      along with other metrics such as its budget and revenue generated in the box office. A user can create a custom watchlist
      and add any movie to it. Within each watchlist, the user can check off whenever they finish watching a movie and also write their notes for each film.
      These functions allow a cinema enthusiast to gather all their movies and thoughts in one place in order to make their watching
      experience more organized and enjoyable.
    <br />
  </p>
  <p>
    This project implements a Go backend and a React.js frontend with the <a href="#Model-View-Controller-Implementation">Model-View-Controller (MVC)</a> architecture pattern. Additionally, it utilizes a Heroku PostgreSQL
    database and is fully hosted on Heroku. 
  </p>
  
  <p>
     <b>NOTE:</b> This project has NOT been optimized for mobile screens. For the best experience, view it on a laptop or computer.
  </p>
  
  <p>
     <a href="https://myflicklist-fa78f7f017a1.herokuapp.com/"> > VIEW LIVE DEMO < </a>
  </p>
  
</div>



<!-- ABOUT THE PROJECT -->
## About The Project

[![My Flick List Screen Shot 1][home-page-screenshot]](https://myflicklist-fa78f7f017a1.herokuapp.com/)



### Built With

* [![Go][Go.dev]][Go-url]
* [![React][React.js]][React-url]
* [![PostgreSQL][PSQL.com]][PSQL-url]
* [![JWT-Go][JWT]][JWT-Go-url]
* [![MUI][MUI.com]][MUI-url]
* [![React Router][ReactRouter.com]][ReactRouter-url]
* [![Framer Motion][Framer.com]][Framer-url]
* [![Postman][Postman]][Postman-url]
* [![Heroku][Heroku.com]][Heroku-url]
* [![TMDbAPI][TMDbAPI.com]][TMDbAPI-url]


[![My Flick List Screen Shot 2][movie-page-screenshot]](https://myflicklist-fa78f7f017a1.herokuapp.com/)


<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Model-View-Controller Implementation

This application implements the Model-View-Controller (MVC) architectural pattern across its entire stack using a Go backend and a React.js frontend.

* **Controller:** The Go backend serves as the Model and Controller layer. I used a routing framework called Chi router to direct incoming API requests to the appropriate controllers. Each controller focuses on a specific RESTful API endpoint or set of related endpoints. The controllers then validate the request and data and makes the respective service function calls. Afterwards, the controllers will return a response for the frontend to update its View for the user accordingly.
* **Model:** The backend utilizes structs (models) to represent data entities. These Models are used by the Controller to decode incoming JSON data. The Models then serve as a layer of abstraction, providing a clean interface for the service functions to handle the core business logic & interactions with the PostgreSQL database.
* **View:** The React.js frontend serves as the View layer, presenting the user interface and facilitating user interactions in sending CRUD requests to the Go backend. It receives and displays the appropriate data and success or error messages based on the backend's response.

By leveraging the MVC pattern, this project aims to promote separation of concerns, modularity, and maintainability throughout the entire application.

[![My Flick List Screen Shot 3][watchlist-page-screenshot]](https://myflicklist-fa78f7f017a1.herokuapp.com/)


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/github_username/repo_name.svg?style=for-the-badge
[contributors-url]: https://github.com/github_username/repo_name/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/github_username/repo_name.svg?style=for-the-badge
[forks-url]: https://github.com/github_username/repo_name/network/members
[stars-shield]: https://img.shields.io/github/stars/github_username/repo_name.svg?style=for-the-badge
[stars-url]: https://github.com/github_username/repo_name/stargazers
[issues-shield]: https://img.shields.io/github/issues/github_username/repo_name.svg?style=for-the-badge
[issues-url]: https://github.com/github_username/repo_name/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/github_username/repo_name/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[home-page-screenshot]: frontend/public/home_page.png
[movie-page-screenshot]: frontend/public/movie_page.png
[watchlist-page-screenshot]: frontend/public/watchlist_page.png


[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Go.dev]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev
[MUI.com]: https://img.shields.io/badge/MUI-%230081CB.svg?style=for-the-badge&logo=mui&logoColor=white
[MUI-url]: https://go.dev](https://mui.com/
[JWT]: https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens
[JWT-Go-url]: https://pkg.go.dev/github.com/golang-jwt/jwt
[PSQL.com]: https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white
[PSQL-url]: https://www.postgresql.org/
[ReactRouter.com]: https://img.shields.io/badge/React_Router-CA4245?style=for-the-badge&logo=react-router&logoColor=white
[ReactRouter-url]: https://reactrouter.com/en/main
[Framer.com]: https://img.shields.io/badge/Framer-black?style=for-the-badge&logo=framer&logoColor=blue
[Framer-url]: https://www.framer.com/motion/
[Heroku.com]: https://img.shields.io/badge/heroku-%23430098.svg?style=for-the-badge&logo=heroku&logoColor=white
[Heroku-url]: https://www.heroku.com/
[TMDbAPI.com]: https://img.shields.io/static/v1?style=for-the-badge&message=The+Movie+Database&color=222222&logo=The+Movie+Database&logoColor=01B4E4&label=TMDb%20API
[TMDbAPI-url]: https://www.themoviedb.org/
[Postman]: https://img.shields.io/static/v1?style=for-the-badge&message=Postman&color=FF6C37&logo=Postman&logoColor=FFFFFF&label=
[Postman-url]: https://www.postman.com/

[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 
