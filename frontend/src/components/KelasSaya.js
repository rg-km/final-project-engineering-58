import React, { useState } from "react";
import "boxicons";
import "../assets/css/KelasSaya.css";
import Avatar from "../assets/images/kelas-1.jpg";
import Avatar2 from "../assets/images/Ani.png";
import Avatar3 from "../assets/images/Adi.png";
import NavSiswa from "./NavSiswa";
import { Link } from "react-router-dom"
import "../assets/css/dashboard.css";

export default function TestiLanding() {

  return (
    <div className="container-all-kelas-1">
        <div className="container-all-kelas-2">
            <NavSiswa />
            <div className="container-all-kelas">
            <h1>Kelas Saya</h1>
                <div className="container-kelas">
                    <Link to="/video" className="link-kelas">
                        <div className="card-kelas">
                            <div className="content-kelas">
                                <div className="image">
                                <img alt="" src={Avatar} width="250"/>
                                </div>
                                <div className="details-kelas">
                                <h2>Lorem Ipsum</h2>
                                </div>
                                <p>Lorem Ipsum Dolor Sit Amet</p>
                            </div>
                        </div>
                    </Link>

                    <div className="card-kelas">
                        <div className="content-kelas">
                            <div className="image">
                            <img alt="" src={Avatar} width="250"/>
                            </div>
                            <div className="details-kelas">
                            <h2>Lorem Ipsum</h2>
                            </div>
                            <p>Lorem Ipsum Dolor Sit Amet</p>
                        </div>
                    </div>

                    <div className="card-kelas">
                        <div className="content-kelas">
                            <div className="image">
                            <img alt="" src={Avatar} width="250"/>
                            </div>
                            <div className="details-kelas">
                            <h2>Lorem Ipsum</h2>
                            </div>
                            <p>Lorem Ipsum Dolor Sit Amet</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
  );
}
