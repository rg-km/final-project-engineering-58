import React from "react";
import Navigation from "../components/Navigation";
import Header from "../components/Header";
import Testimonials from "../components/Testimonials";
import Benefit from "../components/Benefit";
import Footer from "../components/Footer";

export default function LandingPage() {
    return (
        <div>
            <Navigation />
            <Header />
            <Testimonials />
            <Benefit />
            <Footer />
        </div>
    )
}