import axios from "axios";
import React, { Component } from "react";
import web3 from "./Web3";
import RoomFactory from "./contract/RoomFactory";
import styled, { createGlobalStyle } from "styled-components";
import reset from "styled-reset";
import { ArrowLeft, Plus } from "styled-icons/fa-solid";

// 書き出されたRoomFactoryのアドレス
const address = "0xB15FE0CED91C3cCf8D1C276C61c60B94ECa7b932";
const roomFactory = new web3.eth.Contract(RoomFactory.abi, address);

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      deposit: 0,
      rooms: [],
      isShowModal: false,
    };

    this.handleKeyUP = this.handleKeyUP.bind(this);
  }

  async componentDidMount() {
    // Room一覧APIのURL
    const res = await axios.get("http://localhost:8080/rooms");
    console.log(res);
    this.setState({ rooms: res.data });
  }

  handleKeyUP = (e) => {
    console.log(e.target.value);
    this.setState({ deposit: e.target.value });
  };

  handleCreateRoom = async () => {
    const [ethaddress] = await web3.eth.getAccounts();
    console.log(this.state.deposit.toString());
    try {
      roomFactory.methods.createRoom().send({
        from: ethaddress,
        value: web3.utils.toWei("1", "ether"),
      });
    } catch (error) {
      console.log(error);
    }
  };

  handleCancel = () => {
    this.setState({ deposit: "", isShowModal: false });
  };

  handleBack = () => {
    this.setState({ deposit: "", isShowModal: false });
  };

  handleModal = () => {
    this.setState({ isShowModal: true });
  };

  render() {
    const { rooms } = this.state;

    return (
      <>
        <GlobalStyle />
        {!this.state.isShowModal ? (
          // Rooms
          <>
            <Header>
              <h1>Rooms</h1>
            </Header>
            <Subtitle>
              <p>トークルーム一覧</p>
              <ModalButton onClick={this.handleModal}>
                <StyledPlus />
                Create Room
              </ModalButton>
            </Subtitle>
            {rooms &&
              rooms.length !== 0 &&
              rooms.map((room) => (
                <Item>
                  <Code>{room.event_code}</Code>
                  <Deposit>
                    <Label>デポジット額</Label>
                    <Eth>
                      {web3.utils.fromWei(room.wei_balance.toString(), "ether")}
                      ETH
                    </Eth>
                  </Deposit>
                </Item>
              ))}
          </>
        ) : (
          // Create Room
          <>
            <Header>
              <BackButton onClick={this.handleBack}>
                <ArrowLeft />
              </BackButton>
              <h1>Create Room</h1>
            </Header>
            <Contents>
              <p>デポジット額（ETH）を入力</p>
              <input
                type="text"
                onChange={this.handleKeyUp}
                placeholder="Initial Deposit (ETH)"
              />
            </Contents>
            <Footer>
              <button onClick={this.handleCancel} className="cancel">
                キャンセル
              </button>
              <button onClick={this.handleCreateRoom} className="save">
                保存
              </button>
            </Footer>
          </>
        )}
      </>
    );
  }
}

export default App;

// Style
const GlobalStyle = createGlobalStyle`
  ${reset}
  body {
    color: #4a4a4a;
  }
`;

const Header = styled.header`
  background: linear-gradient(to right, #9c56d3, #804aaa);
  padding: 20px;
  h1 {
    color: #fff;
    text-align: center;
  }
`;

const BackButton = styled.div`
  position: absolute;
  width: 18px;
  color: #fff;
  cursor: pointer;
`;

const Contents = styled.div`
  padding: 40px;
  > p {
    margin-bottom: 10px;
  }
  > input {
    align-items: center;
    border-radius: 4px;
    box-shadow: none;
    font-size: 1rem;
    height: 2.25em;
    line-height: 1.5;
    padding: 5px 10px;
    position: relative;
    vertical-align: top;
    border: 1px solid #ccc;
    display: block;
    width: 100%;
    box-sizing: border-box;
  }
`;

const Subtitle = styled.div`
  background: #f8f8f8;
  display: flex;
  padding: 10px;
  justify-content: space-between;
  > p {
    line-height: 3;
    font-size: 12px;
  }
`;

const ModalButton = styled.button`
  background-color: #23d160;
  border-color: transparent;
  cursor: pointer;
  color: #fff;
  padding: 10px 15px;
  border-radius: 50px;
  font-size: 16px;
  font-weight: 100;
`;

const StyledPlus = styled(Plus)`
  width: 16px;
  margin-right: 10px;
`;

const Item = styled.div`
  border-bottom: 1px solid #ccc;
  padding: 20px;
`;

const Code = styled.div`
  color: #363636;
  font-size: 20px;
  margin-bottom: 10px;
`;

const Deposit = styled.div`
  display: flex;
`;

const Label = styled.div`
  font-size: 11px;
  padding: 5px 7px;
  margin-right: 8px;
  border-radius: 5px;
  color: #4a4a4a;
  background: #eee;
  white-space: nowrap;
`;

const Eth = styled.div`
  align-items: center;
  display: inline-flex;
  flex-wrap: wrap;
  font-size: 20px;
  font-weight: 100;
  margin-right: 16px;
  color: #f29d7a;
`;

const Footer = styled.footer`
  border-top: 1px solid #ccc;
  background: #fafafa;
  display: -ms-flexbox;
  display: flex;
  padding: 10px;
  justify-content: center;
  > button {
    color: #fff;
    font-weight: bold;
    min-width: 120px;
    padding: 8px 24px;
    height: auto;
    font-size: 15px;
    border: none;
    border-radius: 4px;

    &.cancel {
      background: #f5a623;
      margin-right: 15px;
    }
    &.save {
      background: #6fbe4e;
    }
  }
`;
