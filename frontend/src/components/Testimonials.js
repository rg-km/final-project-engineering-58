import React from "react";
import "boxicons";
import "../assets/css/Testimonials.css"

export default function Testimonials() {
    return (
        
        <div className="container-all-testi">
            <h1>Testimoni di HaloBelajar</h1>
            <div className="container-testimonials">
                <div className="card-testimonials">
                    <div className="layer-testimonials"></div>
                    <div className="content-testimonials">
                        <div class="image">
                            <img src="img.jpg"></img>
                        </div>
                        <div class="details-testimonials">
                            <h2>Budi<br></br><span>Front End</span></h2>
                        </div>
                        <p>
                            Generate meaningful discussions with your audience and build a
                            strong, loyal community. Think of the insightful conversations you
                            miss out on with a feedback form.
                        </p>
                    </div>
                </div>

                <div className="card-testimonials">
                    <div className="layer-testimonials"></div>
                    <div className="content-testimonials">
                        <div class="image">
                            <img src="img.jpg"></img>
                        </div>
                        <div class="details-testimonials">
                            <h2>Ani<br></br><span>Back End</span></h2>
                        </div>
                        <p>
                            Generate meaningful discussions with your audience and build a
                            strong, loyal community. Think of the insightful conversations you
                            miss out on with a feedback form.
                        </p>
                    </div>
                </div>

                <div className="card-testimonials">
                    <div className="layer-testimonials"></div>
                    <div className="content-testimonials">
                        <div class="image">
                            <img src="img.jpg"></img>
                        </div>
                        <div class="details-testimonials">
                            <h2>Adi<br></br><span>Full Stack</span></h2>
                        </div>
                        <p>
                            Generate meaningful discussions with your audience and build a
                            strong, loyal community. Think of the insightful conversations you
                            miss out on with a feedback form.
                        </p>
                    </div>
                </div>
            </div>
        </div>
    )
}