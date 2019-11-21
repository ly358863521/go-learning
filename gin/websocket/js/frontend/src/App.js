import React from 'react';
import logo from './logo.svg';
import './App.css';
import { throwStatement } from '@babel/types';

class Clock extends React.Component {
  
  
  state = {
    b: 12,
    data: ""
  }
  componentDidMount (){
    let ws = new WebSocket("ws://localhost:8080/time");
    ws.onmessage = (evt)  => {  
      console.log(evt.data)
      this.setState({
        data: evt.data
      })
     };
    setInterval(()=>{
      this.setState({b:this.state.b+1})
    },1000)
    }


  render() {
    return (
      <div>
        <h1>Hello, world!</h1>
        <h2>It is clock.  {this.state.b}</h2>
    <p>{this.state.data}</p>
      </div>
    );
  }
}



// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//       <Clock/>

//     </div>
//   );
// }

function App() {
  return (
    <div className="App">
      
      <Clock/>

    </div>
  );
}



export default App;
