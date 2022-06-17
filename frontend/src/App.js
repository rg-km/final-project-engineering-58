import React from "react";
import Navigation from "./components/Navigation";
import Header from "./components/Header";
import Testimonials from "./components/Testimonials";
import Benefit from "./components/Benefit";
import Footer from "./components/Footer";

const App = () =>{
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

export default App