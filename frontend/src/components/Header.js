import React from "react";
import "boxicons";
import "../assets/css/Header.css"
import Logo from "../assets/images/Logo Text Light.svg";
import LogoText from "../assets/images/Logo Text.svg";

export default function Header() {
    return (
        
        <div className="container">
            <div className="container-header">
                <nav role="navigation">
                    <img className="logo-nav" src={LogoText} alt="Logo HaloBelajar"></img>
                    <ul class="menu-nav" role="list">
                        <li className="log-in"><a href="#">Log In</a></li>
                        <li className="sign-up"><a href="#">Sign Up</a></li>
                    </ul>
                </nav>

                <section className="content-header">
                    <article>
                        <h1>Belajar seru dan mudah dipahami</h1>
                        <p>
                            HaloBelajar merupakan sebuah platform yang dapat membantumu
                            dalam belajar. Ada banyak materi yang dapat kamu akses
                            untuk menambah pengetahuanmu.
                        </p>
                        <button className="btn btn-try">Try for free</button>
                    </article>
                    <img className="logo-header" src={Logo} alt="Logo HaloBelajar"></img>
                </section>
            </div>
        </div>
    )
}