import React from 'react';
import logo from './logo.svg';
import './App.css';
import { throwStatement } from '@babel/types';
import { Table } from 'antd';
import 'antd/dist/antd.css'; // or 'antd/dist/antd.less'
import AddForm from "./addForm";



let data = [
  {
    key: 1,
    name: 'John Brown',
    age: 32,
    address: 'New York No. 1 Lake Park',
    description: 'My name is John Brown, I am 32 years old, living in New York No. 1 Lake Park.',
  },
];


class Clock extends React.Component {
  
  state = {
    b: 12,
    // data: ""
    data
    
  }

   columns = [
    { title: 'Name', dataIndex: 'name', key: 'name' },
    { title: 'Age', dataIndex: 'age', key: 'age' },
    { title: 'Address', dataIndex: 'address', key: 'address' },
    {
      title: 'Action',
      dataIndex: '',
      key: 'x',
      render: (_,r) => <a onClick={(e)=>{console.log(r.key);
      const {data} = this.state;
      this.setState({
        data:data.filter((e)=>r.key!=e.key)
      })
      localStorage.setItem("data",JSON.stringify(this.state.data))
      }}>Delete</a>,
    },
  ];
  flashMark = true
  componentDidMount (){
    // let ws = new WebSocket("ws://localhost:8080/time");
    // ws.onmessage = (evt)  => {  
    //   console.log(evt.data)
    //   this.setState({
    //     data: evt.data
    //   })
    //  };
    setInterval(()=>{
      this.setState({b:this.state.b+1})
    },1000)
    this.readbackFromLS()
    }

  updateToLS(){
    if(this.flashMark){
      this.flashMark=false
      return
    }
    let current = new Date().getTime()
    localStorage.setItem("datatime",current)
    localStorage.setItem("data",JSON.stringify(this.state.data))
  }

  readbackFromLS(){
    let olddata = JSON.parse(localStorage.getItem("data"))
    this.setState({
      data: olddata,
    })
  }

  render() {
    this.updateToLS()
    return (
      <div>
        <h1>Hello, world!</h1>
        <h2>It is clock.  {this.state.b}</h2>
    {/*<p>{this.state.data}</p>*/}
      <Table
      columns={this.columns}
      expandedRowRender={record => <p style={{ margin: 0 }}>{record.description}</p>}
      dataSource={this.state.data}
    />
    <AddForm handleSubmit={(values)=>{
      const { data } = this.state;
      console.log([...data,{
        ...values,
        key: this.state.data.length+1
      }])
     this.setState({
       data: [...data,{
         ...values,
         key: this.state.data.length+1
       }]
     })
     
    }}/>
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
