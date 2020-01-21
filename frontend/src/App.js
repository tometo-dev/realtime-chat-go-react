import React from 'react'
import './App.css'
import {sendMsg} from "./api"

function App() {

    const send = () => {
        console.log("hello")
        sendMsg("hello")
    }

    return (
        <div className="App">
            <button onClick={send}>Hit</button>
        </div>
    )
}

export default App
