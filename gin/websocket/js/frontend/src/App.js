import React from 'react';
import logo from './logo.svg';
import './App.css';
import { throwStatement } from '@babel/types';
import { Table } from 'antd';

const columns = [
  { title: 'Name', dataIndex: 'name', key: 'name' },
  { title: 'Age', dataIndex: 'age', key: 'age' },
  { title: 'Address', dataIndex: 'address', key: 'address' },
  {
    title: 'Action',
    dataIndex: '',
    key: 'x',
    render: () => <a>Delete</a>,
  },
];

const data = [
  {
    key: 1,
    name: 'John Brown',
    age: 32,
    address: 'New York No. 1 Lake Park',
    description: 'My name is John Brown, I am 32 years old, living in New York No. 1 Lake Park.',
  },
  {
    key: 2,
    name: 'Jim Green',
    age: 42,
    address: 'London No. 1 Lake Park',
    description: 'My name is Jim Green, I am 42 years old, living in London No. 1 Lake Park.',
  },
  {
    key: 3,
    name: 'Joe Black',
    age: 32,
    address: 'Sidney No. 1 Lake Park',
    description: 'My name is Joe Black, I am 32 years old, living in Sidney No. 1 Lake Park.',
  },
];


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
      
    //    <Table
    //    columns={columns}
    //    expandedRowRender={record => <p style={{ margin: 0 }}>{record.description}</p>}
    //    dataSource={data}
    //  />

      
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
