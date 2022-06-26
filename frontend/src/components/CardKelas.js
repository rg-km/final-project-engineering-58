import React, { useState } from "react";
import "../assets/css/KelasSaya.css";
import Avatar from "../assets/images/kelas-1.jpg";
import "boxicons";
import { Link } from "react-router-dom"

export default function CardKelas() {

    const [image, setImage] = useState(null);
    const [title, setTitle] = useState("Lorem Ipsum");
    const [subtitle, setSubtitle] = useState("Lorem Ipsum Dolor Sit Amet");

    return(
        <div className="container-kelas">
            <Link to="/video" className="link-kelas">
                <div className="card-kelas">
                    <div className="content-kelas">
                        <div className="image">
                            <img alt="" src={image} width="250"/>
                        </div>
                        <div className="details-kelas">
                            <h2>{title}</h2>
                        </div>
                        <p>{subtitle}</p>
                    </div>
                </div>
            </Link>
        </div>
    )
}