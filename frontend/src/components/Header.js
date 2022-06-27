import React from "react";
import "boxicons";
import "../assets/css/Header.css"
import Logo from "../assets/images/Logo Text Light.svg";
import LogoText from "../assets/images/Logo.png";


export default function Header() {
    return (
        
        
        //     <div className="container-header">
        //         <nav role="navigation">
        //         <div className="container">
                
        //             <img className="logo-nav" src={LogoText} alt="Logo HaloBelajar"></img>
        //             <ul class="menu-nav" role="list">
        //                 <li className="log-in"><a href="#">Log In</a></li>
        //                 <li className="sign-up"><a href="#">Sign Up</a></li>
        //             </ul>
        //             </div>
        //         </nav>
        //         <div className="container">
        //         <section className="content-header">
        //             <article>
        //                 <h1>Belajar seru dan mudah dipahami</h1>
        //                 <p>
        //                     HaloBelajar merupakan sebuah platform yang dapat membantumu
        //                     dalam belajar. Ada berbagai materi yang dapat kamu akses
        //                     untuk menambah pengetahuanmu.
        //                 </p>
        //                 <button className="btn btn-try">Try for free</button>
        //             </article>
        //             <img className="logo-header" src={Logo} alt="Logo HaloBelajar"></img>
        //         </section>
        //     </div>
        // </div>
       <div className="container-header">
         <nav class="navbar navbar-expand-lg bg-dark navbar-dark fixed-top">
  <div class="container">
  <a class="navbar-brand" href="#">
  <img src={Logo} alt="logo"></img>
    </a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav">
        <li class="nav-item">
          <a class="nav-link active" aria-current="page" href="#">Home</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">Kelas</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">Tentang Kami</a>
        </li>
        <li class="nav-item">
          <a class="nav-link disabled">Daftar/Masuk</a>
        </li>
      </ul>
    </div>
  </div>
</nav>
<section className="content-header">
    <div className="container">
    <article className="text-white">
                         <h1><img className="logo-header" src={LogoText} alt="Logo HaloBelajar"></img>Belajar seru dan mudah dipahami</h1>
                         <p>
                             HaloBelajar merupakan sebuah platform yang dapat membantumu
                             dalam belajar. Ada berbagai materi yang dapat kamu akses
                             untuk menambah pengetahuanmu.
                         </p>
                         <button className="btn btn-try">Try for free</button>
                     </article>
                     </div>
                     
                     
                     
                     
                 </section>
       </div>
    )
}