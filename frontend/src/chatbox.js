import React from 'react';

export class ChatBox extends React.Component {
  state = {
    show: true,
  }

  scrollToBottom = () => {
    this.messagesEnd.scrollIntoView({ behavior: "smooth" });
  }
  componentDidMount() {
    this.scrollToBottom();
  }

  componentDidUpdate() {
    this.scrollToBottom();
  }

  renderChatData() {
    let data = this.props.chatData;
    return data.map((s, i)=>{
      return(<p className='chatmsg' key={i}>
        <font size="1" color="green">[{new Date(s.getTime()).toLocaleTimeString()}]</font>
        <font color="blue">{s.getNickname()}: </font>
        {s.getMsg()}
        </p>);
    });
  }

  hideChatBox() {
    this.setState({show: !this.state.show})
    setTimeout(this.scrollToBottom, 400);
  }

  render() {
    return (
      <div>
        <div className="chatbox">
          <h3> Chat room </h3>
          <div className="chatdata" style={{ height: (this.state.show ? 400 : 80) }}>
            <div>
            {this.renderChatData()}
            </div>
            <div style={{ float:"left", clear: "both" }}
              ref={(el) => { this.messagesEnd = el; }}>
            </div>
          </div>
          <input 
            value={this.props.value} 
            onChange={(e)=>{this.props.handleMsgChange(e)}}
            onKeyUp={(e)=>{
              if(e.key==='Enter'){
                this.props.onClick();
              }
            }}
          />
          <button 
            onClick={()=>{this.props.onClick()}}
          >send</button>
        </div>
      <button
        className="hidechatboxbutton"
        onClick={() => { this.hideChatBox() }}
      >{this.state.show ? "hide" : "show"}</button>
      </div>
    );
  }
}