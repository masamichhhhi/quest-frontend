import logo from "./logo.svg";
import "./App.css";
import web3 from "./Web3";
import RoomFactory from "./contract/RoomFactory.json";
import { useEffect, useState } from "react";
import axios from "axios";

const address = "0xF24f828A3F077C416Cb7D9A1E92A1055390b37c9";
const roomFactory = new web3.eth.Contract(RoomFactory.abi, address);

function App() {
  const [rooms, setRooms] = useState([]);
  const [deposit, setDeposit] = useState(0);

  const handleKeyUp = (e) => {
    setDeposit(e.target.value);
  };

  const handleCreateRoom = async () => {
    const [ethaddress] = await web3.eth.getAccounts();

    try {
      const receipt = await roomFactory.methods.createRoom().send({
        from: ethaddress,
        value: web3.utils.toWei(deposit.toString(), "ether"),
      });
    } catch (error) {
      console.log(error);
    }
  };

  const getRooms = async () => {
    const roomlist = await axios.get("");
    setRooms(roomlist);
  };

  useEffect(() => {
    getRooms();
  }, []);

  return (
    <div>
      <input onKeyUp={handleKeyUp} placeholder="Initial Deposit (ETH)" />
      <button onClick={handleCreateRoom}> Create Room</button>
    </div>
  );
}

export default App;
