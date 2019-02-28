import React from 'react';

export class ChatBox extends React.Component {
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
      return(<p className='chatmsg' key={i}>{s}</p>);
    });
  }

  render() {
    return (
      <div className="chatbox">
        <h3> Chat room </h3>
        <div>
          Username:
          <input
            onChange={(e)=>{this.props.handleUserName(e)}}
            placeholder="somebody"
          />
        </div>
        <div className="chatdata">
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
    );
  }
}