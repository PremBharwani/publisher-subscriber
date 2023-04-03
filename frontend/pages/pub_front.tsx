import React, { useEffect } from "react";
import {init} from "./Web3client";


import Button from "../components/CustomButtonComponent";

var clickResponse = () => {
    console.log("clicked");
}
function App() {
    useEffect(()=>{
        init();
    }, []);

  return (
    <>
      <h1>you are a publisher</h1>

      <form method="POST" action="login"> 
        <input type="text" name="address" placeholder="address" />
        <input type="text" name="name" placeholder="name" />
        <input type="text" name="id" placeholder="id" />
        <Button 
            border="solid"
            color="#ffffff"
            height = "50px"
            onClick={()=>{clickResponse}}
            radius = "10%"
            width = "200px"
            children = "Click Here"
        />

     </form>

    </>
  );
}

export default App;